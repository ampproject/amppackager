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

// Icon アイコン
type Icon struct {
	*Resource        // ID
	propAvailability // 有功状態
	propName         // 名称
	propScope        // スコープ
	propTags         // タグ
	propCreatedAt    // 作成日時
	propModifiedAt   // 変更日時

	URL   string `json:",omitempty"` // アイコンURL
	Image string `json:",omitempty"` // 画像データBase64文字列(Sizeパラメータ指定時 or 画像アップロード時に利用)
}

// Image 画像データBASE64文字列
type Image string

// GetURL アイコン画像URL取得
func (icon *Icon) GetURL() string {
	return icon.URL
}

// GetImage アイコン画像データ(base64)取得
func (icon *Icon) GetImage() string {
	return icon.Image
}

// SetImage アイコン画像データ(base64)設定
func (icon *Icon) SetImage(image string) {
	icon.Image = image
}
