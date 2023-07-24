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

package types

// EProxyLBRegion エンハンスドロードバランサ 設置先リージョン
type EProxyLBRegion string

// String EProxyLBRegionの文字列表現
func (r EProxyLBRegion) String() string {
	return string(r)
}

// ProxyLBRegions エンハンスドロードバランサ 設置先リージョン
var ProxyLBRegions = struct {
	// TK1 東京
	TK1 EProxyLBRegion
	// IS1 石狩
	IS1 EProxyLBRegion
	// Anycast エニーキャスト
	Anycast EProxyLBRegion
}{
	TK1:     EProxyLBRegion("tk1"),
	IS1:     EProxyLBRegion("is1"),
	Anycast: EProxyLBRegion("anycast"),
}

// ProxyLBRegionStrings 設置先リージョンを表す文字列
var ProxyLBRegionStrings = []string{"tk1", "is1", "anycast"}
