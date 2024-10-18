package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicZoneRequest Request Object
type ShowPublicZoneRequest struct {

	// 待查询zone的ID。
	ZoneId string `json:"zone_id"`
}

func (o ShowPublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneRequest", string(data)}, " ")
}
