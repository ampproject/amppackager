package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicZoneNameServerResponse Response Object
type ShowPublicZoneNameServerResponse struct {

	// 查询单个公网Zone的名称服务器响应。
	Nameservers    *[]Nameserver `json:"nameservers,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowPublicZoneNameServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneNameServerResponse", string(data)}, " ")
}
