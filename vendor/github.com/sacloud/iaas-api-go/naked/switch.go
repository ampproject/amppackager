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

// Switch スイッチ
type Switch struct {
	ID               types.ID          `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name             string            `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description      string            `yaml:"description"`
	Tags             types.Tags        `yaml:"tags"`
	Icon             *Icon             `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt        *time.Time        `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt       *time.Time        `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Scope            types.EScope      `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	Subnet           *Subnet           `json:",omitempty" yaml:"subnet,omitempty" structs:",omitempty"`
	UserSubnet       *UserSubnet       `json:",omitempty" yaml:"user_subnet,omitempty" structs:",omitempty"`
	Zone             *Zone             `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	Internet         *Internet         `json:",omitempty" yaml:"internet,omitempty" structs:",omitempty"`
	Subnets          []*Subnet         `json:",omitempty" yaml:"subnets,omitempty" structs:",omitempty"`
	IPv6Nets         []*IPv6Net        `json:",omitempty" yaml:"ipv6nets,omitempty" structs:",omitempty"`
	Bridge           *Bridge           `json:",omitempty" yaml:"bridge,omitempty" structs:",omitempty"`
	ServerCount      int               `json:",omitempty" yaml:"server_count,omitempty" structs:",omitempty"`
	HybridConnection *HybridConnection `json:",omitempty" yaml:"hybrid_connection,omitempty" structs:",omitempty"`
}
