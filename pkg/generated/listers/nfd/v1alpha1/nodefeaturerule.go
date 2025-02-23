/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "openshift/node-feature-discovery/pkg/apis/nfd/v1alpha1"
)

// NodeFeatureRuleLister helps list NodeFeatureRules.
// All objects returned here must be treated as read-only.
type NodeFeatureRuleLister interface {
	// List lists all NodeFeatureRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeFeatureRule, err error)
	// Get retrieves the NodeFeatureRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeFeatureRule, error)
	NodeFeatureRuleListerExpansion
}

// nodeFeatureRuleLister implements the NodeFeatureRuleLister interface.
type nodeFeatureRuleLister struct {
	indexer cache.Indexer
}

// NewNodeFeatureRuleLister returns a new NodeFeatureRuleLister.
func NewNodeFeatureRuleLister(indexer cache.Indexer) NodeFeatureRuleLister {
	return &nodeFeatureRuleLister{indexer: indexer}
}

// List lists all NodeFeatureRules in the indexer.
func (s *nodeFeatureRuleLister) List(selector labels.Selector) (ret []*v1alpha1.NodeFeatureRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeFeatureRule))
	})
	return ret, err
}

// Get retrieves the NodeFeatureRule from the index for a given name.
func (s *nodeFeatureRuleLister) Get(name string) (*v1alpha1.NodeFeatureRule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nodefeaturerule"), name)
	}
	return obj.(*v1alpha1.NodeFeatureRule), nil
}
