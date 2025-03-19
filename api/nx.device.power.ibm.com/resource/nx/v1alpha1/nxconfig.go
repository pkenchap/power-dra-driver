/*
Copyright 2025.

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

package v1alpha1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NxConfig holds the set of parameters for configuring a Nest Accelerator.
type NxConfig struct {
	metav1.TypeMeta `json:",inline"`
	Sharing         *NxSharing `json:"sharing,omitempty"`
}

// DefaultNxConfig provides the default NX configuration.
func DefaultNxConfig() *NxConfig {
	return &NxConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupName + "/" + Version,
			Kind:       NxConfigKind,
		},
		Sharing: &NxSharing{
			Strategy: TimeSlicingStrategy,
			TimeSlicingConfig: &TimeSlicingConfig{
				Interval: "Default",
			},
		},
	}
}

// Normalize updates a NxConfig config with implied default values based on other settings.
func (c *NxConfig) Normalize() error {
	if c == nil {
		return fmt.Errorf("config is 'nil'")
	}
	if c.Sharing == nil {
		c.Sharing = &NxSharing{
			Strategy: TimeSlicingStrategy,
		}
	}
	if c.Sharing.Strategy == TimeSlicingStrategy && c.Sharing.TimeSlicingConfig == nil {
		c.Sharing.TimeSlicingConfig = &TimeSlicingConfig{
			Interval: "Default",
		}
	}
	if c.Sharing.Strategy == SpacePartitioningStrategy && c.Sharing.SpacePartitioningConfig == nil {
		c.Sharing.SpacePartitioningConfig = &SpacePartitioningConfig{
			PartitionCount: 1,
		}
	}
	return nil
}
