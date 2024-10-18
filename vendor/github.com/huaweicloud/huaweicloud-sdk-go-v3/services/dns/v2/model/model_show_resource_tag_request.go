package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagRequest Request Object
type ShowResourceTagRequest struct {

	// 资源的类型：DNS-public_zone，DNS-private_zone，DNS-public_recordset，DNS-private_recordset，DNS-ptr_record。
	ResourceType string `json:"resource_type"`

	// 资源id。
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagRequest", string(data)}, " ")
}
