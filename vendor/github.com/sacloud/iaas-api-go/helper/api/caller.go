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
	"strings"
	"time"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/defaults"
	"github.com/sacloud/iaas-api-go/fake"
	"github.com/sacloud/iaas-api-go/trace"
)

func NewCaller() (iaas.APICaller, error) {
	clientOpts, err := client.DefaultOption()
	if err != nil {
		return nil, err
	}
	return NewCallerWithOptions(&CallerOptions{Options: clientOpts}), nil
}

// NewCallerWithOptions 指定のオプションでiaas.APICallerを構築して返す
func NewCallerWithOptions(opts *CallerOptions) iaas.APICaller {
	return newCaller(opts)
}

func newCaller(opts *CallerOptions) iaas.APICaller {
	if opts.UserAgent == "" {
		opts.UserAgent = iaas.DefaultUserAgent
	}

	caller := iaas.NewClientWithOptions(opts.Options)

	defaults.DefaultStatePollingTimeout = 72 * time.Hour

	if opts.TraceAPI {
		// note: exact once
		trace.AddClientFactoryHooks()
	}

	if opts.FakeMode {
		if opts.FakeStorePath != "" {
			fake.DataStore = fake.NewJSONFileStore(opts.FakeStorePath)
		}
		// note: exact once
		fake.SwitchFactoryFuncToFake()

		SetupFakeDefaults()
	}

	if opts.DefaultZone != "" {
		iaas.APIDefaultZone = opts.DefaultZone
	}

	if len(opts.Zones) > 0 {
		iaas.SakuraCloudZones = opts.Zones
	}

	if opts.APIRootURL != "" {
		if strings.HasSuffix(opts.APIRootURL, "/") {
			opts.APIRootURL = strings.TrimRight(opts.APIRootURL, "/")
		}
		iaas.SakuraCloudAPIRoot = opts.APIRootURL
	}

	return caller
}

func SetupFakeDefaults() {
	defaultInterval := 10 * time.Millisecond

	// update default polling intervals: libsacloud/sacloud
	defaults.DefaultStatePollingInterval = defaultInterval
	defaults.DefaultDBStatusPollingInterval = defaultInterval

	// update default polling intervals: libsacloud/helper/setup
	// update default polling intervals: libsacloud/helper/builder
	defaults.DefaultNICUpdateWaitDuration = defaultInterval
	// update default timeouts and span: libsacloud/helper/power
	defaults.DefaultPowerHelperBootRetrySpan = defaultInterval
	defaults.DefaultPowerHelperShutdownRetrySpan = defaultInterval
	defaults.DefaultPowerHelperInitialRequestRetrySpan = defaultInterval
	defaults.DefaultPowerHelperInitialRequestTimeout = defaultInterval * 100

	fake.PowerOnDuration = time.Millisecond
	fake.PowerOffDuration = time.Millisecond
	fake.DiskCopyDuration = time.Millisecond
}
