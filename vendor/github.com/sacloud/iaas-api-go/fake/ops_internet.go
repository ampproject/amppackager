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
	"net"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *InternetOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.InternetFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.Internet
	for _, res := range results {
		dest := &iaas.Internet{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.InternetFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Internet: values,
	}, nil
}

// Create is fake implementation
func (o *InternetOp) Create(ctx context.Context, zone string, param *iaas.InternetCreateRequest) (*iaas.Internet, error) {
	if param.NetworkMaskLen == 0 {
		param.NetworkMaskLen = 28
	}
	if param.BandWidthMbps == 0 {
		param.BandWidthMbps = 100
	}

	result := &iaas.Internet{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	// assign global address
	subnet := pool().nextSubnet(result.NetworkMaskLen)

	// create switch
	swOp := NewSwitchOp()
	sw, err := swOp.Create(ctx, zone, &iaas.SwitchCreateRequest{
		Name:           result.Name,
		NetworkMaskLen: subnet.networkMaskLen,
		DefaultRoute:   subnet.defaultRoute,
	})
	if err != nil {
		return nil, err
	}

	sSubnet := &iaas.SwitchSubnet{
		ID:                   pool().generateID(),
		DefaultRoute:         subnet.defaultRoute,
		NetworkAddress:       subnet.networkAddress,
		NetworkMaskLen:       subnet.networkMaskLen,
		Internet:             result,
		AssignedIPAddressMax: subnet.addresses[len(subnet.addresses)-1],
		AssignedIPAddressMin: subnet.addresses[0],
	}
	sw.Subnets = append(sw.Subnets, sSubnet)

	// for Internet.Switch
	switchInfo := &iaas.SwitchInfo{}
	copySameNameField(sw, switchInfo)

	iSubnet := &iaas.InternetSubnet{
		ID:             sSubnet.ID,
		DefaultRoute:   sSubnet.DefaultRoute,
		NetworkAddress: sSubnet.NetworkAddress,
		NetworkMaskLen: sSubnet.NetworkMaskLen,
	}
	switchInfo.Subnets = []*iaas.InternetSubnet{iSubnet}
	result.Switch = switchInfo

	// for raw subnet
	rSubnet := &iaas.Subnet{}
	copySameNameField(sSubnet, rSubnet)
	for _, ip := range subnet.addresses {
		rSubnet.IPAddresses = append(rSubnet.IPAddresses, &iaas.SubnetIPAddress{IPAddress: ip})
	}
	rSubnet.SwitchID = sw.ID
	rSubnet.InternetID = result.ID

	putSwitch(zone, sw)
	putInternet(zone, result)
	putSubnet(zone, rSubnet)
	return result, nil
}

// Read is fake implementation
func (o *InternetOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.Internet, error) {
	value := getInternetByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.Internet{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *InternetOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.InternetUpdateRequest) (*iaas.Internet, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)

	putInternet(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *InternetOp) Delete(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	// check subnets/ipv6nets
	if len(value.Switch.Subnets) > 1 {
		return newErrorBadRequest(ResourceInternet, value.ID, "Internet resource still have Subnets")
	}
	if len(value.Switch.IPv6Nets) > 1 {
		return newErrorBadRequest(ResourceInternet, value.ID, "Internet resource still have Subnets")
	}

	swOp := NewSwitchOp()
	if err := swOp.Delete(ctx, zone, value.Switch.ID); err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}

// UpdateBandWidth is fake implementation
func (o *InternetOp) UpdateBandWidth(ctx context.Context, zone string, id types.ID, param *iaas.InternetUpdateBandWidthRequest) (*iaas.Internet, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	value.BandWidthMbps = param.BandWidthMbps
	putInternet(zone, value)
	return value, nil
}

// AddSubnet is fake implementation
func (o *InternetOp) AddSubnet(ctx context.Context, zone string, id types.ID, param *iaas.InternetAddSubnetRequest) (*iaas.InternetSubnetOperationResult, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	// assign global address
	subnet := pool().nextSubnetFull(param.NetworkMaskLen, param.NextHop)

	// create switch
	swOp := NewSwitchOp()
	sw, err := swOp.Read(ctx, zone, value.Switch.ID)
	if err != nil {
		return nil, err
	}

	sSubnet := &iaas.SwitchSubnet{
		ID:                   pool().generateID(),
		NetworkAddress:       subnet.networkAddress,
		NetworkMaskLen:       subnet.networkMaskLen,
		NextHop:              param.NextHop,
		StaticRoute:          param.NextHop,
		Internet:             value,
		AssignedIPAddressMax: subnet.addresses[len(subnet.addresses)-1],
		AssignedIPAddressMin: subnet.addresses[0],
	}
	sw.Subnets = append(sw.Subnets, sSubnet)

	// for Internet.Switch
	iSubnet := &iaas.InternetSubnet{
		ID:             sSubnet.ID,
		DefaultRoute:   sSubnet.DefaultRoute,
		NetworkAddress: sSubnet.NetworkAddress,
		NetworkMaskLen: sSubnet.NetworkMaskLen,
		NextHop:        param.NextHop,
		StaticRoute:    param.NextHop,
	}
	value.Switch.Subnets = append(value.Switch.Subnets, iSubnet)

	// for raw subnet
	rSubnet := &iaas.Subnet{}
	copySameNameField(sSubnet, rSubnet)
	for _, ip := range subnet.addresses {
		rSubnet.IPAddresses = append(rSubnet.IPAddresses, &iaas.SubnetIPAddress{IPAddress: ip})
	}
	rSubnet.SwitchID = sw.ID
	rSubnet.InternetID = value.ID

	putSwitch(zone, sw)
	putInternet(zone, value)
	putSubnet(zone, rSubnet)

	return &iaas.InternetSubnetOperationResult{
		ID:             sSubnet.ID,
		NextHop:        param.NextHop,
		StaticRoute:    param.NextHop,
		NetworkMaskLen: sSubnet.NetworkMaskLen,
		NetworkAddress: sSubnet.NetworkAddress,
		IPAddresses:    subnet.addresses,
	}, nil
}

// UpdateSubnet is fake implementation
func (o *InternetOp) UpdateSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID, param *iaas.InternetUpdateSubnetRequest) (*iaas.InternetSubnetOperationResult, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	// create switch
	swOp := NewSwitchOp()
	sw, err := swOp.Read(ctx, zone, value.Switch.ID)
	if err != nil {
		return nil, err
	}

	rSubnet, err := NewSubnetOp().Read(ctx, zone, subnetID)
	if err != nil {
		return nil, err
	}

	rSubnet.NextHop = param.NextHop

	var nwMaskLen int
	var nwAddress, minAddr, maxAddr string
	var addresses []string

	for _, subnet := range sw.Subnets {
		if subnet.ID == subnetID {
			subnet.NextHop = param.NextHop
			subnet.StaticRoute = param.NextHop

			minAddr = subnet.AssignedIPAddressMin
			maxAddr = subnet.AssignedIPAddressMax
			nwMaskLen = subnet.NetworkMaskLen
			nwAddress = subnet.NetworkAddress
		}
	}

	for _, subnet := range value.Switch.Subnets {
		if subnet.ID == subnetID {
			subnet.NextHop = param.NextHop
			subnet.StaticRoute = param.NextHop
		}
	}

	baseIP := net.ParseIP(minAddr).To4()
	min := baseIP[3]
	max := net.ParseIP(maxAddr).To4()[3]

	var i byte
	for (min + i) <= max { // 境界含む
		ip := net.IPv4(baseIP[0], baseIP[1], baseIP[2], baseIP[3]+i)
		addresses = append(addresses, ip.String())
		i++
	}

	putSubnet(zone, rSubnet)
	putSwitch(zone, sw)
	putInternet(zone, value)
	return &iaas.InternetSubnetOperationResult{
		ID:             subnetID,
		NextHop:        param.NextHop,
		StaticRoute:    param.NextHop,
		NetworkMaskLen: nwMaskLen,
		NetworkAddress: nwAddress,
		IPAddresses:    addresses,
	}, nil
}

// DeleteSubnet is fake implementation
func (o *InternetOp) DeleteSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	// create switch
	swOp := NewSwitchOp()
	sw, err := swOp.Read(ctx, zone, value.Switch.ID)
	if err != nil {
		return err
	}

	var sSubnets []*iaas.SwitchSubnet
	for _, subnet := range sw.Subnets {
		if subnet.ID != subnetID {
			sSubnets = append(sSubnets, subnet)
		}
	}
	sw.Subnets = sSubnets

	var iSubnets []*iaas.InternetSubnet
	for _, subnet := range value.Switch.Subnets {
		if subnet.ID != subnetID {
			iSubnets = append(iSubnets, subnet)
		}
	}
	value.Switch.Subnets = iSubnets

	ds().Delete(ResourceSubnet, zone, subnetID)
	putSwitch(zone, sw)
	putInternet(zone, value)
	return nil
}

// Monitor is fake implementation
func (o *InternetOp) Monitor(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.RouterActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.RouterActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorRouterValue{
			Time: now.Add(time.Duration(i*-5) * time.Minute),
			In:   float64(random(1000)),
			Out:  float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorRouter is fake implementation
func (o *InternetOp) MonitorRouter(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.RouterActivity, error) {
	return o.Monitor(ctx, zone, id, condition)
}

// EnableIPv6 is fake implementation
func (o *InternetOp) EnableIPv6(ctx context.Context, zone string, id types.ID) (*iaas.IPv6NetInfo, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	ipv6net := &iaas.IPv6Net{
		ID:            pool().generateID(),
		IPv6Prefix:    "2001:db8:11aa:22bb::/64",
		IPv6PrefixLen: 64,
	}
	putIPv6Net(zone, ipv6net)

	ipv6netInfo := &iaas.IPv6NetInfo{}
	copySameNameField(ipv6net, ipv6netInfo)

	value.Switch.IPv6Nets = []*iaas.IPv6NetInfo{ipv6netInfo}
	putInternet(zone, value)
	return ipv6netInfo, nil
}

// DisableIPv6 is fake implementation
func (o *InternetOp) DisableIPv6(ctx context.Context, zone string, id types.ID, ipv6netID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if len(value.Switch.IPv6Nets) == 0 {
		return nil
	}

	ds().Delete(ResourceIPv6Net, zone, value.Switch.IPv6Nets[0].ID)
	value.Switch.IPv6Nets = []*iaas.IPv6NetInfo{}
	putInternet(zone, value)
	return nil
}
