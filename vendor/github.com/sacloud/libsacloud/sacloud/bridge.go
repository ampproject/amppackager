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

import (
	"encoding/json"
)

// Bridge ブリッジ
type Bridge struct {
	*Resource        // ID
	propName         // 名称
	propDescription  // 説明
	propServiceClass // サービスクラス
	propRegion       // リージョン
	propCreatedAt    // 作成日時

	Info *struct { // インフォ
		Switches []*struct { // 接続スイッチリスト
			*Switch             // スイッチ
			ID      json.Number `json:",omitempty"` // (HACK) ID
		}
	}

	SwitchInZone *struct { // ゾーン内接続スイッチ
		*Resource             // ID
		propScope             // スコープ
		Name           string `json:",omitempty"` // 名称
		ServerCount    int    `json:",omitempty"` // 接続サーバー数
		ApplianceCount int    `json:",omitempty"` // 接続アプライアンス数
	}
}
