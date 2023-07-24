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
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/accessor"
	"github.com/sacloud/iaas-api-go/search"
	"github.com/sacloud/iaas-api-go/types"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func random(max int) int {
	return rand.Intn(max) //nolint:gosec
}

func newErrorNotFound(resourceKey string, id interface{}) error {
	return iaas.NewAPIError("", nil, http.StatusNotFound, &iaas.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "404 NotFound",
		ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
		ErrorMessage: fmt.Sprintf("%s[ID:%s] is not found", resourceKey, id),
	})
}

func newErrorBadRequest(resourceKey string, id interface{}, msgs ...string) error {
	return iaas.NewAPIError("", nil, http.StatusBadRequest, &iaas.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "400 BadRequest",
		ErrorCode:    fmt.Sprintf("%d", http.StatusBadRequest),
		ErrorMessage: fmt.Sprintf("request to %s[ID:%s] is bad: %s", resourceKey, id, strings.Join(msgs, " ")),
	})
}

func newErrorConflict(resourceKey string, id interface{}, msgs ...string) error {
	return iaas.NewAPIError("", nil, http.StatusConflict, &iaas.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "409 Conflict",
		ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
		ErrorMessage: fmt.Sprintf("request to %s[ID:%s] is conflicted: %s", resourceKey, id, strings.Join(msgs, " ")),
	})
}

func newInternalServerError(resourceKey string, id interface{}, msgs ...string) error {
	return iaas.NewAPIError("", nil, http.StatusInternalServerError, &iaas.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "500 Internal Server Error",
		ErrorCode:    fmt.Sprintf("%d", http.StatusInternalServerError),
		ErrorMessage: fmt.Sprintf("request to %s[ID:%s] is failed: %s", resourceKey, id, strings.Join(msgs, " ")),
	})
}

func find(resourceKey, zone string, conditions *iaas.FindCondition) ([]interface{}, error) {
	var results []interface{}
	if conditions == nil {
		conditions = &iaas.FindCondition{}
	}

	targets := ds().List(resourceKey, zone)

FILTER_APPLY_LOOP:
	for i, target := range targets {
		// count
		if conditions.Count != 0 && len(results) >= conditions.Count {
			break
		}

		// from
		if i < conditions.From {
			continue
		}

		// filter
		for key, expression := range conditions.Filter {
			// TODO OpGreater/OpLess is not implemented
			if key.Op == search.OpEqual {
				fieldName := key.String()

				// only ID/Name/Tags
				switch fieldName {
				case "ID", "Name", "Tags.Name", "Scope", "Class":
					exp, ok := expression.(*search.EqualExpression)
					if !ok {
						exp = search.OrEqual(expression)
					}

					if fieldName == "Tags.Name" {
						var tags types.Tags
						if v, ok := target.(accessor.Tags); ok {
							tags = v.GetTags()
						}

						for _, cond := range exp.Conditions {
							condTags, ok := cond.([]string)
							if ok {
								for _, v := range condTags {
									exists := false
									for _, tag := range tags {
										if tag == v {
											exists = true
										}
									}
									if !exists {
										continue FILTER_APPLY_LOOP
									}
								}
							}
						}
					} else {
						var value interface{}
						switch fieldName {
						case "ID":
							if v, ok := target.(accessor.ID); ok {
								value = v.GetID()
							}
						case "Name":
							if v, ok := target.(accessor.Name); ok {
								value = v.GetName()
							}
						case "Scope":
							if v, ok := target.(accessor.Scope); ok {
								value = v.GetScope().String()
							}
						case "Class":
							if v, ok := target.(accessor.Class); ok {
								value = v.GetClass()
							}
						}

						switch exp.Op {
						case search.OpAnd:
							for _, v := range exp.Conditions {
								v1, ok1 := v.(string)
								v2, ok2 := value.(string)
								if !ok1 || !ok2 || !strings.Contains(v2, v1) {
									continue FILTER_APPLY_LOOP
								}
							}
						case search.OpOr:
							match := false
							for _, v := range exp.Conditions {
								if reflect.DeepEqual(value, v) {
									match = true
								}
							}
							if !match {
								continue FILTER_APPLY_LOOP
							}
						}
					}
				}
			}
		}

		results = append(results, target)
	}

	// TODO sort/filter/include/exclude is not implemented
	return results, nil
}

func copySameNameField(source interface{}, dest interface{}) {
	data, _ := json.Marshal(source)
	json.Unmarshal(data, dest) //nolint
}

func fill(target interface{}, fillFuncs ...func(interface{})) {
	for _, f := range fillFuncs {
		f(target)
	}
}

func fillID(target interface{}) {
	if v, ok := target.(accessor.ID); ok {
		id := v.GetID()
		if id.IsEmpty() {
			v.SetID(pool().generateID())
		}
	}
}

func fillAvailability(target interface{}) {
	if v, ok := target.(accessor.Availability); ok {
		value := v.GetAvailability()
		if value == types.Availabilities.Unknown {
			v.SetAvailability(types.Availabilities.Available)
		}
	}
}

func fillScope(target interface{}) {
	if v, ok := target.(accessor.Scope); ok {
		value := v.GetScope()
		if value == types.EScope("") {
			v.SetScope(types.Scopes.User)
		}
	}
}

func fillDiskPlan(target interface{}) {
	if v, ok := target.(accessor.DiskPlan); ok {
		id := v.GetDiskPlanID()
		switch id {
		case types.DiskPlans.HDD:
			v.SetDiskPlanName("標準プラン")
		case types.DiskPlans.SSD:
			v.SetDiskPlanName("SSDプラン")
		}
		v.SetDiskPlanStorageClass("iscsi9999")
	}
}

func fillCreatedAt(target interface{}) {
	if v, ok := target.(accessor.CreatedAt); ok {
		value := v.GetCreatedAt()
		if value.IsZero() {
			v.SetCreatedAt(time.Now())
		}
	}
}

func fillModifiedAt(target interface{}) {
	if v, ok := target.(accessor.ModifiedAt); ok {
		value := v.GetModifiedAt()
		if value.IsZero() {
			v.SetModifiedAt(time.Now())
		}
	}
}
