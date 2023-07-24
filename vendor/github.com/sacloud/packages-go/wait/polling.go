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
	"time"
)

var (
	defaultPollingTimeout  = 20 * time.Minute
	defaultPollingInterval = 5 * time.Second
)

// StateReadFunc PollingWaiterにより利用される、対象リソースの状態を取得するためのfunc
type StateReadFunc func() (state interface{}, err error)

// StateCheckFunc StateReadFuncで得たリソースの情報を元に待ちを継続するか判定するためのfunc
//
// completeがtrueの場合待ち処理を終了する
type StateCheckFunc func(target interface{}) (complete bool, err error)

// ComposeStateCheckFunc 指定のStateCheckFuncを順に適用するするStateCheckFuncを生成する
//
// completeがtrueまたはerrが非nilになったら即時リターンする。
func ComposeStateCheckFunc(funcs ...StateCheckFunc) StateCheckFunc {
	return func(target interface{}) (bool, error) {
		for _, f := range funcs {
			complete, err := f(target)
			if err != nil {
				return false, err
			}
			if complete {
				return complete, nil
			}
		}
		return false, nil
	}
}

// PollingWaiter ポーリングでステート変更を検知するWaiter
type PollingWaiter struct {
	// ReadFunc 対象リソースの状態を取得するためのfunc
	//
	// ReadFuncがerrorを返した場合は待ち処理が即時リターンする。非エラーかつ非nilを返した場合のみStateCheckFuncでの判定処理を行う。
	ReadFunc StateReadFunc

	// StateCheckFunc ReadFuncで得たリソースの情報を元に待ちを継続するかの判定を行うためのfunc
	StateCheckFunc StateCheckFunc

	// Timeout タイムアウト
	Timeout time.Duration // タイムアウト

	// Interval ポーリング間隔
	Interval time.Duration
}

// WaitForState リソースが指定の状態になるまで待つ
func (w *PollingWaiter) WaitForState(ctx context.Context) (interface{}, error) {
	c, p, e := w.WaitForStateAsync(ctx)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case lastState := <-c:
			return lastState, nil
		case <-p:
			// noop
		case err := <-e:
			return nil, err
		}
	}
}

// WaitForStateAsync リソースが指定の状態になるまで待つ
func (w *PollingWaiter) WaitForStateAsync(ctx context.Context) (<-chan interface{}, <-chan interface{}, <-chan error) {
	w.validateFields()
	w.defaults()

	compCh := make(chan interface{})
	progressCh := make(chan interface{})
	errCh := make(chan error)

	ticker := time.NewTicker(w.Interval)

	go func() {
		ctx, cancel := context.WithTimeout(ctx, w.Timeout)
		defer cancel()

		defer ticker.Stop()

		defer close(compCh)
		defer close(progressCh)
		defer close(errCh)

		for {
			select {
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			case <-ticker.C:
				state, err := w.ReadFunc()
				if err != nil {
					errCh <- err
					return
				}
				if state != nil {
					complete, err := w.StateCheckFunc(state)
					if complete {
						compCh <- state
						return
					}

					if err != nil {
						errCh <- err
						return
					}
				}
				// note: nilの場合もあり得る
				progressCh <- state
			}
		}
	}()

	return compCh, progressCh, errCh
}

func (w *PollingWaiter) validateFields() {
	if w.ReadFunc == nil {
		panic("required: PollingWaiter.ReadFunc")
	}

	if w.StateCheckFunc == nil {
		panic("required: PollingWaiter.StateCheckFunc")
	}
}

func (w *PollingWaiter) defaults() {
	if w.Timeout == time.Duration(0) {
		w.Timeout = defaultPollingTimeout
	}
	if w.Interval == time.Duration(0) {
		w.Interval = defaultPollingInterval
	}
}
