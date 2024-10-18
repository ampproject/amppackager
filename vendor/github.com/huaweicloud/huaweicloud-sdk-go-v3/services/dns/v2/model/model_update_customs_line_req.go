package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCustomsLineReq struct {

	// 解析线路名称。
	Name *string `json:"name,omitempty"`

	// P地址段。  以“-”分隔，小IP地址在前，大IP地址在后。IP段之间不能有交叉。当只有一个IP时，填写IP1-IP1。 目前只支持IPV4。
	IpSegments *[]string `json:"ip_segments,omitempty"`

	// 自定义线路的描述信息。长度不超过255个字符。
	Description *string `json:"description,omitempty"`
}

func (o UpdateCustomsLineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomsLineReq struct{}"
	}

	return strings.Join([]string{"UpdateCustomsLineReq", string(data)}, " ")
}
