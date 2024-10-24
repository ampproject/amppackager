package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicZoneReq 创建公网zone请求
type CreatePublicZoneReq struct {

	// Zone名称
	Name string `json:"name"`

	// 描述
	Description *string `json:"description,omitempty"`

	// Zone类型,取值public。
	ZoneType *string `json:"zone_type,omitempty"`

	// 管理该zone的管理员邮箱
	Email *string `json:"email,omitempty"`

	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位.
	Ttl *int32 `json:"ttl,omitempty"`

	// 域名关联的企业项目ID，长度不超过36个字符.
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreatePublicZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneReq", string(data)}, " ")
}
