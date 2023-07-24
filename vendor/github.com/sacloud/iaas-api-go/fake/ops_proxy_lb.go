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
func (o *ProxyLBOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.ProxyLBFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.ProxyLB
	for _, res := range results {
		dest := &iaas.ProxyLB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.ProxyLBFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		ProxyLBs: values,
	}, nil
}

// Create is fake implementation
func (o *ProxyLBOp) Create(ctx context.Context, param *iaas.ProxyLBCreateRequest) (*iaas.ProxyLB, error) {
	result := &iaas.ProxyLB{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available

	vip := pool().nextSharedIP()
	vipNet := net.IPNet{IP: vip, Mask: []byte{255, 255, 255, 0}}
	result.ProxyNetworks = []string{vipNet.String()}
	if param.UseVIPFailover {
		result.FQDN = "fake.proxylb.sakura.ne.jp"
	}
	if result.SorryServer == nil {
		result.SorryServer = &iaas.ProxyLBSorryServer{}
	}
	if result.Timeout == nil {
		result.Timeout = &iaas.ProxyLBTimeout{}
	}
	if result.Timeout.InactiveSec == 0 {
		result.Timeout.InactiveSec = 10
	}

	status := &iaas.ProxyLBHealth{
		ActiveConn: 10,
		CPS:        10,
		CurrentVIP: vip.String(),
	}
	for _, server := range param.Servers {
		status.Servers = append(status.Servers, &iaas.LoadBalancerServerStatus{
			ActiveConn: 10,
			Status:     types.ServerInstanceStatuses.Up,
			IPAddress:  server.IPAddress,
			Port:       types.StringNumber(server.Port),
			CPS:        10,
		})
	}
	ds().Put(ResourceProxyLB+"Status", iaas.APIDefaultZone, result.ID, status)

	putProxyLB(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *ProxyLBOp) Read(ctx context.Context, id types.ID) (*iaas.ProxyLB, error) {
	value := getProxyLBByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.ProxyLB{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *ProxyLBOp) Update(ctx context.Context, id types.ID, param *iaas.ProxyLBUpdateRequest) (*iaas.ProxyLB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	if value.SorryServer == nil {
		value.SorryServer = &iaas.ProxyLBSorryServer{}
	}
	if value.Timeout == nil {
		value.Timeout = &iaas.ProxyLBTimeout{}
	}
	if value.Timeout.InactiveSec == 0 {
		value.Timeout.InactiveSec = 10
	}
	putProxyLB(iaas.APIDefaultZone, value)

	status := ds().Get(ResourceProxyLB+"Status", iaas.APIDefaultZone, id).(*iaas.ProxyLBHealth)
	status.Servers = []*iaas.LoadBalancerServerStatus{}
	for _, server := range param.Servers {
		status.Servers = append(status.Servers, &iaas.LoadBalancerServerStatus{
			ActiveConn: 10,
			Status:     types.ServerInstanceStatuses.Up,
			IPAddress:  server.IPAddress,
			Port:       types.StringNumber(server.Port),
			CPS:        10,
		})
	}
	ds().Put(ResourceProxyLB+"Status", iaas.APIDefaultZone, id, status)

	return value, nil
}

// UpdateSettings is fake implementation
func (o *ProxyLBOp) UpdateSettings(ctx context.Context, id types.ID, param *iaas.ProxyLBUpdateSettingsRequest) (*iaas.ProxyLB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	if value.SorryServer == nil {
		value.SorryServer = &iaas.ProxyLBSorryServer{}
	}
	if value.Timeout == nil {
		value.Timeout = &iaas.ProxyLBTimeout{}
	}
	if value.Timeout.InactiveSec == 0 {
		value.Timeout.InactiveSec = 10
	}
	putProxyLB(iaas.APIDefaultZone, value)

	status := ds().Get(ResourceProxyLB+"Status", iaas.APIDefaultZone, id).(*iaas.ProxyLBHealth)
	status.Servers = []*iaas.LoadBalancerServerStatus{}
	for _, server := range param.Servers {
		status.Servers = append(status.Servers, &iaas.LoadBalancerServerStatus{
			ActiveConn: 10,
			Status:     types.ServerInstanceStatuses.Up,
			IPAddress:  server.IPAddress,
			Port:       types.StringNumber(server.Port),
			CPS:        10,
		})
	}
	ds().Put(ResourceProxyLB+"Status", iaas.APIDefaultZone, id, status)

	return value, nil
}

// Delete is fake implementation
func (o *ProxyLBOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(ResourceProxyLB+"Status", iaas.APIDefaultZone, id)
	ds().Delete(ResourceProxyLB+"Certs", iaas.APIDefaultZone, id)
	ds().Delete(o.key, iaas.APIDefaultZone, id)

	return nil
}

// ChangePlan is fake implementation
func (o *ProxyLBOp) ChangePlan(ctx context.Context, id types.ID, param *iaas.ProxyLBChangePlanRequest) (*iaas.ProxyLB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	value.Plan = types.ProxyLBPlanFromServiceClass(param.ServiceClass)
	putProxyLB(iaas.APIDefaultZone, value)

	return value, err
}

// GetCertificates is fake implementation
func (o *ProxyLBOp) GetCertificates(ctx context.Context, id types.ID) (*iaas.ProxyLBCertificates, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	v := ds().Get(ResourceProxyLB+"Certs", iaas.APIDefaultZone, id)
	if v != nil {
		return v.(*iaas.ProxyLBCertificates), nil
	}

	return nil, err
}

// SetCertificates is fake implementation
func (o *ProxyLBOp) SetCertificates(ctx context.Context, id types.ID, param *iaas.ProxyLBSetCertificatesRequest) (*iaas.ProxyLBCertificates, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	cert := &iaas.ProxyLBCertificates{}
	copySameNameField(param, cert)
	cert.PrimaryCert.CertificateCommonName = "dummy-common-name.org"
	cert.PrimaryCert.CertificateEndDate = time.Now().Add(365 * 24 * time.Hour)

	ds().Put(ResourceProxyLB+"Certs", iaas.APIDefaultZone, id, cert)
	return cert, nil
}

// DeleteCertificates is fake implementation
func (o *ProxyLBOp) DeleteCertificates(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	v := ds().Get(ResourceProxyLB+"Certs", iaas.APIDefaultZone, id)
	if v != nil {
		ds().Delete(ResourceProxyLB+"Certs", iaas.APIDefaultZone, id)
	}
	return nil
}

// RenewLetsEncryptCert is fake implementation
func (o *ProxyLBOp) RenewLetsEncryptCert(ctx context.Context, id types.ID) error {
	return nil
}

// HealthStatus is fake implementation
func (o *ProxyLBOp) HealthStatus(ctx context.Context, id types.ID) (*iaas.ProxyLBHealth, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return ds().Get(ResourceProxyLB+"Status", iaas.APIDefaultZone, id).(*iaas.ProxyLBHealth), nil
}

// MonitorConnection is fake implementation
func (o *ProxyLBOp) MonitorConnection(ctx context.Context, id types.ID, condition *iaas.MonitorCondition) (*iaas.ConnectionActivity, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.ConnectionActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorConnectionValue{
			Time:              now.Add(time.Duration(i*-5) * time.Minute),
			ConnectionsPerSec: float64(random(1000)),
			ActiveConnections: float64(random(1000)),
		})
	}

	return res, nil
}
