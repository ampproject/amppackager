package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePublicZoneStatusRequestBody struct {

	// Zone状态。  取值范围：  ENABLE：启用解析 DISABLE：暂停解析
	Status string `json:"status"`
}

func (o UpdatePublicZoneStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneStatusRequestBody", string(data)}, " ")
}
