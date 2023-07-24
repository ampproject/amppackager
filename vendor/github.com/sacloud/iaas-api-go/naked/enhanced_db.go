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

// EnhancedDB コンテナレジストリ
type EnhancedDB struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Tags         types.Tags          `yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Provider     *Provider           `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Config       *EnhancedDBConfig   `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *EnhancedDBStatus   `json:",omitempty" yaml:"status" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}

// EnhancedDBConfig コンフィグ
type EnhancedDBConfig struct {
	MaxConnections  int      `json:"max_connections,omitempty" yaml:"max_connections,omitempty" structs:",omitempty"`
	AllowedNetworks []string `json:"allowed_networks" yaml:"allowed_networks"`
}

// EnhancedDBStatus ステータス
type EnhancedDBStatus struct {
	DatabaseName string                 `json:"database_name,omitempty" yaml:"database_name,omitempty" structs:",omitempty"`
	DatabaseType types.EnhancedDBType   `json:"database_type,omitempty" yaml:"database_type,omitempty" structs:",omitempty"`
	Region       types.EnhancedDBRegion `json:"region,omitempty" yaml:"region,omitempty" structs:",omitempty"`
	HostName     string                 `json:"hostname,omitempty" yaml:"hostname,omitempty" structs:",omitempty"`
	Port         int                    `json:"port,omitempty" yaml:"port,omitempty" structs:",omitempty"`
}

type EnhancedDBConfigSettings struct {
	EnhancedDB *EnhancedDBConfig `json:",omitempty" yaml:"enhanced_db,omitempty" structs:",omitempty"`
}

// EnhancedDBPasswordSettings セッティング
type EnhancedDBPasswordSettings struct {
	EnhancedDB *EnhancedDBPasswordSetting `json:",omitempty" yaml:"enhanced_db,omitempty" structs:",omitempty"`
}

// EnhancedDBPasswordSetting .
type EnhancedDBPasswordSetting struct {
	Password string `json:"password,omitempty" yaml:"password,omitempty" structs:",omitempty"`
}
