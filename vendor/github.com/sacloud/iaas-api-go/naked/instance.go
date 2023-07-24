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

// Instance サーバなどの起動情報
type Instance struct {
	Host            *Host                       `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`
	Status          types.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
	BeforeStatus    types.EServerInstanceStatus `json:",omitempty" yaml:"before_status,omitempty" structs:",omitempty"`
	StatusChangedAt *time.Time                  `json:",omitempty" yaml:"status_changed_at,omitempty" structs:",omitempty"`
	ModifiedAt      *time.Time                  `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Warnings        string                      `json:",omitempty" yaml:"warnings,omitempty" structs:",omitempty"`
	WarningsValue   int                         `json:",omitempty" yaml:"warnings_value,omitempty" structs:",omitempty"`
	CDROM           *CDROM                      `json:",omitempty" yaml:"cdrom,omitempty" structs:",omitempty"`
}
