// Copyright 2022-2023 The sacloud/packages-go Authors
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

package wait

import (
	"context"
)

// StateWaiter リソースの状態が変わるまで待機する
type StateWaiter interface {
	// WaitForState リソースが指定の状態になるまで待つ
	WaitForState(context.Context) (interface{}, error)
	// WaitForStateAsync リソースが指定の状態になるまで待つ(非同期)
	WaitForStateAsync(context.Context) (completeCh <-chan interface{}, progressCh <-chan interface{}, errorCh <-chan error)
}
