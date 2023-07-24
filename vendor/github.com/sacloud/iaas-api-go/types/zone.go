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

const (
	// ZoneTk1aID 東京第1ゾーン
	ZoneTk1aID = ID(21001)
	// ZoneTk1bID 東京第2ゾーン
	ZoneTk1bID = ID(21002)
	// ZoneIs1aID 石狩第1ゾーン
	ZoneIs1aID = ID(31001)
	// ZoneIs1bID 石狩第1ゾーン
	ZoneIs1bID = ID(31002)
	// ZoneTk1vID サンドボックスゾーン
	ZoneTk1vID = ID(29001)
)

// ZoneNames 利用できるゾーンの一覧
var ZoneNames = []string{"tk1a", "tk1b", "is1a", "is1b", "tk1v"}

// ZoneIDs ゾーンIDと名称のマップ
var ZoneIDs = map[string]ID{
	"tk1a": ZoneTk1aID,
	"tk1b": ZoneTk1bID,
	"is1a": ZoneIs1aID,
	"is1b": ZoneIs1bID,
	"tk1v": ZoneTk1vID,
}
