package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEipRecordSetResponse Response Object
type CreateEipRecordSetResponse struct {

	// PTR记录的ID，格式形如{region}:{floatingip_id}。
	Id *string `json:"id,omitempty"`

	// PTR记录对应的域名。
	Ptrdname *string `json:"ptrdname,omitempty"`

	// 对PTR记录的描述。
	Description *string `json:"description,omitempty"`

	// PTR记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 弹性IP的IP地址。
	Address *string `json:"address,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 对该资源的当前操作。取值范围：CREATE，UPDATE，DELETE，NONE CREATE：表示创建，UPDATE：表示更新，DELETE：表示删除，NONE：表示无操作
	Action *string `json:"action,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 反向解析关联的企业项目ID，长度不超过36个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateEipRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEipRecordSetResponse struct{}"
	}

	return strings.Join([]string{"CreateEipRecordSetResponse", string(data)}, " ")
}
