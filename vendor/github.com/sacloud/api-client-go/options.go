// Copyright 2022-2023 The sacloud/api-client-go Authors
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

package client

import (
	"context"
	"net/http"
	"strings"

	"github.com/sacloud/api-client-go/profile"
	sacloudhttp "github.com/sacloud/go-http"
	"github.com/sacloud/packages-go/envvar"
)

// Options sacloudhttp.Clientを作成する際のオプション
type Options struct {
	// AccessToken APIキー:トークン
	AccessToken string
	// AccessTokenSecret APIキー:シークレット
	AccessTokenSecret string

	// AcceptLanguage APIリクエスト時のAccept-Languageヘッダーの値
	AcceptLanguage string

	// Gzip APIリクエストでgzipを有効にするかのフラグ
	Gzip bool

	// HttpClient APIリクエストで使用されるHTTPクライアント
	//
	// 省略した場合はhttp.DefaultClientが使用される
	HttpClient *http.Client

	// HttpRequestTimeout HTTPリクエストのタイムアウト秒数
	HttpRequestTimeout int
	// HttpRequestRateLimit 1秒あたりの上限リクエスト数
	HttpRequestRateLimit int

	// RetryMax リトライ上限回数
	RetryMax int

	// RetryWaitMax リトライ待ち秒数(最大)
	RetryWaitMax int
	// RetryWaitMin リトライ待ち秒数(最小)
	RetryWaitMin int

	// UserAgent ユーザーエージェント
	UserAgent string

	// Trace HTTPリクエスト/レスポンスのトレースログ(ダンプ)出力
	Trace bool
	// TraceOnlyError HTTPリクエスト/レスポンスのトレースログ(ダンプ)出力で非200番台のレスポンス時のみ出力する
	TraceOnlyError bool

	// RequestCustomizers リクエスト前に*http.Requestのカスタマイズを行うためのfunc
	RequestCustomizers []sacloudhttp.RequestCustomizer

	// CheckRetryFunc リトライすべきか判定するためのfunc
	//
	// CheckRetryStatusCodesより優先される
	CheckRetryFunc func(ctx context.Context, resp *http.Response, err error) (bool, error)

	// CheckRetryStatusCodes リトライすべきステータスコード
	//
	// CheckRetryFuncが指定されていない、かつこの値が指定されている場合、指定のステータスコードを持つレスポンスを受け取ったらリトライする
	CheckRetryStatusCodes []int

	// profileConfigValue プロファイルから読み込んだ値を保持する
	profileConfigValue *profile.ConfigValue
}

// ProfileConfigValue プロファイルから読み込んだprofile.ConfigValueを返す
func (o *Options) ProfileConfigValue() *profile.ConfigValue {
	return o.profileConfigValue
}

// DefaultOption 環境変数、プロファイルからCallerOptionsを組み立てて返す
//
// プロファイルは環境変数`SAKURACLOUD_PROFILE`または`USACLOUD_PROFILE`でプロファイル名が指定されていればそちらを優先し、
// 未指定の場合は通常のプロファイル処理(~/.usacloud/currentファイルから読み込み)される。
// 同じ項目を複数箇所で指定していた場合、環境変数->プロファイルの順で上書きされたものが返される
func DefaultOption() (*Options, error) {
	return DefaultOptionWithProfile("")
}

// DefaultOptionWithProfile 環境変数、プロファイルからCallerOptionsを組み立てて返す
//
// プロファイルは引数を優先し、空の場合は環境変数`SAKURACLOUD_PROFILE`または`USACLOUD_PROFILE`が利用され、
// それも空の場合は通常のプロファイル処理(~/.usacloud/currentファイルから読み込み)される。
// 同じ項目を複数箇所で指定していた場合、環境変数->プロファイルの順で上書きされたものが返される
func DefaultOptionWithProfile(profileName string) (*Options, error) {
	fromProfile, err := OptionsFromProfile(profileName)
	if err != nil {
		return nil, err
	}
	return MergeOptions(defaultOption, OptionsFromEnv(), fromProfile), nil
}

var defaultOption = &Options{
	HttpRequestTimeout:   300,
	HttpRequestRateLimit: 5,
	RetryMax:             sacloudhttp.DefaultRetryMax,
	RetryWaitMax:         int(sacloudhttp.DefaultRetryWaitMax.Seconds()),
	RetryWaitMin:         int(sacloudhttp.DefaultRetryWaitMin.Seconds()),
}

// MergeOptions 指定のCallerOptionsの非ゼロ値フィールドをoのコピーにマージして返す
func MergeOptions(opts ...*Options) *Options {
	merged := &Options{}
	for _, opt := range opts {
		if opt.AccessToken != "" {
			merged.AccessToken = opt.AccessToken
		}
		if opt.AccessTokenSecret != "" {
			merged.AccessTokenSecret = opt.AccessTokenSecret
		}
		if opt.AcceptLanguage != "" {
			merged.AcceptLanguage = opt.AcceptLanguage
		}
		if opt.HttpClient != nil {
			merged.HttpClient = opt.HttpClient
		}
		if opt.HttpRequestTimeout > 0 {
			merged.HttpRequestTimeout = opt.HttpRequestTimeout
		}
		if opt.HttpRequestRateLimit > 0 {
			merged.HttpRequestRateLimit = opt.HttpRequestRateLimit
		}
		if opt.RetryMax > 0 {
			merged.RetryMax = opt.RetryMax
		}
		if opt.RetryWaitMax > 0 {
			merged.RetryWaitMax = opt.RetryWaitMax
		}
		if opt.RetryWaitMin > 0 {
			merged.RetryWaitMin = opt.RetryWaitMin
		}
		if opt.UserAgent != "" {
			merged.UserAgent = opt.UserAgent
		}

		if opt.profileConfigValue != nil {
			merged.profileConfigValue = opt.profileConfigValue
		}

		// Note: bool値は一度trueにしたらMergeでfalseになることがない
		if opt.Gzip {
			merged.Gzip = true
		}
		if opt.Trace {
			merged.Trace = true
		}
		if opt.TraceOnlyError {
			merged.TraceOnlyError = true
		}
		if len(opt.RequestCustomizers) > 0 {
			merged.RequestCustomizers = opt.RequestCustomizers
		}
		if opt.CheckRetryFunc != nil {
			merged.CheckRetryFunc = opt.CheckRetryFunc
		}
		if len(opt.CheckRetryStatusCodes) > 0 {
			merged.CheckRetryStatusCodes = opt.CheckRetryStatusCodes
		}
	}
	return merged
}

// OptionsFromEnv 環境変数からCallerOptionsを組み立てて返す
func OptionsFromEnv() *Options {
	return &Options{
		AccessToken:       envvar.StringFromEnv("SAKURACLOUD_ACCESS_TOKEN", ""),
		AccessTokenSecret: envvar.StringFromEnv("SAKURACLOUD_ACCESS_TOKEN_SECRET", ""),

		AcceptLanguage: envvar.StringFromEnv("SAKURACLOUD_ACCEPT_LANGUAGE", ""),
		Gzip:           envvar.StringFromEnv("SAKURACLOUD_GZIP", "") != "",

		HttpRequestTimeout:   envvar.IntFromEnv("SAKURACLOUD_API_REQUEST_TIMEOUT", 0),
		HttpRequestRateLimit: envvar.IntFromEnv("SAKURACLOUD_API_REQUEST_RATE_LIMIT", 0),

		RetryMax:     envvar.IntFromEnv("SAKURACLOUD_RETRY_MAX", 0),
		RetryWaitMax: envvar.IntFromEnv("SAKURACLOUD_RETRY_WAIT_MAX", 0),
		RetryWaitMin: envvar.IntFromEnv("SAKURACLOUD_RETRY_WAIT_MIN", 0),

		Trace:          envvar.StringFromEnv("SAKURACLOUD_TRACE", "") != "",
		TraceOnlyError: strings.ToLower(envvar.StringFromEnv("SAKURACLOUD_TRACE", "")) == "error",
	}
}

// OptionsFromProfile 指定のプロファイルからCallerOptionsを組み立てて返す
//
// プロファイルは引数を優先し、空の場合は環境変数`SAKURACLOUD_PROFILE`または`USACLOUD_PROFILE`が利用され、
// それも空の場合は通常のプロファイル処理(~/.usacloud/currentファイルから読み込み)される。
func OptionsFromProfile(profileName string) (*Options, error) {
	// 引数がからの場合はまず環境変数から
	if profileName == "" {
		profileName = envvar.StringFromEnvMulti([]string{"SAKURACLOUD_PROFILE", "USACLOUD_PROFILE"}, "")
	}
	// それも空ならプロファイルのcurrentファイルから
	if profileName == "" {
		current, err := profile.CurrentName()
		if err != nil {
			return nil, err
		}
		profileName = current
	}

	config := profile.ConfigValue{}
	if err := profile.Load(profileName, &config); err != nil {
		return nil, err
	}

	return &Options{
		AccessToken:          config.AccessToken,
		AccessTokenSecret:    config.AccessTokenSecret,
		AcceptLanguage:       config.AcceptLanguage,
		Gzip:                 config.Gzip,
		HttpRequestTimeout:   config.HTTPRequestTimeout,
		HttpRequestRateLimit: config.HTTPRequestRateLimit,
		RetryMax:             config.RetryMax,
		RetryWaitMax:         config.RetryWaitMax,
		RetryWaitMin:         config.RetryWaitMin,
		Trace:                config.EnableHTTPTrace(),
		TraceOnlyError:       strings.ToLower(config.TraceMode) == "error",
		profileConfigValue:   &config,
	}, nil
}
