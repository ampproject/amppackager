package types

import (
	"bytes"
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

// 0 is null
type NullablePositiveInt32 int32

func (t *NullablePositiveInt32) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, []byte("null")) {
		return nil
	}
	var i32 int32
	if err := json.Unmarshal(bs, &i32); err != nil {
		return err
	}
	*t = NullablePositiveInt32(i32)
	return nil
}

func (t NullablePositiveInt32) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	return jsoniter.Marshal(int32(t))
}

type NullablePositiveInt64 int64

func (t *NullablePositiveInt64) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, []byte("null")) {
		return nil
	}
	var i64 int64
	if err := json.Unmarshal(bs, &i64); err != nil {
		return err
	}
	*t = NullablePositiveInt64(i64)
	return nil
}

func (t NullablePositiveInt64) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	return jsoniter.Marshal(int64(t))
}
