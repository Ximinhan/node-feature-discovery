/*
Copyright 2020-2021 The Kubernetes Authors.

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

package rules

import (
	"fmt"

	nfdv1alpha1 "openshift/node-feature-discovery/pkg/apis/nfd/v1alpha1"
	"openshift/node-feature-discovery/source"
	"openshift/node-feature-discovery/source/kernel"
)

// LoadedKModRule matches loaded kernel modules in the system
type LoadedKModRule struct {
	nfdv1alpha1.MatchExpressionSet
}

// Match loaded kernel modules on provided list of kernel modules
func (r *LoadedKModRule) Match() (bool, error) {
	modules, ok := source.GetFeatureSource("kernel").GetFeatures().Keys[kernel.LoadedModuleFeature]
	if !ok {
		return false, fmt.Errorf("info about loaded modules not available")
	}

	return r.MatchKeys(modules.Elements)
}
