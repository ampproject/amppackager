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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"

	client "github.com/sacloud/api-client-go"
	sacloudhttp "github.com/sacloud/go-http"
	"github.com/sacloud/iaas-api-go/types"
)

var (
	// SakuraCloudAPIRoot APIリクエスト送信先ルートURL(末尾にスラッシュを含まない)
	SakuraCloudAPIRoot = "https://secure.sakura.ad.jp/cloud/zone"

	// SakuraCloudZones 利用可能なゾーンのデフォルト値
	SakuraCloudZones = types.ZoneNames
)

var (
	// APIDefaultZone デフォルトゾーン、グローバルリソースなどで利用される
	APIDefaultZone = "is1a"
	// DefaultUserAgent デフォルトのユーザーエージェント
	DefaultUserAgent = fmt.Sprintf(
		"sacloud/iaas-api-go/v%s (%s/%s; +https://github.com/sacloud/iaas-api-go) %s",
		Version,
		runtime.GOOS,
		runtime.GOARCH,
		sacloudhttp.DefaultUserAgent,
	)

	defaultCheckRetryStatusCodes = []int{
		http.StatusServiceUnavailable,
		http.StatusLocked,
	}
)

const (
	// APIAccessTokenEnvKey APIアクセストークンの環境変数名
	APIAccessTokenEnvKey = "SAKURACLOUD_ACCESS_TOKEN" //nolint:gosec
	// APIAccessSecretEnvKey APIアクセスシークレットの環境変数名
	APIAccessSecretEnvKey = "SAKURACLOUD_ACCESS_TOKEN_SECRET" //nolint:gosec
)

// APICaller API呼び出し時に利用するトランスポートのインターフェース iaas.Clientなどで実装される
type APICaller interface {
	Do(ctx context.Context, method, uri string, body interface{}) ([]byte, error)
}

// Client APIクライアント、APICallerインターフェースを実装する
//
// レスポンスステータスコード423、または503を受け取った場合、RetryMax回リトライする
// リトライ間隔はRetryMinからRetryMaxまで指数的に増加する(Exponential Backoff)
//
// リトライ時にcontext.Canceled、またはcontext.DeadlineExceededの場合はリトライしない
type Client struct {
	factory *client.Factory
}

// NewClient APIクライアント作成
func NewClient(token, secret string) *Client {
	opts := &client.Options{
		AccessToken:       token,
		AccessTokenSecret: secret,
	}
	return NewClientWithOptions(opts)
}

// NewClientFromEnv 環境変数からAPIキーを取得してAPIクライアントを作成する
func NewClientFromEnv() *Client {
	return NewClientWithOptions(client.OptionsFromEnv())
}

// NewClientWithOptions 指定のオプションでAPIクライアントを作成する
func NewClientWithOptions(opts *client.Options) *Client {
	if len(opts.CheckRetryStatusCodes) == 0 {
		opts.CheckRetryStatusCodes = defaultCheckRetryStatusCodes
	}
	factory := client.NewFactory(opts)
	return &Client{factory: factory}
}

// Do APIコール実施
func (c *Client) Do(ctx context.Context, method, uri string, body interface{}) ([]byte, error) {
	req, err := c.newRequest(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}

	// API call
	resp, err := c.factory.NewHttpRequestDoer().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !c.isOkStatus(resp.StatusCode) {
		errResponse := &APIErrorResponse{}
		err := json.Unmarshal(data, errResponse)
		if err != nil {
			return nil, fmt.Errorf("error in response: %s", string(data))
		}
		return nil, NewAPIError(req.Method, req.URL, resp.StatusCode, errResponse)
	}

	return data, nil
}

func (c *Client) newRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
	// setup url and body
	var url = uri
	var bodyReader io.ReadSeeker
	if body != nil {
		var bodyJSON []byte
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		if method == "GET" {
			url = fmt.Sprintf("%s?%s", url, bytes.NewBuffer(bodyJSON))
		} else {
			bodyReader = bytes.NewReader(bodyJSON)
		}
	}
	return http.NewRequestWithContext(ctx, method, url, bodyReader)
}

func (c *Client) isOkStatus(code int) bool {
	codes := map[int]bool{
		http.StatusOK:        true,
		http.StatusCreated:   true,
		http.StatusAccepted:  true,
		http.StatusNoContent: true,
	}
	_, ok := codes[code]
	return ok
}
