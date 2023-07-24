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

// AutoScale オートスケール
type AutoScale struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Tags         types.Tags          `yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Provider     *Provider           `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *AutoScaleSettings  `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *AutoScaleStatus    `json:",omitempty" yaml:"status" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}

// AutoScaleSettingsUpdate オートスケール更新パラメータ
type AutoScaleSettingsUpdate struct {
	Settings     *AutoScaleSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string             `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
}

// AutoScaleSettings セッティング
type AutoScaleSettings struct {
	TriggerType            types.EAutoScaleTriggerType
	CPUThresholdScaling    *AutoScaleCPUThresholdScaling    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RouterThresholdScaling *AutoScaleRouterThresholdScaling `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ScheduleScaling        []*AutoScaleScheduleScaling
	Zones                  []string `json:"SakuraCloudZones"`
	Config                 string   `json:",omitempty" yaml:",omitempty"`
	Disabled               bool
}

type AutoScaleCPUThresholdScaling struct {
	ServerPrefix string `json:",omitempty" yaml:",omitempty"`
	Up           int    `json:",omitempty" yaml:",omitempty"`
	Down         int    `json:",omitempty" yaml:",omitempty"`
}

type AutoScaleRouterThresholdScaling struct {
	RouterPrefix string `json:",omitempty" yaml:",omitempty"`
	Direction    string `json:",omitempty" yaml:",omitempty"`
	Mbps         int    `json:",omitempty" yaml:",omitempty"`
}

type AutoScaleScheduleScaling struct {
	Action    types.EAutoScaleAction
	Hour      int
	Minute    int
	DayOfWeek []types.EDayOfTheWeek
}

// AutoScaleStatus ステータス
type AutoScaleStatus struct {
	RegisteredBy string           `json:",omitempty" yaml:",omitempty"`
	APIKey       *AutoScaleAPIKey `json:",omitempty" yaml:",omitempty"`
}

type AutoScaleAPIKey struct {
	ID string `json:",omitempty" yaml:",omitempty"`
}

// AutoScaleRunningStatus /statusの戻り値
type AutoScaleRunningStatus struct {
	LatestLogs    []string
	ResourcesText string
}
