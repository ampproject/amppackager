package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRecordSetRequestBody struct {

	// 域名，后缀需以zone name结束且为FQDN（即以“.”号结束的完整主机名）。
	Name string `json:"name"`

	// 可选配置，对域名的描述。  长度不超过255个字符。  默认值为空。
	Description *string `json:"description,omitempty"`

	// Record Set的类型。  公网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。  内网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、SRV。
	Type string `json:"type"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。
	Ttl *int32 `json:"ttl,omitempty"`

	// 解析记录的值。不同类型解析记录对应的值的规则不同。
	Records []string `json:"records"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateRecordSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequestBody", string(data)}, " ")
}
