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
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// Internet ルータ+スイッチのルータ部分
type Internet struct {
	ID             types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name           string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description    string       `yaml:"description"`
	Tags           types.Tags   `yaml:"tags"`
	Icon           *Icon        `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt      *time.Time   `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Scope          types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	ServiceClass   string       `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Switch         *Switch      `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	BandWidthMbps  int          `json:",omitempty" yaml:"band_width_mbps,omitempty" structs:",omitempty"`
	NetworkMaskLen int          `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
}

// SubnetOperationRequest サブネット追加時のリクエストパラメータ
type SubnetOperationRequest struct {
	NetworkMaskLen int    `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
	NextHop        string `json:",omitempty" yaml:"next_hop,omitempty" structs:",omitempty"`
}
