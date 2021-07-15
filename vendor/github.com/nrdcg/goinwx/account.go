package goinwx

import "github.com/mitchellh/mapstructure"

const (
	methodAccountLogin  = "account.login"
	methodAccountLogout = "account.logout"
	methodAccountLock   = "account.lock"
	methodAccountUnlock = "account.unlock"
)

// AccountService API access to Account.
type AccountService service

// Login Account login.
func (s *AccountService) Login() (*LoginResponse, error) {
	req := s.client.NewRequest(methodAccountLogin, map[string]interface{}{
		"user": s.client.username,
		"pass": s.client.password,
	})

	resp, err := s.client.Do(*req)
	if err != nil {
		return nil, err
	}

	var result LoginResponse
	err = mapstructure.Decode(*resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Logout Account logout.
func (s *AccountService) Logout() error {
	req := s.client.NewRequest(methodAccountLogout, nil)

	_, err := s.client.Do(*req)
	return err
}

// Lock Account lock.
func (s *AccountService) Lock() error {
	req := s.client.NewRequest(methodAccountLock, nil)

	_, err := s.client.Do(*req)
	return err
}

// Unlock Account unlock.
func (s *AccountService) Unlock(tan string) error {
	req := s.client.NewRequest(methodAccountUnlock, map[string]interface{}{
		"tan": tan,
	})

	_, err := s.client.Do(*req)
	return err
}

// LoginResponse API model.
type LoginResponse struct {
	CustomerID int64  `mapstructure:"customerId"`
	AccountID  int64  `mapstructure:"accountId"`
	TFA        string `mapstructure:"tfa"`
	BuildDate  string `mapstructure:"builddate"`
	Version    string `mapstructure:"version"`
}
