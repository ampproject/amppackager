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
	"fmt"
	"strconv"
)

// ID さくらのクラウド上のリソースのIDを示す
//
// APIリクエスト/レスポンスに文字列/数値が混在するためここで吸収する
type ID int64

// UnmarshalJSON implements unmarshal from both of JSON number and JSON string
func (i *ID) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "" || s == "null" {
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	id, err := n.Int64()
	if err != nil {
		return err
	}
	*i = ID(id)
	return nil
}

// IsEmpty return true if ID is empty or zero value
func (i ID) IsEmpty() bool {
	return i.Int64() == 0
}

// String returns the literal text of the number.
func (i ID) String() string {
	if i.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%d", i)
}

// Int64 returns the number as an int64.
func (i ID) Int64() int64 {
	return int64(i)
}

// Int64ID creates new ID from int64
func Int64ID(id int64) ID {
	return ID(id)
}

// StringID creates new ID from string
func StringID(id string) ID {
	intID, _ := strconv.ParseInt(id, 10, 64)
	return ID(intID)
}

// IDs IDのコレクション型
type IDs []ID

// StringSlice stringスライスを返す
func (i IDs) StringSlice() []string {
	var ret []string
	for _, id := range i {
		ret = append(ret, id.String())
	}
	return ret
}

// Int64Slice Int64スライスを返す
func (i IDs) Int64Slice() []int64 {
	var ret []int64
	for _, id := range i {
		ret = append(ret, id.Int64())
	}
	return ret
}

// IsEmptyAll すべてゼロ値な場合にtrueを返す
func (i IDs) IsEmptyAll() bool {
	for _, id := range i {
		if !id.IsEmpty() {
			return false
		}
	}
	return true
}

// IsEmptyAny 1つでもゼロ値な場合にtrueを返す
func (i IDs) IsEmptyAny() bool {
	for _, id := range i {
		if id.IsEmpty() {
			return true
		}
	}
	return false
}
