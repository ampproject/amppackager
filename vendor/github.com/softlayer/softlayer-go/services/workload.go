/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// no documentation yet
type Workload_Citrix_Client struct {
	Session *session.Session
	Options sl.Options
}

// GetWorkloadCitrixClientService returns an instance of the Workload_Citrix_Client SoftLayer service
func GetWorkloadCitrixClientService(sess *session.Session) Workload_Citrix_Client {
	return Workload_Citrix_Client{Session: sess}
}

func (r Workload_Citrix_Client) Id(id int) Workload_Citrix_Client {
	r.Options.Id = &id
	return r
}

func (r Workload_Citrix_Client) Mask(mask string) Workload_Citrix_Client {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Workload_Citrix_Client) Filter(filter string) Workload_Citrix_Client {
	r.Options.Filter = filter
	return r
}

func (r Workload_Citrix_Client) Limit(limit int) Workload_Citrix_Client {
	r.Options.Limit = &limit
	return r
}

func (r Workload_Citrix_Client) Offset(offset int) Workload_Citrix_Client {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Workload_Citrix_Client) CreateResourceLocation(citrixCredentials *datatypes.Workload_Citrix_Request) (resp datatypes.Workload_Citrix_Client_Response, err error) {
	params := []interface{}{
		citrixCredentials,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Client", "createResourceLocation", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Workload_Citrix_Client) GetResourceLocations(citrixCredentials *datatypes.Workload_Citrix_Request) (resp datatypes.Workload_Citrix_Client_Response, err error) {
	params := []interface{}{
		citrixCredentials,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Client", "getResourceLocations", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Workload_Citrix_Client) ValidateCitrixCredentials(citrixCredentials *datatypes.Workload_Citrix_Request) (resp datatypes.Workload_Citrix_Client_Response, err error) {
	params := []interface{}{
		citrixCredentials,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Client", "validateCitrixCredentials", params, &r.Options, &resp)
	return
}

// no documentation yet
type Workload_Citrix_Deployment struct {
	Session *session.Session
	Options sl.Options
}

// GetWorkloadCitrixDeploymentService returns an instance of the Workload_Citrix_Deployment SoftLayer service
func GetWorkloadCitrixDeploymentService(sess *session.Session) Workload_Citrix_Deployment {
	return Workload_Citrix_Deployment{Session: sess}
}

func (r Workload_Citrix_Deployment) Id(id int) Workload_Citrix_Deployment {
	r.Options.Id = &id
	return r
}

func (r Workload_Citrix_Deployment) Mask(mask string) Workload_Citrix_Deployment {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Workload_Citrix_Deployment) Filter(filter string) Workload_Citrix_Deployment {
	r.Options.Filter = filter
	return r
}

func (r Workload_Citrix_Deployment) Limit(limit int) Workload_Citrix_Deployment {
	r.Options.Limit = &limit
	return r
}

func (r Workload_Citrix_Deployment) Offset(offset int) Workload_Citrix_Deployment {
	r.Options.Offset = &offset
	return r
}

// Creates a new Citrix Virtual Apps and Desktops deployment.
func (r Workload_Citrix_Deployment) CreateObject(templateObject *datatypes.Workload_Citrix_Deployment) (resp datatypes.Workload_Citrix_Deployment, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "createObject", params, &r.Options, &resp)
	return
}

// Retrieve The [[SoftLayer_Account]] to which the deployment belongs.
func (r Workload_Citrix_Deployment) GetAccount() (resp datatypes.Account, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getAccount", nil, &r.Options, &resp)
	return
}

// Get all Citrix Virtual Apps And Desktop deployments.
func (r Workload_Citrix_Deployment) GetAllObjects() (resp []datatypes.Workload_Citrix_Deployment, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getAllObjects", nil, &r.Options, &resp)
	return
}

// Returns a response object [[SoftLayer_Workload_Citrix_Deployment_Response]] which represents the CVAD deployment [[SoftLayer_Workload_Citrix_Deployment]] together with all the resources ordered under the CVAD order.
//
// The deployment resources are represented by object [[SoftLayer_Workload_Citrix_Deployment_Resource_Response]].
func (r Workload_Citrix_Deployment) GetDeployment(deploymentId *int) (resp datatypes.Workload_Citrix_Deployment_Response, err error) {
	params := []interface{}{
		deploymentId,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getDeployment", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Workload_Citrix_Deployment) GetObject() (resp datatypes.Workload_Citrix_Deployment, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve It contains a collection of items under the CVAD deployment.
func (r Workload_Citrix_Deployment) GetResources() (resp []datatypes.Workload_Citrix_Deployment_Resource, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getResources", nil, &r.Options, &resp)
	return
}

// Retrieve Current Status of the CVAD deployment.
func (r Workload_Citrix_Deployment) GetStatus() (resp datatypes.Workload_Citrix_Deployment_Status, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getStatus", nil, &r.Options, &resp)
	return
}

// Retrieve It shows if the deployment is for Citrix Hypervisor or VMware.
func (r Workload_Citrix_Deployment) GetType() (resp datatypes.Workload_Citrix_Deployment_Type, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getType", nil, &r.Options, &resp)
	return
}

// Retrieve It is the [[SoftLayer_User_Customer]] who placed the order for CVAD.
func (r Workload_Citrix_Deployment) GetUser() (resp datatypes.User_Customer, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getUser", nil, &r.Options, &resp)
	return
}

// Retrieve It is the VLAN resource for the CVAD deployment.
func (r Workload_Citrix_Deployment) GetVlan() (resp datatypes.Network_Vlan, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment", "getVlan", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Workload_Citrix_Deployment_Resource type contains the information of the resource such as the Deployment ID, resource's Billing Item ID, Order ID and Role of the resource in the CVAD deployment.
type Workload_Citrix_Deployment_Resource struct {
	Session *session.Session
	Options sl.Options
}

// GetWorkloadCitrixDeploymentResourceService returns an instance of the Workload_Citrix_Deployment_Resource SoftLayer service
func GetWorkloadCitrixDeploymentResourceService(sess *session.Session) Workload_Citrix_Deployment_Resource {
	return Workload_Citrix_Deployment_Resource{Session: sess}
}

func (r Workload_Citrix_Deployment_Resource) Id(id int) Workload_Citrix_Deployment_Resource {
	r.Options.Id = &id
	return r
}

func (r Workload_Citrix_Deployment_Resource) Mask(mask string) Workload_Citrix_Deployment_Resource {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Workload_Citrix_Deployment_Resource) Filter(filter string) Workload_Citrix_Deployment_Resource {
	r.Options.Filter = filter
	return r
}

func (r Workload_Citrix_Deployment_Resource) Limit(limit int) Workload_Citrix_Deployment_Resource {
	r.Options.Limit = &limit
	return r
}

func (r Workload_Citrix_Deployment_Resource) Offset(offset int) Workload_Citrix_Deployment_Resource {
	r.Options.Offset = &offset
	return r
}

// This will add the resource into CVAD deployment.
func (r Workload_Citrix_Deployment_Resource) CreateObject(templateObject *datatypes.Workload_Citrix_Deployment_Resource) (resp datatypes.Workload_Citrix_Deployment_Resource, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Resource", "createObject", params, &r.Options, &resp)
	return
}

// Get all the resources of Citrix Deployments.
func (r Workload_Citrix_Deployment_Resource) GetAllObjects() (resp []datatypes.Workload_Citrix_Deployment_Resource, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Resource", "getAllObjects", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Workload_Citrix_Deployment_Resource) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Resource", "getBillingItem", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Workload_Citrix_Deployment_Resource) GetDeployment() (resp datatypes.Workload_Citrix_Deployment, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Resource", "getDeployment", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Workload_Citrix_Deployment_Resource object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Workload_Citrix_Deployment_Resource service. You can only retrieve resources that are assigned to your portal user's account.
func (r Workload_Citrix_Deployment_Resource) GetObject() (resp datatypes.Workload_Citrix_Deployment_Resource, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Resource", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Workload_Citrix_Deployment_Resource) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Resource", "getOrder", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Workload_Citrix_Deployment_Resource) GetRole() (resp datatypes.Workload_Citrix_Deployment_Resource_Role, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Resource", "getRole", nil, &r.Options, &resp)
	return
}

// no documentation yet
type Workload_Citrix_Deployment_Type struct {
	Session *session.Session
	Options sl.Options
}

// GetWorkloadCitrixDeploymentTypeService returns an instance of the Workload_Citrix_Deployment_Type SoftLayer service
func GetWorkloadCitrixDeploymentTypeService(sess *session.Session) Workload_Citrix_Deployment_Type {
	return Workload_Citrix_Deployment_Type{Session: sess}
}

func (r Workload_Citrix_Deployment_Type) Id(id int) Workload_Citrix_Deployment_Type {
	r.Options.Id = &id
	return r
}

func (r Workload_Citrix_Deployment_Type) Mask(mask string) Workload_Citrix_Deployment_Type {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Workload_Citrix_Deployment_Type) Filter(filter string) Workload_Citrix_Deployment_Type {
	r.Options.Filter = filter
	return r
}

func (r Workload_Citrix_Deployment_Type) Limit(limit int) Workload_Citrix_Deployment_Type {
	r.Options.Limit = &limit
	return r
}

func (r Workload_Citrix_Deployment_Type) Offset(offset int) Workload_Citrix_Deployment_Type {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Workload_Citrix_Deployment_Type) GetObject() (resp datatypes.Workload_Citrix_Deployment_Type, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Deployment_Type", "getObject", nil, &r.Options, &resp)
	return
}

// no documentation yet
type Workload_Citrix_Workspace_Order struct {
	Session *session.Session
	Options sl.Options
}

// GetWorkloadCitrixWorkspaceOrderService returns an instance of the Workload_Citrix_Workspace_Order SoftLayer service
func GetWorkloadCitrixWorkspaceOrderService(sess *session.Session) Workload_Citrix_Workspace_Order {
	return Workload_Citrix_Workspace_Order{Session: sess}
}

func (r Workload_Citrix_Workspace_Order) Id(id int) Workload_Citrix_Workspace_Order {
	r.Options.Id = &id
	return r
}

func (r Workload_Citrix_Workspace_Order) Mask(mask string) Workload_Citrix_Workspace_Order {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Workload_Citrix_Workspace_Order) Filter(filter string) Workload_Citrix_Workspace_Order {
	r.Options.Filter = filter
	return r
}

func (r Workload_Citrix_Workspace_Order) Limit(limit int) Workload_Citrix_Workspace_Order {
	r.Options.Limit = &limit
	return r
}

func (r Workload_Citrix_Workspace_Order) Offset(offset int) Workload_Citrix_Workspace_Order {
	r.Options.Offset = &offset
	return r
}

// This method will cancel the resources associated with the provided VLAN and have a 'cvad' tag reference.
func (r Workload_Citrix_Workspace_Order) CancelWorkspaceResources(vlanIdentifier *string, cancelImmediately *bool, customerNote *string) (resp datatypes.Workload_Citrix_Workspace_Response_Result, err error) {
	params := []interface{}{
		vlanIdentifier,
		cancelImmediately,
		customerNote,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Workspace_Order", "cancelWorkspaceResources", params, &r.Options, &resp)
	return
}

// This method will return the list of names of VLANs which have a 'cvad' tag reference.  This name can be used with the cancelWorkspaceOrders method.
func (r Workload_Citrix_Workspace_Order) GetWorkspaceNames() (resp datatypes.Workload_Citrix_Workspace_Response_Result, err error) {
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Workspace_Order", "getWorkspaceNames", nil, &r.Options, &resp)
	return
}

// This method will return the list of resources which could be cancelled using cancelWorkspaceResources
func (r Workload_Citrix_Workspace_Order) GetWorkspaceResources(vlanIdentifier *string) (resp datatypes.Workload_Citrix_Workspace_Response_Result, err error) {
	params := []interface{}{
		vlanIdentifier,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Workspace_Order", "getWorkspaceResources", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Workload_Citrix_Workspace_Order) PlaceWorkspaceOrder(orderContainer *datatypes.Workload_Citrix_Workspace_Order_Container) (resp datatypes.Workload_Citrix_Workspace_Response, err error) {
	params := []interface{}{
		orderContainer,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Workspace_Order", "placeWorkspaceOrder", params, &r.Options, &resp)
	return
}

// This service is used to verify that an order meets all the necessary requirements to purchase Citrix Virtual Apps and Desktops running on IBM cloud.
func (r Workload_Citrix_Workspace_Order) VerifyWorkspaceOrder(orderContainer *datatypes.Workload_Citrix_Workspace_Order_Container) (resp datatypes.Workload_Citrix_Workspace_Response, err error) {
	params := []interface{}{
		orderContainer,
	}
	err = r.Session.DoRequest("SoftLayer_Workload_Citrix_Workspace_Order", "verifyWorkspaceOrder", params, &r.Options, &resp)
	return
}

// no documentation yet
type Workload_Citrix_Workspace_Response struct {
	Session *session.Session
	Options sl.Options
}

// GetWorkloadCitrixWorkspaceResponseService returns an instance of the Workload_Citrix_Workspace_Response SoftLayer service
func GetWorkloadCitrixWorkspaceResponseService(sess *session.Session) Workload_Citrix_Workspace_Response {
	return Workload_Citrix_Workspace_Response{Session: sess}
}

func (r Workload_Citrix_Workspace_Response) Id(id int) Workload_Citrix_Workspace_Response {
	r.Options.Id = &id
	return r
}

func (r Workload_Citrix_Workspace_Response) Mask(mask string) Workload_Citrix_Workspace_Response {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Workload_Citrix_Workspace_Response) Filter(filter string) Workload_Citrix_Workspace_Response {
	r.Options.Filter = filter
	return r
}

func (r Workload_Citrix_Workspace_Response) Limit(limit int) Workload_Citrix_Workspace_Response {
	r.Options.Limit = &limit
	return r
}

func (r Workload_Citrix_Workspace_Response) Offset(offset int) Workload_Citrix_Workspace_Response {
	r.Options.Offset = &offset
	return r
}
