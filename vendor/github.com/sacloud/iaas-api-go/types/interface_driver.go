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

// EInterfaceDriver インターフェースドライバ
type EInterfaceDriver string

func (d EInterfaceDriver) String() string {
	return string(d)
}

var (
	// InterfaceDrivers インターフェースドライバ
	InterfaceDrivers = struct {
		VirtIO EInterfaceDriver // virtio
		E1000  EInterfaceDriver // e1000
	}{
		VirtIO: EInterfaceDriver("virtio"),
		E1000:  EInterfaceDriver("e1000"),
	}

	// InterfaceDriverMap インターフェースドライバと文字列表現のマップ
	InterfaceDriverMap = map[string]EInterfaceDriver{
		InterfaceDrivers.VirtIO.String(): InterfaceDrivers.VirtIO,
		InterfaceDrivers.E1000.String():  InterfaceDrivers.E1000,
	}

	// InterfaceDriverValues インターフェースドライバが取りうる有効値
	InterfaceDriverValues = []string{
		InterfaceDrivers.VirtIO.String(),
		InterfaceDrivers.E1000.String(),
	}
)

// InterfaceDriverStrings インターフェースドライバを表す文字列
var InterfaceDriverStrings = []string{"virtio", "e1000"}
