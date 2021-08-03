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

// propConnectedSwitches 接続先スイッチ内包型
type propConnectedSwitches struct {
	ConnectedSwitches []interface{} `json:",omitempty" libsacloud:"requestOnly"` // サーバー作成時の接続先スイッチ指定用パラメーター
}

// GetConnectedSwitches 接続先スイッチ 取得
func (p *propConnectedSwitches) GetConnectedSwitches() []interface{} {
	return p.ConnectedSwitches
}

// SetConnectedSwitches 接続先スイッチ 設定
func (p *propConnectedSwitches) SetConnectedSwitches(switches []interface{}) {
	p.ConnectedSwitches = switches
}

// ClearConnectedSwitches 接続先スイッチ指定パラメータークリア
func (p *propConnectedSwitches) ClearConnectedSwitches() {
	p.ConnectedSwitches = []interface{}{}
}

// AddPublicNWConnectedParam 共有セグメントへ接続したNIC追加
func (p *propConnectedSwitches) AddPublicNWConnectedParam() {
	if p.ConnectedSwitches == nil {
		p.ClearConnectedSwitches()
	}
	p.ConnectedSwitches = append(p.ConnectedSwitches, map[string]interface{}{"Scope": "shared"})
}

// AddExistsSwitchConnectedParam スイッチへ接続したNIC追加
func (p *propConnectedSwitches) AddExistsSwitchConnectedParam(switchID string) {
	if p.ConnectedSwitches == nil {
		p.ClearConnectedSwitches()
	}
	p.ConnectedSwitches = append(p.ConnectedSwitches, map[string]interface{}{"ID": switchID})
}

// AddEmptyConnectedParam 未接続なNIC追加
func (p *propConnectedSwitches) AddEmptyConnectedParam() {
	if p.ConnectedSwitches == nil {
		p.ClearConnectedSwitches()
	}
	p.ConnectedSwitches = append(p.ConnectedSwitches, nil)
}
