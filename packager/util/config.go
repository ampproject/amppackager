// Copyright 2018 Google LLC
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

package util

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

type Config struct {
	LocalOnly bool
	Port      int
	CertFile  string // This must be the full certificate chain.
	KeyFile   string // Just for the first cert, obviously.
	OCSPCache string
	URLSet    []URLSet
}

type URLSet struct {
	Fetch *URLPattern
	Sign  *URLPattern
}

type URLPattern struct {
	Scheme                 []string
	DomainRE               string
	Domain                 string
	PathRE                 *string
	PathExcludeRE          []string
	QueryRE                *string
	ErrorOnStatefulHeaders bool
	MaxLength              int
	SamePath               *bool
}

// TODO(twifkak): Extract default values into a function separate from the one
// that does the parsing and validation. This would make signer_test and
// validation_test less brittle.

var emptyRegexp = ""
var defaultPathRegexp = ".*"

// Also sets defaults.
func validateURLPattern(pattern *URLPattern) error {
	if pattern.PathRE == nil {
		pattern.PathRE = &defaultPathRegexp
	} else if _, err := regexp.Compile(*pattern.PathRE); err != nil {
		return errors.New("PathRE must be a valid regexp")
	}
	for _, exclude := range pattern.PathExcludeRE {
		if _, err := regexp.Compile(exclude); err != nil {
			return errors.Errorf("PathExcludeRE contains invalid regexp %q", exclude)
		}
	}
	if pattern.QueryRE == nil {
		pattern.QueryRE = &emptyRegexp
	} else if _, err := regexp.Compile(*pattern.QueryRE); err != nil {
		return errors.New("QueryRE must be a valid regexp")
	}
	if pattern.MaxLength == 0 {
		pattern.MaxLength = 2000
	}
	return nil
}

func validateSignURLPattern(pattern *URLPattern) error {
	if pattern == nil {
		return errors.New("This section must be specified")
	}
	if pattern.Scheme != nil {
		return errors.New("Scheme not allowed here")
	}
	if pattern.Domain == "" {
		return errors.New("Domain must be specified")
	}
	if pattern.DomainRE != "" {
		return errors.New("DomainRE not allowed here")
	}
	if pattern.SamePath != nil {
		return errors.New("SamePath not allowed here")
	}
	if err := validateURLPattern(pattern); err != nil {
		return err
	}
	return nil
}

var allowedFetchSchemes = map[string]bool{"http": true, "https": true}

func validateFetchURLPattern(pattern *URLPattern) error {
	if pattern == nil {
		return nil
	}
	if len(pattern.Scheme) == 0 {
		// Default Scheme to the list of keys in allowedFetchSchemes.
		pattern.Scheme = make([]string, len(allowedFetchSchemes))
		i := 0
		for scheme := range allowedFetchSchemes {
			pattern.Scheme[i] = scheme
			i++
		}
	} else {
		for _, scheme := range pattern.Scheme {
			if !allowedFetchSchemes[scheme] {
				return errors.Errorf("Scheme contains invalid value %q", scheme)
			}
		}
	}
	if pattern.Domain == "" && pattern.DomainRE == "" {
		return errors.New("Domain or DomainRE must be specified")
	}
	if pattern.Domain != "" && pattern.DomainRE != "" {
		return errors.New("Only one of Domain or DomainRE should be specified")
	}
	if pattern.SamePath == nil {
		// Default SamePath to true.
		pattern.SamePath = new(bool)
		*pattern.SamePath = true
	}
	if pattern.ErrorOnStatefulHeaders {
		return errors.New("ErrorOnStatefulHeaders not allowed here")
	}
	if err := validateURLPattern(pattern); err != nil {
		return err
	}
	return nil
}

// ReadConfig reads the config file specified at --config and validates it.
func ReadConfig(configBytes []byte) (*Config, error) {
	tree, err := toml.LoadBytes(configBytes)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse TOML")
	}
	config := Config{}
	if err = tree.Unmarshal(&config); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal TOML")
	}
	// TODO(twifkak): Return an error if the TOML includes any fields that aren't part of the Config struct.

	if config.Port == 0 {
		config.Port = 8080
	}
	if config.CertFile == "" {
		return nil, errors.New("must specify CertFile")
	}
	if config.KeyFile == "" {
		return nil, errors.New("must specify KeyFile")
	}
	if config.OCSPCache == "" {
		return nil, errors.New("must specify OCSPCache")
	}
	ocspDir := filepath.Dir(config.OCSPCache)
	if stat, err := os.Stat(ocspDir); os.IsNotExist(err) || !stat.Mode().IsDir() {
		return nil, errors.Errorf("OCSPCache parent directory must exist: %s", ocspDir)
	}
	// TODO(twifkak): Verify OCSPCache is writable by the current user.
	if len(config.URLSet) == 0 {
		return nil, errors.New("must specify one or more [[URLSet]]")
	}
	for i := range config.URLSet {
		if config.URLSet[i].Fetch != nil {
			if err := validateFetchURLPattern(config.URLSet[i].Fetch); err != nil {
				return nil, errors.Wrapf(err, "parsing URLSet.%d.Fetch", i)
			}
		}
		if err := validateSignURLPattern(config.URLSet[i].Sign); err != nil {
			return nil, errors.Wrapf(err, "parsing URLSet.%d.Sign", i)
		}
	}
	return &config, nil
}
