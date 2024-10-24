package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateZoneNameServerRequest Request Object
type ShowPrivateZoneNameServerRequest struct {

	// 待查询内网zone的ID。
	ZoneId string `json:"zone_id"`
}

func (o ShowPrivateZoneNameServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerRequest", string(data)}, " ")
}
