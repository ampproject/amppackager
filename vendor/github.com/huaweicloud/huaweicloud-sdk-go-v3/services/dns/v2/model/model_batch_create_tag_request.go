package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateTagRequest Request Object
type BatchCreateTagRequest struct {

	// 资源的类型：DNS-public_zone，DNS-private_zone，DNS-public_recordset，DNS-private_recordset，DNS-ptr_record。
	ResourceType string `json:"resource_type"`

	// 资源id。
	ResourceId string `json:"resource_id"`

	Body *BatchHandTags `json:"body,omitempty"`
}

func (o BatchCreateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateTagRequest", string(data)}, " ")
}
