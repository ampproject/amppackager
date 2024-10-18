package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestorePtrRecordRequest Request Object
type RestorePtrRecordRequest struct {

	// 域名所属的区域。
	Region string `json:"region"`

	// 弹性公网IP（EIP）的ID。
	FloatingipId string `json:"floatingip_id"`

	Body *RestorePtrReq `json:"body,omitempty"`
}

func (o RestorePtrRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePtrRecordRequest struct{}"
	}

	return strings.Join([]string{"RestorePtrRecordRequest", string(data)}, " ")
}
