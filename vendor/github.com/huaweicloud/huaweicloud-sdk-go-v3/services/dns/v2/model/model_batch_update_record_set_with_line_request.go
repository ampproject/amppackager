package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateRecordSetWithLineRequest Request Object
type BatchUpdateRecordSetWithLineRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id"`

	Body *BatchUpdateRecordSetWithLineReq `json:"body,omitempty"`
}

func (o BatchUpdateRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateRecordSetWithLineRequest", string(data)}, " ")
}
