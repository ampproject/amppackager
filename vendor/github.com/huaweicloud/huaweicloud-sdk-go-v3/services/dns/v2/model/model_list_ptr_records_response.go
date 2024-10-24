package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPtrRecordsResponse Response Object
type ListPtrRecordsResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	// 弹性IP的PTR记录ID列表信息。
	Floatingips    *[]ListPtrRecordsFloatingResp `json:"floatingips,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPtrRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPtrRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPtrRecordsResponse", string(data)}, " ")
}
