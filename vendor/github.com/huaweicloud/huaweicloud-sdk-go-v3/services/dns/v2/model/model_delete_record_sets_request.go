package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordSetsRequest Request Object
type DeleteRecordSetsRequest struct {

	// Record Set所属的zone_id。
	ZoneId string `json:"zone_id"`

	// Record Set的id信息。
	RecordsetId string `json:"recordset_id"`
}

func (o DeleteRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetsRequest", string(data)}, " ")
}
