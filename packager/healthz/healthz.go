// Copyright 2019 Google LLC
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

package healthz

import (
	"net/http"

	"github.com/ampproject/amppackager/packager/certcache"
)

type Healthz struct {
	certHandler certcache.CertHandler
}

func New(certHandler certcache.CertHandler) (*Healthz, error) {
	return &Healthz{certHandler}, nil
}

func (this *Healthz) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//curl -k --raw -i -v  --http1.1  https://127.0.0.1:8080/healthz

	// flusher, ok := resp.(http.Flusher)
	// if !ok {
	// 	panic("expected http.ResponseWriter to be an http.Flusher")
	// }
	// // resp.Header().Set("Transfer-Encoding", "chunked") // Also removed, per RFC 2616.
	// resp.Header().Set("X-Content-Type-Options", "nosniff")
	// for i := 1; i <= 10; i++ {
	// 	fmt.Fprintf(resp, "Chunk #%d\n", i)
	// 	flusher.Flush() // Trigger "chunked" encoding and send a chunk...
	// 	// time.Sleep(500 * time.Millisecond)
	// }
}
