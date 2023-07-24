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
)

// WebUI データベースアプライアンスでのUI設定
type WebUI string

// String WebUIの文字列表現
func (w *WebUI) String() string {
	return string(*w)
}

// Bool WebUIの有効/無効
func (w *WebUI) Bool() bool {
	s := w.String()
	if s == "" || s == "false" {
		return false
	}
	return true
}

// ToWebUI bool値からWebUI型へ変換
func ToWebUI(v bool) WebUI {
	return WebUI(fmt.Sprintf("%t", v))
}

// MarshalJSON boolとstring両方に対応するための実装
func (w WebUI) MarshalJSON() ([]byte, error) {
	s := string(w)
	switch s {
	case "true", `1`, `"1"`:
		return json.Marshal(true)
	case "", `""`, `0`, `"0"`, "false":
		return json.Marshal(false)
	default:
		return json.Marshal(s)
	}
}

// UnmarshalJSON 文字列/boolが混在することへの対応
func (w *WebUI) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "true", `1`, `"1"`:
		*w = WebUI("true")
	case "", `""`, `0`, `"0"`, "false":
		*w = WebUI("false")
	default:
		*w = WebUI(s)
	}
	return nil
}
