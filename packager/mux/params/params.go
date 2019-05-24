// A separate module to avoid import cycles between mux and the various
// handlers.

package params

import (
	"context"
	"net/http"
)

type paramsKeyType struct{}
var paramsKey = paramsKeyType{}

// Gets the params from the request context, injected by the mux. Guaranteed to
// be non-nil.
func Params(req *http.Request) map[string]string {
	params := req.Context().Value(paramsKey)
	switch v := params.(type) {
	case map[string]string:
		return v
	default:
		// This should never happen, but just in case, let's not panic.
		return map[string]string{}
	}
}

// Returns a copy of req annotated with the given params. (Params is stored by
// reference, and so may be mutated afterwards.) To be called only by this
// library itself or by tests.
func WithParams(req *http.Request, params map[string]string) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), paramsKey, params))
}
