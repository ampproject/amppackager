package api

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/mimuret/golang-iij-dpf/pkg/meta"
)

var JSON = NewJSONAPIAdapter()

func UnmarshalRead(bs []byte, o interface{}) error {
	return JSON.UnmarshalRead(bs, o)
}

func MarshalCreate(body interface{}) ([]byte, error) {
	return JSON.MarshalCreate(body)
}

func MarshalUpdate(body interface{}) ([]byte, error) {
	return JSON.MarshalUpdate(body)
}

func MarshalApply(body interface{}) ([]byte, error) {
	return JSON.MarshalApply(body)
}

func MarshalOutput(spec Spec) ([]byte, error) {
	return JSON.MarshalOutput(spec)
}

func UnMarshalInput(bs []byte, obj Object) error {
	return JSON.UnMarshalInput(bs, obj)
}

type JSONAPIInterface interface {
	UnmarshalRead(bs []byte, o interface{}) error
	MarshalCreate(body interface{}) ([]byte, error)
	MarshalUpdate(body interface{}) ([]byte, error)
	MarshalApply(body interface{}) ([]byte, error)
}

type JSONFileInterface interface {
	MarshalOutput(spec Spec) ([]byte, error)
	UnMarshalInput(bs []byte, obj Object) error
}

type JSONAPIAdapter struct {
	Read   jsoniter.API
	Update jsoniter.API
	Create jsoniter.API
	Apply  jsoniter.API
	JSON   jsoniter.API
}

func NewJSONAPIAdapter() *JSONAPIAdapter {
	return &JSONAPIAdapter{
		Read: jsoniter.Config{
			EscapeHTML:             true,
			SortMapKeys:            false,
			ValidateJsonRawMessage: true,
			OnlyTaggedField:        true,
			TagKey:                 "read",
		}.Froze(),
		Create: jsoniter.Config{
			EscapeHTML:             true,
			SortMapKeys:            true,
			ValidateJsonRawMessage: true,
			OnlyTaggedField:        true,
			TagKey:                 "create",
		}.Froze(),
		Update: jsoniter.Config{
			EscapeHTML:             true,
			SortMapKeys:            true,
			ValidateJsonRawMessage: true,
			OnlyTaggedField:        true,
			TagKey:                 "update",
		}.Froze(),
		Apply: jsoniter.Config{
			EscapeHTML:             true,
			SortMapKeys:            true,
			ValidateJsonRawMessage: true,
			OnlyTaggedField:        true,
			TagKey:                 "apply",
		}.Froze(),
		JSON: jsoniter.Config{
			EscapeHTML:             true,
			SortMapKeys:            true,
			ValidateJsonRawMessage: true,
			OnlyTaggedField:        false,
			TagKey:                 "json",
		}.Froze(),
	}
}

// Unmarshal api response.
func (j *JSONAPIAdapter) UnmarshalRead(bs []byte, o interface{}) error {
	return j.Read.Unmarshal(bs, o)
}

// Marshal for create request.
func (j *JSONAPIAdapter) MarshalCreate(body interface{}) ([]byte, error) {
	return j.Create.Marshal(body)
}

// Marshal for update request.
func (j *JSONAPIAdapter) MarshalUpdate(body interface{}) ([]byte, error) {
	return j.Update.Marshal(body)
}

// Marshal for apply request.
func (j *JSONAPIAdapter) MarshalApply(body interface{}) ([]byte, error) {
	return j.Apply.Marshal(body)
}

// file format frame.
type OutputFrame struct {
	meta.KindVersion `json:",inline"`
	Resource         Object `json:"resource"`
}

// Marshal for file format.
func (j *JSONAPIAdapter) MarshalOutput(spec Spec) ([]byte, error) {
	t := reflect.TypeOf(spec)
	t = t.Elem()
	out := &OutputFrame{
		KindVersion: meta.KindVersion{
			Kind:       t.Name(),
			APIVersion: spec.GetGroup(),
		},
		Resource: spec,
	}
	return j.JSON.Marshal(out)
}

// UnMarshal for file format.
func (j *JSONAPIAdapter) UnMarshalInput(bs []byte, obj Object) error {
	out := &OutputFrame{
		Resource: obj,
	}
	return j.JSON.Unmarshal(bs, out)
}
