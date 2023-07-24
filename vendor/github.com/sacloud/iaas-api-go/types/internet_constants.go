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

// InternetBandWidths 設定可能な帯域幅の値リスト
var InternetBandWidths = []int{
	100, 250, 500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000,
	5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000,
}

// InternetNetworkMaskLengths 設定可能なネットワークマスク長の値リスト
var InternetNetworkMaskLengths = []int{26, 27, 28}
