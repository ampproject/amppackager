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

// StringNumber 数値型を文字列で表す型
type StringNumber float64

// MarshalJSON implements json.Marshaler
func (n *StringNumber) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, n.String())), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (n *StringNumber) UnmarshalJSON(b []byte) error {
	if string(b) == `""` {
		*n = StringNumber(0)
		return nil
	}

	var num json.Number
	if err := json.Unmarshal(b, &num); err != nil {
		return err
	}
	number, err := num.Float64()
	if err != nil {
		return err
	}
	*n = StringNumber(number)
	return nil
}

// String returns the literal text of the number.
func (n StringNumber) String() string {
	if n.Int64() == 0 {
		return ""
	}
	return strconv.FormatFloat(n.Float64(), 'f', -1, 64)
}

// Int returns the number as an int.
func (n StringNumber) Int() int {
	return int(n)
}

// Int64 returns the number as an int64.
func (n StringNumber) Int64() int64 {
	return int64(n)
}

// Float64 returns the number as an float64.
func (n StringNumber) Float64() float64 {
	return float64(n)
}

// ParseStringNumber 文字列からStringNumberへの変換
func ParseStringNumber(s string) (StringNumber, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return StringNumber(0), err
	}
	return StringNumber(n), nil
}
