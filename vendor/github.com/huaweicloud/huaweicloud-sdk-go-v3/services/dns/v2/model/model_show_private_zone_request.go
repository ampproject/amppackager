package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateZoneRequest Request Object
type ShowPrivateZoneRequest struct {

	// 待查询zone的ID。
	ZoneId string `json:"zone_id"`
}

func (o ShowPrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneRequest", string(data)}, " ")
}
