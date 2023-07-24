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
 MigratedMB - MigratedGB
************************************************/

// MigratedMB is accessor interface of MigratedMB field
type MigratedMB interface {
	GetMigratedMB() int
	SetMigratedMB(size int)
}

// GetMigratedGB returns GB
func GetMigratedGB(target MigratedMB) int {
	sizeMB := target.GetMigratedMB()
	if sizeMB == 0 {
		return 0
	}
	return size.MiBToGiB(sizeMB)
}

// SetMigratedGB sets MigratedMB from GB
func SetMigratedGB(target MigratedMB, sizeGB int) {
	target.SetMigratedMB(size.GiBToMiB(sizeGB))
}
