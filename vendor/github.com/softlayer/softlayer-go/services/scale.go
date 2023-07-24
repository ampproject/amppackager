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

	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// no documentation yet
type Scale_Asset struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleAssetService returns an instance of the Scale_Asset SoftLayer service
func GetScaleAssetService(sess *session.Session) Scale_Asset {
	return Scale_Asset{Session: sess}
}

func (r Scale_Asset) Id(id int) Scale_Asset {
	r.Options.Id = &id
	return r
}

func (r Scale_Asset) Mask(mask string) Scale_Asset {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_Asset) Filter(filter string) Scale_Asset {
	r.Options.Filter = filter
	return r
}

func (r Scale_Asset) Limit(limit int) Scale_Asset {
	r.Options.Limit = &limit
	return r
}

func (r Scale_Asset) Offset(offset int) Scale_Asset {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
type Scale_Asset_Hardware struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleAssetHardwareService returns an instance of the Scale_Asset_Hardware SoftLayer service
func GetScaleAssetHardwareService(sess *session.Session) Scale_Asset_Hardware {
	return Scale_Asset_Hardware{Session: sess}
}

func (r Scale_Asset_Hardware) Id(id int) Scale_Asset_Hardware {
	r.Options.Id = &id
	return r
}

func (r Scale_Asset_Hardware) Mask(mask string) Scale_Asset_Hardware {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_Asset_Hardware) Filter(filter string) Scale_Asset_Hardware {
	r.Options.Filter = filter
	return r
}

func (r Scale_Asset_Hardware) Limit(limit int) Scale_Asset_Hardware {
	r.Options.Limit = &limit
	return r
}

func (r Scale_Asset_Hardware) Offset(offset int) Scale_Asset_Hardware {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
type Scale_Asset_Virtual_Guest struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleAssetVirtualGuestService returns an instance of the Scale_Asset_Virtual_Guest SoftLayer service
func GetScaleAssetVirtualGuestService(sess *session.Session) Scale_Asset_Virtual_Guest {
	return Scale_Asset_Virtual_Guest{Session: sess}
}

func (r Scale_Asset_Virtual_Guest) Id(id int) Scale_Asset_Virtual_Guest {
	r.Options.Id = &id
	return r
}

func (r Scale_Asset_Virtual_Guest) Mask(mask string) Scale_Asset_Virtual_Guest {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_Asset_Virtual_Guest) Filter(filter string) Scale_Asset_Virtual_Guest {
	r.Options.Filter = filter
	return r
}

func (r Scale_Asset_Virtual_Guest) Limit(limit int) Scale_Asset_Virtual_Guest {
	r.Options.Limit = &limit
	return r
}

func (r Scale_Asset_Virtual_Guest) Offset(offset int) Scale_Asset_Virtual_Guest {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
type Scale_Group struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleGroupService returns an instance of the Scale_Group SoftLayer service
func GetScaleGroupService(sess *session.Session) Scale_Group {
	return Scale_Group{Session: sess}
}

func (r Scale_Group) Id(id int) Scale_Group {
	r.Options.Id = &id
	return r
}

func (r Scale_Group) Mask(mask string) Scale_Group {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_Group) Filter(filter string) Scale_Group {
	r.Options.Filter = filter
	return r
}

func (r Scale_Group) Limit(limit int) Scale_Group {
	r.Options.Limit = &limit
	return r
}

func (r Scale_Group) Offset(offset int) Scale_Group {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
type Scale_LoadBalancer struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleLoadBalancerService returns an instance of the Scale_LoadBalancer SoftLayer service
func GetScaleLoadBalancerService(sess *session.Session) Scale_LoadBalancer {
	return Scale_LoadBalancer{Session: sess}
}

func (r Scale_LoadBalancer) Id(id int) Scale_LoadBalancer {
	r.Options.Id = &id
	return r
}

func (r Scale_LoadBalancer) Mask(mask string) Scale_LoadBalancer {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_LoadBalancer) Filter(filter string) Scale_LoadBalancer {
	r.Options.Filter = filter
	return r
}

func (r Scale_LoadBalancer) Limit(limit int) Scale_LoadBalancer {
	r.Options.Limit = &limit
	return r
}

func (r Scale_LoadBalancer) Offset(offset int) Scale_LoadBalancer {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
type Scale_Member struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleMemberService returns an instance of the Scale_Member SoftLayer service
func GetScaleMemberService(sess *session.Session) Scale_Member {
	return Scale_Member{Session: sess}
}

func (r Scale_Member) Id(id int) Scale_Member {
	r.Options.Id = &id
	return r
}

func (r Scale_Member) Mask(mask string) Scale_Member {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_Member) Filter(filter string) Scale_Member {
	r.Options.Filter = filter
	return r
}

func (r Scale_Member) Limit(limit int) Scale_Member {
	r.Options.Limit = &limit
	return r
}

func (r Scale_Member) Offset(offset int) Scale_Member {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
type Scale_Member_Virtual_Guest struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleMemberVirtualGuestService returns an instance of the Scale_Member_Virtual_Guest SoftLayer service
func GetScaleMemberVirtualGuestService(sess *session.Session) Scale_Member_Virtual_Guest {
	return Scale_Member_Virtual_Guest{Session: sess}
}

func (r Scale_Member_Virtual_Guest) Id(id int) Scale_Member_Virtual_Guest {
	r.Options.Id = &id
	return r
}

func (r Scale_Member_Virtual_Guest) Mask(mask string) Scale_Member_Virtual_Guest {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_Member_Virtual_Guest) Filter(filter string) Scale_Member_Virtual_Guest {
	r.Options.Filter = filter
	return r
}

func (r Scale_Member_Virtual_Guest) Limit(limit int) Scale_Member_Virtual_Guest {
	r.Options.Limit = &limit
	return r
}

func (r Scale_Member_Virtual_Guest) Offset(offset int) Scale_Member_Virtual_Guest {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
type Scale_Network_Vlan struct {
	Session *session.Session
	Options sl.Options
}

// GetScaleNetworkVlanService returns an instance of the Scale_Network_Vlan SoftLayer service
func GetScaleNetworkVlanService(sess *session.Session) Scale_Network_Vlan {
	return Scale_Network_Vlan{Session: sess}
}

func (r Scale_Network_Vlan) Id(id int) Scale_Network_Vlan {
	r.Options.Id = &id
	return r
}

func (r Scale_Network_Vlan) Mask(mask string) Scale_Network_Vlan {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Scale_Network_Vlan) Filter(filter string) Scale_Network_Vlan {
	r.Options.Filter = filter
	return r
}

func (r Scale_Network_Vlan) Limit(limit int) Scale_Network_Vlan {
	r.Options.Limit = &limit
	return r
}

func (r Scale_Network_Vlan) Offset(offset int) Scale_Network_Vlan {
	r.Options.Offset = &offset
	return r
}
