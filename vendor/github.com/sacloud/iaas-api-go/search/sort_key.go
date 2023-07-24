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

package search

import "encoding/json"

// SortOrder ソート順
type SortOrder int

const (
	// SortAsc 昇順(デフォルト)
	SortAsc SortOrder = iota
	// SortDesc 降順
	SortDesc
)

// SortKeys ソート順指定
type SortKeys []SortKey

// SortKey ソート順指定対象のフィールド名
type SortKey struct {
	Key   string
	Order SortOrder
}

// SortKeyAsc 昇順ソートキー
func SortKeyAsc(key string) SortKey {
	return SortKey{
		Key:   key,
		Order: SortAsc,
	}
}

// SortKeyDesc 降順ソートキー
func SortKeyDesc(key string) SortKey {
	return SortKey{
		Key:   key,
		Order: SortDesc,
	}
}

// MarshalJSON キーの文字列表現
func (k SortKey) MarshalJSON() ([]byte, error) {
	s := k.Key
	if k.Order == SortDesc {
		s = "-" + k.Key
	}
	return json.Marshal(s)
}
