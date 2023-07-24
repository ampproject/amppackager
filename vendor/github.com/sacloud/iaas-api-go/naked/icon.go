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

// Icon アイコン
type Icon struct {
	ID           types.ID            `yaml:"id"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Tags         types.Tags          `yaml:"tags"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Scope        types.EScope        `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	URL          string              `json:",omitempty" yaml:"url,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`

	Image string `json:",omitempty" yaml:"image,omitempty" structs:",omitempty"` // 画像データBase64文字列(画像アップロード時に利用)
}
