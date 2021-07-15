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

// Interface インターフェース(NIC)
type Interface struct {
	*Resource                   // ID
	propServer                  // サーバー
	propSwitch                  // スイッチ
	MACAddress    string        `json:",omitempty"` // MACアドレス
	IPAddress     string        `json:",omitempty"` // IPアドレス
	UserIPAddress string        `json:",omitempty"` // ユーザー指定IPアドレス
	HostName      string        `json:",omitempty"` // ホスト名
	PacketFilter  *PacketFilter `json:",omitempty"` // 適用パケットフィルタ
}

// GetMACAddress MACアドレス 取得
func (i *Interface) GetMACAddress() string {
	return i.MACAddress
}

//GetIPAddress IPアドレス 取得
func (i *Interface) GetIPAddress() string {
	return i.IPAddress
}

// SetUserIPAddress ユーザー指定IPアドレス 設定
func (i *Interface) SetUserIPAddress(ip string) {
	i.UserIPAddress = ip
}

//GetUserIPAddress ユーザー指定IPアドレス 取得
func (i *Interface) GetUserIPAddress() string {
	return i.UserIPAddress
}

// GetHostName ホスト名 取得
func (i *Interface) GetHostName() string {
	return i.HostName
}

// GetPacketFilter 適用パケットフィルタ 取得
func (i *Interface) GetPacketFilter() *PacketFilter {
	return i.PacketFilter
}

// UpstreamType 上流ネットワーク種別
func (i *Interface) UpstreamType() EUpstreamNetworkType {
	sw := i.Switch
	if sw == nil {
		return EUpstreamNetworkNone
	}

	if sw.Subnet == nil {
		return EUpstreamNetworkSwitch
	}

	if sw.Scope == ESCopeShared {
		return EUpstreamNetworkShared
	}

	return EUpstreamNetworkRouter
}
