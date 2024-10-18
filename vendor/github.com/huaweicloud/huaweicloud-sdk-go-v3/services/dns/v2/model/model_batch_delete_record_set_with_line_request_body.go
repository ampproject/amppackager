package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRecordSetWithLineRequestBody struct {

	// Record Set ID列表。最多支持100个。
	RecordsetIds []string `json:"recordset_ids"`
}

func (o BatchDeleteRecordSetWithLineRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetWithLineRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetWithLineRequestBody", string(data)}, " ")
}
