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

import (
	"github.com/sacloud/iaas-api-go/types"
)

// DiskPlan ディスクプラン
type DiskPlan struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	StorageClass string              `json:",omitempty" yaml:"storage_class,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Size         []*DiskPlanSizeInfo `json:",omitempty" yaml:"size,omitempty" structs:",omitempty"`
}

// DiskPlanSizeInfo ディスクプランに含まれる利用可能なサイズ情報
type DiskPlanSizeInfo struct {
	Availability  types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	DisplaySize   int                 `json:",omitempty" yaml:"display_size,omitempty" structs:",omitempty"`
	DisplaySuffix string              `json:",omitempty" yaml:"display_suffix,omitempty" structs:",omitempty"`
	ServiceClass  string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	SizeMB        int                 `json:",omitempty" yaml:"size_mb,omitempty" structs:",omitempty"`
}
