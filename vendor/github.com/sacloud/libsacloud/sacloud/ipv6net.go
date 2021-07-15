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

// IPv6Net IPv6ネットワーク(サブネット)
type IPv6Net struct {
	*Resource        // ID
	propScope        // スコープ
	propServiceClass // サービスクラス
	propCreatedAt    // 作成日時

	IPv6Prefix         string    `json:",omitempty"` // IPv6プレフィックス
	IPv6PrefixLen      int       `json:",omitempty"` // IPv6プレフィックス長
	IPv6PrefixTail     string    `json:",omitempty"` // IPv6プレフィックス末尾
	IPv6Table          *Resource `json:",omitempty"` // IPv6テーブル
	NamedIPv6AddrCount int       `json:",omitempty"` // 名前付きIPv6アドレス数
	ServiceID          ID        `json:",omitempty"` // サービスID
	Switch             *Switch   `json:",omitempty"` // 接続先スイッチ

}
