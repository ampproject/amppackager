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

import "net/http"

// RequestCustomizer リクエスト前に*http.Requestのカスタマイズを行うためのfunc
type RequestCustomizer func(r *http.Request) error

// ComposeRequestCustomizer 任意の個数のRequestCustomizerを合成してRequestCustomizerを返す
//
// 複数のRequestCustomizerを指定した場合は先頭から呼びだされ、エラーを返したら即時returnする
func ComposeRequestCustomizer(funcs ...RequestCustomizer) RequestCustomizer {
	return func(r *http.Request) error {
		for _, fn := range funcs {
			if err := fn(r); err != nil {
				return err
			}
		}
		return nil
	}
}
