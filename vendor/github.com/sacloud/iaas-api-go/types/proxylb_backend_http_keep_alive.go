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

// EProxyLBBackendHttpKeepAlive エンハンスドロードバランサ 実サーバとのHTTP持続接続
type EProxyLBBackendHttpKeepAlive string

// String EProxyLBBackendHttpKeepAlive
func (p EProxyLBBackendHttpKeepAlive) String() string {
	return string(p)
}

// ProxyLBBackendHttpKeepAlive エンハンスドロードバランサ 実サーバとのHTTP持続接続
var ProxyLBBackendHttpKeepAlive = struct {
	// SAFE クライアントからの接続に応じて実サーバとの持続的接続が使われ、クライアント間での接続の共有は行われにくい設定
	Safe EProxyLBBackendHttpKeepAlive
	// AGGRESSIVE 再使用可能な接続があればより積極的に使用する設定
	Aggressive EProxyLBBackendHttpKeepAlive
}{
	Safe:       EProxyLBBackendHttpKeepAlive("safe"),
	Aggressive: EProxyLBBackendHttpKeepAlive("aggressive"),
}

// ProxyLBBackendHttpKeepAliveStrings 実サーバとのHTTP持続接続を表す文字列
var ProxyLBBackendHttpKeepAliveStrings = []string{"safe", "aggressive"}
