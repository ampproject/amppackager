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

package api

import (
	"fmt"
	"os"
	"runtime"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/api-client-go/profile"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/packages-go/envvar"
)

var UserAgent = fmt.Sprintf(
	"sacloud/iaas-api-go/v%s (%s/%s; +https://github.com/sacloud/iaas-api-go) %s",
	iaas.Version,
	runtime.GOOS,
	runtime.GOARCH,
	client.DefaultUserAgent,
)

// CallerOptions iaas.APICallerを作成する際のオプション
type CallerOptions struct {
	*client.Options

	APIRootURL  string
	DefaultZone string
	Zones       []string

	TraceAPI      bool
	FakeMode      bool
	FakeStorePath string
}

// DefaultOption 環境変数、プロファイルからCallerOptionsを組み立てて返す
//
// プロファイルは環境変数`SAKURACLOUD_PROFILE`または`USACLOUD_PROFILE`でプロファイル名が指定されていればそちらを優先し、
// 未指定の場合は通常のプロファイル処理(~/.usacloud/currentファイルから読み込み)される。
// 同じ項目を複数箇所で指定していた場合、環境変数->プロファイルの順で上書きされたものが返される
func DefaultOption() (*CallerOptions, error) {
	return DefaultOptionWithProfile("")
}

// DefaultOptionWithProfile 環境変数、プロファイルからCallerOptionsを組み立てて返す
//
// プロファイルは引数を優先し、空の場合は環境変数`SAKURACLOUD_PROFILE`または`USACLOUD_PROFILE`が利用され、
// それも空の場合は通常のプロファイル処理(~/.usacloud/currentファイルから読み込み)される。
// 同じ項目を複数箇所で指定していた場合、環境変数->プロファイルの順で上書きされたものが返される
func DefaultOptionWithProfile(profileName string) (*CallerOptions, error) {
	options, err := client.DefaultOptionWithProfile(profileName)
	if err != nil {
		return nil, err
	}

	fromEnv := OptionsFromEnv()
	fromEnv.Options = options

	fromProfile, err := OptionsFromProfile(profileName)
	if err != nil {
		return nil, err
	}
	fromProfile.Options = options

	defaults := &CallerOptions{
		APIRootURL:  iaas.SakuraCloudAPIRoot,
		DefaultZone: iaas.APIDefaultZone,
		Zones:       iaas.SakuraCloudZones,
		Options: &client.Options{
			UserAgent: UserAgent,
		},
	}

	return MergeOptions(defaults, fromEnv, fromProfile), nil
}

// OptionsFromEnv 環境変数からCallerOptionsを組み立てて返す
func OptionsFromEnv() *CallerOptions {
	return &CallerOptions{
		Options:     client.OptionsFromEnv(),
		APIRootURL:  envvar.StringFromEnv("SAKURACLOUD_API_ROOT_URL", ""),
		DefaultZone: envvar.StringFromEnv("SAKURACLOUD_DEFAULT_ZONE", ""),
		Zones:       envvar.StringSliceFromEnv("SAKURACLOUD_ZONES", []string{}),

		TraceAPI: profile.EnableAPITrace(envvar.StringFromEnv("SAKURACLOUD_TRACE", "")),

		FakeMode:      os.Getenv("SAKURACLOUD_FAKE_MODE") != "",
		FakeStorePath: envvar.StringFromEnv("SAKURACLOUD_FAKE_STORE_PATH", ""),
	}
}

// OptionsFromProfile 指定のプロファイルからCallerOptionsを組み立てて返す
// プロファイル名に空文字が指定された場合はカレントプロファイルが利用される
func OptionsFromProfile(profileName string) (*CallerOptions, error) {
	options, err := client.OptionsFromProfile(profileName)
	if err != nil {
		return nil, err
	}
	config := options.ProfileConfigValue()
	return &CallerOptions{
		Options:       options,
		APIRootURL:    config.APIRootURL,
		DefaultZone:   config.DefaultZone,
		Zones:         config.Zones,
		TraceAPI:      config.EnableAPITrace(),
		FakeMode:      config.FakeMode,
		FakeStorePath: config.FakeStorePath,
	}, nil
}

// MergeOptions 指定のCallerOptionsの非ゼロ値フィールドをoのコピーにマージして返す
func MergeOptions(opts ...*CallerOptions) *CallerOptions {
	merged := &CallerOptions{}
	for _, opt := range opts {
		if opt.Options != nil {
			var opts []*client.Options
			if merged.Options != nil {
				opts = append(opts, merged.Options)
			}
			opts = append(opts, opt.Options)
			merged.Options = client.MergeOptions(opts...)
		}
		if opt.APIRootURL != "" {
			merged.APIRootURL = opt.APIRootURL
		}
		if opt.DefaultZone != "" {
			merged.DefaultZone = opt.DefaultZone
		}
		if len(opt.Zones) > 0 {
			merged.Zones = opt.Zones
		}

		// Note: bool値は一度trueにしたらMergeでfalseになることがない
		if opt.TraceAPI {
			merged.TraceAPI = true
		}
		if opt.FakeMode {
			merged.FakeMode = true
		}
		if opt.FakeStorePath != "" {
			merged.FakeStorePath = opt.FakeStorePath
		}
	}
	return merged
}
