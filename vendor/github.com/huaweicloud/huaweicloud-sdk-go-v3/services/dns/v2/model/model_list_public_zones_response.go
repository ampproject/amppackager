package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicZonesResponse Response Object
type ListPublicZonesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// 查询公网Zone的列表响应。
	Zones *[]PublicZoneResp `json:"zones,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPublicZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPublicZonesResponse", string(data)}, " ")
}
