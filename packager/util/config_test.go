package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func errorFrom(_ *Config, err error) string {
	return err.Error()
}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func TestInvalid(t *testing.T) {
	assert.NotEqual(t, "", errorFrom(ReadConfig([]byte(``))))
	assert.Contains(t, errorFrom(ReadConfig([]byte(`abc`))), "failed to parse TOML")
	assert.Contains(t, errorFrom(ReadConfig([]byte(`[Port] X=5`))), "failed to unmarshal TOML")
}

func TestMinimalValidConfig(t *testing.T) {
	config, err := ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
	`))
	require.NoError(t, err)
	assert.Equal(t, Config{
		Port:      8080,
		CertFile:  "cert.pem",
		KeyFile:   "key.pem",
		OCSPCache: "/tmp/ocsp",
		URLSet: []URLSet{{
			Sign: &URLPattern{
				Domain:  "example.com",
				PathRE:  stringPtr(".*"),
				QueryRE: stringPtr(""),
				MaxLength: 2000,
			},
		}},
	}, *config)
}

func TestOCSPDirDoesntExist(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/this/does/not/exist/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
	`))), "OCSPCache parent directory must exist")
}

func TestInvalidPathRE(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		    PathRE = "["
	`))), "PathRE must be a valid regexp")
}

func TestInvalidPathExcludeRE(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		    PathExcludeRE = ["["]
	`))), "PathExcludeRE contains invalid regexp")
}

func TestInvalidQueryRE(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		    QueryRE = "["
	`))), "QueryRE must be a valid regexp")
}

func TestSignMissing(t *testing.T) {
	msg := errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
	`)))
	assert.Contains(t, msg, "This section must be specified")
	assert.Contains(t, msg, "URLSet.0.Sign")
}

func TestSignScheme(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		    Scheme = ["http"]
	`))), "Scheme not allowed here")
}

func TestSignDomainRE(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		    DomainRE = ".*"
	`))), "DomainRE not allowed here")
}

func TestSignSamePath(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		    SamePath = true
	`))), "SamePath not allowed here")
}

func TestSignOverrides(t *testing.T) {
	config, err := ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		    PathRE = "/amp/.*"
		    PathExcludeRE = ["/amp/signin", "/amp/settings(/.*)?"]
		    QueryRE = ""
		    ErrorOnStatefulHeaders = true
		    MaxLength = 8000
	`))
	require.NoError(t, err)
	require.Equal(t, 1, len(config.URLSet))
	// TODO(twifkak): Don't depend on scheme order.
	assert.Equal(t, URLPattern{
		Domain:                 "example.com",
		PathRE:                 stringPtr("/amp/.*"),
		PathExcludeRE:          []string{"/amp/signin", "/amp/settings(/.*)?"},
		QueryRE:                stringPtr(""),
		ErrorOnStatefulHeaders: true,
		MaxLength:              8000,
	}, *config.URLSet[0].Sign)
}

func TestFetchDefaults(t *testing.T) {
	config, err := ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		  [URLSet.Fetch]
		    Domain = "example.com"
	`))
	require.NoError(t, err)
	require.Equal(t, 1, len(config.URLSet))
	fetch := *config.URLSet[0].Fetch
	assert.ElementsMatch(t, []string{"http", "https"}, fetch.Scheme)
	assert.Equal(t, "", fetch.DomainRE)
	assert.Equal(t, "example.com", fetch.Domain)
	assert.Equal(t, stringPtr(".*"), fetch.PathRE)
	assert.Nil(t, fetch.PathExcludeRE)
	assert.Equal(t, stringPtr(""), fetch.QueryRE)
	assert.Equal(t, false, fetch.ErrorOnStatefulHeaders)
	assert.Equal(t, 2000, fetch.MaxLength)
	assert.Equal(t, boolPtr(true), fetch.SamePath)
}

func TestFetchOverrides(t *testing.T) {
	config, err := ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Sign]
		    Domain = "example.com"
		  [URLSet.Fetch]
		    Scheme = ["http"]
		    DomainRE = ".*"
		    PathRE = "/amp/.*"
		    PathExcludeRE = ["/amp/signin", "/amp/settings(/.*)?"]
		    QueryRE = ""
		    MaxLength = 8000
		    SamePath = false
	`))
	require.NoError(t, err)
	require.Equal(t, 1, len(config.URLSet))
	assert.Equal(t, URLPattern{
		Scheme:        []string{"http"},
		DomainRE:      ".*",
		PathRE:        stringPtr("/amp/.*"),
		PathExcludeRE: []string{"/amp/signin", "/amp/settings(/.*)?"},
		QueryRE:       stringPtr(""),
		MaxLength:     8000,
		SamePath:      boolPtr(false),
	}, *config.URLSet[0].Fetch)
}

func TestFetchInvalidScheme(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Fetch]
		    Scheme = ["gopher"]
	`))), "Scheme contains invalid value")
}

func TestFetchNoDomain(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Fetch]
	`))), "Domain or DomainRE must be specified")
}

func TestFetchDomainAndDomainRE(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Fetch]
		    Domain = "example.com"
		    DomainRE = "example.com"
	`))), "Only one of Domain or DomainRE")
}

func TestFetchErrorOnStatefulHeaders(t *testing.T) {
	assert.Contains(t, errorFrom(ReadConfig([]byte(`
		CertFile = "cert.pem"
		KeyFile = "key.pem"
		OCSPCache = "/tmp/ocsp"
		[[URLSet]]
		  [URLSet.Fetch]
		    Domain = "example.com"
		    ErrorOnStatefulHeaders = true
	`))), "ErrorOnStatefulHeaders not allowed")
}
