package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRecordSetReq struct {

	// 域名，后缀需以zone name结束且为FQDN（即以“.”号结束的完整主机名）。
	Name *string `json:"name,omitempty"`

	// 可选配置，对域名的描述。
	Description *string `json:"description,omitempty"`

	// Record Set的类型。
	Type *string `json:"type,omitempty"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 解析记录的值。不同类型解析记录对应的值的规则不同。
	Records *[]string `json:"records,omitempty"`
}

func (o UpdateRecordSetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetReq struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetReq", string(data)}, " ")
}
