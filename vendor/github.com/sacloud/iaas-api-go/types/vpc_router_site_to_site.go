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

const (
	VPCRouterSiteToSiteVPNEncryptionAlgoAES128 = "aes128"
	VPCRouterSiteToSiteVPNEncryptionAlgoAES256 = "aes256"
)

const (
	VPCRouterSiteToSiteVPNHashAlgoSHA1   = "sha1"
	VPCRouterSiteToSiteVPNHashAlgoSHA256 = "sha256"
)

const (
	VPCRouterSiteToSiteVPNDHGroupModp1024 = "modp1024"
	VPCRouterSiteToSiteVPNDHGroupModp2048 = "modp2048"
	VPCRouterSiteToSiteVPNDHGroupModp3072 = "modp3072"
	VPCRouterSiteToSiteVPNDHGroupModp4096 = "modp4096"
)

var (
	VPCRouterSiteToSiteVPNEncryptionAlgos = []string{
		VPCRouterSiteToSiteVPNEncryptionAlgoAES128,
		VPCRouterSiteToSiteVPNEncryptionAlgoAES256,
	}
	VPCRouterSiteToSiteVPNHashAlgos = []string{
		VPCRouterSiteToSiteVPNHashAlgoSHA1,
		VPCRouterSiteToSiteVPNHashAlgoSHA256,
	}
	VPCRouterSiteToSiteVPNDHGroups = []string{
		VPCRouterSiteToSiteVPNDHGroupModp1024,
		VPCRouterSiteToSiteVPNDHGroupModp2048,
		VPCRouterSiteToSiteVPNDHGroupModp3072,
		VPCRouterSiteToSiteVPNDHGroupModp4096,
	}
)
