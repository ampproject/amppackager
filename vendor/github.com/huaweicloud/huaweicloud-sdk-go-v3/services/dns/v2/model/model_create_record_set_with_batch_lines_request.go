package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordSetWithBatchLinesRequest Request Object
type CreateRecordSetWithBatchLinesRequest struct {

	// 所属Zone的ID。
	ZoneId string `json:"zone_id"`

	Body *CreateRSetBatchLinesReq `json:"body,omitempty"`
}

func (o CreateRecordSetWithBatchLinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithBatchLinesRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithBatchLinesRequest", string(data)}, " ")
}
