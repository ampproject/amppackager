package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecordSetsRequest Request Object
type UpdateRecordSetsRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id"`

	// 待查询recordset的ID信息。
	RecordsetId string `json:"recordset_id"`

	Body *UpdateRecordSetsReq `json:"body,omitempty"`
}

func (o UpdateRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetsRequest", string(data)}, " ")
}
