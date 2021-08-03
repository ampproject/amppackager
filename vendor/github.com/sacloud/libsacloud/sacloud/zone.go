// Copyright 2016-2020 The Libsacloud Authors
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

package sacloud

// Zone ゾーン
type Zone struct {
	*Resource       // ID
	propName        // 名称
	propDescription // 説明
	propRegion      // リージョン

	IsDummy bool `json:",omitempty"` // IsDummy ダミーフラグ

	VNCProxy struct { // VNCProxy VPCプロキシ
		HostName  string `json:",omitempty"` // HostName ホスト名
		IPAddress string `json:",omitempty"` // IPAddress IPアドレス
	} `json:",omitempty"`

	FTPServer struct { // FTPServer FTPサーバー
		HostName  string `json:",omitempty"` // HostName ホスト名
		IPAddress string `json:",omitempty"` // IPAddress IPアドレス
	} `json:",omitempty"`
}

// ZoneIsDummy ダミーフラグ 取得
func (z *Zone) ZoneIsDummy() bool {
	return z.IsDummy
}

// GetVNCProxyHostName VNCプロキシホスト名 取得
func (z *Zone) GetVNCProxyHostName() string {
	return z.VNCProxy.HostName
}

// GetVPCProxyIPAddress VNCプロキシIPアドレス 取得
func (z *Zone) GetVPCProxyIPAddress() string {
	return z.VNCProxy.IPAddress
}

// GetFTPHostName FTPサーバーホスト名 取得
func (z *Zone) GetFTPHostName() string {
	return z.FTPServer.HostName
}

// GetFTPServerIPAddress FTPサーバーIPアドレス 取得
func (z *Zone) GetFTPServerIPAddress() string {
	return z.FTPServer.IPAddress
}
