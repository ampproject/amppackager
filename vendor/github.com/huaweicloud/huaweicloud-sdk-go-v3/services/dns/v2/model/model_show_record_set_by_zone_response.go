package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordSetByZoneResponse Response Object
type ShowRecordSetByZoneResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// recordset列表。
	Recordsets *[]ShowRecordSetByZoneResp `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowRecordSetByZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetByZoneResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordSetByZoneResponse", string(data)}, " ")
}
