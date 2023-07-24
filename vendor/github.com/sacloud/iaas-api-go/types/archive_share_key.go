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

import "strings"

// ArchiveShareKey アーカイブ共有キー
type ArchiveShareKey string

// ArchiveShareKeyDelimiter アーカイブ共有キーのデリミタ
const ArchiveShareKeyDelimiter = ":"

// String ArchiveShareKey全体を表す文字列
func (key ArchiveShareKey) String() string {
	return string(key)
}

// Zone キーに含まれるゾーン名
func (key ArchiveShareKey) Zone() string {
	tokens := key.tokens()
	if len(tokens) > 0 {
		return tokens[0]
	}
	return ""
}

// SourceArchiveID 共有元となるアーカイブのID
func (key ArchiveShareKey) SourceArchiveID() ID {
	tokens := key.tokens()
	if len(tokens) > 1 {
		return StringID(tokens[1])
	}
	return ID(0)
}

// Token 認証キー本体
func (key ArchiveShareKey) Token() string {
	tokens := key.tokens()
	if len(tokens) > 2 {
		return tokens[2]
	}
	return ""
}

// ValidFormat 有効なキーフォーマットか
func (key ArchiveShareKey) ValidFormat() bool {
	return len(key.tokens()) == 3
}

func (key ArchiveShareKey) tokens() []string {
	return strings.Split(key.String(), ArchiveShareKeyDelimiter)
}
