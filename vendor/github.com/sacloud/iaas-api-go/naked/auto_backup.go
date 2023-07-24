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

package naked

import (
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// AutoBackup 自動バックアップ
type AutoBackup struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Tags         types.Tags          `yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Provider     *Provider           `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *AutoBackupSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"setting_hash,omitempty" structs:",omitempty"`
	Status       *AutoBackupStatus   `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// AutoBackupSettingsUpdate 自動バックアップ
type AutoBackupSettingsUpdate struct {
	Settings     *AutoBackupSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"setting_hash,omitempty" structs:",omitempty"`
}

// AutoBackupSettings 自動バックアップ設定
type AutoBackupSettings struct {
	Autobackup *AutoBackupSetting `json:",omitempty" yaml:"autobackup,omitempty" structs:",omitempty"` // HACK: 注: API側がキャメルケースになっていない
}

// AutoBackupSetting 自動バックアップ設定
type AutoBackupSetting struct {
	BackupSpanType          types.EBackupSpanType `json:",omitempty" yaml:"backup_span_type,omitempty" structs:",omitempty"`
	BackupSpanWeekdays      []types.EDayOfTheWeek `json:",omitempty" yaml:"backup_span_weekdays,omitempty" structs:",omitempty"`
	MaximumNumberOfArchives int                   `json:",omitempty" yaml:"maximum_number_of_archives,omitempty" structs:",omitempty"`
}

// AutoBackupStatus 自動バックアップステータス
type AutoBackupStatus struct {
	DiskID    types.ID `json:"DiskId,omitempty" yaml:"disk_id,omitempty" structs:",omitempty"`
	AccountID types.ID `json:"AccountId,omitempty" yaml:"account_id,omitempty" structs:",omitempty"`
	ZoneID    types.ID `json:"ZoneId,omitempty" yaml:"zone_id,omitempty" structs:",omitempty"`
	ZoneName  string   `json:",omitempty" yaml:"zone_name,omitempty" structs:",omitempty"`
}
