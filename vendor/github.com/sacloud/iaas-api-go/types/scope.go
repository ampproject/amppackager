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

// EScope スコープ
type EScope string

// String スコープの文字列表現
func (s EScope) String() string {
	return string(s)
}

// Scopes スコープ
var Scopes = &struct {
	Shared EScope // 共有
	User   EScope // ユーザー
}{
	Shared: EScope("shared"),
	User:   EScope("user"),
}

// ScopeStrings Scopeに指定できる有効な文字列
var ScopeStrings = []string{
	Scopes.Shared.String(),
	Scopes.User.String(),
}
