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

import "sort"

// EDayOfTheWeek 曜日
type EDayOfTheWeek string

// DaysOfTheWeek 曜日
var DaysOfTheWeek = struct {
	Sunday    EDayOfTheWeek
	Monday    EDayOfTheWeek
	Tuesday   EDayOfTheWeek
	Wednesday EDayOfTheWeek
	Thursday  EDayOfTheWeek
	Friday    EDayOfTheWeek
	Saturday  EDayOfTheWeek
}{
	Sunday:    EDayOfTheWeek("sun"),
	Monday:    EDayOfTheWeek("mon"),
	Tuesday:   EDayOfTheWeek("tue"),
	Wednesday: EDayOfTheWeek("wed"),
	Thursday:  EDayOfTheWeek("thu"),
	Friday:    EDayOfTheWeek("fri"),
	Saturday:  EDayOfTheWeek("sat"),
}

func DayOfTheWeekFromString(s string) EDayOfTheWeek {
	switch s {
	case "sun":
		return DaysOfTheWeek.Sunday
	case "mon":
		return DaysOfTheWeek.Monday
	case "tue":
		return DaysOfTheWeek.Tuesday
	case "wed":
		return DaysOfTheWeek.Wednesday
	case "thu":
		return DaysOfTheWeek.Thursday
	case "fri":
		return DaysOfTheWeek.Friday
	case "sat":
		return DaysOfTheWeek.Saturday
	}
	return DaysOfTheWeek.Monday // デフォルト
}

func DayOfTheWeekFromInt(i int) EDayOfTheWeek {
	switch i {
	case 0:
		return DaysOfTheWeek.Sunday
	case 1:
		return DaysOfTheWeek.Monday
	case 2:
		return DaysOfTheWeek.Tuesday
	case 3:
		return DaysOfTheWeek.Wednesday
	case 4:
		return DaysOfTheWeek.Thursday
	case 5:
		return DaysOfTheWeek.Friday
	case 6:
		return DaysOfTheWeek.Saturday
	}

	return DaysOfTheWeek.Monday // デフォルト
}

// String Stringer実装
func (w EDayOfTheWeek) String() string {
	return string(w)
}

// Int 曜日の数値表現
func (w EDayOfTheWeek) Int() int {
	switch w.String() {
	case "sun":
		return 0
	case "mon":
		return 1
	case "tue":
		return 2
	case "wed":
		return 3
	case "thu":
		return 4
	case "fri":
		return 5
	case "sat":
		return 6
	}
	return -1
}

// SortDayOfTheWeekList バックアップ取得曜日のソート(日曜開始)
func SortDayOfTheWeekList(weekdays []EDayOfTheWeek) {
	sort.Slice(weekdays, func(i, j int) bool {
		return weekdays[i].Int() < weekdays[j].Int()
	})
}

// DaysOfTheWeekStrings 有効なバックアップ取得曜日のリスト(文字列)
var DaysOfTheWeekStrings = []string{
	DaysOfTheWeek.Sunday.String(),
	DaysOfTheWeek.Monday.String(),
	DaysOfTheWeek.Tuesday.String(),
	DaysOfTheWeek.Wednesday.String(),
	DaysOfTheWeek.Thursday.String(),
	DaysOfTheWeek.Friday.String(),
	DaysOfTheWeek.Saturday.String(),
}
