package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionItem struct {

	// 版本ID（版本号），如v2。
	Id *string `json:"id,omitempty"`

	// 版本状态，为如下3种： CURRENT：表示该版本为主推版本。 SUPPORTED：表示为老版本，但是现在还继续支持。 DEPRECATED：表示为废弃版本，存在后续删除的可能。
	Status *string `json:"status,omitempty"`

	// API的URL地址。
	Links *[]LinksItem `json:"links,omitempty"`

	// 版本发布时间。
	Updated *string `json:"updated,omitempty"`

	// 支持的最大微版本号。若该版本API不支持微版本，则为空。
	Version *string `json:"version,omitempty"`

	// 支持的最小微版本号。若该版本API不支持微版本，则为空。
	MinVersion *string `json:"min_version,omitempty"`
}

func (o VersionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionItem struct{}"
	}

	return strings.Join([]string{"VersionItem", string(data)}, " ")
}
