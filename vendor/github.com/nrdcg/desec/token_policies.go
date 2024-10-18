package desec

import (
	"context"
	"fmt"
	"net/http"
)

// TokenPolicy represents a policy applied to a token.
type TokenPolicy struct {
	ID              string  `json:"id,omitempty"`
	Domain          *string `json:"domain"`
	SubName         *string `json:"subname"`
	Type            *string `json:"type"`
	WritePermission bool    `json:"perm_write,omitempty"`
}

// TokenPoliciesService handles communication with the token policy related methods of the deSEC API.
//
// https://desec.readthedocs.io/en/latest/auth/tokens.html
type TokenPoliciesService struct {
	client *Client
}

// Get retrieves token rrset's policies.
// https://desec.readthedocs.io/en/latest/auth/tokens.html#token-policy-management
func (s *TokenPoliciesService) Get(ctx context.Context, tokenID string) ([]TokenPolicy, error) {
	endpoint, err := s.client.createEndpoint("auth", "tokens", tokenID, "policies", "rrsets")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, handleError(resp)
	}

	var policies []TokenPolicy
	err = handleResponse(resp, &policies)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

// Create creates token policy.
// https://desec.readthedocs.io/en/latest/auth/tokens.html#create-additional-tokens
func (s *TokenPoliciesService) Create(ctx context.Context, tokenID string, policy TokenPolicy) (*TokenPolicy, error) {
	endpoint, err := s.client.createEndpoint("auth", "tokens", tokenID, "policies", "rrsets")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, endpoint, policy)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusCreated {
		return nil, handleError(resp)
	}

	var tokenPolicy TokenPolicy
	err = handleResponse(resp, &tokenPolicy)
	if err != nil {
		return nil, err
	}

	return &tokenPolicy, nil
}

// Delete deletes a token rrset's policy.
// https://desec.readthedocs.io/en/latest/auth/tokens.html#token-policy-management
func (s *TokenPoliciesService) Delete(ctx context.Context, tokenID, policyID string) error {
	endpoint, err := s.client.createEndpoint("auth", "tokens", tokenID, "policies", "rrsets", policyID)
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusNoContent {
		return handleError(resp)
	}
	return nil
}
