// Package oapi is an internal package containing code generated from the
// Exoscale API OpenAPI specs, as well as helpers and transition types exposed
// in the public-facing package.
package oapi

import "context"

//go:generate oapi-codegen -generate types,client -package oapi -o oapi.gen.go source.json

type oapiClient interface {
	ClientWithResponsesInterface

	GetOperationWithResponse(context.Context, string, ...RequestEditorFn) (*GetOperationResponse, error)
}

// OptionalString returns the dereferenced string value of v if not nil, otherwise an empty string.
func OptionalString(v *string) string {
	if v != nil {
		return *v
	}

	return ""
}

// OptionalInt64 returns the dereferenced int64 value of v if not nil, otherwise 0.
func OptionalInt64(v *int64) int64 {
	if v != nil {
		return *v
	}

	return 0
}

// NilableString returns the input string pointer v if the dereferenced string is non-empty, otherwise nil.
// This helper is intended for use with OAPI types containing nilable string properties.
func NilableString(v *string) *string {
	if v != nil && *v == "" {
		return nil
	}

	return v
}
