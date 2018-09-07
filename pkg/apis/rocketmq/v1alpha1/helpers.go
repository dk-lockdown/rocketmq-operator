/*
Copyright The Kubernetes Authors.

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

const (
	// The default rocketmq version to use if not specified explicitly by user
	defaultVersion = "4.3.0-operator"
	defaultGroups  = 2
	defaultMembers = 2
)

// EnsureDefaults will ensure that if a user omits and fields in the
// spec that are required, we set some sensible defaults.
// For example a user can choose to omit the version
// and number of members.
func (c *BrokerCluster) EnsureDefaults() *BrokerCluster {
	if c.Spec.Version == "" {
		c.Spec.Version = defaultVersion
	}
	if c.Spec.GroupReplica == 0 {
		c.Spec.GroupReplica = defaultGroups
	}

	if c.Spec.MembersPerGroup == 0 {
		c.Spec.MembersPerGroup = defaultMembers
	}
	return c
}