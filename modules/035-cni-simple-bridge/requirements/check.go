/*
Copyright 2024 Flant JSC

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

package requirements

import (
	"errors"

	"github.com/deckhouse/deckhouse/go_lib/dependency/requirements"
)

const (
	cniConfigurationSettledKey             = "cniConfigurationSettled"
	cniConfigurationSettledRequirementsKey = "cniConfigurationSettled"
)

func init() {
	checkCNIConfigurationSettledFunc := func(_ string, getter requirements.ValueGetter) (bool, error) {
		cniConfigurationSettledStatusRaw, exists := getter.Get(cniConfigurationSettledKey)
		if !exists {
			return true, nil
		}

		if cniConfigurationSettledStatus, ok := cniConfigurationSettledStatusRaw.(string); ok {
			if cniConfigurationSettledStatus == "false" {
				return false, errors.New(
					"A problem has been found in the CNI configuration, see ClusterAlerts for details",
				)
			}
		}
		return true, nil
	}
	requirements.RegisterCheck(cniConfigurationSettledRequirementsKey, checkCNIConfigurationSettledFunc)
}
