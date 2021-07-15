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

// propInstance インスタンス内包型
type propInstance struct {
	Instance *Instance `json:",omitempty"` // インスタンス
}

// GetInstance インスタンス 取得
func (p *propInstance) GetInstance() *Instance {
	return p.Instance
}

// IsUp インスタンスが起動しているか判定
func (p *propInstance) IsUp() bool {
	if p.Instance == nil {
		return false
	}
	return p.Instance.IsUp()
}

// IsDown インスタンスがダウンしているか確認
func (p *propInstance) IsDown() bool {
	if p.Instance == nil {
		return false
	}
	return p.Instance.IsDown()
}

// GetInstanceStatus ステータス 取得
func (p *propInstance) GetInstanceStatus() string {
	if p.Instance == nil {
		return ""
	}
	return p.Instance.GetStatus()
}

// GetInstanceBeforeStatus 以前のステータス 取得
func (p *propInstance) GetInstanceBeforeStatus() string {
	if p.Instance == nil {
		return ""
	}
	return p.Instance.GetBeforeStatus()
}

// MaintenanceScheduled メンテナンス予定の有無
func (p *propInstance) MaintenanceScheduled() bool {
	if p.Instance == nil {
		return false
	}
	return p.Instance.MaintenanceScheduled()
}

// GetMaintenanceInfoURL メンテナンス情報 URL取得
func (p *propInstance) GetMaintenanceInfoURL() string {
	if p.Instance == nil {
		return ""
	}
	return p.Instance.Host.InfoURL
}
