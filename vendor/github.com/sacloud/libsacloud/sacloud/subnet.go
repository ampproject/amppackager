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

// Subnet IPv4サブネット
type Subnet struct {
	*Resource        // ID
	propServiceClass // サービスクラス
	propCreatedAt    // 作成日時

	DefaultRoute   string       `json:",omitempty"` // デフォルトルート
	IPAddresses    []*IPAddress `json:",omitempty"` // IPv4アドレス範囲
	NetworkAddress string       `json:",omitempty"` // ネットワークアドレス
	NetworkMaskLen int          `json:",omitempty"` // ネットワークマスク長
	ServiceID      ID           `json:",omitempty"` // サービスID
	StaticRoute    string       `json:",omitempty"` // スタティックルート
	NextHop        string       `json:",omitempty"` // ネクストホップ
	Switch         *Switch      `json:",omitempty"` // スイッチ
	Internet       *Internet    `json:",omitempty"` // ルーター
}
