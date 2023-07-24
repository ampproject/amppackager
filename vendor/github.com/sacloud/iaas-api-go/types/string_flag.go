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
	"strings"
)

var trueStrings = []string{"true", "on", "1"}

var (
	// StringTrue true値
	StringTrue = StringFlag(true)
	// StringFalse false値
	StringFalse = StringFlag(false)
)

// StringFlag bool型のラッパー、文字列(true/false/on/off/1/0)などをbool値として扱う
//
// - 大文字/小文字の区別はしない
// - 空文字だった場合はfalse
// - 小文字にした場合に次のいずれかにマッチしない場合はfalse [ true / on / 1 ]
type StringFlag bool

// String StringFlagの文字列表現
func (f *StringFlag) String() string {
	if f.Bool() {
		return "True"
	}
	return "False"
}

// Bool StringFlagのbool表現
func (f *StringFlag) Bool() bool {
	return f != nil && bool(*f)
}

// MarshalJSON 文字列でのJSON出力に対応するためのMarshalJSON実装
func (f *StringFlag) MarshalJSON() ([]byte, error) {
	if f != nil && bool(*f) {
		return []byte(`"True"`), nil
	}
	return []byte(`"False"`), nil
}

// UnmarshalJSON 文字列に対応するためのUnmarshalJSON実装
func (f *StringFlag) UnmarshalJSON(b []byte) error {
	s := strings.ReplaceAll(strings.ToLower(string(b)), `"`, ``)
	res := false
	for _, strTrue := range trueStrings {
		if s == strTrue {
			res = true
			break
		}
	}
	*f = StringFlag(res)
	return nil
}
