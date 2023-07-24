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

package mapconv

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/sacloud/packages-go/objutil"
)

// DefaultMapConvTag デフォルトのmapconvタグ名
const DefaultMapConvTag = "mapconv"

// DecoderConfig mapconvでの変換の設定
type DecoderConfig struct {
	TagName     string
	FilterFuncs map[string]FilterFunc
}

// FilterFunc mapconvでの変換時に適用するフィルタ
type FilterFunc func(v interface{}) (interface{}, error)

// TagInfo mapconvタグの情報
type TagInfo struct {
	Ignore       bool
	SourceFields []string
	Filters      []string
	DefaultValue interface{}
	OmitEmpty    bool
	Recursive    bool
	Squash       bool
	IsSlice      bool
}

// Decoder mapconvでの変換
type Decoder struct {
	Config *DecoderConfig
}

func (d *Decoder) ConvertTo(source interface{}, dest interface{}) error {
	s := structs.New(source)
	mappedValues := Map(make(map[string]interface{}))

	// recursiveの際に参照するためのdestのmap
	destValues := Map(make(map[string]interface{}))
	if structs.IsStruct(dest) {
		destValues = Map(structs.Map(dest))
	}

	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() {
			continue
		}

		tags := d.ParseMapConvTag(f.Tag(d.Config.TagName))
		if tags.Ignore {
			continue
		}
		for _, key := range tags.SourceFields {
			destKey := f.Name()
			value := f.Value()

			if key != "" {
				destKey = key
			}
			if f.IsZero() {
				if tags.OmitEmpty {
					continue
				}
				if tags.DefaultValue != nil {
					value = tags.DefaultValue
				}
			}

			for _, filter := range tags.Filters {
				filterFunc, ok := d.Config.FilterFuncs[filter]
				if !ok {
					return fmt.Errorf("filter %s not exists", filter)
				}
				filtered, err := filterFunc(value)
				if err != nil {
					return fmt.Errorf("failed to apply the filter: %s", err)
				}
				value = filtered
			}

			if tags.Squash {
				dest := Map(make(map[string]interface{}))
				err := d.ConvertTo(value, &dest)
				if err != nil {
					return err
				}
				for k, v := range dest {
					mappedValues.Set(k, v)
				}
				continue
			}

			if tags.Recursive {
				current, err := destValues.Get(destKey)
				if err != nil {
					return err
				}

				var dest []interface{}
				values := valueToSlice(value)
				currentValues := valueToSlice(current)
				for i, v := range values {
					if structs.IsStruct(v) {
						var currentDest interface{}
						if len(currentValues) > i {
							currentDest = currentValues[i]
						}
						destMap := Map(make(map[string]interface{}))
						if err := d.ConvertTo(v, &destMap); err != nil {
							return err
						}
						// 宛先が存在しstructであれば(map[string]interface{}になっているはずなので)マージする
						if currentDest != nil {
							mv, ok := currentDest.(map[string]interface{})
							// 元の値から空の値を除去する(structs:",omitempty"でも可)
							for k, v := range mv {
								if objutil.IsEmpty(v) {
									delete(mv, k)
								}
							}
							if ok {
								for k, v := range destMap.Map() {
									mv[k] = v
								}
								destMap = Map(mv)
							}
						}
						dest = append(dest, destMap)
					} else {
						dest = append(dest, v)
					}
				}
				if tags.IsSlice || dest == nil || len(dest) > 1 {
					value = dest
				} else {
					value = dest[0]
				}
			}

			mappedValues.Set(destKey, value)
		}
	}

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           dest,
		ZeroFields:       true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(mappedValues.Map())
}

func (d *Decoder) ConvertFrom(source interface{}, dest interface{}) error {
	var sourceMap Map
	if m, ok := source.(map[string]interface{}); ok {
		sourceMap = Map(m)
	} else {
		sourceMap = Map(structs.New(source).Map())
	}
	destMap := Map(make(map[string]interface{}))

	s := structs.New(dest)
	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() {
			continue
		}

		tags := d.ParseMapConvTag(f.Tag(d.Config.TagName))
		if tags.Ignore {
			continue
		}
		if tags.Squash {
			return errors.New("ConvertFrom is not allowed squash")
		}
		for _, key := range tags.SourceFields {
			sourceKey := f.Name()
			if key != "" {
				sourceKey = key
			}

			value, err := sourceMap.Get(sourceKey)
			if err != nil {
				return err
			}
			if value == nil || reflect.ValueOf(value).IsZero() {
				continue
			}

			for _, filter := range tags.Filters {
				filterFunc, ok := d.Config.FilterFuncs[filter]
				if !ok {
					return fmt.Errorf("filter %s not exists", filter)
				}
				filtered, err := filterFunc(value)
				if err != nil {
					return fmt.Errorf("failed to apply the filter: %s", err)
				}
				value = filtered
			}

			if tags.Recursive {
				t := reflect.TypeOf(f.Value())
				if t.Kind() == reflect.Slice {
					t = t.Elem().Elem()
				} else {
					t = t.Elem()
				}

				var dest []interface{}
				values := valueToSlice(value)
				for _, v := range values {
					if v == nil {
						dest = append(dest, v)
						continue
					}
					dt := reflect.New(t).Interface()
					if err := d.ConvertFrom(v, dt); err != nil {
						return err
					}
					dest = append(dest, dt)
				}

				if dest != nil {
					if tags.IsSlice || len(dest) > 1 {
						value = dest
					} else {
						value = dest[0]
					}
				}
			}

			destMap.Set(f.Name(), value)
		}
	}
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           dest,
		ZeroFields:       true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(destMap.Map())
}

// ConvertTo converts struct which input by mapconv to plain models
func ConvertTo(source interface{}, dest interface{}) error {
	decoder := &Decoder{Config: &DecoderConfig{TagName: DefaultMapConvTag}}
	return decoder.ConvertTo(source, dest)
}

// ConvertFrom converts struct which input by mapconv from plain models
func ConvertFrom(source interface{}, dest interface{}) error {
	decoder := &Decoder{Config: &DecoderConfig{TagName: DefaultMapConvTag}}
	return decoder.ConvertFrom(source, dest)
}

// ParseMapConvTag mapconvタグを文字列で受け取りパースしてTagInfoを返す
func (d *Decoder) ParseMapConvTag(tagBody string) TagInfo {
	tokens := strings.Split(tagBody, ",")
	key := strings.TrimSpace(tokens[0])

	keys := strings.Split(key, "/")
	var defaultValue interface{}
	var filters []string
	var ignore, omitEmpty, recursive, squash, isSlice bool

	for _, k := range keys {
		if k == "-" {
			ignore = true
			break
		}
		if strings.Contains(k, "[]") {
			isSlice = true
		}
	}

	for i, token := range tokens {
		if i == 0 {
			continue
		}

		token = strings.TrimSpace(token)

		switch {
		case strings.HasPrefix(token, "omitempty"):
			omitEmpty = true
		case strings.HasPrefix(token, "recursive"):
			recursive = true
		case strings.HasPrefix(token, "squash"):
			squash = true
		case strings.HasPrefix(token, "filters"):
			keyValue := strings.Split(token, "=")
			if len(keyValue) > 1 {
				filters = strings.Split(strings.Join(keyValue[1:], ""), " ")
			}
		case strings.HasPrefix(token, "default"):
			keyValue := strings.Split(token, "=")
			if len(keyValue) > 1 {
				defaultValue = strings.Join(keyValue[1:], "")
			}
		}
	}
	return TagInfo{
		Ignore:       ignore,
		SourceFields: keys,
		DefaultValue: defaultValue,
		OmitEmpty:    omitEmpty,
		Recursive:    recursive,
		Squash:       squash,
		IsSlice:      isSlice,
		Filters:      filters,
	}
}
