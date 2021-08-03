// Package v2 is the new Exoscale client API binding.
// Reference: https://openapi-v2.exoscale.com/
package v2

import (
	"fmt"
	"reflect"
)

// resetFieldName returns the value corresponding to the `reset:""` struct tag
// in the struct res for the specified field, for example:
//
// type MyStruct struct {
//     FieldA string
//     FieldB int `reset:"field-b"`
// }
func resetFieldName(res, field interface{}) (string, error) {
	fieldValue := reflect.ValueOf(field)
	if fieldValue.Kind() != reflect.Ptr || fieldValue.IsNil() {
		return "", fmt.Errorf("field must be a non-nil pointer value")
	}

	structValue := reflect.ValueOf(res).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		structField := structValue.Type().Field(i)

		if structValue.Field(i).UnsafeAddr() == fieldValue.Pointer() {
			resetField, ok := structField.Tag.Lookup("reset")
			if !ok {
				return "", fmt.Errorf("struct field %s.%s is not resettable",
					structValue.Type(), structField.Name)
			}

			return resetField, nil
		}
	}

	return "", fmt.Errorf("field not found in struct %s", structValue.Type())
}
