// +build !dns01

package certfetcher

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/pkg/errors"
)

func DNSProvider(name string) (challenge.Provider, error) {
	return nil, errors.New("amppkg was built without DNS-01 support; please rebuild with `-tags dns01`")
}
