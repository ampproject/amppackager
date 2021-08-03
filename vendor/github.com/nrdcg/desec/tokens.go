package desec

import (
	"fmt"
	"net/http"
	"time"
)

// Token a token representation.
type Token struct {
	ID      string     `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Value   string     `json:"token,omitempty"`
	Created *time.Time `json:"created,omitempty"`
}

// TokensService handles communication with the tokens related methods of the deSEC API.
//
// https://desec.readthedocs.io/en/latest/auth/tokens.html
type TokensService struct {
	client *Client
}

// GetAll retrieving all current tokens.
// https://desec.readthedocs.io/en/latest/auth/tokens.html#retrieving-all-current-tokens
func (s *TokensService) GetAll() ([]Token, error) {
	endpoint, err := s.client.createEndpoint("auth", "tokens")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodGet, endpoint, nil)
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

	var tokens []Token
	err = handleResponse(resp, &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

// Create creates additional tokens.
// https://desec.readthedocs.io/en/latest/auth/tokens.html#create-additional-tokens
func (s *TokensService) Create(name string) (*Token, error) {
	endpoint, err := s.client.createEndpoint("auth", "tokens")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, Token{Name: name})
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

	var token Token
	err = handleResponse(resp, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// Delete deletes tokens.
// https://desec.readthedocs.io/en/latest/auth/tokens.html#delete-tokens
func (s *TokensService) Delete(tokenID string) error {
	endpoint, err := s.client.createEndpoint("auth", "tokens", tokenID)
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodDelete, endpoint, nil)
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
