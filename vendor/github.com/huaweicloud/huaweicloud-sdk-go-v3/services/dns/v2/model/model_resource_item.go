package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceItem struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源详情。资源对象，用于扩展，默认为空。
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]Tag `json:"tags,omitempty"`

	// 资源名称，没有默认为空字符串
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ResourceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceItem struct{}"
	}

	return strings.Join([]string{"ResourceItem", string(data)}, " ")
}
