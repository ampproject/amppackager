// +build dns01

package certfetcher

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns"
)

func DNSProvider(name string) (challenge.Provider, error) {
	return dns.NewDNSChallengeProviderByName(name)
}
