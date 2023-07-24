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
	"fmt"
	"sync"

	"github.com/sacloud/iaas-api-go/types"
)

// InMemoryStore データをメモリ上に保存するためのデータストア
type InMemoryStore struct {
	data map[string]map[string]interface{}
	mu   sync.Mutex
}

// NewInMemoryStore .
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]map[string]interface{}),
	}
}

// Init .
func (s *InMemoryStore) Init() error {
	return nil
}

// NeedInitData .
func (s *InMemoryStore) NeedInitData() bool {
	return true
}

// Put .
func (s *InMemoryStore) Put(resourceKey, zone string, id types.ID, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values == nil {
		values = map[string]interface{}{}
	}
	values[id.String()] = value
	s.data[s.key(resourceKey, zone)] = values
}

// Get .
func (s *InMemoryStore) Get(resourceKey, zone string, id types.ID) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values == nil {
		return nil
	}
	return values[id.String()]
}

// List .
func (s *InMemoryStore) List(resourceKey, zone string) []interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	var ret []interface{}
	for _, v := range values {
		ret = append(ret, v)
	}
	return ret
}

// Delete .
func (s *InMemoryStore) Delete(resourceKey, zone string, id types.ID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values != nil {
		delete(values, id.String())
	}
}

func (s *InMemoryStore) key(resourceKey, zone string) string {
	return fmt.Sprintf("%s/%s", resourceKey, zone)
}

func (s *InMemoryStore) values(resourceKey, zone string) map[string]interface{} {
	return s.data[s.key(resourceKey, zone)]
}
