package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCustomLines struct {

	// 解析线路名称。  长度限制为1-80个字符，只允许包含中文、字母、数字、'-'、'_'、'.'字符。  租户内，解析线路名称是唯一的。
	Name string `json:"name"`

	// IP地址段。  以“-”分隔，小IP地址在前，大IP地址在后。IP段之间不能有交叉。当只有一个IP时，填写IP1-IP1。 目前只支持IPV4。  最多支持50个。
	IpSegments []string `json:"ip_segments"`

	// 自定义线路的描述信息。长度不超过255个字符。  默认值为空。
	Description *string `json:"description,omitempty"`
}

func (o CreateCustomLines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomLines struct{}"
	}

	return strings.Join([]string{"CreateCustomLines", string(data)}, " ")
}
