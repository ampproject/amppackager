package v2

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRequestObj = errors.New("failed to build request")
	ErrNotFound          = errors.New("object not found")
)

type (
	BadResponseError struct {
		ErrorMsg    string `json:"error,omitempty"` //nolint: tagliatelle
		Description string `json:"description,omitempty"`
		Location    string `json:"location,omitempty"`
		Code        int    `json:"code"`
	}
)

func (e BadResponseError) Error() string {
	err := fmt.Sprintf("error response: %v.", e.ErrorMsg)
	if e.Description != "" {
		err += fmt.Sprintf(" Description: %v.", e.Description)
	}
	if e.Location != "" {
		err += fmt.Sprintf(" Location: %v.", e.Location)
	}

	return err
}
