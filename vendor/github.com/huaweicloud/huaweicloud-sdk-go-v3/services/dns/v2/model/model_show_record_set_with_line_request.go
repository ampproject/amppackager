package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordSetWithLineRequest Request Object
type ShowRecordSetWithLineRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id"`

	// 待查询recordset的ID信息。
	RecordsetId string `json:"recordset_id"`
}

func (o ShowRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetWithLineRequest", string(data)}, " ")
}
