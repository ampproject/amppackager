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

import "github.com/sacloud/iaas-api-go/types"

type LoadBalancerVirtualIPAddresses []*LoadBalancerVirtualIPAddress

// AddGSLBServer サーバの追加
func (o *LoadBalancerVirtualIPAddresses) Add(vip *LoadBalancerVirtualIPAddress) {
	if o.Exist(vip) {
		return // noop if already exists
	}
	*o = append(*o, vip)
}

// Exist サーバの存在確認
func (o *LoadBalancerVirtualIPAddresses) Exist(vip *LoadBalancerVirtualIPAddress) bool {
	for _, v := range *o {
		if v.VirtualIPAddress == vip.VirtualIPAddress && v.Port == vip.Port {
			return true
		}
	}
	return false
}

// ExistAt サーバの存在確認
func (o *LoadBalancerVirtualIPAddresses) ExistAt(vip string, port int) bool {
	return o.Exist(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip, Port: types.StringNumber(port)})
}

// Find サーバの検索
func (o *LoadBalancerVirtualIPAddresses) Find(vip *LoadBalancerVirtualIPAddress) *LoadBalancerVirtualIPAddress {
	for _, v := range *o {
		if v.VirtualIPAddress == vip.VirtualIPAddress && v.Port == vip.Port {
			return v
		}
	}
	return nil
}

// FindAt サーバの検索
func (o *LoadBalancerVirtualIPAddresses) FindAt(vip string, port int) *LoadBalancerVirtualIPAddress {
	return o.Find(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip, Port: types.StringNumber(port)})
}

// Update サーバの更新
func (o *LoadBalancerVirtualIPAddresses) Update(old *LoadBalancerVirtualIPAddress, new *LoadBalancerVirtualIPAddress) {
	for _, v := range *o {
		if v.VirtualIPAddress == old.VirtualIPAddress && v.Port == old.Port {
			*v = *new
			return
		}
	}
}

// UpdateAt サーバの更新
func (o *LoadBalancerVirtualIPAddresses) UpdateAt(vip string, port int, new *LoadBalancerVirtualIPAddress) {
	o.Update(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip, Port: types.StringNumber(port)}, new)
}

// Delete サーバの削除
func (o *LoadBalancerVirtualIPAddresses) Delete(vip *LoadBalancerVirtualIPAddress) {
	var res []*LoadBalancerVirtualIPAddress
	for _, v := range *o {
		if !(v.VirtualIPAddress == vip.VirtualIPAddress && v.Port == vip.Port) {
			res = append(res, v)
		}
	}
	*o = res
}

// DeleteAt サーバの削除
func (o *LoadBalancerVirtualIPAddresses) DeleteAt(vip string, port int) {
	o.Delete(&LoadBalancerVirtualIPAddress{VirtualIPAddress: vip, Port: types.StringNumber(port)})
}
