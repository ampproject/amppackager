package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/account"
)

// APIKeysService handles 'account/apikeys' endpoint.
type APIKeysService service

// List returns all api keys in the account.
//
// NS1 API docs: https://ns1.com/api/#apikeys-get
func (s *APIKeysService) List() ([]*account.APIKey, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "account/apikeys", nil)
	if err != nil {
		return nil, nil, err
	}

	kl := []*account.APIKey{}
	resp, err := s.client.Do(req, &kl)
	if err != nil {
		return nil, resp, err
	}

	return kl, resp, nil
}

// Get returns details of an api key, including permissions, for a single API Key.
// Note: do not use the API Key itself as the keyid in the URL — use the id of the key.
//
// NS1 API docs: https://ns1.com/api/#apikeys-id-get
func (s *APIKeysService) Get(keyID string) (*account.APIKey, *http.Response, error) {
	path := fmt.Sprintf("account/apikeys/%s", keyID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var a account.APIKey
	resp, err := s.client.Do(req, &a)
	if err != nil {
		switch err.(type) {
		case *Error:
			if resourceMissingMatch(err.(*Error).Message) {
				return nil, resp, ErrKeyMissing
			}

		}
		return nil, resp, err
	}

	return &a, resp, nil
}

// Create takes a *APIKey and creates a new account apikey.
//
// NS1 API docs: https://ns1.com/api/#apikeys-put
func (s *APIKeysService) Create(a *account.APIKey) (*http.Response, error) {
	var (
		req *http.Request
		err error
	)

	// If this is DDI then the permissions need to be transformed to DDI-compatible permissions.
	if s.client.DDI && a != nil {
		ddiAPIKey := apiKeyToDDIAPIKey(a)
		req, err = s.client.NewRequest("PUT", "account/apikeys", ddiAPIKey)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = s.client.NewRequest("PUT", "account/apikeys", a)
		if err != nil {
			return nil, err
		}
	}

	// Update account fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &a)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == fmt.Sprintf("api key with name \"%s\" exists", a.Name) {
				return resp, ErrKeyExists
			}
		}
		return resp, err
	}

	return resp, nil
}

// Update changes the name or access rights for an API Key.
//
// NS1 API docs: https://ns1.com/api/#apikeys-id-post
func (s *APIKeysService) Update(a *account.APIKey) (*http.Response, error) {
	path := fmt.Sprintf("account/apikeys/%s", a.ID)

	var (
		req *http.Request
		err error
	)

	// If this is DDI then the permissions need to be transformed to DDI-compatible permissions.
	if s.client.DDI && a != nil {
		ddiAPIKey := apiKeyToDDIAPIKey(a)
		req, err = s.client.NewRequest("POST", path, ddiAPIKey)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = s.client.NewRequest("POST", path, a)
		if err != nil {
			return nil, err
		}
	}

	// Update apikey fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &a)
	if err != nil {
		switch err.(type) {
		case *Error:
			if resourceMissingMatch(err.(*Error).Message) {
				return resp, ErrKeyMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete deletes an apikey.
//
// NS1 API docs: https://ns1.com/api/#apikeys-id-delete
func (s *APIKeysService) Delete(keyID string) (*http.Response, error) {
	path := fmt.Sprintf("account/apikeys/%s", keyID)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch err.(type) {
		case *Error:
			if resourceMissingMatch(err.(*Error).Message) {
				return resp, ErrKeyMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

var (
	// ErrKeyExists bundles PUT create error.
	ErrKeyExists = errors.New("key already exists")
	// ErrKeyMissing bundles GET/POST/DELETE error.
	ErrKeyMissing = errors.New("key does not exist")
)

func apiKeyToDDIAPIKey(k *account.APIKey) *ddiAPIKey {
	ddiAPIKey := &ddiAPIKey{
		ID:                k.ID,
		Key:               k.Key,
		LastAccess:        k.LastAccess,
		Name:              k.Name,
		TeamIDs:           k.TeamIDs,
		IPWhitelist:       k.IPWhitelist,
		IPWhitelistStrict: k.IPWhitelistStrict,
		Permissions: ddiPermissionsMap{
			DNS:  k.Permissions.DNS,
			Data: k.Permissions.Data,
			Account: permissionsDDIAccount{
				ManageUsers:           k.Permissions.Account.ManageUsers,
				ManageTeams:           k.Permissions.Account.ManageTeams,
				ManageApikeys:         k.Permissions.Account.ManageApikeys,
				ManageAccountSettings: k.Permissions.Account.ManageAccountSettings,
				ViewActivityLog:       k.Permissions.Account.ViewActivityLog,
			},
		},
	}

	if k.Permissions.Security != nil {
		ddiAPIKey.Permissions.Security = permissionsDDISecurity(*k.Permissions.Security)
	}

	if k.Permissions.DHCP != nil {
		ddiAPIKey.Permissions.DHCP = *k.Permissions.DHCP
	}

	if k.Permissions.IPAM != nil {
		ddiAPIKey.Permissions.IPAM = *k.Permissions.IPAM
	}

	return ddiAPIKey
}
