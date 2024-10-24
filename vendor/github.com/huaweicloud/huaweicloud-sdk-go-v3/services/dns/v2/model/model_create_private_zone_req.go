package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePrivateZoneReq struct {

	// 待创建的域名。
	Name string `json:"name"`

	// 域名的描述信息。
	Description *string `json:"description,omitempty"`

	// 域名类型。取值：private。
	ZoneType string `json:"zone_type"`

	// 管理该zone的管理员邮箱。
	Email *string `json:"email,omitempty"`

	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	Router *Router `json:"router"`

	// 内网Zone的子域名递归解析代理模式。  取值范围：  AUTHORITY：当前Zone未开启递归解析代理 RECURSIVE：当前Zone已开启递归解析代理
	ProxyPattern *string `json:"proxy_pattern,omitempty"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 域名关联的企业项目ID，长度不超过36个字符。  默认值为0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreatePrivateZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePrivateZoneReq", string(data)}, " ")
}
