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

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// EqualExpression Equalで比較する際の条件
type EqualExpression struct {
	Op         LogicalOperator
	Conditions []interface{}
}

// PartialMatch 部分一致(Partial Match)かつAND条件を示すEqualFilterを作成
//
// AndEqualのエイリアス
func PartialMatch(conditions ...string) *EqualExpression {
	return AndEqual(conditions...)
}

// ExactMatch 完全一致(Partial Match)かつOR条件を示すEqualFilterを作成
//
// OrEqualのエイリアス
func ExactMatch(conditions ...string) *EqualExpression {
	var values []interface{}
	for _, p := range conditions {
		values = append(values, p)
	}
	return OrEqual(values...)
}

// AndEqual 部分一致(Partial Match)かつAND条件を示すEqualFilterを作成
func AndEqual(conditions ...string) *EqualExpression {
	var values []interface{}
	for _, p := range conditions {
		values = append(values, p)
	}

	return &EqualExpression{
		Op:         OpAnd,
		Conditions: values,
	}
}

// OrEqual 完全一致(Exact Match)かつOR条件を示すEqualFilterを作成
func OrEqual(conditions ...interface{}) *EqualExpression {
	return &EqualExpression{
		Op:         OpOr,
		Conditions: conditions,
	}
}

// TagsAndEqual タグのAND検索
func TagsAndEqual(tags ...string) *EqualExpression {
	return OrEqual(tags) // 配列のまま渡す
}

// MarshalJSON .
func (eq *EqualExpression) MarshalJSON() ([]byte, error) {
	var conditions []interface{}
	for _, cond := range eq.Conditions {
		var c interface{}
		switch v := cond.(type) {
		case time.Time:
			c = v.Format(time.RFC3339)
		case string:
			c = escapeFilterString(v)
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64:
			c = fmt.Sprintf("%v", v)
		default:
			c = v
		}
		if c != nil {
			conditions = append(conditions, c)
		}
	}

	var value interface{}
	switch eq.Op {
	case OpOr:
		value = conditions
	case OpAnd:
		var strConds []string
		for _, c := range conditions {
			strConds = append(strConds, c.(string)) // Note: OpAndで文字列以外の要素が指定された場合のことは考慮しない
		}
		value = strings.Join(strConds, "%20")
	default:
		return nil, fmt.Errorf("invalid search.LogicalOperator: %v", eq.Op)
	}

	return json.Marshal(value)
}
