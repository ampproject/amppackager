package utils

import (
	"fmt"
)

func ToInt64(val interface{}) (int64, error) {
	var res int64

	switch v := val.(type) {
	case int:
		res = int64(v)
	case int8:
		res = int64(v)
	case int16:
		res = int64(v)
	case int32:
		res = int64(v)
	case int64:
		res = v
	default:
		return 0, fmt.Errorf("failed to convert int64")
	}
	return res, nil
}

func ToString(val interface{}) (string, error) {
	var res string
	switch v := val.(type) {
	case string:
		res = v
	default:
		return "", fmt.Errorf("failed to convert int64")
	}
	return res, nil
}
