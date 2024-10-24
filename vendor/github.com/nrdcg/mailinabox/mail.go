package mailinabox

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// MailUsers Represents a set of users for a domain.
type MailUsers struct {
	Domain string `json:"domain,omitempty"`
	Users  []User `json:"users,omitempty"`
}

// User Represents a user.
type User struct {
	Email      string   `json:"email,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
	Status     string   `json:"status,omitempty"`
	Mailbox    string   `json:"mailbox,omitempty"`
}

// MailAliases Represents a set of aliases for a domain.
type MailAliases struct {
	Domain  string  `json:"domain,omitempty"`
	Aliases []Alias `json:"aliases,omitempty"`
}

// Alias Represents an alias.
type Alias struct {
	Address          string   `json:"address,omitempty"`
	AddressDisplay   string   `json:"address_display,omitempty"`
	ForwardsTo       []string `json:"forwards_to,omitempty"`
	PermittedSenders []string `json:"permitted_senders,omitempty"`
	Required         bool     `json:"required,omitempty"`
}

// MailService Mail API.
// https://mailinabox.email/api-docs.html#tag/Mail
type MailService service

// GetUsers Returns all mail users.
// https://mailinabox.email/api-docs.html#operation/getMailUsers
func (s *MailService) GetUsers(ctx context.Context) ([]MailUsers, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "mail", "users")

	query := endpoint.Query()
	query.Set("format", "json")
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results []MailUsers

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// AddUser Adds a new mail user.
// https://mailinabox.email/api-docs.html#operation/addMailUser
func (s *MailService) AddUser(ctx context.Context, email, password, privilege string) (string, error) {
	if email == "" || password == "" || privilege == "" {
		return "", errors.New("email, password, and privilege are required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "mail", "users", "add")

	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)
	data.Set("privileges", privilege)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// RemoveUser Removes an existing mail user.
// https://mailinabox.email/api-docs.html#operation/removeMailUser
func (s *MailService) RemoveUser(ctx context.Context, email string) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "mail", "users", "remove")

	data := url.Values{}
	data.Set("email", email)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// AddUserPrivilege Adds a privilege to an existing mail user.
// https://mailinabox.email/api-docs.html#operation/addMailUserPrivilege
func (s *MailService) AddUserPrivilege(ctx context.Context, email, privilege string) (string, error) {
	if email == "" || privilege == "" {
		return "", errors.New("email and privilege are required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "mail", "users", "privileges", "add")

	data := url.Values{}
	data.Set("email", email)
	data.Set("privileges", privilege)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// RemoveUserPrivilege Adds a privilege to an existing mail user.
// https://mailinabox.email/api-docs.html#operation/removeMailUserPrivilege
func (s *MailService) RemoveUserPrivilege(ctx context.Context, email, privilege string) (string, error) {
	if email == "" {
		return "", errors.New("email is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "mail", "users", "privileges", "remove")

	data := url.Values{}
	data.Set("email", email)
	data.Set("privileges", privilege)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// SetUserPassword Sets a password for an existing mail user.
// https://mailinabox.email/api-docs.html#operation/setMailUserPassword
func (s *MailService) SetUserPassword(ctx context.Context, email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "mail", "users", "password")

	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetUserPrivileges Returns all privileges for an existing mail user.
// https://mailinabox.email/api-docs.html#operation/getMailUserPrivileges
func (s *MailService) GetUserPrivileges(ctx context.Context, email string) (string, error) {
	if email == "" {
		return "", errors.New("email is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "mail", "users", "privileges")

	query := endpoint.Query()
	query.Set("email", email)
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetDomains Returns all mail domains.
// https://mailinabox.email/api-docs.html#operation/getMailDomains
func (s *MailService) GetDomains(ctx context.Context) ([]string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "mail", "domains")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return nil, err
	}

	domains := strings.Split(strings.TrimSpace(string(resp)), "\n")

	return domains, nil
}

// GetAliases Returns all mail aliases.
// https://mailinabox.email/api-docs.html#operation/getMailAliases
func (s *MailService) GetAliases(ctx context.Context) ([]MailAliases, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "mail", "aliases")

	query := endpoint.Query()
	query.Set("format", "json")
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results []MailAliases

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// UpsertAlias Adds or updates a mail alias. If updating, you need to set update_if_exists: 1.
// https://mailinabox.email/api-docs.html#operation/upsertMailAlias
func (s *MailService) UpsertAlias(ctx context.Context, updateIfExists bool, address string, forwardsTo, permittedSenders []string) (string, error) {
	if address == "" {
		return "", errors.New("address is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "mail", "aliases", "add")

	data := url.Values{}
	data.Set("update_if_exists", boolToIntStr(updateIfExists))
	data.Set("address", address)
	data.Set("forwards_to", strings.Join(forwardsTo, ","))
	data.Set("permitted_senders", strings.Join(permittedSenders, ","))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// RemoveAliases Removes a mail alias.
// https://mailinabox.email/api-docs.html#operation/removeMailAlias
func (s *MailService) RemoveAliases(ctx context.Context, address string) (string, error) {
	if address == "" {
		return "", errors.New("address is required")
	}

	endpoint := s.client.baseURL.JoinPath("admin", "mail", "aliases", "remove")

	data := url.Values{}
	data.Set("address", address)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}
