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

type LoadBalancerServers []*LoadBalancerServer

// AddGSLBServer サーバの追加
func (o *LoadBalancerServers) Add(server *LoadBalancerServer) {
	if o.Exist(server) {
		return // noop if already exists
	}
	*o = append(*o, server)
}

// Exist サーバの存在確認
func (o *LoadBalancerServers) Exist(server *LoadBalancerServer) bool {
	for _, v := range *o {
		if v.IPAddress == server.IPAddress {
			return true
		}
	}
	return false
}

// ExistAt サーバの存在確認
func (o *LoadBalancerServers) ExistAt(ip string) bool {
	return o.Exist(&LoadBalancerServer{IPAddress: ip})
}

// Find サーバの検索
func (o *LoadBalancerServers) Find(server *LoadBalancerServer) *LoadBalancerServer {
	for _, v := range *o {
		if v.IPAddress == server.IPAddress {
			return v
		}
	}
	return nil
}

// FindAt サーバの検索
func (o *LoadBalancerServers) FindAt(ip string) *LoadBalancerServer {
	return o.Find(&LoadBalancerServer{IPAddress: ip})
}

// Update サーバの更新
func (o *LoadBalancerServers) Update(old *LoadBalancerServer, new *LoadBalancerServer) {
	for _, v := range *o {
		if v.IPAddress == old.IPAddress {
			*v = *new
			return
		}
	}
}

// UpdateAt サーバの更新
func (o *LoadBalancerServers) UpdateAt(ip string, new *LoadBalancerServer) {
	o.Update(&LoadBalancerServer{IPAddress: ip}, new)
}

// Delete サーバの削除
func (o *LoadBalancerServers) Delete(server *LoadBalancerServer) {
	var res []*LoadBalancerServer
	for _, v := range *o {
		if v.IPAddress != server.IPAddress {
			res = append(res, v)
		}
	}
	*o = res
}

// DeleteAt サーバの削除
func (o *LoadBalancerServers) DeleteAt(ip string) {
	o.Delete(&LoadBalancerServer{IPAddress: ip})
}
