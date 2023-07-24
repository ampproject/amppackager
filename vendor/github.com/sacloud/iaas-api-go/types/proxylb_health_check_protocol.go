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

// EProxyLBHealthCheckProtocol エンハンスドロードバランサ 監視プロトコル
type EProxyLBHealthCheckProtocol string

// String EProxyLBHealthCheckProtocolの文字列表現
func (p EProxyLBHealthCheckProtocol) String() string {
	return string(p)
}

// ProxyLBProtocols エンハンスドロードバランサ 監視プロトコル
var ProxyLBProtocols = struct {
	// HTTP http
	HTTP EProxyLBHealthCheckProtocol
	// TCP tcp
	TCP EProxyLBHealthCheckProtocol
}{
	HTTP: EProxyLBHealthCheckProtocol("http"),
	TCP:  EProxyLBHealthCheckProtocol("tcp"),
}

// ProxyLBProtocolStrings 監視プロトコルを表す文字列
var ProxyLBProtocolStrings = []string{"http", "tcp"}
