package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AliasTarget 别名记录。
type AliasTarget struct {

	// 资源服务类型，支持别名记录的服务。取值：  cloudsite：云速建站 waf：Web应用防火墙
	ResourceType *string `json:"resource_type,omitempty"`

	// 对应服务下的域名，由各服务提供。
	ResourceDomainName *string `json:"resource_domain_name,omitempty"`
}

func (o AliasTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AliasTarget struct{}"
	}

	return strings.Join([]string{"AliasTarget", string(data)}, " ")
}
