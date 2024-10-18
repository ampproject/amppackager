package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNameServersResponse Response Object
type ListNameServersResponse struct {

	// name server列表对象。
	Nameservers    *[]NameServersResp `json:"nameservers,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListNameServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNameServersResponse struct{}"
	}

	return strings.Join([]string{"ListNameServersResponse", string(data)}, " ")
}
