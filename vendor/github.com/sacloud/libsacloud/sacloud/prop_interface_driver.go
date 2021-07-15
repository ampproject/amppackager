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

// propInterfaceDriver インターフェースドライバ内包型
type propInterfaceDriver struct {
	InterfaceDriver EInterfaceDriver `json:",omitempty"` // NIC
}

// SetInterfaceDriver インターフェースドライバ 設定
func (p *propInterfaceDriver) SetInterfaceDriver(v EInterfaceDriver) {
	p.InterfaceDriver = v
}

// GetInterfaceDriver インターフェースドライバ 取得
func (p *propInterfaceDriver) GetInterfaceDriver() EInterfaceDriver {
	return p.InterfaceDriver
}

// SetInterfaceDriverByString インターフェースドライバ 設定(文字列)
func (p *propInterfaceDriver) SetInterfaceDriverByString(v string) {
	p.InterfaceDriver = EInterfaceDriver(v)
}

// GetInterfaceDriverString インターフェースドライバ 取得(文字列)
func (p *propInterfaceDriver) GetInterfaceDriverString() string {
	return string(p.InterfaceDriver)
}
