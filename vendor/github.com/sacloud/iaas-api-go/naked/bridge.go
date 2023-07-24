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

// Bridge ブリッジ
type Bridge struct {
	ID           types.ID          `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string            `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string            `yaml:"description"`
	ServiceClass string            `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time        `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Region       *Region           `json:",omitempty" yaml:"region,omitempty" structs:",omitempty"`
	Info         *BridgeInfo       `json:",omitempty" yaml:"info,omitempty" structs:",omitempty"`
	SwitchInZone *BridgeSwitchInfo `json:",omitempty" yaml:"switch_in_zone,omitempty" structs:",omitempty"`
}

// BridgeInfo ブリッジに接続されているスイッチの情報
type BridgeInfo struct {
	Switches []*Switch `json:",omitempty" yaml:"switches,omitempty" structs:",omitempty"`
}

// BridgeSwitchInfo ゾーン内での接続スイッチ情報
type BridgeSwitchInfo struct {
	ID             types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Scope          types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	Name           string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	ServerCount    int          `json:",omitempty" yaml:"server_count,omitempty" structs:",omitempty"`
	ApplianceCount int          `json:",omitempty" yaml:"appliance_count,omitempty" structs:",omitempty"`
}
