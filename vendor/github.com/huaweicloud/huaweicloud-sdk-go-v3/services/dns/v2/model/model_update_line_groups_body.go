package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLineGroupsBody struct {

	// 线路分组名称。 不能与自定义线路名称、预制线路名称重复。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name string `json:"name"`

	// 线路分组的描述信息。长度不超过255个字符。默认值为空。
	Description *string `json:"description,omitempty"`

	// 线路列表。
	Lines []string `json:"lines"`
}

func (o UpdateLineGroupsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLineGroupsBody struct{}"
	}

	return strings.Join([]string{"UpdateLineGroupsBody", string(data)}, " ")
}
