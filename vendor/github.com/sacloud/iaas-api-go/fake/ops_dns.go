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

// Find is fake implementation
func (o *DNSOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.DNSFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.DNS
	for _, res := range results {
		dest := &iaas.DNS{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.DNSFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		DNS:   values,
	}, nil
}

// Create is fake implementation
func (o *DNSOp) Create(ctx context.Context, param *iaas.DNSCreateRequest) (*iaas.DNS, error) {
	result := &iaas.DNS{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"
	result.DNSZone = param.Name
	result.DNSNameServers = []string{"ns1.gslb4.sakura.ne.jp", "ns2.gslb4.sakura.ne.jp"}

	putDNS(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *DNSOp) Read(ctx context.Context, id types.ID) (*iaas.DNS, error) {
	value := getDNSByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.DNS{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *DNSOp) Update(ctx context.Context, id types.ID, param *iaas.DNSUpdateRequest) (*iaas.DNS, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putDNS(iaas.APIDefaultZone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *DNSOp) UpdateSettings(ctx context.Context, id types.ID, param *iaas.DNSUpdateSettingsRequest) (*iaas.DNS, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putDNS(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *DNSOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}
