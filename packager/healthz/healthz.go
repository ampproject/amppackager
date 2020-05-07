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
	"fmt"
	"github.com/ampproject/amppackager/packager/certcache"
	"net/http"
)

type Healthz struct {
	certHandler certcache.CertHandler
}

func New(certHandler certcache.CertHandler) (*Healthz, error) {
	return &Healthz{certHandler}, nil
}

func (this *Healthz) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// Follow https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/
	err := this.certHandler.IsHealthy()
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(fmt.Sprintf("not healthy: %v", err)))
	} else {
		resp.WriteHeader(200)
		resp.Write([]byte("ok"))
	}
}
