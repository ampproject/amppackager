package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLineGroupsReq struct {

	// 线路分组名称。 不能与自定义线路名称、预制线路名称重复。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name string `json:"name"`

	// 线路分组的描述信息。 长度不超过255个字符。默认值为空。
	Description *string `json:"description,omitempty"`

	// 线路分组包含的线路列表。最少为2个线路。 解析线路ID。
	Lines []string `json:"lines"`
}

func (o CreateLineGroupsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLineGroupsReq struct{}"
	}

	return strings.Join([]string{"CreateLineGroupsReq", string(data)}, " ")
}
