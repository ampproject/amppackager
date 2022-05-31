package schema

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/meta"
)

var SchemaSet = NewSchemaSet()

type Register struct {
	Group string
}

func NewRegister(groupName string) *Register {
	return &Register{Group: groupName}
}

func (r *Register) Add(items ...apis.Spec) {
	SchemaSet.Add(r.Group, items)
}

func NewSchemaSet() *schemaSet {
	return &schemaSet{}
}

type schemaSet map[string]*schema

func (s schemaSet) Add(groupName string, items []apis.Spec) {
	if _, ok := s[groupName]; !ok {
		s[groupName] = &schema{group: groupName, objectMap: make(map[string]apis.Spec)}
	}
	s[groupName].Add(items)
}

type schema struct {
	group     string
	objectMap map[string]apis.Spec
}

func (s *schema) Add(items []apis.Spec) {
	for _, item := range items {
		st := reflect.TypeOf(item)
		if st.Kind() != reflect.Ptr {
			name := st.Elem().Name()
			panic(fmt.Sprintf("schema.Add name: `%s` is not ptr %v", name, item))
		}
		name := st.Elem().Name()
		if v, ok := s.objectMap[name]; ok {
			panic(fmt.Sprintf("schema.Add name: `%s` is duplicated, old: %v, new: %v", name, v, item))
		}
		s.objectMap[name] = item
	}
}

func (s schemaSet) Parse(bs []byte) (apis.Spec, error) {
	kv := &meta.KindVersion{}
	if err := json.Unmarshal(bs, kv); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}
	if kv.Kind == "" {
		return nil, fmt.Errorf("kind value is not exist")
	}
	if kv.APIVersion == "" {
		return nil, fmt.Errorf("apiVersion value is not exist")
	}
	gs, ok := s[kv.APIVersion]
	if !ok {
		return nil, fmt.Errorf("apiVersion `%s` is not support", kv.APIVersion)
	}

	spec, ok := gs.objectMap[kv.Kind]
	if !ok {
		return nil, fmt.Errorf("kind value `%s` is not supported", kv.Kind)
	}
	obj, ok := spec.DeepCopyObject().(apis.Spec)
	if !ok {
		return nil, fmt.Errorf("kind value `%s` DeepCopyObject is invalid", kv.Kind)
	}
	if err := api.UnMarshalInput(bs, obj); err != nil {
		return nil, fmt.Errorf("failed to parse resource: %w", err)
	}
	return obj, nil
}
