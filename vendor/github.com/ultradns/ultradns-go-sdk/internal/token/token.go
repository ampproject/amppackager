package token

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

const tokenURL = "authorization/token"

type TokenSource struct {
	Ctx      context.Context
	BaseURL  string
	Username string
	Password string
	T        *oauth2.Token
}

func (ts *TokenSource) Token() (*oauth2.Token, error) {
	conf := &oauth2.Config{Endpoint: ts.getTokenEndpoint()}

	if ts.T == nil {
		return ts.PasswordCredentialsToken(conf)
	}

	token, err := conf.TokenSource(ts.Ctx, ts.T).Token()

	if err != nil {
		return ts.PasswordCredentialsToken(conf)
	}

	ts.T = token

	return token, err
}

func (ts *TokenSource) PasswordCredentialsToken(conf *oauth2.Config) (token *oauth2.Token, err error) {
	token, err = conf.PasswordCredentialsToken(ts.Ctx, ts.Username, ts.Password)
	ts.T = token

	return token, err
}

func (ts *TokenSource) getTokenEndpoint() oauth2.Endpoint {
	return oauth2.Endpoint{TokenURL: getTokenURL(ts.BaseURL)}
}

func getTokenURL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, tokenURL)
}
