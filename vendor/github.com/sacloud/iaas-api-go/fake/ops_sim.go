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
	"errors"
	"strings"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *SIMOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.SIMFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.SIM
	for _, res := range results {
		dest := &iaas.SIM{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.SIMFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		SIMs:  values,
	}, nil
}

// Create is fake implementation
func (o *SIMOp) Create(ctx context.Context, param *iaas.SIMCreateRequest) (*iaas.SIM, error) {
	result := &iaas.SIM{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillModifiedAt)

	result.Class = "sim"
	result.Availability = types.Availabilities.Available
	result.Info = &iaas.SIMInfo{
		ICCID:          param.ICCID,
		RegisteredDate: time.Now(),
		Registered:     true,
		ResourceID:     result.ID.String(),
	}

	putSIM(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *SIMOp) Read(ctx context.Context, id types.ID) (*iaas.SIM, error) {
	value := getSIMByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.SIM{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SIMOp) Update(ctx context.Context, id types.ID, param *iaas.SIMUpdateRequest) (*iaas.SIM, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	putSIM(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *SIMOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}

// Activate is fake implementation
func (o *SIMOp) Activate(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.Activated {
		return errors.New("SIM[%d] is already activated")
	}
	value.Info.Activated = true
	value.Info.ActivatedDate = time.Now()
	value.Info.DeactivatedDate = time.Time{}
	putSIM(iaas.APIDefaultZone, value)
	return nil
}

// Deactivate is fake implementation
func (o *SIMOp) Deactivate(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if !value.Info.Activated {
		return errors.New("SIM[%d] is already deactivated")
	}
	value.Info.Activated = false
	value.Info.ActivatedDate = time.Time{}
	value.Info.DeactivatedDate = time.Now()
	putSIM(iaas.APIDefaultZone, value)
	return nil
}

// AssignIP is fake implementation
func (o *SIMOp) AssignIP(ctx context.Context, id types.ID, param *iaas.SIMAssignIPRequest) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.IP != "" {
		return errors.New("SIM[%d] already has IPAddress")
	}
	value.Info.IP = param.IP
	putSIM(iaas.APIDefaultZone, value)

	return nil
}

// ClearIP is fake implementation
func (o *SIMOp) ClearIP(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.IP == "" {
		return errors.New("SIM[%d] doesn't have IPAddress")
	}
	value.Info.IP = ""
	putSIM(iaas.APIDefaultZone, value)
	return nil
}

// IMEILock is fake implementation
func (o *SIMOp) IMEILock(ctx context.Context, id types.ID, param *iaas.SIMIMEILockRequest) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.IMEILock {
		return errors.New("SIM[%d] is already locked with IMEI")
	}
	value.Info.IMEILock = true
	value.Info.IMEI = param.IMEI
	putSIM(iaas.APIDefaultZone, value)
	return nil
}

// IMEIUnlock is fake implementation
func (o *SIMOp) IMEIUnlock(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if !value.Info.IMEILock {
		return errors.New("SIM[%d] is not locked with IMEI")
	}
	value.Info.IMEILock = false
	value.Info.IMEI = ""
	putSIM(iaas.APIDefaultZone, value)
	return nil
}

// Logs is fake implementation
func (o *SIMOp) Logs(ctx context.Context, id types.ID) (*iaas.SIMLogsResult, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &iaas.SIMLogsResult{
		Total: 1,
		From:  0,
		Count: 1,
		Logs: []*iaas.SIMLog{
			{
				Date:          time.Now(),
				SessionStatus: "up",
				ResourceID:    value.ID.String(),
				IMEI:          value.Info.ConnectedIMEI,
				IMSI:          strings.Join(value.Info.IMSI, " "),
			},
		},
	}, nil
}

// GetNetworkOperator is fake implementation
func (o *SIMOp) GetNetworkOperator(ctx context.Context, id types.ID) ([]*iaas.SIMNetworkOperatorConfig, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	v := ds().Get(o.key+"NetworkOperator", iaas.APIDefaultZone, id)
	if v != nil {
		var res []*iaas.SIMNetworkOperatorConfig
		configs := v.(*[]*iaas.SIMNetworkOperatorConfig)
		res = append(res, *configs...)
		return res, nil
	}

	return []*iaas.SIMNetworkOperatorConfig{}, nil
}

// SetNetworkOperator is fake implementation
func (o *SIMOp) SetNetworkOperator(ctx context.Context, id types.ID, configs []*iaas.SIMNetworkOperatorConfig) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Put(o.key+"NetworkOperator", iaas.APIDefaultZone, id, &configs)
	return nil
}

// MonitorSIM is fake implementation
func (o *SIMOp) MonitorSIM(ctx context.Context, id types.ID, condition *iaas.MonitorCondition) (*iaas.LinkActivity, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.LinkActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorLinkValue{
			Time:        now.Add(time.Duration(i*-5) * time.Minute),
			UplinkBPS:   float64(random(1000)),
			DownlinkBPS: float64(random(1000)),
		})
	}

	return res, nil
}

// Status is fake implementation
func (o *SIMOp) Status(ctx context.Context, id types.ID) (*iaas.SIMInfo, error) {
	v, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return v.Info, nil
}
