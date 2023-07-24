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

// EGSLBHealthCheckProtocol GSLB 監視プロトコル
type EGSLBHealthCheckProtocol string

// String EGSLBHealthCheckProtocolの文字列表現
func (p EGSLBHealthCheckProtocol) String() string {
	return string(p)
}

// GSLBHealthCheckProtocols GSLB 監視プロトコル
var GSLBHealthCheckProtocols = struct {
	// HTTP http
	HTTP EGSLBHealthCheckProtocol
	// HTTPS https
	HTTPS EGSLBHealthCheckProtocol
	// TCP tcp
	TCP EGSLBHealthCheckProtocol
	// Ping ping
	Ping EGSLBHealthCheckProtocol
}{
	HTTP:  EGSLBHealthCheckProtocol("http"),
	HTTPS: EGSLBHealthCheckProtocol("https"),
	TCP:   EGSLBHealthCheckProtocol("tcp"),
	Ping:  EGSLBHealthCheckProtocol("ping"),
}

// GSLBHealthCheckProtocolStrings 有効なGSLB監視プロトコルを示す文字列のリスト
var GSLBHealthCheckProtocolStrings = []string{
	GSLBHealthCheckProtocols.HTTP.String(),
	GSLBHealthCheckProtocols.HTTPS.String(),
	GSLBHealthCheckProtocols.TCP.String(),
	GSLBHealthCheckProtocols.Ping.String(),
}
