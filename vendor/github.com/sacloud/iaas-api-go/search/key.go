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

import "fmt"

// FilterKey 検索条件(Filter)のキー
type FilterKey struct {
	// フィールド名
	Field string
	// 演算子
	Op ComparisonOperator
}

// String Keyの文字列表現
func (k *FilterKey) String() string {
	return fmt.Sprintf("%s%s", k.Field, k.Op)
}

// Key キーの作成
func Key(field string) FilterKey {
	return FilterKey{Field: field}
}

// KeyWithOp 演算子を指定してキーを作成
func KeyWithOp(field string, op ComparisonOperator) FilterKey {
	return FilterKey{Field: field, Op: op}
}
