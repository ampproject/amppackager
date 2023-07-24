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

package fake

import (
	"context"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

type ipv6Addr struct {
	ID types.ID
	*iaas.IPv6Addr
}

// Find is fake implementation
func (o *IPv6AddrOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.IPv6AddrFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.IPv6Addr
	for _, res := range results {
		dest := &iaas.IPv6Addr{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.IPv6AddrFindResult{
		Total:     len(results),
		Count:     len(results),
		From:      0,
		IPv6Addrs: values,
	}, nil
}

// Create is fake implementation
func (o *IPv6AddrOp) Create(ctx context.Context, zone string, param *iaas.IPv6AddrCreateRequest) (*iaas.IPv6Addr, error) {
	result := &iaas.IPv6Addr{}
	copySameNameField(param, result)

	ds().Put(ResourceIPv6Addr, zone, pool().generateID(), &ipv6Addr{IPv6Addr: result})
	return result, nil
}

// Read is fake implementation
func (o *IPv6AddrOp) Read(ctx context.Context, zone string, ipv6addr string) (*iaas.IPv6Addr, error) {
	var value *iaas.IPv6Addr

	results := ds().List(o.key, zone)
	for _, res := range results {
		v := res.(*ipv6Addr)
		if v.IPv6Addr.IPv6Addr == ipv6addr {
			value = v.IPv6Addr
			break
		}
	}

	if value == nil {
		return nil, newErrorNotFound(o.key, ipv6addr)
	}
	return value, nil
}

// Update is fake implementation
func (o *IPv6AddrOp) Update(ctx context.Context, zone string, ipv6addr string, param *iaas.IPv6AddrUpdateRequest) (*iaas.IPv6Addr, error) {
	found := false
	results := ds().List(o.key, zone)
	var value *iaas.IPv6Addr
	for _, res := range results {
		v := res.(*ipv6Addr)
		if v.IPv6Addr.IPv6Addr == ipv6addr {
			copySameNameField(param, v.IPv6Addr)
			found = true
			ds().Put(o.key, zone, v.ID, v)
			value = v.IPv6Addr
		}
	}

	if !found {
		return nil, newErrorNotFound(o.key, ipv6addr)
	}

	return value, nil
}

// Delete is fake implementation
func (o *IPv6AddrOp) Delete(ctx context.Context, zone string, ipv6addr string) error {
	found := false
	results := ds().List(o.key, zone)
	for _, res := range results {
		v := res.(*ipv6Addr)
		if v.IPv6Addr.IPv6Addr == ipv6addr {
			found = true
			ds().Delete(o.key, zone, v.ID)
		}
	}

	if !found {
		return newErrorNotFound(o.key, ipv6addr)
	}

	return nil
}
