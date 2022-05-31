package core

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

type LogStatus string

const (
	LogStatusStart   LogStatus = "start"
	LogStatusSuccess LogStatus = "success"
	LogStatusFailure LogStatus = "failure"
	LogStatusRetry   LogStatus = "retry"
)

func (c LogStatus) String() string {
	return string(c)
}

// +k8s:deepcopy-gen=false
type KeywordsLogStatus []LogStatus

type Log struct {
	Time      types.Time `read:"time"`
	LogType   string     `read:"log_type"`
	Operator  string     `read:"operator"`
	Operation string     `read:"operation"`
	Target    string     `read:"target"`
	RequestID string     `read:"request_id"`
	Status    LogStatus  `read:"status"`
}

// +k8s:deepcopy-gen=false
type SearchLogsOffset int32

func (s SearchLogsOffset) Validate() bool {
	if s < 0 || s > 9900 {
		return false
	}
	return true
}

// +k8s:deepcopy-gen=false
type SearchLogsLimit int32

func (s SearchLogsLimit) Validate() bool {
	if s < 1 || s > 100 {
		return false
	}
	return true
}

var _ api.SearchParams = &LogListSearchKeywords{}

// +k8s:deepcopy-gen=false
type LogListSearchKeywords struct {
	api.CommonSearchParams
	FullText  api.KeywordsString `url:"_keywords_full_text[],omitempty"`
	LogType   api.KeywordsString `url:"_keywords_log_type[],omitempty"`
	Operator  api.KeywordsString `url:"_keywords_operator[],omitempty"`
	Operation api.KeywordsString `url:"_keywords_operation[],omitempty"`
	Target    api.KeywordsString `url:"_keywords_target[],omitempty"`
	Detail    api.KeywordsString `url:"_keywords_detail[],omitempty"`
	RequestID api.KeywordsString `url:"_keywords_request_id[],omitempty"`
	Status    KeywordsLogStatus  `url:"_keywords_status[],omitempty"`
}

func (s *LogListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }
