package porkbun

import (
	"encoding/json"
	"fmt"
)

type apiRequest interface{}

type authRequest struct {
	APIKey       string `json:"apikey"`
	SecretAPIKey string `json:"secretapikey"`
	apiRequest
}

func (f authRequest) MarshalJSON() ([]byte, error) {
	type clone authRequest
	c := clone(f)

	root, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	if c.apiRequest == nil {
		return root, nil
	}

	embedded, err := json.Marshal(c.apiRequest)
	if err != nil {
		return nil, err
	}

	return []byte(string(root[:len(root)-1]) + ",   " + string(embedded[1:])), nil
}

// Status the API response status.
type Status struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func (a Status) Error() string {
	return fmt.Sprintf("%s: %s", a.Status, a.Message)
}

// ServerError the API server error.
type ServerError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message,omitempty"`
}

func (a ServerError) Error() string {
	return fmt.Sprintf("status: %d message: %s", a.StatusCode, a.Message)
}

// Record a DNS record.
type Record struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	TTL     string `json:"ttl,omitempty"`
	Prio    string `json:"prio,omitempty"`
	Notes   string `json:"notes,omitempty"`
}

type pingResponse struct {
	Status
	YourIP string `json:"yourIp"`
}

type createResponse struct {
	Status
	ID int `json:"id"`
}

type retrieveResponse struct {
	Status
	Records []Record `json:"records"`
}
