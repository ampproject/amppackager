package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicZoneNameServerRequest Request Object
type ShowPublicZoneNameServerRequest struct {

	// 待查询zone的ID。  可以通过查询公网Zone列表获取。
	ZoneId string `json:"zone_id"`
}

func (o ShowPublicZoneNameServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneNameServerRequest", string(data)}, " ")
}
