package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRecordSetsReq struct {

	// 域名，后缀需以zone name结束且为FQDN（即以“.”号结束的完整主机名）。
	Name string `json:"name"`

	// 可选配置，对域名的描述。  长度不超过255个字符。  如果为空，表示维持原值。  默认值为空。
	Description *string `json:"description,omitempty"`

	// Record Set的类型。  取值范围：A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。
	Type string `json:"type"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 解析记录的值。不同类型解析记录对应的值的规则不同。
	Records *[]string `json:"records,omitempty"`

	// 解析记录的权重。  当weight不填时，表示该解析记录将保持原有设置的权重。 当weight=0，表示该解析记录为备用域名解析记录。 当weight>0，表示该解析记录为主用域名解析记录。 取值范围：0~100  默认值为空。
	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateRecordSetsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordSetsReq struct{}"
	}

	return strings.Join([]string{"UpdateRecordSetsReq", string(data)}, " ")
}
