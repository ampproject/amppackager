// Package lwApi is a minimalist API client to LiquidWeb's (https://www.liquidweb.com) API:
//
// https://cart.liquidweb.com/storm/api/docs/v1
//
// https://cart.liquidweb.com/storm/api/docs/bleed
//
// As you might have guessed from the above API documentation links, there are API versions:
// "v1" and "bleed". As the name suggests, if you always want the latest features and abilities,
// use "bleed". If you want long term compatibility (at the cost of being a little further behind
// sometimes), use "v1".
package lwApi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// A LWAPIConfig holds the configuration details used to call
// the API with the client.
type LWAPIConfig struct {
	Username *string
	Password *string
	Token    *string
	Url      string
	Timeout  uint
	Insecure bool
}

// A Client holds the packages *LWAPIConfig and *http.Client. To get a *Client, call New.
type Client struct {
	Headers    http.Header
	config     *LWAPIConfig
	httpClient *http.Client
	mutex      sync.Mutex
}

// A LWAPIError is used to identify error responses when JSON unmarshalling json from a
// byte slice.
type LWAPIError struct {
	ErrorMsg     string `json:"error,omitempty"`
	ErrorClass   string `json:"error_class,omitempty"`
	ErrorFullMsg string `json:"full_message,omitempty"`
}

// Given a LWAPIError, returns a string containing the ErrorClass and ErrorFullMsg.
func (e LWAPIError) Error() string {
	return fmt.Sprintf("%v: %v", e.ErrorClass, e.ErrorFullMsg)
}

// Given a LWAPIError, returns boolean if ErrorClass was present or not. You can
// use this function to determine if a LWAPIRes response indicates an error or not.
func (e LWAPIError) HadError() bool {
	return e.ErrorClass != ""
}

// LWAPIRes is a convenient interface used (for example) by CallInto to ensure a passed
// struct knows how to indicate whether or not it had an error.
type LWAPIRes interface {
	Error() string
	HadError() bool
}

// New takes a *LWAPIConfig, and gives you a *Client. If there's an error, it is returned.
// When using this package, this should be the first function you call. Below is an example
// that demonstrates creating the config and passing it to New.
//
// Example:
//	username := "ExampleUsername"
//	password := "ExamplePassword"
//
//	config := lwApi.LWAPIConfig{
//		Username: &username,
//		Password: &password,
//		Url:      "api.liquidweb.com",
//	}
//	apiClient, newErr := lwApi.New(&config)
//	if newErr != nil {
//		panic(newErr)
//	}
func New(config *LWAPIConfig) (*Client, error) {
	if err := processConfig(config); err != nil {
		return nil, err
	}

	httpClient := &http.Client{Timeout: time.Duration(time.Duration(config.Timeout) * time.Second)}

	if config.Insecure {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient.Transport = tr
	}

	headers := make(http.Header)
	client := Client{
		config:     config,
		httpClient: httpClient,
		Headers:    headers,
		mutex:      sync.Mutex{},
	}
	return &client, nil
}

// Call takes a path, such as "network/zone/details" and a params structure.
// It is recommended that the params be a map[string]interface{}, but you can use
// anything that serializes to the right json structure.
// A `interface{}` and an error are returned, in typical go fasion.
//
// Example:
//	args := map[string]interface{}{
//		"uniq_id": "ABC123",
//	}
//	got, gotErr := apiClient.Call("bleed/asset/details", args)
//	if gotErr != nil {
//		panic(gotErr)
//	}
func (client *Client) Call(method string, params interface{}) (interface{}, error) {
	bsRb, err := client.CallRaw(method, params)
	if err != nil {
		return nil, err
	}

	var resp interface{}
	resp, err = client.callRawRespToInterface(bsRb)

	return resp, err
}

// CallInto is like call, but instead of returning an interface you pass it a
// struct which is filled, much like the json.Unmarshal function.  The struct
// you pass must satisfy the LWAPIRes interface.  If you embed the LWAPIError
// struct from this package into your struct, this will be taken care of for you.
//
// Example:
//	type ZoneDetails struct {
//		lwApi.LWAPIError
//		AvlZone     string   `json:"availability_zone"`
//		Desc        string   `json:"description"`
//		GatewayDevs []string `json:"gateway_devices"`
//		HvType      string   `json:"hv_type"`
//		ID          int      `json:"id"`
//		Legacy      int      `json:"legacy"`
//		Name        string   `json:"name"`
//		Status      string   `json:"status"`
//		SourceHVs   []string `json:"valid_source_hvs"`
//	}
//	var zone ZoneDetails
//	err = apiClient.CallInto("network/zone/details", paramers, &zone)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Got struct %#v\n", zone)
//
func (client *Client) CallInto(method string, params interface{}, into LWAPIRes) error {
	bsRb, err := client.CallRaw(method, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bsRb, into)
	if err != nil {
		return err
	}

	if into.HadError() {
		// the LWAPIRes satisfies the Error interface, so we can just return it on
		// error.
		return into
	}

	return nil
}

// Similar to CallInto(), but populates an interface without needing to satisfy
// the LWAPIRes interface.
// Example:
//      type ZoneDetails struct {
//              AvlZone     string   `json:"availability_zone"`
//              Desc        string   `json:"description"`
//              GatewayDevs []string `json:"gateway_devices"`
//              HvType      string   `json:"hv_type"`
//              ID          int      `json:"id"`
//              Legacy      int      `json:"legacy"`
//              Name        string   `json:"name"`
//              Status      string   `json:"status"`
//              SourceHVs   []string `json:"valid_source_hvs"`
//      }
//      var zone ZoneDetails
//      err = apiClient.CallIntoInterface("network/zone/details", params, &zone)

func (client *Client) CallIntoInterface(method string, params interface{}, into interface{}) error {
	bsRb, err := client.CallRaw(method, params)
	if err != nil {
		return err
	}

	if _, err = client.callRawRespToInterface(bsRb, &into); err != nil {
		return err
	}

	return nil
}

// CallRaw is just like Call, except it returns the raw json as a byte slice. However, in contrast to
// Call, CallRaw does *not* check the API response for LiquidWeb specific exceptions as defined in
// the type LWAPIError. As such, if calling this function directly, you must check for LiquidWeb specific
// exceptions yourself.
//
// Example:
//	args := map[string]interface{}{
//		"uniq_id": "ABC123",
//	}
//	got, gotErr := apiClient.CallRaw("bleed/asset/details", args)
//	if gotErr != nil {
//		panic(gotErr)
//	}
//	// Check got now for LiquidWeb specific exceptions, as described above.
func (client *Client) CallRaw(method string, params interface{}) ([]byte, error) {
	config := client.config
	//  api wants the "params" prefix key. Do it here so consumers dont have
	// to do this everytime.
	args := map[string]interface{}{
		"params": params,
	}
	encodedArgs, encodeErr := json.Marshal(args)
	if encodeErr != nil {
		return nil, encodeErr
	}
	// formulate the HTTP POST request
	url := fmt.Sprintf("%s/%s", config.Url, method)
	req, reqErr := http.NewRequest("POST", url, bytes.NewReader(encodedArgs))
	if reqErr != nil {
		return nil, reqErr
	}

	// We need a unique copy of the headers map in each request struct, otherwise
	// we can end up with a concurrent map access and a panic.
	client.mutex.Lock()
	for name, value := range client.Headers {
		newvalue := make([]string, len(value))
		copy(newvalue, value)
		req.Header[name] = newvalue
	}
	client.mutex.Unlock()

	if config.Token != nil {
		// Oauth2 token
		req.Header.Add("Authorization", "Bearer "+*config.Token)
	} else if config.Username != nil && config.Password != nil {
		// HTTP basic auth
		req.SetBasicAuth(*config.Username, *config.Password)
	} else {
		return nil, fmt.Errorf("No valid credential provided")
	}

	// make the POST request
	resp, doErr := client.httpClient.Do(req)
	if doErr != nil {
		return nil, doErr
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad HTTP response code [%d] from [%s]", resp.StatusCode, url)
	}
	// read the response body into a byte slice
	bsRb, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	return bsRb, nil
}

/* private */

func processConfig(config *LWAPIConfig) error {
	if config.Url == "" {
		return fmt.Errorf("url is missing from config")
	}
	if config.Timeout == 0 {
		config.Timeout = 20
	}

	if config.Token != nil {
		// Oauth2 token
		if *config.Token == "" {
			return fmt.Errorf("Bearer token provided, but empty")
		}
	} else if config.Username != nil && config.Password != nil {
		// HTTP basic auth
		if *config.Username == "" {
			return fmt.Errorf("provided username is empty")
		}
		if *config.Password == "" {
			return fmt.Errorf("provided password is empty")
		}
	} else {
		return fmt.Errorf("No valid credential provided")
	}

	return nil
}

func (client *Client) callRawRespToInterface(dataBytes []byte, into ...interface{}) (dataInter interface{}, err error) {
	var raw map[string]interface{}
	if err = json.Unmarshal(dataBytes, &raw); err == nil {
		if errorClass, exists := raw["error_class"]; exists {
			errorClassStr := fmt.Sprintf("%s", errorClass)
			if errorClassStr != "" {
				err = LWAPIError{
					ErrorClass:   errorClassStr,
					ErrorFullMsg: fmt.Sprintf("%s", raw["full_message"]),
					ErrorMsg:     fmt.Sprintf("%s", raw["error"]),
				}
				return
			}
		}
	} else {
		return
	}

	if len(into) > 0 {
		dataInter = &into[0]
	}

	err = json.Unmarshal(dataBytes, &dataInter)

	return
}
