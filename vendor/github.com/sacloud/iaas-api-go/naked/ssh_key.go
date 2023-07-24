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

// SSHKey 公開鍵
type SSHKey struct {
	ID          types.ID   `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string     `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string     `yaml:"description"`
	CreatedAt   *time.Time `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	PublicKey   string     `json:",omitempty" yaml:"public_key,omitempty" structs:",omitempty"`  // 公開鍵
	PrivateKey  string     `json:",omitempty" yaml:"public_key,omitempty" structs:",omitempty"`  // 秘密鍵、API側での鍵生成時のみセットされる
	Fingerprint string     `json:",omitempty" yaml:"fingerprint,omitempty" structs:",omitempty"` // フィンガープリント

	GenerateFormat string `json:",omitempty" yaml:"generate_format,omitempty" structs:",omitempty"` // 鍵生成時のみ利用(openssh固定)
	PassPhrase     string `json:",omitempty" yaml:"pass_phrase,omitempty" structs:",omitempty"`     // 鍵生成時のみ利用
}
