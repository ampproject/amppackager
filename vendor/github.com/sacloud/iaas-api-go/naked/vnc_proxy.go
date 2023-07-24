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

// VNCProxy VNCプロキシ
type VNCProxy struct {
	HostName  string `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
}

// VNCProxyInfo サーバに対するVNCProxy
type VNCProxyInfo struct {
	Status       string `json:",omitempty"` // ステータス
	Host         string `json:",omitempty"` // プロキシホスト
	IOServerHost string `json:",omitempty"` // 新プロキシホスト(Hostがlocalhostの場合にこちらを利用する)
	Port         string `json:",omitempty"` // ポート番号
	Password     string `json:",omitempty"` // VNCパスワード
	VNCFile      string `json:",omitempty"` // VNC接続情報ファイル(VNCビューア用)
}
