package errors

import (
	"errors"
	"fmt"
)

var (
	errValidation     = errors.New("Missing required parameters")
	errService        = errors.New("service configuration failed")
	errResponseTarget = errors.New("Unexpected response type received")
	errResponse       = errors.New("Server error Response")
	errProcess        = errors.New("Error while")
	errTypeMismatch   = errors.New("Type mismatch error")
	errInvalid        = errors.New("Invalid input error")
	errNotFound       = errors.New("Resource not found")
	errMultipleData   = errors.New("Returned resources list instead single resource")
)

func ValidationError(key string) error {
	return fmt.Errorf("%w: [%s ]", errValidation, key)
}

func ServiceError(service string) error {
	return fmt.Errorf("%s %w", service, errService)
}

func ServiceConfigError(service string, err error) error {
	return fmt.Errorf("%s %w: %w", service, errService, err)
}

func ResponseTargetError(key string) error {
	return fmt.Errorf("%w: '%s'", errResponseTarget, key)
}

func APIResponseError(err string) error {
	return fmt.Errorf("%w - %s", errResponse, err)
}

func processError(service, process, id string, err error) error {
	return fmt.Errorf("%w %s %s: %w: {key: '%s'}", errProcess, process, service, err, id)
}

func CreateError(service, id string, err error) error {
	return processError(service, "creating", id, err)
}

func UpdateError(service, id string, err error) error {
	return processError(service, "updating", id, err)
}

func PartialUpdateError(service, id string, err error) error {
	return processError(service, "partial updating", id, err)
}

func ReadError(service, id string, err error) error {
	return processError(service, "reading", id, err)
}

func DeleteError(service, id string, err error) error {
	return processError(service, "deleting", id, err)
}

func ListError(service, uri string, err error) error {
	return processError(service, "listing", uri, err)
}

func MigrateError(service, uri string, err error) error {
	return processError(service, "migrating", uri, err)
}

func TypeMismatchError(expected, found string) error {
	return fmt.Errorf("%w: { expected: '%s', found: '%s' }", errTypeMismatch, expected, found)
}

func UnknownDataError(key, found string, data []string) error {
	return fmt.Errorf("%w: { key: '%s', value: '%s', valid_values: %s }", errInvalid, key, found, data)
}

func ResourceTypeNotFoundError(resourceName, resourceType, key string) error {
	return fmt.Errorf("%w: { name: '%s', type: '%s', key:'%s'}", errNotFound, resourceName, resourceType, key)
}

func MultipleResourceFoundError(service, id string) error {
	return processError(service, "reading", id, errMultipleData)
}
