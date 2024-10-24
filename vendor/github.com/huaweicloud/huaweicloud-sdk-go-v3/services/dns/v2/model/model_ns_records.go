package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NsRecords struct {

	// 主机名。  当为内网名称服务器时，此值为空。
	Hostname *string `json:"hostname,omitempty"`

	// 名称服务器地址。  当为公网名称服务器时，此值为空。
	Address *string `json:"address,omitempty"`

	// 优先级。  示例：  如果priority的值为“1”，表示会第一个采用该域名服务器进行解析。
	Priority *int32 `json:"priority,omitempty"`
}

func (o NsRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NsRecords struct{}"
	}

	return strings.Join([]string{"NsRecords", string(data)}, " ")
}
