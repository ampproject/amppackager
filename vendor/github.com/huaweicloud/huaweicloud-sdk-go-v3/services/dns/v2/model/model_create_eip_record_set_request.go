package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEipRecordSetRequest Request Object
type CreateEipRecordSetRequest struct {

	// 租户的区域信息。
	Region string `json:"region"`

	// 弹性公网IP（EIP）的ID。
	FloatingipId string `json:"floatingip_id"`

	Body *CreatePtrReq `json:"body,omitempty"`
}

func (o CreateEipRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEipRecordSetRequest struct{}"
	}

	return strings.Join([]string{"CreateEipRecordSetRequest", string(data)}, " ")
}
