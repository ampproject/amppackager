/*
Copyright © LiquidWeb

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package apiTypes

type PaginatedList struct {
	ItemCount int64                    `json:"item_count" mapstructure:"item_count"`
	ItemTotal int64                    `json:"item_total" mapstructure:"item_total"`
	Items     []map[string]interface{} `json:"items" mapstructure:"items"`
	PageNum   int64                    `json:"page_num" mapstructure:"page_num"`
	PageSize  int64                    `json:"page_size mapstructure:"page_size"`
	PageTotal int64                    `json:"page_total" mapstructure:"page_total"`
}

type MergedPaginatedList struct {
	Items       []map[string]interface{} `json:"items" mapstructure:"items"`
	MergedPages int64                    `json:"merged_pages" mapstructure:"merged_pages"`
	PageSize    int64                    `json:"page_size" mapstructure:"page_size"`
}
