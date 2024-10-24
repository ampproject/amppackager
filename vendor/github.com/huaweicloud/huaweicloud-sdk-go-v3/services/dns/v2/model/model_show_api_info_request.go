package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiInfoRequest Request Object
type ShowApiInfoRequest struct {

	// 待查询版本号。以v开头，如v2。
	Version string `json:"version"`
}

func (o ShowApiInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowApiInfoRequest", string(data)}, " ")
}
