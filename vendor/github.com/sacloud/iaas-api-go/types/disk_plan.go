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

// DiskPlans ディスクプランID 利用可能なサイズはDiskPlanAPIで取得すること
var DiskPlans = struct {
	// SSD ssdプラン
	SSD ID
	// HDD hddプラン
	HDD ID
}{
	SSD: ID(4),
	HDD: ID(2),
}

// DiskPlanStrings ディスクプランを表す文字列
var DiskPlanStrings = []string{"ssd", "hdd"}

// DiskPlanIDMap ディスクプランと文字列のマップ
var DiskPlanIDMap = map[string]ID{
	"ssd": DiskPlans.SSD,
	"hdd": DiskPlans.HDD,
}

// DiskPlanNameMap ディスクプランIDと文字列のマップ
var DiskPlanNameMap = map[ID]string{
	DiskPlans.SSD: "ssd",
	DiskPlans.HDD: "hdd",
}
