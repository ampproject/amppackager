// Copyright 2022-2023 The sacloud/packages-go Authors
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

package objutil

import "reflect"

// ToSlice vがsliceだったら[]interface{}にする
// vが非スライスだったらvを単一の要素とする[]interface{}を返す
func ToSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice {
		return []interface{}{v}
	}

	var results []interface{}
	for i := 0; i < rv.Len(); i++ {
		results = append(results, rv.Index(i).Interface())
	}
	return results
}
