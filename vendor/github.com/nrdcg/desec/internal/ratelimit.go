package internal

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// SmartClient is an Rate Limited HTTP Client.
type SmartClient struct {
	client *http.Client

	rlRead  *rate.Limiter
	rlWrite *rate.Limiter
}

// New Creates a new SmartClient.
func New(client *http.Client, nRead, nWrite int) *SmartClient {
	c := client
	if client == nil {
		c = http.DefaultClient
	}

	nr := nRead
	if nr < 1 {
		nr = 1
	}

	nw := nWrite
	if nw < 1 {
		nw = 1
	}

	return &SmartClient{
		client:  c,
		rlRead:  rate.NewLimiter(rate.Every(time.Second/time.Duration(nr)), 1),
		rlWrite: rate.NewLimiter(rate.Every(time.Second/time.Duration(nw)), 1),
	}
}

// Do sends an HTTP request and returns an HTTP response.
func (s *SmartClient) Do(req *http.Request) (*http.Response, error) {
	switch req.Method {
	// https://github.com/desec-io/desec-stack/blob/bc0b4de7dcfc53ab17ff7823846d214cfc9e7024/api/desecapi/views.py#L102
	// https://www.django-rest-framework.org/api-guide/permissions/#custom-permissions
	case http.MethodGet, http.MethodOptions, http.MethodHead:
		return s.do(req, s.rlRead)
	default:
		return s.do(req, s.rlWrite)
	}
}

func (s *SmartClient) do(req *http.Request, rl *rate.Limiter) (*http.Response, error) {
	err := rl.Wait(req.Context())
	if err != nil {
		return nil, err
	}

	return s.client.Do(req)
}
