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

// NFSPlans NFSプラン
//
// Note: NFS作成時のPlanIDはこの値+サイズでNFSプランを検索、そのIDを指定すること
// NFSプランの検索はutils/nfsのfunc FindPlan(plan types.ID, size int64)を利用する
var NFSPlans = struct {
	// HDD hddプラン
	HDD ID
	// SSD ssdプラン
	SSD ID
}{
	HDD: ID(1),
	SSD: ID(2),
}

// NFSPlanStrings NFSプランを表す文字列
var NFSPlanStrings = []string{"hdd", "ssd"}

// NFSPlanIDMap 文字列とNFSプランIDのマップ
var NFSPlanIDMap = map[string]ID{
	"hdd": NFSPlans.HDD,
	"ssd": NFSPlans.SSD,
}

// NFSPlanNameMap NFSプランIDと名前のマップ
var NFSPlanNameMap = map[ID]string{
	NFSPlans.HDD: "hdd",
	NFSPlans.SSD: "ssd",
}
