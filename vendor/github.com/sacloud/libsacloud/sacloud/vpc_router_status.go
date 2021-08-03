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

// VPCRouterStatus VPCルータのステータス情報
type VPCRouterStatus struct {
	FirewallReceiveLogs []string
	FirewallSendLogs    []string
	VPNLogs             []string
	SessionCount        int
	DHCPServerLeases    []struct {
		IPAddress  string
		MACAddress string
	}
	L2TPIPsecServerSessions []struct {
		User      string
		IPAddress string
		TimeSec   int
	}
	PPTPServerSessions []struct {
		User      string
		IPAddress string
		TimeSec   int
	}
	SiteToSiteIPsecVPNPeers []struct {
		Status string
		Peer   string
	}
}
