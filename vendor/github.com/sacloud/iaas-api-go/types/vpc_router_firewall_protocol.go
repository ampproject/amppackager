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

// VPCRouterFirewallProtocolStrings VPCルータでのファイアウォールプロトコルを表す文字列
var VPCRouterFirewallProtocolStrings = []string{"tcp", "udp", "icmp", "ip"}

// EVPCRouterFirewallProtocol VPCルータでのファイアウォールプロトコルを表す文字列
type EVPCRouterFirewallProtocol string

// VPCRouterFirewallProtocols ファイアアウォール プロトコル
var VPCRouterFirewallProtocols = struct {
	// TCP tcp
	TCP EVPCRouterFirewallProtocol
	// UDP udp
	UDP EVPCRouterFirewallProtocol
	// ICMP udp
	ICMP EVPCRouterFirewallProtocol
	// ICMP udp
	IP EVPCRouterFirewallProtocol
}{
	TCP:  EVPCRouterFirewallProtocol("tcp"),
	UDP:  EVPCRouterFirewallProtocol("udp"),
	ICMP: EVPCRouterFirewallProtocol("icmp"),
	IP:   EVPCRouterFirewallProtocol("ip"),
}
