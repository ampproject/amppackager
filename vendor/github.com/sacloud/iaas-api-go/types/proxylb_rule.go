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

package types

// EProxyLBRuleAction エンハンスドロードバランサでのルール設定: アクション
type EProxyLBRuleAction string

func (a EProxyLBRuleAction) String() string {
	return string(a)
}

// ProxyLBRuleActions エンハンスドロードバランサでのルール設定: アクション
var ProxyLBRuleActions = struct {
	Forward  EProxyLBRuleAction
	Redirect EProxyLBRuleAction
	Fixed    EProxyLBRuleAction
}{
	Forward:  "forward",
	Redirect: "redirect",
	Fixed:    "fixed",
}

func ProxyLBRuleActionStrings() []string {
	return []string{
		ProxyLBRuleActions.Forward.String(),
		ProxyLBRuleActions.Redirect.String(),
		ProxyLBRuleActions.Fixed.String(),
	}
}

// EProxyLBRedirectStatusCode エンハンスドロードバランサでのルール設定: リダイレクト時のステータスコード
type EProxyLBRedirectStatusCode StringNumber

func (c EProxyLBRedirectStatusCode) String() string {
	return StringNumber(c).String()
}

func (c EProxyLBRedirectStatusCode) Int() int {
	return StringNumber(c).Int()
}

// ProxyLBRedirectStatusCodes エンハンスドロードバランサでのルール設定: リダイレクト時のステータスコード
var ProxyLBRedirectStatusCodes = struct {
	MovedPermanently EProxyLBRedirectStatusCode
	Found            EProxyLBRedirectStatusCode
}{
	MovedPermanently: EProxyLBRedirectStatusCode(301),
	Found:            EProxyLBRedirectStatusCode(302),
}

func ProxyLBRedirectStatusCodeStrings() []string {
	return []string{
		ProxyLBRedirectStatusCodes.MovedPermanently.String(),
		ProxyLBRedirectStatusCodes.Found.String(),
	}
}

// EProxyLBFixedStatusCode エンハンスドロードバランサでのルール設定: 固定レスポンス時のステータスコード
type EProxyLBFixedStatusCode StringNumber

func (c EProxyLBFixedStatusCode) String() string {
	return StringNumber(c).String()
}

func (c EProxyLBFixedStatusCode) Int() int {
	return StringNumber(c).Int()
}

var ProxyLBFixedStatusCodes = struct {
	OK                 EProxyLBFixedStatusCode
	Forbidden          EProxyLBFixedStatusCode
	ServiceUnavailable EProxyLBFixedStatusCode
}{
	OK:                 EProxyLBFixedStatusCode(200),
	Forbidden:          EProxyLBFixedStatusCode(403),
	ServiceUnavailable: EProxyLBFixedStatusCode(503),
}

func ProxyLBFixedStatusCodeStrings() []string {
	return []string{
		ProxyLBFixedStatusCodes.OK.String(),
		ProxyLBFixedStatusCodes.Forbidden.String(),
		ProxyLBFixedStatusCodes.ServiceUnavailable.String(),
	}
}

// EProxyLBFixedContentType エンハンスドロードバランサでのルール設定: 固定レスポンス時のContent-Typeヘッダ
type EProxyLBFixedContentType string

func (t EProxyLBFixedContentType) String() string {
	return string(t)
}

// ProxyLBFixedContentTypes エンハンスドロードバランサでのルール設定: 固定レスポンス時のContent-Typeヘッダ
var ProxyLBFixedContentTypes = struct {
	Plain      EProxyLBFixedContentType
	HTML       EProxyLBFixedContentType
	JavaScript EProxyLBFixedContentType
	JSON       EProxyLBFixedContentType
}{
	Plain:      "text/plain",
	HTML:       "text/html",
	JavaScript: "application/javascript",
	JSON:       "application/json",
}

// ProxyLBFixedContentTypeStrings 指定可能なContent-Typeリスト
func ProxyLBFixedContentTypeStrings() []string {
	return []string{
		ProxyLBFixedContentTypes.Plain.String(),
		ProxyLBFixedContentTypes.HTML.String(),
		ProxyLBFixedContentTypes.JavaScript.String(),
		ProxyLBFixedContentTypes.JSON.String(),
	}
}
