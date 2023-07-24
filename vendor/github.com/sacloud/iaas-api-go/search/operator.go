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

package search

// ComparisonOperator フィルター比較演算子
type ComparisonOperator string

const (
	// OpEqual =
	OpEqual = ComparisonOperator("")
	// OpGreaterThan >
	OpGreaterThan = ComparisonOperator(">")
	// OpGreaterEqual >=
	OpGreaterEqual = ComparisonOperator(">=")
	// OpLessThan <
	OpLessThan = ComparisonOperator("<")
	// OpLessEqual <=
	OpLessEqual = ComparisonOperator("<=")
)

// LogicalOperator フィルター論理演算子
type LogicalOperator int

const (
	// OpAnd AND
	OpAnd LogicalOperator = iota
	// OpOr OR
	OpOr
)
