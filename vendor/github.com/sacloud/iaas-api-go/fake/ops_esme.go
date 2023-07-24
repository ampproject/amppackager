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
	"fmt"
	"math/rand"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *ESMEOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.ESMEFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.ESME
	for _, res := range results {
		dest := &iaas.ESME{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.ESMEFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		ESME:  values,
	}, nil
}

// Create is fake implementation
func (o *ESMEOp) Create(ctx context.Context, param *iaas.ESMECreateRequest) (*iaas.ESME, error) {
	result := &iaas.ESME{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)
	result.Availability = types.Availabilities.Available

	putESME(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *ESMEOp) Read(ctx context.Context, id types.ID) (*iaas.ESME, error) {
	value := getESMEByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.ESME{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *ESMEOp) Update(ctx context.Context, id types.ID, param *iaas.ESMEUpdateRequest) (*iaas.ESME, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	putESME(iaas.APIDefaultZone, value)

	return value, nil
}

// Delete is fake implementation
func (o *ESMEOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}

// randomName testutilパッケージからのコピー(循環参照を防ぐため) TODO パッケージ構造の見直し
func (o *ESMEOp) randomName(strlen int) string {
	charSetNumber := "012346789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSetNumber[rand.Intn(len(charSetNumber))] //nolint:gosec
	}
	return string(result)
}

func (o *ESMEOp) generateMessageID() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		o.randomName(8),
		o.randomName(4),
		o.randomName(4),
		o.randomName(4),
		o.randomName(12),
	)
}

func (o *ESMEOp) SendMessageWithGeneratedOTP(ctx context.Context, id types.ID, param *iaas.ESMESendMessageWithGeneratedOTPRequest) (*iaas.ESMESendMessageResult, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	result := &iaas.ESMESendMessageResult{
		MessageID: o.generateMessageID(),
		Status:    "Accepted", // Note: 現在のfakeドライバでは"Delivered"に変更する処理は未実装
		OTP:       o.randomName(6),
	}

	logs, err := o.Logs(ctx, id)
	if err != nil {
		return nil, err
	}
	logs = append(logs, &iaas.ESMELogs{
		MessageID:   result.MessageID,
		Status:      result.Status,
		OTP:         result.OTP,
		Destination: param.Destination,
		SentAt:      time.Now(),
		RetryCount:  0,
	})
	ds().Put(o.key+"Logs", iaas.APIDefaultZone, id, logs)

	return result, nil
}

func (o *ESMEOp) SendMessageWithInputtedOTP(ctx context.Context, id types.ID, param *iaas.ESMESendMessageWithInputtedOTPRequest) (*iaas.ESMESendMessageResult, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	result := &iaas.ESMESendMessageResult{
		MessageID: o.generateMessageID(),
		Status:    "Accepted", // Note: 現在のfakeドライバでは"Delivered"に変更する処理は未実装
		OTP:       param.OTP,
	}

	logs, err := o.Logs(ctx, id)
	if err != nil {
		return nil, err
	}
	logs = append(logs, &iaas.ESMELogs{
		MessageID:   result.MessageID,
		Status:      result.Status,
		OTP:         result.OTP,
		Destination: param.Destination,
		SentAt:      time.Now(),
		RetryCount:  0,
	})
	ds().Put(o.key+"Logs", iaas.APIDefaultZone, id, logs)

	return result, nil
}

func (o *ESMEOp) Logs(ctx context.Context, id types.ID) ([]*iaas.ESMELogs, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	v := ds().Get(o.key+"Logs", iaas.APIDefaultZone, id)
	if v == nil {
		return nil, nil
	}
	return v.([]*iaas.ESMELogs), nil
}
