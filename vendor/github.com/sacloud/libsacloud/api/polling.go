// Copyright 2016-2020 The Libsacloud Authors
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
	"fmt"
	"time"
)

type pollingHandler func() (exit bool, state interface{}, err error)

func poll(handler pollingHandler, timeout time.Duration) (chan (interface{}), chan (interface{}), chan (error)) {

	compChan := make(chan interface{})
	progChan := make(chan interface{})
	errChan := make(chan error)

	tick := time.Tick(5 * time.Second)
	bomb := time.After(timeout)

	go func() {
		for {
			select {
			case <-tick:
				exit, state, err := handler()
				if state != nil {
					progChan <- state
				}
				if err != nil {
					errChan <- fmt.Errorf("Failed: poll: %s", err)
					return
				}
				if exit {
					compChan <- state
					return
				}
			case <-bomb:
				errChan <- fmt.Errorf("Timeout")
				return
			}
		}
	}()
	return compChan, progChan, errChan
}

func blockingPoll(handler pollingHandler, timeout time.Duration) error {
	c, p, e := poll(handler, timeout)
	for {
		select {
		case <-c:
			return nil
		case <-p:
			// noop
		case err := <-e:
			return err
		}
	}
}

type hasAvailable interface {
	IsAvailable() bool
}
type hasFailed interface {
	IsFailed() bool
}

func waitingForAvailableFunc(readFunc func() (hasAvailable, error), maxRetry int) func() (bool, interface{}, error) {
	counter := 0
	return func() (bool, interface{}, error) {
		v, err := readFunc()
		if err != nil {
			counter++
			if maxRetry > 0 && counter < maxRetry {
				return false, nil, nil
			}
			return false, nil, err
		}
		if v == nil {
			return false, nil, fmt.Errorf("readFunc returns nil")
		}

		if v.IsAvailable() {
			return true, v, nil
		}
		if f, ok := v.(hasFailed); ok && f.IsFailed() {
			return false, v, fmt.Errorf("InstanceState is failed: %#v", v)
		}

		return false, v, nil
	}
}

type hasUpDown interface {
	IsUp() bool
	IsDown() bool
}

func waitingForUpFunc(readFunc func() (hasUpDown, error), maxRetry int) func() (bool, interface{}, error) {
	counter := 0
	return func() (bool, interface{}, error) {
		v, err := readFunc()
		if err != nil {
			counter++
			if maxRetry > 0 && counter < maxRetry {
				return false, nil, nil
			}
			return false, nil, err
		}
		if v == nil {
			return false, nil, fmt.Errorf("readFunc returns nil")
		}

		if v.IsUp() {
			return true, v, nil
		}
		return false, v, nil
	}
}

func waitingForDownFunc(readFunc func() (hasUpDown, error), maxRetry int) func() (bool, interface{}, error) {
	counter := 0
	return func() (bool, interface{}, error) {
		v, err := readFunc()
		if err != nil {
			counter++
			if maxRetry > 0 && counter < maxRetry {
				return false, nil, nil
			}
			return false, nil, err
		}
		if v == nil {
			return false, nil, fmt.Errorf("readFunc returns nil")
		}

		if v.IsDown() {
			return true, v, nil
		}
		return false, v, nil
	}
}

func waitingForReadFunc(readFunc func() (interface{}, error), maxRetry int) func() (bool, interface{}, error) {
	counter := 0
	return func() (bool, interface{}, error) {
		v, err := readFunc()
		if err != nil {
			counter++
			if maxRetry > 0 && counter < maxRetry {
				return false, nil, nil
			}
			return false, nil, err
		}
		if v != nil {
			return true, v, nil
		}
		return false, v, nil
	}
}
