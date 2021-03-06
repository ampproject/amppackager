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

// DescribeRecordStatisticsHistory invokes the alidns.DescribeRecordStatisticsHistory API synchronously
// api document: https://help.aliyun.com/api/alidns/describerecordstatisticshistory.html
func (client *Client) DescribeRecordStatisticsHistory(request *DescribeRecordStatisticsHistoryRequest) (response *DescribeRecordStatisticsHistoryResponse, err error) {
	response = CreateDescribeRecordStatisticsHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordStatisticsHistoryWithChan invokes the alidns.DescribeRecordStatisticsHistory API asynchronously
// api document: https://help.aliyun.com/api/alidns/describerecordstatisticshistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecordStatisticsHistoryWithChan(request *DescribeRecordStatisticsHistoryRequest) (<-chan *DescribeRecordStatisticsHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordStatisticsHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordStatisticsHistory(request)
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

// DescribeRecordStatisticsHistoryWithCallback invokes the alidns.DescribeRecordStatisticsHistory API asynchronously
// api document: https://help.aliyun.com/api/alidns/describerecordstatisticshistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecordStatisticsHistoryWithCallback(request *DescribeRecordStatisticsHistoryRequest, callback func(response *DescribeRecordStatisticsHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordStatisticsHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordStatisticsHistory(request)
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

// DescribeRecordStatisticsHistoryRequest is the request struct for api DescribeRecordStatisticsHistory
type DescribeRecordStatisticsHistoryRequest struct {
	*requests.RpcRequest
	Rr           string `position:"Query" name:"Rr"`
	EndDate      string `position:"Query" name:"EndDate"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	StartDate    string `position:"Query" name:"StartDate"`
}

// DescribeRecordStatisticsHistoryResponse is the response struct for api DescribeRecordStatisticsHistory
type DescribeRecordStatisticsHistoryResponse struct {
	*responses.BaseResponse
	RequestId  string                                      `json:"RequestId" xml:"RequestId"`
	Statistics StatisticsInDescribeRecordStatisticsHistory `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeRecordStatisticsHistoryRequest creates a request to invoke DescribeRecordStatisticsHistory API
func CreateDescribeRecordStatisticsHistoryRequest() (request *DescribeRecordStatisticsHistoryRequest) {
	request = &DescribeRecordStatisticsHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeRecordStatisticsHistory", "Alidns", "openAPI")
	return
}

// CreateDescribeRecordStatisticsHistoryResponse creates a response to parse from DescribeRecordStatisticsHistory response
func CreateDescribeRecordStatisticsHistoryResponse() (response *DescribeRecordStatisticsHistoryResponse) {
	response = &DescribeRecordStatisticsHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
