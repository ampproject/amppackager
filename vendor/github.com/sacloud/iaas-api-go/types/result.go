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

import "encoding/json"

// APIResult APIからの戻り値"Success"の別表現
//
// Successにはbool以外にも"Accepted"などの文字列が返ることがある(例:アプライアンス)
// このためAPIResultでUnmarshalJSONを実装してラップする
type APIResult int

const (
	// ResultUnknown 不明
	ResultUnknown APIResult = iota
	// ResultSuccess 成功
	ResultSuccess
	// ResultAccepted 受付成功
	ResultAccepted
	// ResultFailed 失敗
	ResultFailed
)

// UnmarshalJSON bool/string混在型に対応するためのUnmarshalJSON実装
func (r *APIResult) UnmarshalJSON(data []byte) error {
	*r = ResultUnknown

	// try bool
	var b bool
	if err := json.Unmarshal(data, &b); err != nil {
		// try string
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}

		if s == "Accepted" {
			*r = ResultAccepted
		}
		return nil
	}

	if b {
		*r = ResultSuccess
	} else {
		*r = ResultFailed
	}
	return nil
}
