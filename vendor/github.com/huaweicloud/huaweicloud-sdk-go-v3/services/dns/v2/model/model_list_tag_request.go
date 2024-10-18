package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagRequest Request Object
type ListTagRequest struct {

	// 资源的类型：DNS-public_zone，DNS-private_zone，DNS-public_recordset，DNS-private_recordset，DNS-ptr_record。
	ResourceType string `json:"resource_type"`

	Body *ListTagReq `json:"body,omitempty"`
}

func (o ListTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagRequest struct{}"
	}

	return strings.Join([]string{"ListTagRequest", string(data)}, " ")
}
