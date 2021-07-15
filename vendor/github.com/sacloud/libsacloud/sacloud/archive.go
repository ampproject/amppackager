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

// AllowArchiveSizes 作成できるアーカイブのサイズ
func AllowArchiveSizes() []string {
	return []string{"20", "40", "60", "80", "100", "250", "500", "750", "1024"}
}

// Archive アーカイブ
type Archive struct {
	*Resource             // ID
	propAvailability      // 有功状態
	propName              // 名称
	propDescription       // 説明
	propSizeMB            // サイズ(MB単位)
	propMigratedMB        // コピー済みデータサイズ(MB単位)
	propScope             // スコープ
	propCopySource        // コピー元情報
	propServiceClass      // サービスクラス
	propPlanID            // プランID
	propJobStatus         // マイグレーションジョブステータス
	propOriginalArchiveID // オリジナルアーカイブID
	propStorage           // ストレージ
	propBundleInfo        // バンドル情報
	propTags              // タグ
	propIcon              // アイコン
	propCreatedAt         // 作成日時
}
