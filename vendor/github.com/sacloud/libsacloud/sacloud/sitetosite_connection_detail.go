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

// SiteToSiteConnectionDetail サイト間VPN接続詳細情報
type SiteToSiteConnectionDetail struct {
	ESP struct {
		AuthenticationProtocol string
		DHGroup                string
		EncryptionProtocol     string
		Lifetime               string
		Mode                   string
		PerfectForwardSecrecy  string
	}
	IKE struct {
		AuthenticationProtocol string
		EncryptionProtocol     string
		Lifetime               string
		Mode                   string
		PerfectForwardSecrecy  string
		PreSharedSecret        string
	}
	Peer struct {
		ID               string
		InsideNetworks   []string
		OutsideIPAddress string
	}
	VPCRouter struct {
		ID               string
		InsideNetworks   []string
		OutsideIPAddress string
	}
}

// SiteToSiteConnectionInfo サイト間VPN接続情報
type SiteToSiteConnectionInfo struct {
	Details struct {
		Config []SiteToSiteConnectionDetail
	}
}
