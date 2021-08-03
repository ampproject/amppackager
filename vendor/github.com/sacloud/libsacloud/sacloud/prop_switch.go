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

// propSwitch スイッチ内包型
type propSwitch struct {
	Switch *Switch `json:",omitempty"` // スイッチ
}

// GetSwitch スイッチ 取得
func (p *propSwitch) GetSwitch() *Switch {
	return p.Switch
}

// SetSwitch スイッチ 設定
func (p *propSwitch) SetSwitch(sw *Switch) {
	p.Switch = sw
}

// SetSwitchID スイッチID 設定
func (p *Interface) SetSwitchID(id ID) {
	p.Switch = &Switch{Resource: &Resource{ID: id}}
}
