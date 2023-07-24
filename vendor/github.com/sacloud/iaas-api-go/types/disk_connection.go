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

// EDiskConnection ディスク接続方法
type EDiskConnection string

// String EDiskConnectionの文字列表現
func (c EDiskConnection) String() string {
	return string(c)
}

var (
	// DiskConnections ディスク接続方法
	DiskConnections = struct {
		// VirtIO virtio
		VirtIO EDiskConnection
		// IDE ide
		IDE EDiskConnection
	}{
		VirtIO: EDiskConnection("virtio"),
		IDE:    EDiskConnection("ide"),
	}

	// DiskConnectionMap 文字列とDiskConnectionのマップ
	DiskConnectionMap = map[string]EDiskConnection{
		DiskConnections.VirtIO.String(): DiskConnections.VirtIO,
		DiskConnections.IDE.String():    DiskConnections.IDE,
	}

	// DiskConnectionStrings DiskConnectionに指定できる有効な文字列
	DiskConnectionStrings = []string{
		DiskConnections.VirtIO.String(),
		DiskConnections.IDE.String(),
	}
)
