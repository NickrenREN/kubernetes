// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResource) DeepCopyInto(out *ExtendedResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResource.
func (in *ExtendedResource) DeepCopy() *ExtendedResource {
	if in == nil {
		return nil
	}
	out := new(ExtendedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtendedResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResourceClaim) DeepCopyInto(out *ExtendedResourceClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResourceClaim.
func (in *ExtendedResourceClaim) DeepCopy() *ExtendedResourceClaim {
	if in == nil {
		return nil
	}
	out := new(ExtendedResourceClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtendedResourceClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResourceClaimList) DeepCopyInto(out *ExtendedResourceClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExtendedResourceClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResourceClaimList.
func (in *ExtendedResourceClaimList) DeepCopy() *ExtendedResourceClaimList {
	if in == nil {
		return nil
	}
	out := new(ExtendedResourceClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtendedResourceClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResourceClaimSpec) DeepCopyInto(out *ExtendedResourceClaimSpec) {
	*out = *in
	in.MetadataRequirements.DeepCopyInto(&out.MetadataRequirements)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResourceClaimSpec.
func (in *ExtendedResourceClaimSpec) DeepCopy() *ExtendedResourceClaimSpec {
	if in == nil {
		return nil
	}
	out := new(ExtendedResourceClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResourceClaimStatus) DeepCopyInto(out *ExtendedResourceClaimStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResourceClaimStatus.
func (in *ExtendedResourceClaimStatus) DeepCopy() *ExtendedResourceClaimStatus {
	if in == nil {
		return nil
	}
	out := new(ExtendedResourceClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResourceList) DeepCopyInto(out *ExtendedResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExtendedResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResourceList.
func (in *ExtendedResourceList) DeepCopy() *ExtendedResourceList {
	if in == nil {
		return nil
	}
	out := new(ExtendedResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtendedResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResourceSpec) DeepCopyInto(out *ExtendedResourceSpec) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(ResourceNodeAffinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResourceSpec.
func (in *ExtendedResourceSpec) DeepCopy() *ExtendedResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ExtendedResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedResourceStatus) DeepCopyInto(out *ExtendedResourceStatus) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
	out.Allocatable = in.Allocatable.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedResourceStatus.
func (in *ExtendedResourceStatus) DeepCopy() *ExtendedResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ExtendedResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceNodeAffinity) DeepCopyInto(out *ResourceNodeAffinity) {
	*out = *in
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.NodeSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceNodeAffinity.
func (in *ResourceNodeAffinity) DeepCopy() *ResourceNodeAffinity {
	if in == nil {
		return nil
	}
	out := new(ResourceNodeAffinity)
	in.DeepCopyInto(out)
	return out
}
