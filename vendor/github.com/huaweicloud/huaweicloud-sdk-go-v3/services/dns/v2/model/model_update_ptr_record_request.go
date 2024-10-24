package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePtrRecordRequest Request Object
type UpdatePtrRecordRequest struct {

	// 域名所属的区域。
	Region string `json:"region"`

	// 弹性公网IP（EIP）的ID。
	FloatingipId string `json:"floatingip_id"`

	Body *UpdatePtrReq `json:"body,omitempty"`
}

func (o UpdatePtrRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePtrRecordRequest struct{}"
	}

	return strings.Join([]string{"UpdatePtrRecordRequest", string(data)}, " ")
}
