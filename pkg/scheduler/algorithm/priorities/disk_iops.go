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

package priorities

import (
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"

	"github.com/golang/glog"
)

func NewDiskIOPSPriority() algorithm.PriorityFunction {
	return CalculateDiskIOPSPriority
}

var DiskIOPSLeft map[string]int64

func CalculateDiskIOPSPriority(pod *v1.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo, nodes []*v1.Node) (schedulerapi.HostPriorityList, error) {
	DiskIOPSLeft = make(map[string]int64, len(nodes))
	for nodeName, nodeInfo := range nodeNameToInfo {
		DiskIOPSLeft[nodeName] = nodeInfo.AllocatableResource().DiskIOPS - nodeInfo.RequestedResource().DiskIOPS
	}

	var maxLeft int64
	var minLeft int64

	for _, node := range nodes {
		if DiskIOPSLeft[node.Name] > maxLeft {
			maxLeft = DiskIOPSLeft[node.Name]
		}
		if DiskIOPSLeft[node.Name] < minLeft {
			minLeft = DiskIOPSLeft[node.Name]
		}
	}

	// calculate final priority score for each node
	result := make(schedulerapi.HostPriorityList, 0, len(nodes))
	for _, node := range nodes {
		fScore := float64(0)
		if (maxLeft - minLeft) > 0 {
			fScore = float64(schedulerapi.MaxPriority) * float64((DiskIOPSLeft[node.Name]-minLeft)/(maxLeft-minLeft))
		}
		result = append(result, schedulerapi.HostPriority{Host: node.Name, Score: int(fScore)})
		if glog.V(10) {
			glog.Infof("%v -> %v: DiskIOPSPriority, Score: (%d)", pod.Name, node.Name, int(fScore))
		}
	}
	return result, nil
}
