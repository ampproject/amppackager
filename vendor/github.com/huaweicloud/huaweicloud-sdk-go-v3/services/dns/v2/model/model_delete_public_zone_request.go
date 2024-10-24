package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublicZoneRequest Request Object
type DeletePublicZoneRequest struct {

	// 待删除zone的ID
	ZoneId string `json:"zone_id"`
}

func (o DeletePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicZoneRequest", string(data)}, " ")
}
