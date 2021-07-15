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

// propScope スコープ内包型
type propScope struct {
	Scope string `json:",omitempty"` // スコープ
}

// GetScope スコープ 取得
func (p *propScope) GetScope() string {
	return p.Scope
}

// SetScope スコープ 設定
func (p *propScope) SetScope(scope string) {
	p.Scope = scope
}

// SetSharedScope 共有スコープに設定
func (p *propScope) SetSharedScope() {
	p.Scope = string(ESCopeShared)
}

// SetUserScope ユーザースコープに設定
func (p *propScope) SetUserScope() {
	p.Scope = string(ESCopeUser)
}

// IsSharedScope 共有スコープか判定
func (p *propScope) IsSharedScope() bool {
	if p == nil {
		return false
	}
	return p.Scope == string(ESCopeShared)
}

// IsUserScope ユーザースコープか判定
func (p *propScope) IsUserScope() bool {
	if p == nil {
		return false
	}
	return p.Scope == string(ESCopeUser)
}
