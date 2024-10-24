package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePublicZoneInfo struct {

	// 域名的描述信息。长度不超过255个字符。
	Description *string `json:"description,omitempty"`

	// 管理该zone的管理员邮箱。  如果为空，表示维持原值。  默认值为空。
	Email *string `json:"email,omitempty"`

	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o UpdatePublicZoneInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneInfo struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneInfo", string(data)}, " ")
}
