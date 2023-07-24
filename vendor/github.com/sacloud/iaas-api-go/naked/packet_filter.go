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
	"encoding/json"
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// PacketFilter パケットフィルタ
type PacketFilter struct {
	ID                  types.ID                `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name                string                  `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description         string                  `yaml:"description"`
	RequiredHostVersion types.StringNumber      `json:",omitempty" yaml:"require_host_version,omitempty" structs:",omitempty"`
	Expression          PacketFilterExpressions `yaml:"expression"`
	ExpressionHash      string                  `json:",omitempty" yaml:"expression_hash,omitempty" structs:",omitempty"`
	CreatedAt           time.Time               `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	// Notice              interface{}               `json:"Notice"`
}

// PacketFilterExpressions パケットフィルターのルール
type PacketFilterExpressions []*PacketFilterExpression

// MarshalJSON nullの場合に空配列を出力するための実装
func (p *PacketFilterExpressions) MarshalJSON() ([]byte, error) {
	if *p == nil {
		*p = make([]*PacketFilterExpression, 0)
	}
	type alias PacketFilterExpressions
	tmp := alias(*p)
	return json.Marshal(&tmp)
}

// PacketFilterExpression パケットフィルタのルール
type PacketFilterExpression struct {
	Protocol        types.Protocol            `yaml:"protocol"`
	SourceNetwork   types.PacketFilterNetwork `yaml:"source_network"`
	DestinationPort types.PacketFilterPort    `yaml:"destination_port"`
	Action          types.Action              `yaml:"action"`
	SourcePort      types.PacketFilterPort    `yaml:"source_port"`
	Description     string                    `yaml:"description"`
}

// PacketFilterInfo パケットフィルタ - Interface配下などでの参照用
type PacketFilterInfo struct {
	ID                  types.ID           `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name                string             `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	RequiredHostVersion types.StringNumber `json:",omitempty" yaml:"require_host_version,omitempty" structs:",omitempty"`
}
