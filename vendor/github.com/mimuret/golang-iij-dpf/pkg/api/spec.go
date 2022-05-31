package api

import (
	"reflect"
)

type Object interface {
	DeepCopyObject() Object
}

type Spec interface {
	Object
	GetName() string
	GetGroup() string
	GetPathMethod(Action) (string, string)
}

type ListSpec interface {
	Spec
	Initializer
	GetItems() interface{}
	Len() int
	Index(int) interface{}
}

type CountableListSpec interface {
	ListSpec
	SetCount(int32)
	GetCount() int32
	GetMaxLimit() int32
	AddItem(interface{}) bool
	ClearItems()
}

type Initializer interface {
	Init()
}

func DeepCopySpec(s Spec) Spec {
	if s == nil || reflect.ValueOf(s).IsNil() {
		return nil
	}
	ret, ok := s.DeepCopyObject().(Spec)
	if !ok {
		panic("s is not Spec")
	}
	return ret
}

func DeepCopyListSpec(s ListSpec) ListSpec {
	if s == nil || reflect.ValueOf(s).IsNil() {
		return nil
	}
	ret, ok := s.DeepCopyObject().(ListSpec)
	if !ok {
		panic("s is not ListSpec")
	}
	return ret
}

func DeepCopyCountableListSpec(s CountableListSpec) CountableListSpec {
	if s == nil || reflect.ValueOf(s).IsNil() {
		return nil
	}
	ret, ok := s.DeepCopyObject().(CountableListSpec)
	if !ok {
		panic("s is not CountableListSpec")
	}
	return ret
}
