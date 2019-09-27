// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package certloader

import (
	"crypto"
	"crypto/x509"
        "encoding/pem"
	"io/ioutil"
	"log"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/pkg/errors"

	"github.com/ampproject/amppackager/packager/certcache"
	"github.com/ampproject/amppackager/packager/certfetcher"
	"github.com/ampproject/amppackager/packager/util"
)

// Creates cert cache by loading certs and keys from disk, doing validation
// and populating the cert cache with current set of certificate related information.
// If development mode is true, prints a warning for certs that can't sign HTTP exchanges.
func PopulateCertCache(config *util.Config, key crypto.PrivateKey,
	developmentMode bool, autoRenewCert bool) (*certcache.CertCache, error) {

	certs, err := loadCertsFromFile(config, developmentMode)
	if err != nil {
		log.Println(errors.Wrap(err, "Can't load cert file."))
		certs = nil
	}
	domain := ""
	for _, urlSet := range config.URLSet {
		domain = urlSet.Sign.Domain
		if certs != nil {
			if err := util.CertificateMatches(certs[0], key, domain); err != nil {
				return nil, errors.Wrapf(err, "checking %s", config.CertFile)
			}
		}
	}

	certFetcher, err := createCertFetcher(config, key, domain, developmentMode, autoRenewCert)
	if err != nil {
		return nil, errors.Wrap(err, "creating cert fetcher from config.")
	}
	certCache := certcache.New(certs, certFetcher, config.OCSPCache)

	return certCache, nil
}

func createCertFetcher(config *util.Config, key crypto.PrivateKey, domain string,
	developmentMode bool, autoRenewCert bool) (*certfetcher.CertFetcher, error) {
	if autoRenewCert {
		if config.ACMEConfig == nil {
			return nil, errors.New("missing ACMEConfig")
		}
		if config.ACMEConfig.Production == nil {
			return nil, errors.New("missing ACMEConfig.Production")
		}
		if config.ACMEConfig.Production.EmailAddress == "" {
			return nil, errors.New("missing email address")
		}
		emailAddress := config.ACMEConfig.Production.EmailAddress
		if config.ACMEConfig.Production.DiscoURL == "" {
			return nil, errors.New("missing acme disco url")
		}
		acmeDiscoveryURL := config.ACMEConfig.Production.DiscoURL
		if (config.ACMEConfig.Production.HttpChallengePort == 0 &&
			config.ACMEConfig.Production.HttpWebRootDir == "" &&
			config.ACMEConfig.Production.TlsChallengePort == 0 &&
			config.ACMEConfig.Production.DnsProvider == "") {
			return nil, errors.New("One of HttpChallengePort, HttpWebRootDir, TlsChallengePort and DnsProvider must be present.")
		}
		httpChallengePort := config.ACMEConfig.Production.HttpChallengePort
		httpWebRootDir := config.ACMEConfig.Production.HttpWebRootDir
		tlsChallengePort := config.ACMEConfig.Production.TlsChallengePort
		dnsProvider := config.ACMEConfig.Production.DnsProvider
		if developmentMode {
			if config.ACMEConfig.Development == nil {
				return nil, errors.New("missing ACMEConfig.Development")
			}
			if config.ACMEConfig.Development.EmailAddress == "" {
				return nil, errors.New("missing email address")
			}
			emailAddress = config.ACMEConfig.Development.EmailAddress

			if config.ACMEConfig.Development.DiscoURL == "" {
				return nil, errors.New("missing acme disco url")
			}
			acmeDiscoveryURL = config.ACMEConfig.Development.DiscoURL

			if (config.ACMEConfig.Development.HttpChallengePort == 0 &&
				config.ACMEConfig.Development.HttpWebRootDir == "" &&
				config.ACMEConfig.Development.TlsChallengePort == 0 &&
				config.ACMEConfig.Development.DnsProvider == "") {
				return nil, errors.New("One of HttpChallengePort, HttpWebRootDir, TlsChallengePort and DnsProvider must be present.")
			}
			httpChallengePort = config.ACMEConfig.Development.HttpChallengePort
			httpWebRootDir = config.ACMEConfig.Development.HttpWebRootDir
			tlsChallengePort = config.ACMEConfig.Development.TlsChallengePort
			dnsProvider = config.ACMEConfig.Development.DnsProvider
		}

		csr, err := loadCSRFromFile(config)
		if err != nil {
			return nil, errors.Wrap(err, "missing CSR")
		}

		// Create the cert fetcher that will auto-renew the cert.
		certFetcher, err := certfetcher.NewFetcher(emailAddress, csr, key, acmeDiscoveryURL,
			[]string{domain}, httpChallengePort, httpWebRootDir, tlsChallengePort, dnsProvider, true)
		if err != nil {
			return nil, errors.Wrap(err, "creating certfetcher")
		}
		log.Println("Certfetcher created successfully.")
		return certFetcher, nil
	}
	// Certfetcher can be nil, if auto renew is off.
	return nil, nil
}

// Loads X509 certificates from disk.
// Returns appropriate errors if:
//	The file can't be read.
//	The certificate can't be parsed.
//	No certificates found in the file.
//	Certificates cannot be used to sign HTTP exchanges.
//	 (if developmentMode, print a warning that certs can't
//	 be used to sign HTTP exchanges).
// If there are no errors, the array of certificates is returned.
func loadCertsFromFile(config *util.Config, developmentMode bool) ([]*x509.Certificate, error) {
	// TODO(twifkak): Document what cert/key storage formats this accepts.
	certPem, err := ioutil.ReadFile(config.CertFile)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", config.CertFile)
	}
	certs, err := signedexchange.ParseCertificates(certPem)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing %s", config.CertFile)
	}
	if certs == nil || len(certs) == 0 {
		return nil, errors.Errorf("no cert found in %s", config.CertFile)
	}
	if err := util.CanSignHttpExchanges(certs[0]); err != nil {
		if developmentMode {
			log.Println("WARNING:", err)
		} else {
			return nil, err
		}
	}

	return certs, nil
}

func loadCSRFromFile(config *util.Config) (*x509.CertificateRequest, error) {
	data, err := ioutil.ReadFile(config.CSRFile)
        if err != nil {
                return nil, errors.Wrapf(err, "reading %s", config.CSRFile)
        }
        block, _ := pem.Decode(data)
        if block == nil {
                return nil, errors.Errorf("pem decode: no key found in %s", config.CSRFile)
        }
        csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing CSR %s", config.CSRFile)
	}
	return csr, nil
}

// Loads private key from file.
// Returns appropriate errors if:
//	The file can't be read.
//	The key can't be parsed.
// If there are no errors, the key is returned.
func LoadKeyFromFile(config *util.Config) (crypto.PrivateKey, error) {
	keyPem, err := ioutil.ReadFile(config.KeyFile)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", config.KeyFile)
	}

	key, err := util.ParsePrivateKey(keyPem)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing %s", config.KeyFile)
	}

	return key, nil
}
