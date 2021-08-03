/*
Copyright © LiquidWeb

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package validate

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cast"

	"github.com/liquidweb/liquidweb-cli/utils"
)

var ValidationFailure = errors.New("validation failed")

type InputTypes struct {
	UniqId                          InputTypeUniqId
	IP                              InputTypeIP
	PositiveInt64                   InputTypePositiveInt64
	PositiveInt                     InputTypePositiveInt
	NonEmptyString                  InputTypeNonEmptyString
	LoadBalancerStrategy            InputTypeLoadBalancerStrategyString
	HttpsLiquidwebUrl               InputTypeHttpsLiquidwebUrl
	NetworkPortPair                 InputTypeNetworkPortPair
	NetworkPort                     InputTypeNetworkPort
	LoadBalancerHttpCodeRange       InputTypeLoadBalancerHttpCodeRange
	LoadBalancerHealthCheckProtocol InputTypeLoadBalancerHealthCheckProtocol
}

// UniqId

type InputTypeUniqId struct {
	UniqId string
}

func (x InputTypeUniqId) Validate() error {
	// must be uppercase
	allUpper := strings.ToUpper(x.UniqId)
	if allUpper != x.UniqId {
		return fmt.Errorf("a uniq_id must be uppercase")
	}

	// must be 6 characters
	if len(x.UniqId) != 6 {
		return fmt.Errorf("a uniq_id must be 6 characters long")
	}

	// must be alphanumeric
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return err
	}
	regexStr := reg.ReplaceAllString(x.UniqId, "")
	if regexStr != x.UniqId {
		return fmt.Errorf("a uniq_id must be alphanumeric")
	}

	return nil
}

// IP

type InputTypeIP struct {
	IP string
}

func (x InputTypeIP) Validate() error {

	if !utils.IpIsValid(x.IP) {
		return fmt.Errorf("ip [%s] is not a valid IP address", x.IP)
	}

	return nil
}

// PositiveInt64

type InputTypePositiveInt64 struct {
	PositiveInt64 int64
}

func (x InputTypePositiveInt64) Validate() error {
	if x.PositiveInt64 < 0 {
		return fmt.Errorf("PositiveInt64 is not > 0")
	}

	return nil
}

// PositiveInt

type InputTypePositiveInt struct {
	PositiveInt int
}

func (x InputTypePositiveInt) Validate() error {
	if x.PositiveInt < 0 {
		return fmt.Errorf("PositiveInt is not > 0")
	}

	return nil
}

// NonEmptyString

type InputTypeNonEmptyString struct {
	NonEmptyString string
}

func (x InputTypeNonEmptyString) Validate() error {
	if x.NonEmptyString == "" {
		return fmt.Errorf("NonEmptyString cannot be empty")
	}

	return nil
}

// LoadBalancerStrategy

type InputTypeLoadBalancerStrategyString struct {
	LoadBalancerStrategy string
}

func (x InputTypeLoadBalancerStrategyString) Validate() error {
	strategies := map[string]int{
		"roundrobin":  1,
		"connections": 1,
		"cells":       1,
	}

	if _, exists := strategies[x.LoadBalancerStrategy]; !exists {
		var slice []string
		slice = append(slice, fmt.Sprintf("LoadBalancer strategy [%s] is invalid. Valid strategies: ",
			x.LoadBalancerStrategy))
		for strategy, _ := range strategies {
			slice = append(slice, fmt.Sprintf("%s ", strategy))
		}
		return fmt.Errorf("%s", strings.Join(slice[:], ""))
	}

	return nil
}

// HttpsLiquidwebUrl

type InputTypeHttpsLiquidwebUrl struct {
	HttpsLiquidwebUrl string
}

func (x InputTypeHttpsLiquidwebUrl) Validate() error {
	if !strings.HasPrefix(x.HttpsLiquidwebUrl, "https://") {
		return fmt.Errorf("given url [%s] appears invalid; should start with 'https://'",
			x.HttpsLiquidwebUrl)
	}

	if !strings.Contains(x.HttpsLiquidwebUrl, "liquidweb.com") {
		return fmt.Errorf("given url [%s] appears invalid; should contain 'liquidweb.com'",
			x.HttpsLiquidwebUrl)
	}

	if _, err := url.ParseRequestURI(x.HttpsLiquidwebUrl); err != nil {
		return fmt.Errorf("given url [%s] appears invalid; %s", x.HttpsLiquidwebUrl, err)
	}

	return nil
}

// NetworkPortPair

type InputTypeNetworkPortPair struct {
	NetworkPortPair string
}

func (x InputTypeNetworkPortPair) Validate() error {
	if !strings.Contains(x.NetworkPortPair, ":") {
		return fmt.Errorf("given NetworkPortPair [%s] contains no ':' which is invalid",
			x.NetworkPortPair)
	}

	splitPair := strings.Split(x.NetworkPortPair, ":")

	if len(splitPair) != 2 {
		return fmt.Errorf(
			"A NetworkPortPair must contain exactly one source/destination port pair")
	}

	for _, portStr := range splitPair {
		if _, err := strconv.Atoi(portStr); err != nil {
			return fmt.Errorf("port [%s] in port pair [%s] doesnt look numeric", portStr,
				x.NetworkPortPair)
		}

		obj := InputTypeNetworkPort{NetworkPort: cast.ToInt(portStr)}
		if err := obj.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NetworkPort

type InputTypeNetworkPort struct {
	NetworkPort int
}

func (x InputTypeNetworkPort) Validate() error {
	if x.NetworkPort <= 0 || x.NetworkPort > 65535 {
		return fmt.Errorf("NetworkPort [%d] is invalid; must be between 1 and 65535", x.NetworkPort)
	}

	return nil
}

// LoadBalancerHealthCheckProtocol

type InputTypeLoadBalancerHealthCheckProtocol struct {
	LoadBalancerHealthCheckProtocol string
}

func (x InputTypeLoadBalancerHealthCheckProtocol) Validate() error {
	if x.LoadBalancerHealthCheckProtocol != "tcp" && x.LoadBalancerHealthCheckProtocol != "http" {
		return fmt.Errorf("A LoadBalancerHealthCheckProtocol must be one of [http, tcp]")
	}

	return nil
}

// LoadBalancerHttpCodeRange

type InputTypeLoadBalancerHttpCodeRange struct {
	LoadBalancerHttpCodeRange string
}

func (x InputTypeLoadBalancerHttpCodeRange) Validate() error {
	if x.LoadBalancerHttpCodeRange == "" {
		return fmt.Errorf("A LoadBalancerHttpCodeRange cannot be blank")
	}

	// format example: 200,201,400-404,205
	codes := strings.Split(x.LoadBalancerHttpCodeRange, ",")
	for _, code := range codes {
		var eachCode []string
		if strings.Contains(code, "-") {
			splitCodes := strings.Split(code, "-")
			if len(splitCodes) != 2 {
				return fmt.Errorf("a LoadBalancerHttpCodeRange with '-' for ranges must be in form '200-202' for example")
			}
			eachCode = append(eachCode, splitCodes[0])
			eachCode = append(eachCode, splitCodes[1])
		} else {
			eachCode = append(eachCode, code)
		}

		for _, code := range eachCode {
			if _, err := strconv.Atoi(code); err != nil {
				return fmt.Errorf("http code [%s] in [%s] doesnt look numeric", code, x.LoadBalancerHttpCodeRange)
			}

			reg, err := regexp.Compile(`^\d\d\d$`)
			if err != nil {
				return err
			}
			matched := reg.MatchString(code)
			if !matched {
				return fmt.Errorf("http code [%s] must be a three digit number", code)
			}
		}
	}

	return nil
}
