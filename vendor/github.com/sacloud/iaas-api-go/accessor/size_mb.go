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

import "github.com/sacloud/packages-go/size"

/************************************************
 SizeMB - SizeGB
************************************************/

// SizeMB is accessor interface of SizeMB field
type SizeMB interface {
	GetSizeMB() int
	SetSizeMB(size int)
}

// GetSizeGB returns GB
func GetSizeGB(target SizeMB) int {
	sizeMB := target.GetSizeMB()
	if sizeMB == 0 {
		return 0
	}
	return size.MiBToGiB(sizeMB)
}

// SetSizeGB sets SizeMB from GB
func SetSizeGB(target SizeMB, sizeGB int) {
	target.SetSizeMB(size.GiBToMiB(sizeGB))
}
