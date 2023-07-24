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

// Protocol パケットフィルタやロードバランサ、VPCルータなどで利用するプロトコル
type Protocol string

// Protocols パケットフィルタやロードバランサ、VPCルータなどで利用するプロトコル
var Protocols = &struct {
	HTTP     Protocol
	HTTPS    Protocol
	TCP      Protocol
	UDP      Protocol
	ICMP     Protocol
	Fragment Protocol
	IP       Protocol
}{
	HTTP:     "http",
	HTTPS:    "https",
	TCP:      "tcp",
	UDP:      "udp",
	ICMP:     "icmp",
	Fragment: "fragment",
	IP:       "ip",
}

// String Protocolの文字列表現
func (p Protocol) String() string {
	return string(p)
}

// PacketFilterProtocolStrings 有効なパケットフィルタプロトコルを示す文字列のリスト
var PacketFilterProtocolStrings = []string{
	Protocols.HTTP.String(),
	Protocols.HTTPS.String(),
	Protocols.TCP.String(),
	Protocols.UDP.String(),
	Protocols.ICMP.String(),
	Protocols.Fragment.String(),
	Protocols.IP.String(),
}
