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

import "github.com/sacloud/iaas-api-go/types"

// DiskEdit ディスクの修正パラメータ
type DiskEdit struct {
	Password            string            `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // パスワード
	SSHKey              *DiskEditSSHKey   `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // 公開鍵(単体)
	SSHKeys             []*DiskEditSSHKey `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // 公開鍵(複数)
	DisablePWAuth       bool              `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // パスワード認証無効化フラグ
	EnableDHCP          bool              `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // DHCPの有効化
	ChangePartitionUUID bool              `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // パーティションのUUID変更
	HostName            string            `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // ホスト名
	Notes               []*DiskEditNote   `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // スタートアップスクリプト
	UserIPAddress       string            `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // IPアドレス
	UserSubnet          *UserSubnet       `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // デフォルトルート/サブネットマスク長
	Background          bool              `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // バックグラウンド実行
}

// DiskEditSSHKey ディスク修正時のSSHキー
type DiskEditSSHKey struct {
	ID        types.ID `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PublicKey string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// DiskEditNote ディスクの修正で指定するスタートアップスクリプト
type DiskEditNote struct {
	ID        types.ID               `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	APIKey    *APIKey                `json:",omitempty" yaml:"api_key,omitempty" structs:",omitempty"`
	Variables map[string]interface{} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

type APIKey struct {
	ID types.ID `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}
