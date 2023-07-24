// Copyright 2022-2023 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package naked

import "github.com/sacloud/iaas-api-go/types"

// Storage ストレージ
type Storage struct {
	ID          types.ID  `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string    `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string    `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Generation  int       `json:",omitempty" yaml:"generation,omitempty" structs:",omitempty"`
	Class       string    `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	DiskPlan    *DiskPlan `json:",omitempty" yaml:"disk_plan,omitempty" structs:",omitempty"`
	Zone        *Zone     `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
}
