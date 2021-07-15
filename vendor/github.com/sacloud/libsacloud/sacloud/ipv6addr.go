// Copyright 2016-2020 The Libsacloud Authors
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

package sacloud

// IPv6Addr IPアドレス(IPv6)
type IPv6Addr struct {
	HostName  string     `json:",omitempty"` // ホスト名
	IPv6Addr  string     `json:",omitempty"` // IPv6アドレス
	Interface *Interface `json:",omitempty"` // インターフェース
	IPv6Net   *IPv6Net   `json:",omitempty"` // IPv6サブネット

}

// GetIPv6NetID IPv6アドレスが所属するIPv6NetのIDを取得
func (a *IPv6Addr) GetIPv6NetID() ID {
	if a.IPv6Net != nil {
		return a.IPv6Net.ID
	}
	return 0
}

// GetInternetID IPv6アドレスを所有するルータ+スイッチ(Internet)のIDを取得
func (a *IPv6Addr) GetInternetID() ID {
	if a.IPv6Net != nil && a.IPv6Net.Switch != nil && a.IPv6Net.Switch.Internet != nil {
		return a.IPv6Net.Switch.Internet.ID
	}
	return 0
}

// CreateNewIPv6Addr IPv6アドレス作成
func CreateNewIPv6Addr() *IPv6Addr {
	return &IPv6Addr{
		IPv6Net: &IPv6Net{
			Resource: &Resource{},
		},
	}
}
