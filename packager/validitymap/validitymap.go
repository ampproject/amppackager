// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validitymap

import (
	"bytes"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type ValidityMap struct {
	validityMap []byte
}

func New() (*ValidityMap, error) {
	this := new(ValidityMap)
	// https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.6
	// This is an empty validity map `{}`.
	this.validityMap = []byte("\xA0")
	return this, nil
}

func (this *ValidityMap) ServeHTTP(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	resp.Header().Set("Content-Type", "application/cbor")
	resp.Header().Set("Cache-Control", "public, max-age=604800")
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	http.ServeContent(resp, req, "", time.Time{}, bytes.NewReader(this.validityMap))
}
