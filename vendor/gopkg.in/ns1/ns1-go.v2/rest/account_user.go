package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/account"
)

// UsersService handles 'account/users' endpoint.
type UsersService service

// List returns all users in the account.
//
// NS1 API docs: https://ns1.com/api/#users-get
func (s *UsersService) List() ([]*account.User, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "account/users", nil)
	if err != nil {
		return nil, nil, err
	}

	ul := []*account.User{}
	resp, err := s.client.Do(req, &ul)
	if err != nil {
		return nil, resp, err
	}

	return ul, resp, nil
}

// Get returns details of a single user.
//
// NS1 API docs: https://ns1.com/api/#users-user-get
func (s *UsersService) Get(username string) (*account.User, *http.Response, error) {
	path := fmt.Sprintf("account/users/%s", username)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var u account.User
	resp, err := s.client.Do(req, &u)
	if err != nil {
		switch err.(type) {
		case *Error:
			if resourceMissingMatch(err.(*Error).Message) {
				return nil, resp, ErrUserMissing
			}
		}
		return nil, resp, err
	}

	return &u, resp, nil
}

// Create takes a *User and creates a new account user.
//
// NS1 API docs: https://ns1.com/api/#users-put
func (s *UsersService) Create(u *account.User) (*http.Response, error) {
	var (
		req *http.Request
		err error
	)

	// If this is DDI then the permissions need to be transformed to DDI-compatible permissions.
	if s.client.DDI && u != nil {
		ddiUser := userToDDIUser(u)
		req, err = s.client.NewRequest("PUT", "account/users", ddiUser)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = s.client.NewRequest("PUT", "account/users", u)
		if err != nil {
			return nil, err
		}
	}

	// Update user fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &u)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "request failed:Login Name is already in use." {
				return resp, ErrUserExists
			}
		}
		return resp, err
	}

	return resp, nil
}

// Update change contact details, notification settings, or access rights for a user.
//
// NS1 API docs: https://ns1.com/api/#users-user-post
func (s *UsersService) Update(u *account.User) (*http.Response, error) {
	path := fmt.Sprintf("account/users/%s", u.Username)

	var (
		req *http.Request
		err error
	)

	// If this is DDI then the permissions need to be transformed to DDI-compatible permissions.
	if s.client.DDI && u != nil {
		ddiUser := userToDDIUser(u)
		req, err = s.client.NewRequest("POST", path, ddiUser)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = s.client.NewRequest("POST", path, u)
		if err != nil {
			return nil, err
		}
	}

	// Update user fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &u)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "Unknown user" {
				return resp, ErrUserMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete deletes a user.
//
// NS1 API docs: https://ns1.com/api/#users-user-delete
func (s *UsersService) Delete(username string) (*http.Response, error) {
	path := fmt.Sprintf("account/users/%s", username)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "Unknown user" {
				return resp, ErrUserMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

var (
	// ErrUserExists bundles PUT create error.
	ErrUserExists = errors.New("user already exists")
	// ErrUserMissing bundles GET/POST/DELETE error.
	ErrUserMissing = errors.New("user does not exist")
)

func userToDDIUser(u *account.User) *ddiUser {
	ddiUser := &ddiUser{
		LastAccess:        u.LastAccess,
		Name:              u.Name,
		Username:          u.Username,
		Email:             u.Email,
		TeamIDs:           u.TeamIDs,
		Notify:            u.Notify,
		IPWhitelist:       u.IPWhitelist,
		IPWhitelistStrict: u.IPWhitelistStrict,
		Permissions: ddiPermissionsMap{
			DNS:  u.Permissions.DNS,
			Data: u.Permissions.Data,
			Account: permissionsDDIAccount{
				ManageUsers:           u.Permissions.Account.ManageUsers,
				ManageTeams:           u.Permissions.Account.ManageTeams,
				ManageApikeys:         u.Permissions.Account.ManageApikeys,
				ManageAccountSettings: u.Permissions.Account.ManageAccountSettings,
				ViewActivityLog:       u.Permissions.Account.ViewActivityLog,
			},
		},
	}

	if u.Permissions.Security != nil {
		ddiUser.Permissions.Security = permissionsDDISecurity(*u.Permissions.Security)
	}

	if u.Permissions.DHCP != nil {
		ddiUser.Permissions.DHCP = *u.Permissions.DHCP
	}

	if u.Permissions.IPAM != nil {
		ddiUser.Permissions.IPAM = *u.Permissions.IPAM
	}

	return ddiUser
}
