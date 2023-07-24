// Copyright 2022-2023 The sacloud/packages-go Authors
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

// Package size さくらのクラウドでのサイズ(MiB/GiB)
//
// MBを起点としてGiBへの変換などを行う
package size

const (
	// MiB 1024KiB
	MiB = 1
	// GiB 1024MiB
	GiB = 1024 * MiB
	// TiB 1024GiB
	TiB = 1024 * GiB
	// PiB 1024TiB
	PiB = 1024 * TiB
)

// GiBToMiB GiBからMiB
func GiBToMiB(sizeGiB int) int {
	return convertUnit(sizeGiB, GiB, MiB)
}

// MiBToGiB MiBからGiB
func MiBToGiB(sizeMiB int) int {
	return convertUnit(sizeMiB, MiB, GiB)
}

func convertUnit(size int, sourceUnit int64, desiredUnit int64) int {
	return int(int64(size) * sourceUnit / desiredUnit)
}
