package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPtrRecordSetRequest Request Object
type ShowPtrRecordSetRequest struct {

	// 租户的区域信息。
	Region string `json:"region"`

	// 弹性公网IP（EIP）的ID。
	FloatingipId string `json:"floatingip_id"`
}

func (o ShowPtrRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPtrRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowPtrRecordSetRequest", string(data)}, " ")
}
