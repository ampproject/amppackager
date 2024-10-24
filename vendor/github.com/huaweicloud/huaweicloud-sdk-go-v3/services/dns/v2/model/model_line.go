package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Line struct {

	// 解析线路ID。
	LineId *string `json:"line_id,omitempty"`

	// 解析线路名称。
	Name *string `json:"name,omitempty"`

	// IP地址段。
	IpSegments *[]string `json:"ip_segments,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 自定义线路的描述信息。
	Description *string `json:"description,omitempty"`
}

func (o Line) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Line struct{}"
	}

	return strings.Join([]string{"Line", string(data)}, " ")
}
