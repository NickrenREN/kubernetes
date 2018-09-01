/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package suningpod

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apiserver/pkg/admission"
	api "k8s.io/kubernetes/pkg/apis/core"
	informers "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion"
	corelisters "k8s.io/kubernetes/pkg/client/listers/core/internalversion"
	storagelisters "k8s.io/kubernetes/pkg/client/listers/storage/internalversion"
	kubeapiserveradmission "k8s.io/kubernetes/pkg/kubeapiserver/admission"
	"k8s.io/kubernetes/pkg/volume/util"
)

const (
	// PluginName is the name of this admission controller plugin
	PluginName = "SuningPod"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		plugin := newPlugin()
		return plugin, nil
	})
}

// suningPodPlugin holds state for and implements the admission plugin.
type suningPodPlugin struct {
	*admission.Handler

	pvcLister corelisters.PersistentVolumeClaimLister
	scLister  storagelisters.StorageClassLister
}

var _ admission.Interface = &suningPodPlugin{}
var _ = kubeapiserveradmission.WantsInternalKubeInformerFactory(&suningPodPlugin{})

// newPlugin creates a new admission plugin.
func newPlugin() *suningPodPlugin {
	return &suningPodPlugin{
		Handler: admission.NewHandler(admission.Create),
	}
}

func (c *suningPodPlugin) SetInternalKubeInformerFactory(f informers.SharedInformerFactory) {
	pvcInformer := f.Core().InternalVersion().PersistentVolumeClaims()
	c.pvcLister = pvcInformer.Lister()
	scInformer := f.Storage().InternalVersion().StorageClasses()
	c.scLister = scInformer.Lister()
	c.SetReadyFunc(func() bool {
		return pvcInformer.Informer().HasSynced() && scInformer.Informer().HasSynced()
	})
}

// ValidateInitialization ensures lister is set.
func (c *suningPodPlugin) ValidateInitialization() error {
	if c.pvcLister == nil {
		return fmt.Errorf("missing PVC lister")
	}
	if c.scLister == nil {
		return fmt.Errorf("missing SC lister")
	}
	return nil
}

// Admit calculates the local storage request and set it to the pod's annotation
// The key of the annotation is: localstoragerequest
func (c *suningPodPlugin) Admit(a admission.Attributes) error {
	if len(a.GetSubresource()) != 0 || a.GetResource().GroupResource() != api.Resource("pods") || a.GetOperation() != admission.Create {
		return nil
	}

	pod, ok := a.GetObject().(*api.Pod)
	if !ok {
		return errors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}

	localPVCRequest, err := c.getLocalStorageRequest(pod)
	if err != nil {
		return errors.NewBadRequest("get local storage request from Pod error")
	}

	if pod.Annotations == nil {
		pod.Annotations = make(map[string]string)
	}
	pvcRequestBytes, err := json.Marshal(localPVCRequest)
	if err != nil {
		return errors.NewBadRequest("marshal pvc request map error")
	}
	pod.Annotations["localstoragerequest"] = string(pvcRequestBytes)
	return nil
}

func (c *suningPodPlugin) getLocalStorageRequest(pod *api.Pod) (map[string]string, error) {
	localPVCRequestMap := map[string]string{}
	for _, vol := range pod.Spec.Volumes {
		if vol.PersistentVolumeClaim != nil {
			pvc, err := c.pvcLister.PersistentVolumeClaims(pod.Namespace).Get(vol.PersistentVolumeClaim.ClaimName)
			if err != nil {
				glog.Errorf("get PVC error: %v", err)
				return nil, err
			}
			isLocalVolume, pvcrequest, err := c.getLocalStorageReqeust(pvc)
			if err != nil {
				return nil, err
			}
			if isLocalVolume {
				localPVCRequestMap[string(pvc.UID)] = pvcrequest.String()
			}
			//request.Add(pvcrequest)
		}
	}

	return localPVCRequestMap, nil
}

func (c *suningPodPlugin) getLocalStorageReqeust(pvc *api.PersistentVolumeClaim) ( /*isLocalVolume*/ bool, resource.Quantity, error) {
	// only calculate local storage dynamic provisioning request
	if pvc.Spec.StorageClassName == nil {
		return false, resource.Quantity{}, nil
	}
	sc, err := c.scLister.Get(*pvc.Spec.StorageClassName)
	if err != nil {
		return false, resource.Quantity{}, err
	}
	if sc.Provisioner != "local-volume-provisioner" {
		// not request local storage
		return false, resource.Quantity{}, nil
	}

	// round up to GiB
	request := util.RoundUpToGiB(pvc.Spec.Resources.Requests[api.ResourceStorage])

	return true, *resource.NewQuantity(request, resource.BinarySI), nil
}
