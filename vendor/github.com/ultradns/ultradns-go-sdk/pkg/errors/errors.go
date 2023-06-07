package errors

import (
	"errors"
	"fmt"
)

var (
	errValidation     = errors.New("config validation failure")
	errService        = errors.New("service is not properly configured")
	errResponseTarget = errors.New("response target type mismatched : returned type")
	errResponse       = errors.New("error from api response -")
	errTypeMismatch   = errors.New("type mismatched")
	errUnknownData    = errors.New("should be any of the following data")
	errNotFound       = errors.New("not found")
)

func ValidationError(key string) error {
	return fmt.Errorf("%w: %s is missing", errValidation, key)
}

func ServiceError(service string) error {
	return fmt.Errorf("%s %w", service, errService)
}

func ServiceConfigError(service string, err error) error {
	return fmt.Errorf("config error while creating %s service : %w", service, err)
}

func ResponseTargetError(key string) error {
	return fmt.Errorf("%w - %s", errResponseTarget, key)
}

func APIResponseError(err string) error {
	return fmt.Errorf("%w %s", errResponse, err)
}

func TypeMismatchError(expected, found string) error {
	return fmt.Errorf("%w : expected - %s : found - %s", errTypeMismatch, expected, found)
}

func UnknownDataError(key, found string, data []string) error {
	return fmt.Errorf("%s %w %s : found - %s", key, errUnknownData, data, found)
}

func CreateError(service, id string, err error) error {
	return fmt.Errorf("error while creating %s - %s : %w", service, id, err)
}

func UpdateError(service, id string, err error) error {
	return fmt.Errorf("error while updating %s - %s : %w", service, id, err)
}

func PartialUpdateError(service, id string, err error) error {
	return fmt.Errorf("error while partial updating %s - %s : %w", service, id, err)
}

func ReadError(service, id string, err error) error {
	return fmt.Errorf("error while reading %s - %s : %w", service, id, err)
}

func DeleteError(service, id string, err error) error {
	return fmt.Errorf("error while deleting %s - %s : %w", service, id, err)
}

func ListError(service, uri string, err error) error {
	return fmt.Errorf("error while listing %s : uri - %s : %w", service, uri, err)
}

func ResourceTypeNotFoundError(resourceName, resourceType, key string) error {
	return fmt.Errorf("%s resource of type %s - %s %w", resourceName, resourceType, key, errNotFound)
}
