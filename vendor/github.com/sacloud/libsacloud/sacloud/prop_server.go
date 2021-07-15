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

// propServer 接続先サーバー内包型
type propServer struct {
	Server *Server `json:",omitempty"` // 接続先サーバー
}

// GetServer 接続先サーバー 取得
func (p *propServer) GetServer() *Server {
	return p.Server
}

// SetServer 接続先サーバー 設定
func (p *propServer) SetServer(server *Server) {
	p.Server = server
}

// SetServerID サーバーIDの設定
func (p *propServer) SetServerID(id ID) {
	p.Server = &Server{Resource: &Resource{ID: id}}
}
