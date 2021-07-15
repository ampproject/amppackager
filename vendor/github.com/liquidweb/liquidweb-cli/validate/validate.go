/*
Copyright © LiquidWeb

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package validate

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

func Validate(chk map[interface{}]interface{}) (err error) {
	defer func() {
		if paniced := recover(); paniced != nil {
			err = fmt.Errorf("%w %s", ValidationFailure, paniced)
		}
	}()

	for inputFieldValue, inputField := range chk {

		inputFieldVal := reflect.ValueOf(inputField)

		// by default, assume input field passed by user cannot be empty
		inputFieldOptional := false
		if inputFieldVal.Kind() == reflect.Map {
			iface := inputFieldVal.Interface()
			inputFieldType := iface.(map[string]string)["type"]
			if iface.(map[string]string)["optional"] == "true" {
				inputFieldOptional = true
			}
			inputFieldVal = reflect.ValueOf(inputFieldType)
		}

		inputFieldStr := cast.ToString(inputFieldVal)

		// inputField must be defined in InputTypes struct
		defined, shouldBeType, fieldVal := inputTypeDefined(inputFieldStr)
		if !defined {
			err = fmt.Errorf("%w for input field [%+v] type [%s] is not valid", ValidationFailure,
				inputFieldValue, inputFieldStr)
			return
		}

		// inputFieldValue must be of the correct type
		reflectValue := reflect.TypeOf(inputFieldValue).Name()
		if reflectValue != shouldBeType {
			err = fmt.Errorf("%w for input field [%+v] type [%s] has an invalid type of [%s] wanted [%s]",
				ValidationFailure, inputFieldValue, inputFieldStr, reflectValue, shouldBeType)
			return
		}

		// if the input field wasn't passed, and allow optional is true, continue
		if inputFieldOptional {
			// if inputFieldValue is a zero value, return without error
			inputFieldValueVal := reflect.ValueOf(inputFieldValue)

			if inputFieldStr == "NonEmptyString" {
				// since we check by going by the zero value for the type if the input field was passed,
				// we can't enforce the string type to not be empty optionally for NonEmptyString. Since
				// we can't differentiate between not passed and its zero value.
				err = fmt.Errorf("NonEmptyString input fields cannot be optional")
				return
			}

			if inputFieldValueVal.IsZero() {
				continue
			}
		}

		// if there's a Validate method call it
		iface := fieldVal.Interface()
		if interfaceHasMethod(iface, "Validate") {
			if validateErr := interfaceInputTypeValidate(iface, inputFieldValue); validateErr != nil {
				err = fmt.Errorf("%w for input field [%+v] %s", ValidationFailure, inputFieldValue,
					validateErr)
				return
			}
		}

	}

	return
}

func interfaceInputTypeValidate(iface, inputFieldValue interface{}) error {
	switch iface.(type) {
	case InputTypeUniqId:
		var obj InputTypeUniqId
		obj.UniqId = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeIP:
		var obj InputTypeIP
		obj.IP = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypePositiveInt64:
		var obj InputTypePositiveInt64
		obj.PositiveInt64 = cast.ToInt64(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypePositiveInt:
		var obj InputTypePositiveInt
		obj.PositiveInt = cast.ToInt(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeNonEmptyString:
		var obj InputTypeNonEmptyString
		obj.NonEmptyString = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeLoadBalancerStrategyString:
		var obj InputTypeLoadBalancerStrategyString
		obj.LoadBalancerStrategy = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeHttpsLiquidwebUrl:
		var obj InputTypeHttpsLiquidwebUrl
		obj.HttpsLiquidwebUrl = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeNetworkPortPair:
		var obj InputTypeNetworkPortPair
		obj.NetworkPortPair = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeNetworkPort:
		var obj InputTypeNetworkPort
		obj.NetworkPort = cast.ToInt(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeLoadBalancerHealthCheckProtocol:
		var obj InputTypeLoadBalancerHealthCheckProtocol
		obj.LoadBalancerHealthCheckProtocol = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	case InputTypeLoadBalancerHttpCodeRange:
		var obj InputTypeLoadBalancerHttpCodeRange
		obj.LoadBalancerHttpCodeRange = cast.ToString(inputFieldValue)
		if err := obj.Validate(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("bug: validation missing entry for %s", inputFieldValue)
	}

	return nil
}

func interfaceHasMethod(iface interface{}, methodName string) bool {
	ifaceVal := reflect.ValueOf(iface)

	if !ifaceVal.IsValid() {
		// not valid, so we already know its false
		return false
	}

	if ifaceVal.Type().Kind() != reflect.Ptr {
		ifaceVal = reflect.New(reflect.TypeOf(iface))
	}

	method := ifaceVal.MethodByName(methodName)

	if method.IsValid() {
		return true
	}

	return false
}

func inputTypeDefined(inputType string) (bool, string, reflect.Value) {
	var validTypes InputTypes

	err, fieldType, fieldVal := structHasField(validTypes, inputType)
	if err != nil {
		return false, fieldType, fieldVal
	}

	return true, fieldType, fieldVal
}

func structHasField(data interface{}, fieldName string) (error, string, reflect.Value) {
	dataVal := reflect.ValueOf(data)

	if !dataVal.IsValid() {
		return fmt.Errorf("failed fetching value for fieldName [%s]", fieldName), "",
			reflect.Value{}
	}

	if dataVal.Type().Kind() != reflect.Ptr {
		dataVal = reflect.New(reflect.TypeOf(data))
	}

	fieldVal := dataVal.Elem().FieldByName(fieldName)
	if !fieldVal.IsValid() {
		return fmt.Errorf("[%s] has no field [%s]", dataVal.Type(), fieldName), "", fieldVal
	}

	fieldValKindStr := fieldVal.Kind().String()

	if fieldValKindStr == "struct" {
		fieldValKindStr = fieldVal.Field(0).Kind().String()
	}

	return nil, fieldValKindStr, fieldVal
}
