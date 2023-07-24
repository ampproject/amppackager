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

import (
	"encoding/json"
	"sort"
)

// Tags タグ
type Tags []string

// Sort 昇順でソートする
func (t Tags) Sort() {
	sort.Strings([]string(t))
}

// MarshalJSON タグを空にする場合への対応
func (t Tags) MarshalJSON() ([]byte, error) {
	tags := t
	if tags == nil {
		tags = make([]string, 0)
	}

	tags.Sort()
	type alias Tags
	tmp := alias(tags)
	return json.Marshal(tmp)
}

// UnmarshalJSON タグを空にする場合への対応
func (t *Tags) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		*t = make([]string, 0)
		return nil
	}
	type alias Tags
	tmp := alias(*t)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = Tags(tmp)
	t.Sort()
	return nil
}
