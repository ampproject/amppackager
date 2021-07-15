package desec

import (
	"fmt"
	"net/http"
	"time"
)

// Account an account representation.
type Account struct {
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	LimitDomains int        `json:"limit_domains,omitempty"`
	Created      *time.Time `json:"created,omitempty"`
}

// Captcha a captcha representation.
type Captcha struct {
	ID        string `json:"id,omitempty"`
	Challenge string `json:"challenge,omitempty"`
	Solution  string `json:"solution,omitempty"`
}

// Registration a registration representation.
type Registration struct {
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	NewEmail string   `json:"new_email,omitempty"`
	Captcha  *Captcha `json:"captcha,omitempty"`
}

// AccountService handles communication with the account related methods of the deSEC API.
//
// https://desec.readthedocs.io/en/latest/auth/account.html
type AccountService struct {
	client *Client
}

// Login Log in.
// https://desec.readthedocs.io/en/latest/auth/account.html#log-in
func (s *AccountService) Login(email, password string) (*Token, error) {
	endpoint, err := s.client.createEndpoint("auth", "login")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, Account{Email: email, Password: password})
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

	var token Token
	err = handleResponse(resp, &token)
	if err != nil {
		return nil, err
	}

	s.client.token = token.Value

	return &token, nil
}

// Logout log out (= delete current token).
// https://desec.readthedocs.io/en/latest/auth/account.html#log-out
func (s *AccountService) Logout() error {
	endpoint, err := s.client.createEndpoint("auth", "logout")
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, nil)
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

	s.client.token = ""

	return nil
}

// ObtainCaptcha Obtain a captcha.
// https://desec.readthedocs.io/en/latest/auth/account.html#obtain-a-captcha
func (s *AccountService) ObtainCaptcha() (*Captcha, error) {
	endpoint, err := s.client.createEndpoint("captcha")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, nil)
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

	var captcha Captcha
	err = handleResponse(resp, &captcha)
	if err != nil {
		return nil, err
	}

	return &captcha, nil
}

// Register register account.
// https://desec.readthedocs.io/en/latest/auth/account.html#register-account
func (s *AccountService) Register(registration Registration) error {
	endpoint, err := s.client.createEndpoint("auth")
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, registration)
	if err != nil {
		return err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusAccepted {
		return handleError(resp)
	}

	return nil
}

// RetrieveInformation retrieve account information.
// https://desec.readthedocs.io/en/latest/auth/account.html#retrieve-account-information
func (s *AccountService) RetrieveInformation() (*Account, error) {
	endpoint, err := s.client.createEndpoint("auth", "account")
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, nil)
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

	var account Account
	err = handleResponse(resp, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// PasswordReset password reset and password change.
// https://desec.readthedocs.io/en/latest/auth/account.html#password-reset
// https://desec.readthedocs.io/en/latest/auth/account.html#password-change
func (s *AccountService) PasswordReset(email string, captcha Captcha) error {
	endpoint, err := s.client.createEndpoint("auth", "account", "reset-password")
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, Registration{Email: email, Captcha: &captcha})
	if err != nil {
		return err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusAccepted {
		return handleError(resp)
	}

	return nil
}

// ChangeEmail changes email address.
// https://desec.readthedocs.io/en/latest/auth/account.html#change-email-address
func (s *AccountService) ChangeEmail(email, password, newEmail string) error {
	endpoint, err := s.client.createEndpoint("auth", "account", "change-email")
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, Registration{Email: email, Password: password, NewEmail: newEmail})
	if err != nil {
		return err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusAccepted {
		return handleError(resp)
	}

	return nil
}

// Delete deletes account.
// https://desec.readthedocs.io/en/latest/auth/account.html#delete-account
func (s *AccountService) Delete(email, password string) error {
	endpoint, err := s.client.createEndpoint("auth", "account", "delete")
	if err != nil {
		return fmt.Errorf("failed to create endpoint: %w", err)
	}

	req, err := s.client.newRequest(http.MethodPost, endpoint, Account{Email: email, Password: password})
	if err != nil {
		return err
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusAccepted {
		return handleError(resp)
	}

	return nil
}
