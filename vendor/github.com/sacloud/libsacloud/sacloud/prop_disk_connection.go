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

// propDiskConnection ディスク接続情報内包型
type propDiskConnection struct {
	Connection      EDiskConnection `json:",omitempty"` // ディスク接続方法
	ConnectionOrder int             `json:",omitempty"` // コネクション順序

}

// GetDiskConnection ディスク接続方法 取得
func (p *propDiskConnection) GetDiskConnection() EDiskConnection {
	return p.Connection
}

// SetDiskConnection ディスク接続方法 設定
func (p *propDiskConnection) SetDiskConnection(conn EDiskConnection) {
	p.Connection = conn
}

// GetDiskConnectionByStr ディスク接続方法 取得
func (p *propDiskConnection) GetDiskConnectionByStr() string {
	return string(p.Connection)
}

// SetDiskConnectionByStr ディスク接続方法 設定
func (p *propDiskConnection) SetDiskConnectionByStr(conn string) {
	p.Connection = EDiskConnection(conn)
}

// GetDiskConnectionOrder コネクション順序 取得
func (p *propDiskConnection) GetDiskConnectionOrder() int {
	return p.ConnectionOrder
}
