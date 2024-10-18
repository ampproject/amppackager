package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateRouterResponse Response Object
type DisassociateRouterResponse struct {

	// 关联VPC的ID。
	RouterId *string `json:"router_id,omitempty"`

	// 关联VPC所在的region。
	RouterRegion *string `json:"router_region,omitempty"`

	// 资源状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouterResponse struct{}"
	}

	return strings.Join([]string{"DisassociateRouterResponse", string(data)}, " ")
}
