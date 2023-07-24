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

//go:build !wasm
// +build !wasm

package fake

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func (s *JSONFileStore) startWatcher() {
	ctx := s.Ctx
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	log.Printf("file watch start: %q", s.Path)

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				switch {
				case event.Op&fsnotify.Write == fsnotify.Write,
					event.Op&fsnotify.Create == fsnotify.Create,
					event.Op&fsnotify.Rename == fsnotify.Rename:

					if err := s.load(); err != nil {
						log.Printf("reloading %q is failed: %s\n", s.Path, err)
					}

					if event.Op&fsnotify.Rename == fsnotify.Rename {
						if err := watcher.Add(s.Path); err != nil {
							panic(err)
						}
					}
					log.Printf("reloaded: %q\n", s.Path)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				panic(err)
			case <-ctx.Done():
				return
			}
		}
	}()
	if err := watcher.Add(s.Path); err != nil {
		panic(err)
	}
}
