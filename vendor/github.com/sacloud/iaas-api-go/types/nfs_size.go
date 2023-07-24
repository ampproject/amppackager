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

package types

import "fmt"

// ENFSSize NFSサイズ
type ENFSSize int

func (s ENFSSize) Int() int {
	return int(s)
}

func (s ENFSSize) Int64() int64 {
	return int64(s)
}

func (s ENFSSize) String() string {
	return fmt.Sprintf("%d", s.Int())
}

// NFSHDDSizes NFSのHDDプランで指定可能なサイズ
var NFSHDDSizes = struct {
	Size100GB ENFSSize
	Size500GB ENFSSize
	Size1TB   ENFSSize
	Size2TB   ENFSSize
	Size4TB   ENFSSize
	Size8TB   ENFSSize
	Size12TB  ENFSSize
}{
	Size100GB: ENFSSize(100),
	Size500GB: ENFSSize(500),
	Size1TB:   ENFSSize(1024 * 1),
	Size2TB:   ENFSSize(1024 * 2),
	Size4TB:   ENFSSize(1024 * 4),
	Size8TB:   ENFSSize(1024 * 8),
	Size12TB:  ENFSSize(1024 * 12),
}

// NFSSSDSizes NFSのSSDプランで指定可能なサイズ
var NFSSSDSizes = struct {
	Size20GB  ENFSSize
	Size100GB ENFSSize
	Size500GB ENFSSize
	Size1TB   ENFSSize
	Size2TB   ENFSSize
	Size4TB   ENFSSize
}{
	Size20GB:  ENFSSize(20),
	Size100GB: ENFSSize(100),
	Size500GB: ENFSSize(500),
	Size1TB:   ENFSSize(1024 * 1),
	Size2TB:   ENFSSize(1024 * 2),
	Size4TB:   ENFSSize(1024 * 4),
}

// NFSIntSizes NFSで使用可能なサイズの一覧
var NFSIntSizes = []int{
	int(NFSHDDSizes.Size100GB),
	int(NFSHDDSizes.Size500GB),
	int(NFSHDDSizes.Size1TB),
	int(NFSHDDSizes.Size2TB),
	int(NFSHDDSizes.Size4TB),
	int(NFSHDDSizes.Size8TB),
	int(NFSHDDSizes.Size12TB),
}
