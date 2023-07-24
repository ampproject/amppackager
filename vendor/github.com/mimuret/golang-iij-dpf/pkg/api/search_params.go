package api

import (
	"net/url"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

type SearchParams interface {
	GetValues() (url.Values, error)
	GetOffset() int32
	SetOffset(int32)
	GetLimit() int32
	SetLimit(int32)
}

type RowSearchParams struct {
	url.Values
}

func NewRowSearchParams(queryString string) (*RowSearchParams, error) {
	values, err := url.ParseQuery(queryString)
	if err != nil {
		return nil, err
	}
	return &RowSearchParams{values}, nil
}

func (r RowSearchParams) GetOffset() int32 {
	v := r.Values.Get("offset")
	if v == "" {
		return 0
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

func (r RowSearchParams) SetOffset(offset int32) {
	str := strconv.FormatInt(int64(offset), 10)
	r.Set("offset", str)
}

func (r RowSearchParams) GetLimit() int32 {
	v := r.Values.Get("limit")
	if v == "" {
		return 100
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 100
	}
	return int32(i)
}

func (r RowSearchParams) SetLimit(limit int32) {
	str := strconv.FormatInt(int64(limit), 10)
	r.Set("limit", str)
}

func (r *RowSearchParams) GetValues() (url.Values, error) { return r.Values, nil }

type CommonSearchParams struct {
	Type   SearchType `url:"type,omitempty"`
	Offset int32      `url:"offset,omitempty"`
	Limit  int32      `url:"limit,omitempty"`
}

func (s *CommonSearchParams) GetValues() (url.Values, error) { return query.Values(s) }

func (k *CommonSearchParams) GetType() SearchType    { return k.Type }
func (k *CommonSearchParams) SetType(t SearchType)   { k.Type = t }
func (k *CommonSearchParams) GetOffset() int32       { return k.Offset }
func (k *CommonSearchParams) SetOffset(offset int32) { k.Offset = offset }
func (k *CommonSearchParams) GetLimit() int32 {
	if k.Limit == 0 {
		return 100
	}
	return k.Limit
}
func (k *CommonSearchParams) SetLimit(limit int32) { k.Limit = limit }

// +k8s:deepcopy-gen=false
type SearchType string

const (
	SearchTypeAND SearchType = "AND"
	SearchTypeOR  SearchType = "OR"
)

func (s SearchType) Validate() bool {
	switch s {
	case SearchTypeAND, SearchTypeOR:
	default:
		return false
	}
	return true
}

// +k8s:deepcopy-gen=false
type SearchOffset int32

func (s SearchOffset) Validate() bool {
	if s < 0 || s > 10000000 {
		return false
	}
	return true
}

// +k8s:deepcopy-gen=false
type SearchLimit int32

func (s SearchLimit) Validate() bool {
	if s < 1 || s > 10000 {
		return false
	}
	return true
}

// +k8s:deepcopy-gen=false
type SearchDate time.Time

// +k8s:deepcopy-gen=false
type SearchOrder string

const (
	SearchOrderASC  SearchOrder = "ASC"
	SearchOrderDESC SearchOrder = "DESC"
)

func (s SearchOrder) Validate() bool {
	switch s {
	case SearchOrderASC, SearchOrderDESC:
	default:
		return false
	}
	return true
}

// +k8s:deepcopy-gen=false
type KeywordsString []string

func (s KeywordsString) Validate() bool {
	for _, v := range s {
		if len(v) > 255 {
			return false
		}
	}
	return true
}

// +k8s:deepcopy-gen=false
type KeywordsID []int64

func (s KeywordsID) Validate() bool {
	for _, v := range s {
		if v < 0 {
			return false
		}
	}
	return true
}

// +k8s:deepcopy-gen=false
type KeywordsBoolean []types.Boolean

func (c KeywordsBoolean) EncodeValues(key string, v *url.Values) error {
	for _, plan := range c {
		v.Add(key, strconv.Itoa(int(plan)))
	}
	return nil
}

// +k8s:deepcopy-gen=false
type KeywordsState []types.State

func (c KeywordsState) EncodeValues(key string, v *url.Values) error {
	for _, plan := range c {
		v.Add(key, strconv.Itoa(int(plan)))
	}
	return nil
}

// +k8s:deepcopy-gen=false
type KeywordsFavorite []types.Favorite

func (c KeywordsFavorite) EncodeValues(key string, v *url.Values) error {
	for _, plan := range c {
		v.Add(key, strconv.Itoa(int(plan)))
	}
	return nil
}

// +k8s:deepcopy-gen=false
type KeywordsLabels []Label

func (c KeywordsLabels) EncodeValues(key string, v *url.Values) error {
	for _, plan := range c {
		v.Add(key, plan.String())
	}
	return nil
}

type Label struct {
	Key   string
	Value string
}

func (l Label) String() string {
	return l.Key + "=" + l.Value
}
