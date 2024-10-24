package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRouterRequest Request Object
type AssociateRouterRequest struct {

	// 关联VPC的Zone ID。
	ZoneId string `json:"zone_id"`

	Body *AssociateRouterRequestBody `json:"body,omitempty"`
}

func (o AssociateRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouterRequest struct{}"
	}

	return strings.Join([]string{"AssociateRouterRequest", string(data)}, " ")
}
