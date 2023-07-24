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

package accessor

import "github.com/sacloud/iaas-api-go/types"

/************************************************
 ID - StringID
************************************************/

// ID is accessor interface of ID field
type ID interface {
	GetID() types.ID
	SetID(id types.ID)
}

// GetStringID returns string id
func GetStringID(target ID) string {
	return target.GetID().String()
}

// SetStringID sets id from string
func SetStringID(target ID, id string) {
	target.SetID(types.StringID(id))
}

// GetInt64ID returns int64 id
func GetInt64ID(target ID) int64 {
	return target.GetID().Int64()
}

// SetInt64ID sets id from int64
func SetInt64ID(target ID, id int64) {
	target.SetID(types.ID(id))
}
