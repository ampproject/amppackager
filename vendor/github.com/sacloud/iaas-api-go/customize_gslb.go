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

package iaas

type GSLBServers []*GSLBServer

// NewGSLBServer GSLB実サーバの作成
func NewGSLBServer(ip string) *GSLBServer {
	return &GSLBServer{
		IPAddress: ip,
		Enabled:   true,
		Weight:    1,
	}
}

// AddGSLBServer サーバの追加
func (o *GSLBServers) Add(server *GSLBServer) {
	if o.Exist(server) {
		return // noop if already exists
	}
	*o = append(*o, server)
}

// Exist サーバの存在確認
func (o *GSLBServers) Exist(server *GSLBServer) bool {
	for _, s := range *o {
		if s.IPAddress == server.IPAddress {
			return true
		}
	}
	return false
}

// ExistAt サーバの存在確認
func (o *GSLBServers) ExistAt(ip string) bool {
	return o.Exist(NewGSLBServer(ip))
}

// Find サーバの検索
func (o *GSLBServers) Find(server *GSLBServer) *GSLBServer {
	for _, s := range *o {
		if s.IPAddress == server.IPAddress {
			return s
		}
	}
	return nil
}

// FindAt サーバの検索
func (o *GSLBServers) FindAt(ip string) *GSLBServer {
	return o.Find(NewGSLBServer(ip))
}

// Update サーバの更新
func (o *GSLBServers) Update(old *GSLBServer, new *GSLBServer) {
	for _, s := range *o {
		if s.IPAddress == old.IPAddress {
			*s = *new
			return
		}
	}
}

// UpdateAt サーバの更新
func (o *GSLBServers) UpdateAt(ip string, new *GSLBServer) {
	o.Update(NewGSLBServer(ip), new)
}

// Delete サーバの削除
func (o *GSLBServers) Delete(server *GSLBServer) {
	var res []*GSLBServer
	for _, s := range *o {
		if s.IPAddress != server.IPAddress {
			res = append(res, s)
		}
	}
	*o = res
}

// DeleteAt サーバの削除
func (o *GSLBServers) DeleteAt(ip string) {
	o.Delete(NewGSLBServer(ip))
}
