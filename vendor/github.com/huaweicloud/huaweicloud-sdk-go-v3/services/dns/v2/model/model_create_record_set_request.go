package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordSetRequest Request Object
type CreateRecordSetRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id"`

	Body *CreateRecordSetRequestBody `json:"body,omitempty"`
}

func (o CreateRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequest", string(data)}, " ")
}
