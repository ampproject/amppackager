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

package amppackager

import (
	"regexp"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

type Config struct {
	LocalOnly    bool
	Port         int
	PackagerBase string // The base URL under which /amppkg/ URLs will be served on the internet.
	CertFile     string // This must be the full certificate chain.
	KeyFile      string // Just for the first cert, obviously.
	URLSet       []URLSet
}

type URLSet struct {
	SamePath bool
	Fetch    URLPattern
	Sign     URLPattern
}

type URLPattern struct {
	Scheme                 []string
	Domain                 string
	PathRE                 *string
	PathExcludeRE          []string
	QueryRE                *string
	ErrorOnStatefulHeaders bool
}

var dotStarRegexp = ".*"

// Also sets defaults.
func validateURLPattern(pattern *URLPattern, allowedSchemes map[string]bool) error {
	if len(pattern.Scheme) == 0 {
		// Default Scheme to the list of keys in allowedSchemes.
		pattern.Scheme = make([]string, len(allowedSchemes))
		i := 0
		for scheme := range allowedSchemes {
			pattern.Scheme[i] = scheme
			i++
		}
	} else {
		for _, scheme := range pattern.Scheme {
			if !allowedSchemes[scheme] {
				return errors.Errorf("Scheme contains invalid value %q", scheme)
			}
		}
	}
	if pattern.Domain == "" {
		return errors.New("Domain must be specified")
	}
	if pattern.PathRE == nil {
		pattern.PathRE = &dotStarRegexp
	} else if _, err := regexp.Compile(*pattern.PathRE); err != nil {
		return errors.New("PathRE must be a valid regexp")
	}
	for _, exclude := range pattern.PathExcludeRE {
		if _, err := regexp.Compile(exclude); err != nil {
			return errors.Errorf("PathExcludeRE contains be invalid regexp %q", exclude)
		}
	}
	if pattern.QueryRE == nil {
		pattern.QueryRE = &dotStarRegexp
	} else if _, err := regexp.Compile(*pattern.QueryRE); err != nil {
		return errors.New("QueryRE must be a valid regexp")
	}
	return nil
}

var allowedFetchSchemes = map[string]bool{"http": true, "https": true}
var allowedSignSchemes = map[string]bool{"https": true}

// ReadConfig reads the config file specified at --config and validates it.
func ReadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, errors.New("must specify --config")
	}
	tree, err := toml.LoadFile(configPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse config at: %s", configPath)
	}
	config := Config{}
	if err = tree.Unmarshal(&config); err != nil {
		return nil, errors.Wrapf(err, "failed to parse config at: %s", configPath)
	}
	// TODO(twifkak): Return an error if the TOML includes any fields that aren't part of the Config struct.

	if config.Port == 0 {
		config.Port = 8080
	}
	if !strings.HasSuffix(config.PackagerBase, "/") {
		// This ensures that the ResolveReference call doesn't replace the last path component.
		config.PackagerBase += "/"
	}
	if config.CertFile == "" {
		return nil, errors.New("must specify CertFile")
	}
	if config.KeyFile == "" {
		return nil, errors.New("must specify KeyFile")
	}
	if len(config.URLSet) == 0 {
		return nil, errors.New("must specify one or more [[URLSet]]")
	}
	for i := range config.URLSet {
		if err := validateURLPattern(&config.URLSet[i].Fetch, allowedFetchSchemes); err != nil {
			return nil, errors.Wrapf(err, "parsing URLSet.%d.Fetch", i)
		}
		if err := validateURLPattern(&config.URLSet[i].Sign, allowedSignSchemes); err != nil {
			return nil, errors.Wrapf(err, "parsing URLSet.%d.Sign", i)
		}
		if config.URLSet[i].Sign.ErrorOnStatefulHeaders {
			return nil, errors.Errorf("URLSet.%d.Sign.ErrorOnStatefulHeaders is not allowed; perhaps you meant to put this in the Fetch section?", i)
		}
	}
	return &config, nil
}
