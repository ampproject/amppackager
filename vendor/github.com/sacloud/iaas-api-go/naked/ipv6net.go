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

// IPv6Net InternetリソースでのIPv6アドレス帯を表す
type IPv6Net struct {
	ID                 types.ID   `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	ServiceID          types.ID   `json:",omitempty" yaml:"service_id,omitempty" structs:",omitempty"`
	IPv6Prefix         string     `json:",omitempty" yaml:"ipv6prefix,omitempty" structs:",omitempty"`
	IPv6PrefixLen      int        `json:",omitempty" yaml:"ipv6prefix_len,omitempty" structs:",omitempty"`
	IPv6PrefixTail     string     `json:",omitempty" yaml:"ipv6prefix_tail,omitempty" structs:",omitempty"`
	ServiceClass       string     `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	IPv6Table          *IPv6Table `json:",omitempty" yaml:"ipv6table,omitempty" structs:",omitempty"`
	NamedIPv6AddrCount int        `json:",omitempty" yaml:"named_ipv6addr_count,omitempty" structs:",omitempty"`
	CreatedAt          *time.Time `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Switch             *Switch    `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
}

// IPv6Table IPv6テーブル
type IPv6Table struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}
