package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRecordSetWithLineRequest Request Object
type BatchDeleteRecordSetWithLineRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id"`

	Body *BatchDeleteRecordSetWithLineRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetWithLineRequest", string(data)}, " ")
}
