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

import "github.com/sacloud/iaas-api-go/types"

// Store fakeドライバーでのバックエンド(永続化)を担当するドライバーインターフェース
type Store interface {
	Init() error
	NeedInitData() bool
	Put(resourceKey, zone string, id types.ID, value interface{})
	Get(resourceKey, zone string, id types.ID) interface{}
	List(resourceKey, zone string) []interface{}
	Delete(resourceKey, zone string, id types.ID)
}
