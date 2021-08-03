// +build dns01

package certfetcher

import (
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns"
)

func DNSProvider(name string) (challenge.Provider, error) {
	return dns.NewDNSChallengeProviderByName(name)
}
