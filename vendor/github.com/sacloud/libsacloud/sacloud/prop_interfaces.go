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

// propInterfaces インターフェース(NIC)配列内包型
type propInterfaces struct {
	Interfaces []Interface `json:",omitempty"` // インターフェース
}

// GetInterfaces インターフェース(NIC)配列 取得
func (p *propInterfaces) GetInterfaces() []Interface {
	return p.Interfaces
}

// GetFirstInterface インターフェース(NIC)配列の先頭要素を返す
func (p *propInterfaces) GetFirstInterface() *Interface {
	if len(p.Interfaces) == 0 {
		return nil
	}
	return &p.Interfaces[0]
}
