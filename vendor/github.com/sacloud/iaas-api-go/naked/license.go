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

// License ライセンス
type License struct {
	ID          types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string       `yaml:"description"`
	CreatedAt   *time.Time   `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt  *time.Time   `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	LicenseInfo *LicenseInfo `json:",omitempty" yaml:"license_info,omitempty" structs:",omitempty"` // ライセンス情報
}

// LicenseInfo ライセンスプラン
type LicenseInfo struct {
	ID           types.ID   `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string     `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	ServiceClass string     `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	TermsOfUse   string     `json:",omitempty" yaml:"terms_of_use,omitempty" structs:",omitempty"` // 利用規約
}
