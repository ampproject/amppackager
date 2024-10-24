package mailinabox

import (
	"context"
	"fmt"
	"net/http"
)

// Session Represents user session data.
type Session struct {
	APIKey     string   `json:"api_key,omitempty"`
	Email      string   `json:"email,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
	Status     string   `json:"status,omitempty"`
	Reason     string   `json:"reason,omitempty"`
}

// UserService User API.
// https://mailinabox.email/api-docs.html#tag/User
type UserService service

// Login Returns user information and a session API key.
// https://mailinabox.email/api-docs.html#operation/login
func (s *UserService) Login(ctx context.Context) (*Session, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "login")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var result Session

	err = s.client.doJSON(req, &result)
	if err != nil {
		return nil, err
	}

	if result.Status != "ok" {
		return nil, fmt.Errorf("%s: %s", result.Status, result.Reason)
	}

	return &result, nil
}

// Logout Invalidates a session API key so that it cannot be used after this API call.
// https://mailinabox.email/api-docs.html#operation/logout
func (s *UserService) Logout(ctx context.Context) (*Session, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "logout")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var result Session

	err = s.client.doJSON(req, &result)
	if err != nil {
		return nil, err
	}

	if result.Status != "ok" {
		return nil, fmt.Errorf("%s: %s", result.Status, result.Reason)
	}

	return &result, nil
}
