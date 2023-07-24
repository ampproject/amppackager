// Copyright 2021-2023 The sacloud/go-http authors
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

package http

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

// TracingRoundTripper リクエスト/レスポンスのトレースログを出力するためのhttp.RoundTripper実装
//
// Client.Gzipがtrueの場合でも関知しないため利用者側で制御する必要がある
type TracingRoundTripper struct {
	// Transport 親となるhttp.RoundTripper、nilの場合http.DefaultTransportが利用される
	Transport http.RoundTripper
	// OutputOnlyError trueの場合レスポンスのステータスコードが200番台の時はリクエスト/レスポンスのトレースを出力しない
	OutputOnlyError bool
}

// RoundTrip http.RoundTripperの実装
func (r *TracingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if r.Transport == nil {
		r.Transport = http.DefaultTransport
	}

	var bodyBytes []byte
	if req.Body != nil {
		bb, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		bodyBytes = bb
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	res, err := r.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if r.OutputOnlyError && res.StatusCode < 300 {
		return res, err
	}

	if req.Body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	data, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("[TRACE] \trequest: %s %s\n==============================\n%s\n============================\n", req.Method, req.URL.String(), string(data))

	data, err = httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}
	log.Printf("[TRACE] \tresponse: %s %s\n==============================\n%s\n============================\n", req.Method, req.URL.String(), string(data))

	return res, err
}
