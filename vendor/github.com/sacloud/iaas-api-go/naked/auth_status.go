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
	"github.com/sacloud/iaas-api-go/types"
)

// AuthStatus 現在の認証状態
type AuthStatus struct {
	Account            *Account                 // アカウント
	Member             *Member                  // 会員情報
	AuthClass          types.EAuthClass         `json:",omitempty" yaml:"auth_class,omitempty" structs:",omitempty"`          // 認証クラス
	AuthMethod         types.EAuthMethod        `json:",omitempty" yaml:"auth_method,omitempty" structs:",omitempty"`         // 認証方法
	ExternalPermission types.ExternalPermission `json:",omitempty" yaml:"external_permission,omitempty" structs:",omitempty"` // 他サービスへのアクセス権
	IsAPIKey           bool                     `yaml:"is_api_key"`                                                           // APIキーでのアクセスフラグ
	OperationPenalty   types.EOperationPenalty  `json:",omitempty" yaml:"operation_penalty,omitempty" structs:",omitempty"`   // オペレーションペナルティ
	Permission         types.EPermission        `json:",omitempty" yaml:"permission,omitempty" structs:",omitempty"`          // 権限
}
