package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Nameserver struct {

	// 主机名。
	Hostname *string `json:"hostname,omitempty"`

	// 优先级。
	Priority *int32 `json:"priority,omitempty"`
}

func (o Nameserver) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nameserver struct{}"
	}

	return strings.Join([]string{"Nameserver", string(data)}, " ")
}
