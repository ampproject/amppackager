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

package util

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPError encodes an internal message to be logged and an HTTP status code
// to be used for the external error message. External errors should only be
// used to signal misconfiguration of the packager. For errors that are
// transient or a result of downstream server errors, the signer should fall
// back to proxying the content unsigned.
type HTTPError struct {
	InternalMsg string
	StatusCode  int
}

func NewHTTPError(statusCode int, msg ...interface{}) *HTTPError {
	return &HTTPError{fmt.Sprint(msg...), statusCode}
}

func (e *HTTPError) LogAndRespond(resp http.ResponseWriter) {
	log.Println(e.InternalMsg)
	resp.Header().Set("Cache-Control", "no-store")
	http.Error(resp, http.StatusText(e.StatusCode), e.StatusCode)
}
