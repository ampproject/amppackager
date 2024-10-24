package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLineGroupRequest Request Object
type ShowLineGroupRequest struct {

	// 待查询的线路分组ID。
	LinegroupId string `json:"linegroup_id"`
}

func (o ShowLineGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLineGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowLineGroupRequest", string(data)}, " ")
}
