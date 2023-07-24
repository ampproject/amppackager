// Copyright 2022-2023 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iaas

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// IsNotFoundError 指定のerrorがAPI呼び出し時の404エラーであるか判定
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	if apiError, ok := err.(APIError); ok {
		return apiError.ResponseCode() == http.StatusNotFound
	}

	return false
}

// IsNoResultsError 指定のerrorがNoResultErrorであるか判定
func IsNoResultsError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*NoResultsError)
	return ok
}

// IsStillCreatingError 指定のerrorがAPI呼び出し時の409エラー、かつエラーコードがstill_creatingであるか判定
func IsStillCreatingError(err error) bool {
	if err == nil {
		return false
	}

	if apiError, ok := err.(APIError); ok {
		return apiError.ResponseCode() == http.StatusConflict && apiError.Code() == "still_creating"
	}

	return false
}

// NoResultsError APIが返した応答に処理すべきデータが含まれていない場合を示すエラー型
type NoResultsError struct {
	error
}

// NewNoResultsError NoResultErrorを返す
func NewNoResultsError() *NoResultsError {
	return &NoResultsError{error: errors.New("no results")}
}

// APIErrorResponse APIエラー型
type APIErrorResponse struct {
	IsFatal      bool   `json:"is_fatal,omitempty"`   // IsFatal
	Serial       string `json:"serial,omitempty"`     // Serial
	Status       string `json:"status,omitempty"`     // Status
	ErrorCode    string `json:"error_code,omitempty"` // ErrorCode
	ErrorMessage string `json:"error_msg,omitempty"`  // ErrorMessage
}

// APIError APIコール時のエラー情報
type APIError interface {
	// errorインターフェースを内包
	error

	// エラー発生時のレスポンスコード
	ResponseCode() int

	// エラーコード
	Code() string

	// エラー発生時のメッセージ
	Message() string

	// エラー追跡用シリアルコード
	Serial() string

	// エラー(オリジナル)
	OrigErr() *APIErrorResponse
}

// NewAPIError APIコール時のエラー情報
func NewAPIError(requestMethod string, requestURL *url.URL, responseCode int, err *APIErrorResponse) APIError {
	return &apiError{
		responseCode: responseCode,
		method:       requestMethod,
		url:          requestURL,
		origErr:      err,
	}
}

type apiError struct {
	responseCode int
	method       string
	url          *url.URL
	origErr      *APIErrorResponse
}

// Error errorインターフェース
func (e *apiError) Error() string {
	return fmt.Sprintf("Error in response: %#v", e.origErr)
}

// ResponseCode エラー発生時のレスポンスコード
func (e *apiError) ResponseCode() int {
	return e.responseCode
}

// Code エラーコード
func (e *apiError) Code() string {
	if e.origErr != nil {
		return e.origErr.ErrorCode
	}
	return ""
}

// Message エラー発生時のメッセージ
func (e *apiError) Message() string {
	if e.origErr != nil {
		return e.origErr.ErrorMessage
	}
	return ""
}

// Serial エラー追跡用シリアルコード
func (e *apiError) Serial() string {
	if e.origErr != nil {
		return e.origErr.Serial
	}
	return ""
}

// OrigErr エラー(オリジナル)
func (e *apiError) OrigErr() *APIErrorResponse {
	return e.origErr
}
