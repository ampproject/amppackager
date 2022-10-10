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

// SetDomainRecordStatus invokes the alidns.SetDomainRecordStatus API synchronously
func (client *Client) SetDomainRecordStatus(request *SetDomainRecordStatusRequest) (response *SetDomainRecordStatusResponse, err error) {
	response = CreateSetDomainRecordStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SetDomainRecordStatusWithChan invokes the alidns.SetDomainRecordStatus API asynchronously
func (client *Client) SetDomainRecordStatusWithChan(request *SetDomainRecordStatusRequest) (<-chan *SetDomainRecordStatusResponse, <-chan error) {
	responseChan := make(chan *SetDomainRecordStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDomainRecordStatus(request)
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

// SetDomainRecordStatusWithCallback invokes the alidns.SetDomainRecordStatus API asynchronously
func (client *Client) SetDomainRecordStatusWithCallback(request *SetDomainRecordStatusRequest, callback func(response *SetDomainRecordStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDomainRecordStatusResponse
		var err error
		defer close(result)
		response, err = client.SetDomainRecordStatus(request)
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

// SetDomainRecordStatusRequest is the request struct for api SetDomainRecordStatus
type SetDomainRecordStatusRequest struct {
	*requests.RpcRequest
	RecordId     string `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

// SetDomainRecordStatusResponse is the response struct for api SetDomainRecordStatus
type SetDomainRecordStatusResponse struct {
	*responses.BaseResponse
	Status    string `json:"Status" xml:"Status"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	RecordId  string `json:"RecordId" xml:"RecordId"`
}

// CreateSetDomainRecordStatusRequest creates a request to invoke SetDomainRecordStatus API
func CreateSetDomainRecordStatusRequest() (request *SetDomainRecordStatusRequest) {
	request = &SetDomainRecordStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "SetDomainRecordStatus", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDomainRecordStatusResponse creates a response to parse from SetDomainRecordStatus response
func CreateSetDomainRecordStatusResponse() (response *SetDomainRecordStatusResponse) {
	response = &SetDomainRecordStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
