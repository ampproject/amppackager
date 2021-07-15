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

// PropZone ゾーン内包型
type propZone struct {
	Zone *Zone `json:",omitempty"` // ゾーン
}

// GetZone ゾーン 取得
func (p *propZone) GetZone() *Zone {
	return p.Zone
}

// GetZoneID ゾーンID 取得
func (p *propZone) GetZoneID() ID {
	if p.Zone == nil {
		return -1
	}
	return p.Zone.GetID()
}

// GetZoneName ゾーン名 取得
func (p *propZone) GetZoneName() string {
	if p.Zone == nil {
		return ""
	}

	return p.Zone.GetName()
}

// GetZoneDescription ゾーン説明 取得
func (p *propZone) GetZoneDescription() string {
	if p.Zone == nil {
		return ""
	}
	return p.Zone.GetDescription()
}

// GetRegion リージョン 取得
func (p *propZone) GetRegion() *Region {
	if p.Zone == nil {
		return nil
	}
	return p.Zone.GetRegion()
}

// GetRegionID リージョンID 取得
func (p *propZone) GetRegionID() ID {
	if p.Zone == nil {
		return -1
	}
	return p.Zone.GetRegionID()
}

// GetRegionName リージョン名 取得
func (p *propZone) GetRegionName() string {
	if p.Zone == nil {
		return ""
	}
	return p.Zone.GetRegionName()
}

// GetRegionDescription リージョン説明 取得
func (p *propZone) GetRegionDescription() string {
	if p.Zone == nil {
		return ""
	}
	return p.Zone.GetRegionDescription()
}

// GetRegionNameServers リージョンのネームサーバー(のIPアドレス)取得
func (p *propZone) GetRegionNameServers() []string {
	if p.Zone == nil {
		return []string{}
	}

	return p.Zone.GetRegionNameServers()
}

// ZoneIsDummy ダミーフラグ 取得
func (p *propZone) ZoneIsDummy() bool {
	if p.Zone == nil {
		return false
	}
	return p.Zone.ZoneIsDummy()
}

// GetVNCProxyHostName VNCプロキシホスト名 取得
func (p *propZone) GetVNCProxyHostName() string {
	if p.Zone == nil {
		return ""
	}

	return p.Zone.GetVNCProxyHostName()
}

// GetVPCProxyIPAddress VNCプロキシIPアドレス 取得
func (p *propZone) GetVPCProxyIPAddress() string {
	if p.Zone == nil {
		return ""
	}
	return p.Zone.GetVPCProxyIPAddress()
}

// GetFTPHostName FTPサーバーホスト名 取得
func (p *propZone) GetFTPHostName() string {
	if p.Zone == nil {
		return ""
	}
	return p.Zone.GetFTPHostName()
}

// GetFTPServerIPAddress FTPサーバーIPアドレス 取得
func (p *propZone) GetFTPServerIPAddress() string {
	if p.Zone == nil {
		return ""
	}
	return p.Zone.GetFTPServerIPAddress()
}
