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

// License ライセンス
type License struct {
	*Resource       // ID
	propName        // 名称
	propDescription // 説明
	propCreatedAt   // 作成日時
	propModifiedAt  // 変更日時

	LicenseInfo *ProductLicense `json:",omitempty"` // ライセンス情報
}

// GetLicenseInfo ライセンス情報 取得
func (l *License) GetLicenseInfo() *ProductLicense {
	return l.LicenseInfo
}

// SetLicenseInfo ライセンス情報 設定
func (l *License) SetLicenseInfo(license *ProductLicense) {
	l.LicenseInfo = license
}

// SetLicenseInfoByID ライセンス情報 設定
func (l *License) SetLicenseInfoByID(id ID) {
	l.LicenseInfo = &ProductLicense{Resource: NewResource(id)}
}
