package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRouterResponse Response Object
type AssociateRouterResponse struct {

	// 关联VPC的ID。
	RouterId *string `json:"router_id,omitempty"`

	// 关联VPC所在的region。
	RouterRegion *string `json:"router_region,omitempty"`

	// 资源状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouterResponse struct{}"
	}

	return strings.Join([]string{"AssociateRouterResponse", string(data)}, " ")
}
