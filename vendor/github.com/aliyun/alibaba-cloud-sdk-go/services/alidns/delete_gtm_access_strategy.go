package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteGtmAccessStrategy invokes the alidns.DeleteGtmAccessStrategy API synchronously
// api document: https://help.aliyun.com/api/alidns/deletegtmaccessstrategy.html
func (client *Client) DeleteGtmAccessStrategy(request *DeleteGtmAccessStrategyRequest) (response *DeleteGtmAccessStrategyResponse, err error) {
	response = CreateDeleteGtmAccessStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGtmAccessStrategyWithChan invokes the alidns.DeleteGtmAccessStrategy API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletegtmaccessstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteGtmAccessStrategyWithChan(request *DeleteGtmAccessStrategyRequest) (<-chan *DeleteGtmAccessStrategyResponse, <-chan error) {
	responseChan := make(chan *DeleteGtmAccessStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGtmAccessStrategy(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteGtmAccessStrategyWithCallback invokes the alidns.DeleteGtmAccessStrategy API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletegtmaccessstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteGtmAccessStrategyWithCallback(request *DeleteGtmAccessStrategyRequest, callback func(response *DeleteGtmAccessStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGtmAccessStrategyResponse
		var err error
		defer close(result)
		response, err = client.DeleteGtmAccessStrategy(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteGtmAccessStrategyRequest is the request struct for api DeleteGtmAccessStrategy
type DeleteGtmAccessStrategyRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	StrategyId   string `position:"Query" name:"StrategyId"`
	Lang         string `position:"Query" name:"Lang"`
}

// DeleteGtmAccessStrategyResponse is the response struct for api DeleteGtmAccessStrategy
type DeleteGtmAccessStrategyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteGtmAccessStrategyRequest creates a request to invoke DeleteGtmAccessStrategy API
func CreateDeleteGtmAccessStrategyRequest() (request *DeleteGtmAccessStrategyRequest) {
	request = &DeleteGtmAccessStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DeleteGtmAccessStrategy", "Alidns", "openAPI")
	return
}

// CreateDeleteGtmAccessStrategyResponse creates a response to parse from DeleteGtmAccessStrategy response
func CreateDeleteGtmAccessStrategyResponse() (response *DeleteGtmAccessStrategyResponse) {
	response = &DeleteGtmAccessStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
