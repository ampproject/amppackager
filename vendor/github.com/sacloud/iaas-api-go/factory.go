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

var clientFactory = make(map[string]func(APICaller) interface{})

// SetClientFactoryFunc リソースごとのクライアントファクトリーを登録する
func SetClientFactoryFunc(resourceName string, factoryFunc func(caller APICaller) interface{}) {
	clientFactory[resourceName] = factoryFunc
}

var clientFactoryHooks = make(map[string][]func(interface{}) interface{})

// AddClientFacotyHookFunc クライアントファクトリーのフックを登録する
func AddClientFacotyHookFunc(resourceName string, hookFunc func(interface{}) interface{}) {
	clientFactoryHooks[resourceName] = append(clientFactoryHooks[resourceName], hookFunc)
}

// GetClientFactoryFunc リソースごとのクライアントファクトリーを取得する
//
// resourceNameに対するファクトリーが登録されてない場合はpanicする
func GetClientFactoryFunc(resourceName string) func(APICaller) interface{} {
	f, ok := clientFactory[resourceName]
	if !ok {
		panic(resourceName + " is not found in clientFactory")
	}
	if hooks, ok := clientFactoryHooks[resourceName]; ok {
		return func(caller APICaller) interface{} {
			ret := f(caller)
			for _, hook := range hooks {
				ret = hook(ret)
			}
			return ret
		}
	}
	return f
}
