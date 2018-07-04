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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
)

// ExtendedResourceLister helps list ExtendedResources.
type ExtendedResourceLister interface {
	// List lists all ExtendedResources in the indexer.
	List(selector labels.Selector) (ret []*extensions.ExtendedResource, err error)
	// Get retrieves the ExtendedResource from the index for a given name.
	Get(name string) (*extensions.ExtendedResource, error)
	ExtendedResourceListerExpansion
}

// extendedResourceLister implements the ExtendedResourceLister interface.
type extendedResourceLister struct {
	indexer cache.Indexer
}

// NewExtendedResourceLister returns a new ExtendedResourceLister.
func NewExtendedResourceLister(indexer cache.Indexer) ExtendedResourceLister {
	return &extendedResourceLister{indexer: indexer}
}

// List lists all ExtendedResources in the indexer.
func (s *extendedResourceLister) List(selector labels.Selector) (ret []*extensions.ExtendedResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*extensions.ExtendedResource))
	})
	return ret, err
}

// Get retrieves the ExtendedResource from the index for a given name.
func (s *extendedResourceLister) Get(name string) (*extensions.ExtendedResource, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(extensions.Resource("extendedresource"), name)
	}
	return obj.(*extensions.ExtendedResource), nil
}
