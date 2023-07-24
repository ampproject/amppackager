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

// ESIMOperatorName SIMキャリア名
type ESIMOperatorName string

// String ESIMOperatorNameの文字列表現
func (o ESIMOperatorName) String() string {
	return string(o)
}

// SIMOperators SIMキャリア名
var SIMOperators = struct {
	KDDI     ESIMOperatorName
	Docomo   ESIMOperatorName
	SoftBank ESIMOperatorName
}{
	KDDI:     ESIMOperatorName("KDDI"),
	Docomo:   ESIMOperatorName("NTT DOCOMO"),
	SoftBank: ESIMOperatorName("SoftBank"),
}

// SIMOperatorShortNameMap 省略名をキーとするESIMOperatorNameのマップ
var SIMOperatorShortNameMap = map[string]ESIMOperatorName{
	"kddi":     SIMOperators.KDDI,
	"docomo":   SIMOperators.Docomo,
	"softbank": SIMOperators.SoftBank,
}

// SIMOperatorShortNames SIMOperatorの省略名リスト
func SIMOperatorShortNames() []string {
	var keys []string
	for k := range SIMOperatorShortNameMap {
		keys = append(keys, k)
	}
	return keys
}
