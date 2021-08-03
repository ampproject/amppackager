// Copyright 2016-2020 The Libsacloud Authors
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

package sacloud

// PublicPrice 料金
type PublicPrice struct {
	DisplayName      string `json:",omitempty"` // 表示名
	IsPublic         bool   `json:",omitempty"` // 公開フラグ
	ServiceClassID   int    `json:",omitempty"` // サービスクラスID
	ServiceClassName string `json:",omitempty"` // サービスクラス名
	ServiceClassPath string `json:",omitempty"` // サービスクラスパス

	Price struct { // 価格
		Base    int    `json:",omitempty"` // 基本料金
		Daily   int    `json:",omitempty"` // 日単位料金
		Hourly  int    `json:",omitempty"` // 時間単位料金
		Monthly int    `json:",omitempty"` // 分単位料金
		Zone    string `json:",omitempty"` // ゾーン
	}
}
