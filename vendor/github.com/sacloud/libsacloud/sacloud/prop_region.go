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

// propRegion リージョン内包型
type propRegion struct {
	Region *Region `json:",omitempty"` // リージョン
}

// GetRegion リージョン 取得
func (p *propRegion) GetRegion() *Region {
	return p.Region
}

// GetRegionID リージョンID 取得
func (p *propRegion) GetRegionID() ID {
	if p.Region == nil {
		return -1
	}
	return p.Region.GetID()
}

// GetRegionName リージョン名 取得
func (p *propRegion) GetRegionName() string {
	if p.Region == nil {
		return ""
	}
	return p.Region.GetName()
}

// GetRegionDescription リージョン説明 取得
func (p *propRegion) GetRegionDescription() string {
	if p.Region == nil {
		return ""
	}
	return p.Region.GetDescription()
}

// GetRegionNameServers リージョンのネームサーバー(のIPアドレス)取得
func (p *propRegion) GetRegionNameServers() []string {
	if p.Region == nil {
		return []string{}
	}

	return p.Region.GetNameServers()
}
