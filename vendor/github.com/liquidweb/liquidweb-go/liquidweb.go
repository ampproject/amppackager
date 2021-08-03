package liquidweb

import (
	"github.com/liquidweb/liquidweb-go/types"
)

// Backend is an interface for calls against Liquid Web's API.
type Backend interface {
	CallIntoInterface(string, interface{}, interface{}) error
	CallRaw(string, interface{}) ([]byte, error)
}

// ListMeta handles Liquid Web's pagination in HTTP responses.
type ListMeta struct {
	ItemCount types.FlexInt `json:"item_count,omitempty"`
	ItemTotal types.FlexInt `json:"item_total,omitempty"`
	PageNum   types.FlexInt `json:"page_num,omitempty"`
	PageSize  types.FlexInt `json:"page_size,omitempty"`
	PageTotal types.FlexInt `json:"page_total,omitempty"`
}

// PageParams support pagination parameters in parameter types.
type PageParams struct {
	PageNum  types.FlexInt `json:"page_num,omitempty"`
	PageSize types.FlexInt `json:"page_size,omitempty"`
}
