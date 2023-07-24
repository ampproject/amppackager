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
	"fmt"
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// Disk ディスク
type Disk struct {
	ID              types.ID              `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name            string                `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description     string                `yaml:"description"`
	Tags            types.Tags            `yaml:"tags"`
	Icon            *Icon                 `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt       *time.Time            `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt      *time.Time            `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability    types.EAvailability   `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass    string                `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	SizeMB          int                   `json:",omitempty" yaml:"size_mb,omitempty" structs:",omitempty"`
	MigratedMB      int                   `json:",omitempty" yaml:"migrated_mb,omitempty" structs:",omitempty"`
	Connection      types.EDiskConnection `json:",omitempty" yaml:"connection,omitempty" structs:",omitempty"`
	ConnectionOrder int                   `json:",omitempty" yaml:"connection_order,omitempty" structs:",omitempty"`
	ReinstallCount  int                   `json:",omitempty" yaml:"reinstall_count,omitempty" structs:",omitempty"`
	JobStatus       *MigrationJobStatus   `json:",omitempty" yaml:"job_status,omitempty" structs:",omitempty"`
	Plan            *DiskPlan             `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	SourceDisk      *Disk                 `json:",omitempty" yaml:"source_disk,omitempty" structs:",omitempty"`
	SourceArchive   *Archive              `json:",omitempty" yaml:"source_archive,omitempty" structs:",omitempty"`
	BundleInfo      *BundleInfo           `json:",omitempty" yaml:"bundle_info,omitempty" structs:",omitempty"`
	Storage         *Storage              `json:",omitempty" yaml:"storage,omitempty" structs:",omitempty"`
	Server          *Server               `json:",omitempty" yaml:"server,omitempty" structs:",omitempty"`
}

// MigrationJobStatus マイグレーションジョブステータス
type MigrationJobStatus struct {
	Status      string          `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"` // ステータス
	ConfigError *JobConfigError `json:",omitempty" yaml:"config_error,omitempty" structs:",omitempty"`
	Delays      *struct {       // Delays
		Start *struct { // 開始
			Max int `json:",omitempty" yaml:"max,omitempty" structs:",omitempty"` // 最大
			Min int `json:",omitempty" yaml:"min,omitempty" structs:",omitempty"` // 最小
		} `json:",omitempty" yaml:"start,omitempty" structs:",omitempty"`

		Finish *struct { // 終了
			Max int `json:",omitempty" yaml:"max,omitempty" structs:",omitempty"` // 最大
			Min int `json:",omitempty" yaml:"min,omitempty" structs:",omitempty"` // 最小
		} `json:",omitempty" yaml:"finish,omitempty" structs:",omitempty"`
	}
}

// JobConfigError マイグレーションジョブのエラー
type JobConfigError struct {
	ErrorCode string `json:",omitempty" yaml:"error_code,omitempty" structs:",omitempty"`
	ErrorMsg  string `json:",omitempty" yaml:"error_msg,omitempty" structs:",omitempty"`
	Status    string `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// String マイグレーションジョブエラーの文字列表現
func (e *JobConfigError) String() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorMsg)
}

// ResizePartitionRequest リサイズ時のオプション
type ResizePartitionRequest struct {
	Background bool `yaml:"background"`
}
