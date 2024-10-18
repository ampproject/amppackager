package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePtrReq struct {

	// PTR记录对应的域名。
	Ptrdname string `json:"ptrdname"`

	// 对PTR记录的描述。
	Description *string `json:"description,omitempty"`

	// PTR记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdatePtrReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePtrReq struct{}"
	}

	return strings.Join([]string{"UpdatePtrReq", string(data)}, " ")
}
