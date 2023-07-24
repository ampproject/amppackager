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

// DatabasePlans データベースプラン
var DatabasePlans = struct {
	// DB10GB 10GB
	DB10GB ID
	// DB30GB 30GB
	DB30GB ID
	// DB90GB 90GB
	DB90GB ID
	// DB240GB 240GB
	DB240GB ID
	// DB500GB 500GB
	DB500GB ID
	// DB1TB 1TB
	DB1TB ID
}{
	DB10GB:  ID(10),
	DB30GB:  ID(30),
	DB90GB:  ID(90),
	DB240GB: ID(240),
	DB500GB: ID(500),
	DB1TB:   ID(1000),
}

// SlaveDatabasePlanID マスター側のプランIDからスレーブのプランIDを算出
func SlaveDatabasePlanID(masterPlanID ID) ID {
	return ID(int64(masterPlanID) + 1)
}

// DatabasePlanIDs データベースプランのID
var DatabasePlanIDs = []ID{
	DatabasePlans.DB10GB,
	DatabasePlans.DB30GB,
	DatabasePlans.DB90GB,
	DatabasePlans.DB240GB,
	DatabasePlans.DB500GB,
	DatabasePlans.DB1TB,
}

// DatabasePlanStrings データベースプランを示す文字列
var DatabasePlanStrings = []string{"10g", "30g", "90g", "240g", "500g", "1t"}

// DatabasePlanIDMap 文字列とデータベースプランのマップ、キーはDatabasePlanStringsから参照すること
var DatabasePlanIDMap = map[string]ID{
	"10g":  DatabasePlans.DB10GB,
	"30g":  DatabasePlans.DB30GB,
	"90g":  DatabasePlans.DB90GB,
	"240g": DatabasePlans.DB240GB,
	"500g": DatabasePlans.DB500GB,
	"1t":   DatabasePlans.DB1TB,
}

// DatabasePlanNameMap プランIDと名称のマップ
var DatabasePlanNameMap = map[ID]string{
	DatabasePlans.DB10GB:  "10g",
	DatabasePlans.DB30GB:  "30g",
	DatabasePlans.DB90GB:  "90g",
	DatabasePlans.DB240GB: "240g",
	DatabasePlans.DB500GB: "500g",
	DatabasePlans.DB1TB:   "1t",
}
