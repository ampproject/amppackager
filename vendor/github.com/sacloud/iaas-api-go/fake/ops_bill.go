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
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// ByContract is fake implementation
func (o *BillOp) ByContract(ctx context.Context, accountID types.ID) (*iaas.BillByContractResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, nil)
	var values []*iaas.Bill
	for _, res := range results {
		dest := &iaas.Bill{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.BillByContractResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Bills: values,
	}, nil
}

// ByContractYear is fake implementation
func (o *BillOp) ByContractYear(ctx context.Context, accountID types.ID, year int) (*iaas.BillByContractYearResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, nil)
	var values []*iaas.Bill
	for _, res := range results {
		dest := &iaas.Bill{}
		copySameNameField(res, dest)
		if dest.Date.Year() == year {
			values = append(values, dest)
		}
	}
	return &iaas.BillByContractYearResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Bills: values,
	}, nil
}

// ByContractYearMonth is fake implementation
func (o *BillOp) ByContractYearMonth(ctx context.Context, accountID types.ID, year int, month int) (*iaas.BillByContractYearMonthResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, nil)
	var values []*iaas.Bill
	for _, res := range results {
		dest := &iaas.Bill{}
		copySameNameField(res, dest)
		if dest.Date.Year() == year && int(dest.Date.Month()) == month {
			values = append(values, dest)
		}
	}
	return &iaas.BillByContractYearMonthResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Bills: values,
	}, nil
}

// Read is fake implementation
func (o *BillOp) Read(ctx context.Context, id types.ID) (*iaas.BillReadResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, nil)
	var values []*iaas.Bill
	for _, res := range results {
		dest := &iaas.Bill{}
		copySameNameField(res, dest)
		if dest.ID == id {
			values = append(values, dest)
		}
	}
	return &iaas.BillReadResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Bills: values,
	}, nil
}

// Details is fake implementation
func (o *BillOp) Details(ctx context.Context, memberCode string, id types.ID) (*iaas.BillDetailsResult, error) {
	rawResults := ds().Get(o.key+"Details", iaas.APIDefaultZone, id)
	if rawResults == nil {
		return nil, newErrorNotFound(o.key+"Details", id)
	}

	results := rawResults.(*[]*iaas.BillDetail)
	var values []*iaas.BillDetail
	for _, res := range *results {
		dest := &iaas.BillDetail{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}

	return &iaas.BillDetailsResult{
		Total:       len(*results),
		Count:       len(*results),
		From:        0,
		BillDetails: values,
	}, nil
}

// DetailsCSV is fake implementation
func (o *BillOp) DetailsCSV(ctx context.Context, memberCode string, id types.ID) (*iaas.BillDetailCSV, error) {
	rawResults := ds().Get(o.key+"Details", iaas.APIDefaultZone, id)
	if rawResults == nil {
		return nil, newErrorNotFound(o.key+"Details", id)
	}

	results := rawResults.(*[]*iaas.BillDetail)
	for _, res := range *results {
		dest := &iaas.BillDetail{}
		copySameNameField(res, dest)
	}

	return &iaas.BillDetailCSV{
		Count:       len(*results),
		ResponsedAt: time.Now(),
		Filename:    "sakura_cloud_20yy_mm.csv",
		RawBody:     "this,is,dummy,header\r\nthis,is,dummy,body",
		HeaderRow:   []string{"this", "is", "dummy", "header"},
		BodyRows: [][]string{
			{
				"this", "is", "dummy", "body",
			},
		},
	}, nil
}
