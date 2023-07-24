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
	"net/url"
	"time"
)

// Filter 検索系APIでの検索条件
//
// Note: libsacloudではリクエスト時に`X-Sakura-Bigint-As-Int`ヘッダを指定することで
// 文字列で表されているBitintをintとして取得している。
// このため、libsacloud側では数値型に見える項目でもさくらのクラウド側では文字列となっている場合がある。
// これらの項目ではOpEqual以外の演算子は利用できない。
// また、これらの項目でスカラ値を検索条件に与えた場合は部分一致ではなく完全一致となるため注意。
type Filter map[FilterKey]interface{}

// MarshalJSON 検索系APIコール時のGETパラメータを出力するためのjson.Marshaler実装
func (f Filter) MarshalJSON() ([]byte, error) {
	result := make(map[string]interface{})

	for key, expression := range f {
		if expression == nil {
			continue
		}

		exp := expression
		switch key.Op {
		case OpEqual:
			if _, ok := exp.(*EqualExpression); !ok {
				exp = OrEqual(expression)
			}
		default:
			v, err := convertToValidFilterCondition(exp)
			if err != nil {
				return nil, err
			}
			exp = v
		}

		result[key.String()] = exp
	}

	return json.Marshal(result)
}

func convertToValidFilterCondition(v interface{}) (string, error) {
	switch v := v.(type) {
	case time.Time:
		return v.Format(time.RFC3339), nil
	case string:
		return escapeFilterString(v), nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprintf("%v", v), nil
	}
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func escapeFilterString(s string) string {
	// HACK さくらのクラウド側でqueryStringでの+エスケープに対応していないため、
	// %20にエスケープされるurl.Pathを利用する。
	// http://qiita.com/shibukawa/items/c0730092371c0e243f62
	//
	// UPDATE: https://github.com/sacloud/libsacloud/issues/657#issuecomment-733467472
	// (&url.URL{Path:s}).String()だと、MACAddressが"./00:00:5E:00:53:00"のようになってしまう。
	// このためurl.PathEscapeを利用する。
	return url.PathEscape(s)
}
