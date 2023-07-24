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

	"github.com/sacloud/iaas-api-go/types"
)

// Interface サーバなどに接続されているNICの情報
type Interface struct {
	ID            types.ID          `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	MACAddress    string            `json:",omitempty" yaml:"mac_address,omitempty" structs:",omitempty"`
	IPAddress     string            `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	UserIPAddress string            `json:",omitempty" yaml:"user_ip_address,omitempty" structs:",omitempty"`
	HostName      string            `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	Switch        *Switch           `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	PacketFilter  *PacketFilterInfo `json:",omitempty" yaml:"packet_filter,omitempty" structs:",omitempty"`
	Server        *Server           `json:",omitempty" yaml:"server,omitempty" structs:",omitempty"`

	// Index 仮想フィールド、VPCルータなどでInterfaces(実体は[]*Interface)を扱う場合にUnmarshalJSONの中で設定される
	//
	// Findした際のAPIからの応答にも同名のフィールドが含まれるが無関係。
	Index int

	// UpstreamType 上流ネットワーク種別 UnmarshalJSONの中で算出される
	UpstreamType types.EUpstreamNetworkType
}

// Interfaces Interface配列
//
// 配列中にnullが返ってくる(VPCルータなど)への対応のためのtype
type Interfaces []*Interface

// UnmarshalJSON 配列中にnullが返ってくる(VPCルータなど)への対応
func (i *Interfaces) UnmarshalJSON(b []byte) error {
	type alias Interfaces
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	var dest []*Interface
	for i, v := range a {
		if v != nil {
			if v.Index == 0 {
				v.Index = i
			}
			dest = append(dest, v)
		}
	}

	*i = Interfaces(dest)
	return nil
}

// MarshalJSON 配列中にnullが入る場合(VPCルータなど)への対応
func (i Interfaces) MarshalJSON() ([]byte, error) {
	max := 0
	for _, iface := range i {
		if max < iface.Index {
			max = iface.Index
		}
	}

	var dest = make([]*Interface, max+1)
	for _, iface := range i {
		dest[iface.Index] = iface
	}

	return json.Marshal(dest)
}

// MarshalJSON Indexフィールドを出力しないための実装
func (i *Interface) MarshalJSON() ([]byte, error) {
	type alias struct {
		ID            types.ID          `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
		MACAddress    string            `json:",omitempty" yaml:"mac_address,omitempty" structs:",omitempty"`
		IPAddress     string            `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
		UserIPAddress string            `json:",omitempty" yaml:"user_ip_address,omitempty" structs:",omitempty"`
		HostName      string            `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
		Switch        *Switch           `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
		PacketFilter  *PacketFilterInfo `json:",omitempty" yaml:"packet_filter,omitempty" structs:",omitempty"`
		Server        *Server           `json:",omitempty" yaml:"server,omitempty" structs:",omitempty"`
	}

	tmp := alias{
		ID:            i.ID,
		MACAddress:    i.MACAddress,
		IPAddress:     i.IPAddress,
		UserIPAddress: i.UserIPAddress,
		HostName:      i.HostName,
		Switch:        i.Switch,
		PacketFilter:  i.PacketFilter,
		Server:        i.Server,
	}
	return json.Marshal(tmp)
}

// UnmarshalJSON 仮想フィールド UpstreamType を表現するための実装
func (i *Interface) UnmarshalJSON(b []byte) error {
	type alias Interface
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	// calculate UpstreamType
	var upstreamType types.EUpstreamNetworkType
	sw := tmp.Switch
	switch {
	case sw == nil:
		upstreamType = types.UpstreamNetworkTypes.None
	case sw.Subnet == nil:
		upstreamType = types.UpstreamNetworkTypes.Switch
	case sw.Scope == types.Scopes.Shared:
		upstreamType = types.UpstreamNetworkTypes.Shared
	default:
		upstreamType = types.UpstreamNetworkTypes.Router
	}
	tmp.UpstreamType = upstreamType

	*i = Interface(tmp)
	return nil
}
