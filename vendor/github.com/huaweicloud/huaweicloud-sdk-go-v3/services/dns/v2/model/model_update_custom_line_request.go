package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomLineRequest Request Object
type UpdateCustomLineRequest struct {

	// 解析线路ID。
	LineId string `json:"line_id"`

	Body *UpdateCustomsLineReq `json:"body,omitempty"`
}

func (o UpdateCustomLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomLineRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomLineRequest", string(data)}, " ")
}
