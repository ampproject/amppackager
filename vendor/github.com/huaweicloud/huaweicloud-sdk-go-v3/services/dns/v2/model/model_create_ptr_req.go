package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePtrReq struct {

	// PTR记录对应的域名。
	Ptrdname string `json:"ptrdname"`

	// 对PTR记录的描述。
	Description *string `json:"description,omitempty"`

	// PTR记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。取值范围：1～2147483647
	Ttl *int32 `json:"ttl,omitempty"`

	// 反向解析关联的企业项目ID，长度不超过36个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreatePtrReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePtrReq struct{}"
	}

	return strings.Join([]string{"CreatePtrReq", string(data)}, " ")
}
