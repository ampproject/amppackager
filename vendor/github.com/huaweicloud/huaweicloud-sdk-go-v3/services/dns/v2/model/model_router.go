package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Router struct {

	// 关联VPC的ID。
	RouterId string `json:"router_id"`

	// 关联VPC所在的region。
	RouterRegion *string `json:"router_region,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`
}

func (o Router) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Router struct{}"
	}

	return strings.Join([]string{"Router", string(data)}, " ")
}
