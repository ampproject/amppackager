package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateRecordSetWithLineReq struct {

	// RecordSet 列表。
	Recordsets []BatchUpdateRecordSet `json:"recordsets"`
}

func (o BatchUpdateRecordSetWithLineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRecordSetWithLineReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateRecordSetWithLineReq", string(data)}, " ")
}
