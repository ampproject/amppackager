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

// AutoBackup 自動バックアップ(CommonServiceItem)
type AutoBackup struct {
	*Resource        // ID
	propName         // 名称
	propDescription  // 説明
	propServiceClass // サービスクラス
	propIcon         // アイコン
	propTags         // タグ
	propCreatedAt    // 作成日時
	propModifiedAt   // 変更日時

	Status   *AutoBackupStatus   `json:",omitempty"` // ステータス
	Provider *AutoBackupProvider `json:",omitempty"` // プロバイダ
	Settings *AutoBackupSettings `json:",omitempty"` // 設定

}

// AutoBackupSettings 自動バックアップ設定
type AutoBackupSettings struct {
	AccountID  ID                    `json:"AccountId,omitempty"` // アカウントID
	DiskID     ID                    `json:"DiskId,omitempty"`    // ディスクID
	ZoneID     ID                    `json:"ZoneId,omitempty"`    // ゾーンID
	ZoneName   string                `json:",omitempty"`          // ゾーン名称
	Autobackup *AutoBackupRecordSets `json:",omitempty"`          // 自動バックアップ定義

}

// AutoBackupStatus 自動バックアップステータス
type AutoBackupStatus struct {
	AccountID ID     `json:"AccountId,omitempty"` // アカウントID
	DiskID    ID     `json:"DiskId,omitempty"`    // ディスクID
	ZoneID    ID     `json:"ZoneId,omitempty"`    // ゾーンID
	ZoneName  string `json:",omitempty"`          // ゾーン名称
}

// AutoBackupProvider 自動バックアッププロバイダ
type AutoBackupProvider struct {
	Class string `json:",omitempty"` // クラス
}

// CreateNewAutoBackup 自動バックアップ 作成(CommonServiceItem)
func CreateNewAutoBackup(backupName string, diskID ID) *AutoBackup {
	return &AutoBackup{
		Resource: &Resource{},
		propName: propName{Name: backupName},
		Status: &AutoBackupStatus{
			DiskID: diskID,
		},
		Provider: &AutoBackupProvider{
			Class: "autobackup",
		},
		Settings: &AutoBackupSettings{
			Autobackup: &AutoBackupRecordSets{
				BackupSpanType: "weekdays",
			},
		},
	}
}

// AllowAutoBackupWeekdays 自動バックアップ実行曜日リスト
func AllowAutoBackupWeekdays() []string {
	return []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
}

// AutoBackupRecordSets 自動バックアップ定義
type AutoBackupRecordSets struct {
	BackupSpanType          string   // バックアップ間隔タイプ
	BackupSpanWeekdays      []string // バックアップ実施曜日
	MaximumNumberOfArchives int      // 世代数

}

// SetBackupSpanWeekdays バックアップ実行曜日設定
func (a *AutoBackup) SetBackupSpanWeekdays(weekdays []string) {
	a.Settings.Autobackup.BackupSpanWeekdays = weekdays
}

// SetBackupMaximumNumberOfArchives 世代数設定
func (a *AutoBackup) SetBackupMaximumNumberOfArchives(max int) {
	a.Settings.Autobackup.MaximumNumberOfArchives = max
}
