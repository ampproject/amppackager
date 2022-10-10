package auroradns

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// TokenTransport HTTP transport for API authentication.
type TokenTransport struct {
	apiKey string
	secret string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// NewTokenTransport Creates a  new TokenTransport.
func NewTokenTransport(apiKey, secret string) (*TokenTransport, error) {
	if apiKey == "" || secret == "" {
		return nil, errors.New("credentials missing")
	}

	return &TokenTransport{apiKey: apiKey, secret: secret}, nil
}

// RoundTrip executes a single HTTP transaction.
func (t *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	enrichedReq := &http.Request{}
	*enrichedReq = *req

	enrichedReq.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		enrichedReq.Header[k] = append([]string(nil), s...)
	}

	if t.apiKey != "" && t.secret != "" {
		timestamp := time.Now().UTC()

		fmtTime := timestamp.Format("20060102T150405Z")
		enrichedReq.Header.Set("X-AuroraDNS-Date", fmtTime)

		token, err := newToken(t.apiKey, t.secret, req.Method, req.URL.Path, timestamp)
		if err == nil {
			enrichedReq.Header.Set("Authorization", fmt.Sprintf("AuroraDNSv1 %s", token))
		}
	}

	return t.transport().RoundTrip(enrichedReq)
}

// Wrap Wraps an HTTP client Transport with the TokenTransport.
func (t *TokenTransport) Wrap(client *http.Client) *http.Client {
	backup := client.Transport
	t.Transport = backup
	client.Transport = t

	return client
}

// Client Creates a new HTTP client.
func (t *TokenTransport) Client() *http.Client {
	return &http.Client{
		Transport: t,
		Timeout:   30 * time.Second,
	}
}

func (t *TokenTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}

// newToken generates a token for accessing a specific method of the API.
func newToken(apiKey, secret, method, action string, timestamp time.Time) (string, error) {
	fmtTime := timestamp.Format("20060102T150405Z")
	message := method + action + fmtTime

	signatureHmac := hmac.New(sha256.New, []byte(secret))
	_, err := signatureHmac.Write([]byte(message))
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(signatureHmac.Sum(nil))

	apiKeyAndSignature := fmt.Sprintf("%s:%s", apiKey, signature)

	token := base64.StdEncoding.EncodeToString([]byte(apiKeyAndSignature))

	return token, nil
}
