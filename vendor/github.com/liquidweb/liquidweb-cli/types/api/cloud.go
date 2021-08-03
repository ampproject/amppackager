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
package apiTypes

import (
	"fmt"
	"strings"
)

type CloudServerStatus struct {
	DetailedStatus string                         `json:"detailed_status" mapstructure:"detailed_status"`
	Progress       float64                        `json:"progress" mapstructure:"progress"`
	Running        []CloudServerStatusRunningData `json:"running" mapstructure:"running"`
	Status         string                         `json:"status" mapstructure:"status"`
}

type CloudServerStatusRunningData struct {
	CurrentStep    int64  `json:"current_step" mapstructure:"current_step"`
	DetailedStatus string `json:"detailed_status" mapstructure:"detailed_status"`
	Name           string `json:"name" mapstructure:"name"`
	Status         string `json:"status" mapstructure:"status"`
}

type CloudServerRebootResponse struct {
	Rebooted string `json:"rebooted" mapstructure:"rebooted"`
}

type CloudServerDetails struct {
	Accnt               int64                  `json:"accnt" mapstructure:"accnt"`
	ConfigId            int64                  `json:"config_id" mapstructure:"config_id"`
	Memory              int64                  `json:"memory" mapstructure:"memory"`
	Template            string                 `json:"template" mapstructure:"template"`
	Type                string                 `json:"type" mapstructure:"type"`
	BackupEnabled       int64                  `json:"backup_enabled" mapstructure:"backup_enabled"`
	BackupSize          float64                `json:"backup_size" mapstructure:"backup_size"`
	UniqId              string                 `json:"uniq_id" mapstructure:"uniq_id"`
	Vcpu                int64                  `json:"vcpu" mapstructure:"vcpu"`
	BackupPlan          string                 `json:"backup_plan" mapstructure:"backup_plan"`
	BandwidthQuota      string                 `json:"bandwidth_quota" mapstructure:"bandwidth_quota"`
	Ip                  string                 `json:"ip" mapstructure:"ip"`
	IpCount             int64                  `json:"ip_count" mapstructure:"ip_count"`
	ManageLevel         string                 `json:"manage_level" mapstructure:"manage_level"`
	CreateDate          string                 `json:"create_date" mapstructure:"create_date"`
	DiskSpace           int64                  `json:"diskspace" mapstructure:"diskspace"`
	Domain              string                 `json:"domain" mapstructure:"domain"`
	Active              int64                  `json:"active" mapstructure:"active"`
	BackupQuota         int64                  `json:"backup_quota" mapstructure:"backup_quota"`
	Zone                CloudServerDetailsZone `json:"zone" mapstructure:"zone"`
	ConfigDescription   string                 `json:"config_description" mapstructure:"config_description"`
	TemplateDescription string                 `json:"template_description" mapstructure:"template_description"`
	PrivateParent       string                 `json:"parent" mapstructure:"parent"`
}

type CloudServerDetailsZone struct {
	Id     int64                        `json:"id" mapstructure:"id"`
	Name   string                       `json:"name" mapstructure:"name"`
	Region CloudServerDetailsZoneRegion `json:"region" mapstructure:"region"`
}

type CloudServerDetailsZoneRegion struct {
	Id   int64  `json:"id" mapstructure:"id"`
	Name string `json:"name" mapstructure:"name"`
}

func (x CloudServerDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Domain: %s UniqId: %s\n", x.Domain, x.UniqId))

	slice = append(slice, fmt.Sprintf("\tIp: %s\n", x.Ip))
	slice = append(slice, fmt.Sprintf("\tIpCount: %d\n", x.IpCount))
	slice = append(slice, fmt.Sprintf("\tRegion: %s (id %d) Zone: %s (id %d)\n", x.Zone.Region.Name,
		x.Zone.Region.Id, x.Zone.Name, x.Zone.Id))

	if x.PrivateParent != "" {
		slice = append(slice, fmt.Sprintf("\tPrivate Parent Child on Private Parent [%s]\n", x.PrivateParent))
	} else {
		slice = append(slice, fmt.Sprintf("\tConfigId: %d\n", x.ConfigId))
	}
	slice = append(slice, fmt.Sprintf("\tConfigDescription: %s\n", x.ConfigDescription))
	slice = append(slice, fmt.Sprintf("\tVcpus: %d\n", x.Vcpu))
	slice = append(slice, fmt.Sprintf("\tMemory: %d\n", x.Memory))
	slice = append(slice, fmt.Sprintf("\tDiskSpace: %d\n", x.DiskSpace))
	slice = append(slice, fmt.Sprintf("\tTemplate: %s\n", x.Template))
	slice = append(slice, fmt.Sprintf("\tTemplateDescription: %s\n", x.TemplateDescription))
	slice = append(slice, fmt.Sprintf("\tType: %s\n", x.Type))
	slice = append(slice, fmt.Sprintf("\tBackupEnabled: %d\n", x.BackupEnabled))
	if x.BackupEnabled == 1 {
		slice = append(slice, fmt.Sprintf("\tBackupPlan: %s\n", x.BackupPlan))
		slice = append(slice, fmt.Sprintf("\tBackupSize: %.0f\n", x.BackupSize))
		if x.BackupQuota != 0 {
			slice = append(slice, fmt.Sprintf("\tBackupQuota: %d\n", x.BackupQuota))
		}
	}
	slice = append(slice, fmt.Sprintf("\tBandwidthQuota: %s\n", x.BandwidthQuota))
	slice = append(slice, fmt.Sprintf("\tManageLevel: %s\n", x.ManageLevel))
	slice = append(slice, fmt.Sprintf("\tActive: %d\n", x.Active))
	slice = append(slice, fmt.Sprintf("\tAccnt: %d\n", x.Accnt))

	return strings.Join(slice[:], "")
}

type CloudPrivateParentDetails struct {
	Accnt             int64                                     `json:"accnt" mapstructure:"accnt"`
	BucketUniqId      string                                    `json:"bucket_uniq_id" mapstructure:"bucket_uniq_id"`
	ConfigDescription string                                    `json:"config_description" mapstructure:"config_description"`
	ConfigId          int64                                     `json:"config_id" mapstructure:"config_id"`
	CreateDate        string                                    `json:"create_date" mapstructure:"create_date"`
	DiskDetails       CloudPrivateParentDetailsEntryDiskDetails `json:"diskDetails" mapstructure:"diskDetails"`
	Domain            string                                    `json:"domain" mapstructure:"domain"`
	Id                int64                                     `json:"id" mapstructure:"id"`
	LicenseState      string                                    `json:"license_state" mapstructure:"license_state"`
	RegionId          int64                                     `json:"region_id" mapstructure:"region_id"`
	Resources         CloudPrivateParentDetailsEntryResource    `json:"resources" mapstructure:"resources"`
	SalesforceAsset   string                                    `json:"salesforce_asset" mapstructure:"salesforce_asset"`
	Status            string                                    `json:"status" mapstructure:"status"`
	Subaccnt          int64                                     `json:"subaccnt" mapstructure:"subaccnt"`
	Type              string                                    `json:"type" mapstructure:"type"`
	UniqId            string                                    `json:"uniq_id" mapstructure:"uniq_id"`
	Vcpu              int64                                     `json:"vcpu" mapstructure:"vcpu"`
	Zone              CloudPrivateParentDetailsEntryZone        `json:"zone" mapstructure:"zone"`
}

type CloudPrivateParentDetailsEntryResource struct {
	DiskSpace CloudPrivateParentDetailsEntryResourceEntry `json:"diskspace" mapstructure:"diskspace"`
	Memory    CloudPrivateParentDetailsEntryResourceEntry `json:"memory" mapstructure:"memory"`
}

type CloudPrivateParentDetailsEntryResourceEntry struct {
	Free  int64 `json:"free" mapstructure:"free"`
	Total int64 `json:"total" mapstructure:"total"`
	Used  int64 `json:"used" mapstructure:"used"`
}

type CloudPrivateParentDetailsEntryDiskDetails struct {
	Allocated int64 `json:"allocated" mapstructure:"allocated"`
	Snapshots int64 `json:"snapshots" mapstructure:"snapshots"`
}

type CloudPrivateParentDetailsEntryZone struct {
	AvailabilityZone string                                   `json:"availability_zone" mapstructure:"availability_zone"`
	Description      string                                   `json:"description" mapstructure:"description"`
	HvType           string                                   `json:"hv_type" mapstructure:"hv_type"`
	Id               int64                                    `json:"id" mapstructure:"id"`
	Legacy           int64                                    `json:"legacy" mapstructure:"legacy"`
	Name             string                                   `json:"name" mapstructure:"name"`
	Region           CloudPrivateParentDetailsEntryZoneRegion `json:"region" mapstructure:"region"`
	Status           string                                   `json:"status" mapstructure:"status"`
	ValidSourceHvs   []string                                 `json:"valid_source_hvs" mapstructure:"valid_source_hvs"`
}

type CloudPrivateParentDetailsEntryZoneRegion struct {
	Id   int64  `json:"id" mapstructure:"id"`
	Name string `json:"name" mapstructure:"name"`
}

func (x CloudPrivateParentDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Private Parent: %s\n", x.Domain))
	slice = append(slice, fmt.Sprintf("\tUniqId: %s\n", x.UniqId))
	slice = append(slice, fmt.Sprintf("\tStatus: %s\n", x.Status))
	slice = append(slice, fmt.Sprintf("\tConfigId: %d\n", x.ConfigId))
	slice = append(slice, fmt.Sprintf("\tConfigDescription: %s\n", x.ConfigDescription))
	slice = append(slice, fmt.Sprintf("\tVcpus: %d\n", x.Vcpu))

	// resources
	slice = append(slice, fmt.Sprintf("\tResource Usage:\n"))
	// diskspace
	slice = append(slice, fmt.Sprintf("\t\tDiskSpace:\n"))
	slice = append(slice, fmt.Sprintf("\t\t\t%d out of %d used; free %d\n", x.Resources.DiskSpace.Used,
		x.Resources.DiskSpace.Total, x.Resources.DiskSpace.Free))
	// memory
	slice = append(slice, fmt.Sprintf("\t\tMemory:\n"))
	slice = append(slice, fmt.Sprintf("\t\t\t%d out of %d used; free %d\n", x.Resources.Memory.Used,
		x.Resources.Memory.Total, x.Resources.Memory.Free))

	slice = append(slice, fmt.Sprintf("\tRegion %s (id %d) - %s (id %d)\n", x.Zone.Region.Name, x.Zone.Region.Id,
		x.Zone.Description, x.Zone.Id))

	slice = append(slice, fmt.Sprintf("\tHypervisor: %s\n", x.Zone.HvType))
	slice = append(slice, fmt.Sprintf("\tCreateDate: %s\n", x.CreateDate))
	slice = append(slice, fmt.Sprintf("\tLicenseState: %s\n", x.LicenseState))

	return strings.Join(slice[:], "")
}

type CloudConfigDetails struct {
	Id                int64         `json:"id" mapstructure:"id"`
	Active            int64         `json:"active" mapstructure:"active"`
	Available         int64         `json:"available" mapstructure:"available"`
	Category          string        `json:"category" mapstructure:"category"`
	Description       string        `json:"description" mapstructure:"description"`
	Disk              int64         `json:"disk,omitempty" mapstructure:"disk"`
	Featured          int64         `json:"featured" mapstructure:"featured"`
	Memory            int64         `json:"memory,omitempty" mapstructure:"memory"`
	Vcpu              int64         `json:"vcpu,omitempty" mapstructure:"vcpu"`
	ZoneAvailability  []map[int]int `json:"zone_availability" mapstructure:"zone_availability"`
	Retired           int64         `json:"retired,omitempty" mapstructure:"retired"`
	RamTotal          int64         `json:"ram_total,omitempty" mapstructure:"ram_total"`
	RamAvailable      int64         `json:"ram_available,omitempty" mapstructure:"ram_available"`
	RaidLevel         int64         `json:"raid_level,omitempty" mapstructure:"raid_level"`
	DiskType          string        `json:"disk_type,omitempty" mapstructure:"disk_type"`
	DiskTotal         int64         `json:"disk_total,omitempty" mapstructure:"disk_total"`
	DiskCount         int64         `json:"disk_count,omitempty" mapstructure:"disk_count"`
	CpuSpeed          int64         `json:"cpu_speed,omitempty" mapstructure:"cpu_speed"`
	CpuModel          string        `json:"cpu_model,omitempty" mapstructure:"cpu_model"`
	CpuHyperthreading int64         `json:"cpu_hyperthreading,omitempty" mapstructure:"cpu_hyperthreading"`
	CpuCount          int64         `json:"cpu_count,omitempty" mapstructure:"cpu_count"`
	CpuCores          int64         `json:"cpu_cores,omitempty" mapstructure:"cpu_cores"`
}

type CloudServerDestroyResponse struct {
	Destroyed string `json:"destroyed" mapstructure:"destroyed"`
}

type CloudServerShutdownResponse struct {
	Shutdown string `json:"shutdown" mapstructure:"shutdown"`
}

type CloudServerStartResponse struct {
	Started string `json:"started" mapstructure:"started"`
}

type CloudPrivateParentDeleteResponse struct {
	Deleted string `json:"deleted" mapstructure:"deleted"`
}

type CloudImageCreateResponse struct {
	Created string `json:"created" mapstructure:"created"`
}

type CloudImageRestoreResponse struct {
	Reimaged string `json"reimaged" mapstructure:"reimaged"`
}

type CloudBackupRestoreResponse struct {
	Restored string `json:"restored" mapstructure:"restored"`
}

type CloudImageDeleteResponse struct {
	Deleted int64 `json:"deleted" mapstructure:"deleted"`
}

type CloudTemplateRestoreResponse struct {
	Reimaged string `json:"reimaged" mapstructure:"reimaged"`
}

type CloudServerIsBlockStorageOptimized struct {
	IsOptimized bool `json:"is_optimized" mapstructure:"is_optimized"`
}

type CloudServerIsBlockStorageOptimizedSetResponse struct {
	Updated string `json:"updated" mapstructure:"updated"`
}

type CloudServerCloneResponse struct {
	Accnt               int64                        `json:"accnt" mapstructure:"accnt"`
	Active              int64                        `json:"active" mapstructure:"active"`
	BackupEnabled       int64                        `json:"backup_enabled" mapstructure:"backup_enabled"`
	BackupPlan          string                       `json:"backup_plan" mapstructure:"backup_plan"`
	BackupQuota         int64                        `json:"backup_quota" mapstructure:"backup_quota"`
	BackupSize          float64                      `json:"backup_size" mapstructure:"backup_size"`
	BandwidthQuota      string                       `json:"bandwidth_quota" mapstructure:"bandwidth_quota"`
	Categories          []interface{}                `json:"categories" mapstructure:"categories"`
	ConfigDescription   string                       `json:"config_description" mapstructure:"config_description"`
	ConfigId            int64                        `json:"config_id" mapstructure:"config_id"`
	CreateDate          string                       `json:"create_date" mapstructure:"create_date"`
	Description         string                       `json:"description" mapstructure:"description"`
	Diskspace           int64                        `json:"diskspace" mapstructure:"diskspace"`
	Domain              string                       `json:"domain" mapstructure:"domain"`
	HvType              string                       `json:"hv_type" mapstructure:"hv_type"`
	Instance            interface{}                  `json:"instance" mapstructure:"instance"`
	Ip                  string                       `json:"ip" mapstructure:"ip"`
	IpCount             int64                        `json:"ip_count" mapstructure:"ip_count"`
	ManageLevel         string                       `json:"manage_level" mapstructure:"manage_level"`
	Memory              int64                        `json:"memory" mapstructure:"memory"`
	Parent              string                       `json:"parent" mapstructure:"parent"`
	RegionId            int64                        `json:"region_id" mapstructure:"region_id"`
	ShortDescription    string                       `json:"shortDescription" mapstructure:"shortDescription"`
	Status              string                       `json:"status" mapstructure:"status"`
	Template            string                       `json:"template" mapstructure:"template"`
	TemplateDescription string                       `json:"template_description" mapstructure:"template_description"`
	Type                string                       `json:"type" mapstructure:"type"`
	UniqId              string                       `json:"uniq_id" mapstructure:"uniq_id"`
	ValidSourceHvs      map[string]int64             `json:"valid_source_hvs" mapstructure:"valid_source_hvs"`
	Vcpu                int64                        `json:"vcpu" mapstructure:"vcpu"`
	Zone                CloudServerCloneResponseZone `json:"zone" mapstructure:"zone"`
}

type CloudServerCloneResponseZone struct {
	Id     int64                              `json:"id" mapstructure:"id"`
	Name   string                             `json:"name" mapstructure:"name"`
	Region CloudServerCloneResponseZoneRegion `json:"region" mapstructure:"region"`
}

type CloudServerCloneResponseZoneRegion struct {
	HostPrefix string `json:"host_prefix" mapstructure:"host_prefix"`
	Id         int64  `json:"id" mapstructure:"id"`
	Name       string `json:"name" mapstructure:"name"`
}

type CloudImageDetails struct {
	Accnt               int64                    `json:"accnt" mapstructure:"accnt"`
	Features            []map[string]interface{} `json:"features" mapstructure:"features"`
	HvType              string                   `json:"hv_type" mapstructure:"hv_type"`
	Id                  int64                    `json:"id" mapstructure:"id"`
	Name                string                   `json:"name" mapstructure:"name"`
	Size                float64                  `json:"size" mapstructure:"size"`
	SourceHostname      string                   `json:"source_hostname" mapstructure:"source_hostname"`
	SourceUniqId        string                   `json:"source_uniq_id" mapstructure:"source_uniq_id"`
	Template            string                   `json:"template" mapstructure:"template"`
	TemplateDescription string                   `json:"template_description" mapstructure:"template_description"`
	TimeTaken           string                   `json:"time_taken" mapstructure:"time_taken"`
}

func (x CloudImageDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Cloud Image: %s\n", x.Name))
	slice = append(slice, fmt.Sprintf("\tId: %d\n", x.Id))
	slice = append(slice, fmt.Sprintf("\tsize: %.2f\n", x.Size))
	if x.SourceHostname != "" {
		slice = append(slice, fmt.Sprintf("\tSource Hostname: %s\n", x.SourceHostname))
	}
	if x.SourceUniqId != "" {
		slice = append(slice, fmt.Sprintf("\tSource UniqId: %s\n", x.SourceUniqId))
	}
	slice = append(slice, fmt.Sprintf("\tTemplate: %s\n", x.Template))
	slice = append(slice, fmt.Sprintf("\tTemplate Description: %s\n", x.TemplateDescription))
	slice = append(slice, fmt.Sprintf("\tTime Taken: %s\n", x.TimeTaken))
	slice = append(slice, fmt.Sprintf("\tHypervisor: %s\n", x.HvType))
	slice = append(slice, fmt.Sprintf("\tFeatures: %+v\n", x.Features))
	slice = append(slice, fmt.Sprintf("\tAccount: %d\n", x.Accnt))

	return strings.Join(slice[:], "")
}

type CloudBackupDetails struct {
	Accnt     int64                    `json:"accnt" mapstructure:"accnt"`
	Features  []map[string]interface{} `json:"features" mapstructure:"features"`
	HvType    string                   `json:"hv_type" mapstructure:"hv_type"`
	Id        int64                    `json:"id" mapstructure:"id"`
	Name      string                   `json:"name" mapstructure:"name"`
	Size      float64                  `json:"size" mapstructure:"size"`
	Template  string                   `json:"template" mapstructure:"template"`
	TimeTaken string                   `json:"time_taken" mapstructure:"time_taken"`
	UniqId    string                   `json:"uniq_id" mapstructure:"uniq_id"`
}

func (x CloudBackupDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Cloud Backup Id: %d\n", x.Id))
	slice = append(slice, fmt.Sprintf("\tUniqId: %s\n", x.UniqId))
	slice = append(slice, fmt.Sprintf("\tName: %s\n", x.Name))
	slice = append(slice, fmt.Sprintf("\tsize: %.2f\n", x.Size))
	slice = append(slice, fmt.Sprintf("\tTemplate: %s\n", x.Template))
	slice = append(slice, fmt.Sprintf("\tTime Taken: %s\n", x.TimeTaken))
	slice = append(slice, fmt.Sprintf("\tHypervisor: %s\n", x.HvType))
	slice = append(slice, fmt.Sprintf("\tFeatures: %+v\n", x.Features))
	slice = append(slice, fmt.Sprintf("\tAccount: %d\n", x.Accnt))

	return strings.Join(slice[:], "")
}

type CloudNetworkVipDetails struct {
	Active       int64    `json:"active" mapstructure:"active"`
	ActiveStatus string   `json:"activeStatus" mapstructure:"activeStatus"`
	Domain       string   `json:"domain" mapstructure:"domain"`
	UniqId       string   `json:"uniq_id" mapstructure:"uniq_id"`
	Ip           string   `json:"ip" mapstructure:"ip"`
	PrivateIp    []string `json:"private_ip" mapstructure:"private_ip"`
}

func (x CloudNetworkVipDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("VIP Details:\n"))
	slice = append(slice, fmt.Sprintf("\tActive: %d\n", x.Active))
	slice = append(slice, fmt.Sprintf("\tActive Status: %s\n", x.ActiveStatus))
	slice = append(slice, fmt.Sprintf("\tName: %s\n", x.Domain))
	slice = append(slice, fmt.Sprintf("\tUniqId: %s\n", x.UniqId))
	slice = append(slice, fmt.Sprintf("\tIP: %s\n", x.Ip))
	slice = append(slice, fmt.Sprintf("\tPrivate IP: %+v\n", x.PrivateIp))

	return strings.Join(slice[:], "")
}

type CloudNetworkVipDestroyResponse struct {
	Destroyed string `json:"destroyed" mapstructure:"destroyed"`
}

type CloudNetworkVipAssetListAlsoWithZoneResponse struct {
	Active   int64                  `json:"active" mapstructure:"active"`
	Domain   string                 `json:"domain" mapstructure:"domain"`
	Ip       string                 `json:"ip" mapstructure:"ip"`
	RegionId int64                  `json:"region_id" mapstructure:"region_id"`
	Status   string                 `json:"status" mapstructure:"status"`
	Type     string                 `json:"type" mapstructure:"type"`
	UniqId   string                 `json:"uniq_id" mapstructure:"uniq_id"`
	Zone     CloudServerDetailsZone `json:"zone" mapstructure:"zone"`
}

type CloudNetworkPrivateAttachResponse struct {
	Attached string `json:"attached" mapstructure:"attached"`
}

type CloudNetworkPrivateDetachResponse struct {
	Detached string `json:"detached" mapstructure:"detached"`
}

type CloudNetworkPrivateGetIpResponse struct {
	UniqId string `json:"uniq_id" mapstructure:"uniq_id"`
	Legacy bool   `json:"legacy" mapstructure:"legacy"`
	Ip     string `json:"ip" mapstructure:"ip"`
}

func (self CloudNetworkPrivateGetIpResponse) String() string {
	var slice []string

	if self.Ip == "" {
		slice = append(slice, fmt.Sprintf("Cloud Server [%s] is not attached to a Private Network\n", self.UniqId))
	} else {
		slice = append(slice, fmt.Sprintf("Cloud Server [%s] is attached to a Private Network\n", self.UniqId))
		slice = append(slice, fmt.Sprintf("\tIP: %s\n", self.Ip))
		slice = append(slice, fmt.Sprintf("\tLegacy: %t\n", self.Legacy))
	}

	return strings.Join(slice[:], "")
}

type CloudNetworkPrivateIsAttachedResponse struct {
	IsAttached bool `json:"is_attached" mapstructure:"is_attached"`
}

type CloudBlockStorageVolumeDetails struct {
	AttachedTo       []CloudBlockStorageVolumeDetailsAttachedTo `json:"attachedTo" mapstructure:"attachedTo"`
	CrossAttach      bool                                       `json:"cross_attach" mapstructure:"cross_attach"`
	Domain           string                                     `json:"domain" mapstructure:"domain"`
	Label            string                                     `json:"label" mapstructure:"label"`
	Size             int64                                      `json:"size" mapstructure:"size"`
	Status           string                                     `json:"status" mapstructure:"status"`
	UniqId           string                                     `json:"uniq_id" mapstructure:"uniq_id"`
	ZoneAvailability []int64                                    `json:"zoneAvailability" mapstructure:"zoneAvailability"`
}

type CloudBlockStorageVolumeDetailsAttachedTo struct {
	Device   string `json:"device" mapstructure:"device"`
	Resource string `json:"resource" mapstructure:"resource"`
}

func (x CloudBlockStorageVolumeDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Volume: %s\n", x.Domain))
	slice = append(slice, fmt.Sprintf("\tStatus: %s\n", x.Status))
	slice = append(slice, fmt.Sprintf("\tUniqId: %s\n", x.UniqId))
	slice = append(slice, fmt.Sprintf("\tAttaches:\n"))
	for _, entry := range x.AttachedTo {
		if entry.Resource != "" {
			slice = append(slice, fmt.Sprintf("\t\tAttached to %s as device %s\n", entry.Resource, entry.Device))
		}
	}

	slice = append(slice, fmt.Sprintf("\tCross Attach Enabled: %t\n", x.CrossAttach))
	slice = append(slice, fmt.Sprintf("\tSize: %d\n", x.Size))
	slice = append(slice, fmt.Sprintf("\tLabel: %s\n", x.Label))
	slice = append(slice, fmt.Sprintf("\tZone Availability: %+v\n", x.ZoneAvailability))

	return strings.Join(slice[:], "")
}

type CloudBlockStorageVolumeDelete struct {
	Deleted string `json:"deleted" mapstructure:"deleted"`
}

type CloudBlockStorageVolumeAttach struct {
	Attached string `json:"attached" mapstructure:"attached"`
	To       string `json:"to" mapstructure:"to"`
}

type CloudBlockStorageVolumeDetach struct {
	Detached     string `json:"detached" mapstructure:"detached"`
	DetachedFrom string `json:"detached_from" mapstructure:"detached_from"`
}

type CloudBlockStorageVolumeResize struct {
	NewSize int64  `json:"new_size" mapstructure:"new_size"`
	OldSize int64  `json:"old_size" mapstructure:"old_size"`
	UniqId  string `json:"uniq_id" mapstructure:"uniq_id"`
}

type CloudServerResizeExpectation struct {
	DiskDifference   int64    `json:"diskDifference" mapstructure:"diskDifference"`
	MemoryDifference int64    `json:"memoryDifference" mapstructure:"memoryDifference"`
	VcpuDifference   int64    `json:"vcpuDifference" mapstructure:"vcpuDifference"`
	RebootRequired   FlexBool `json:"rebootRequired" mapstructure:"rebootRequired"`
}

type FlexBool bool

func (self *FlexBool) UnmarshalJSON(data []byte) error {
	str := string(data)

	if str == "1" || str == "true" {
		*self = true
	} else if str == "0" || str == "false" {
		*self = false
	} else {
		return fmt.Errorf("Boolean unmarshal error: invalid input %s", str)
	}

	return nil
}

type CloudObjectStoreDetails struct {
	Accnt       int64                              `json:"accnt" mapstructure:"accnt"`
	Caps        []CloudObjectStoreDetailsCapsEntry `json:"caps" mapstructure:"caps"`
	DisplayName string                             `json:"display_name" mapstructure:"display_name"`
	Host        string                             `json:"host" mapstructure:"host"`
	Keys        []CloudObjectStoreKeyDetails       `json:"keys" mapstructure:"keys"`
	MaxBuckets  int64                              `json:"max_buckets" mapstructure:"max_buckets"`
	Suspended   bool                               `json:"suspended" mapstructure:"suspended"`
	UniqId      string                             `json:"uniq_id" mapstructure:"uniq_id"`
	UserId      string                             `json:"user_id" mapstructure:"user_id"`
}

type CloudObjectStoreKeyDetails struct {
	AccessKey string `json:"access_key" mapstructure:"access_key"`
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
	User      string `json:"user" mapstructure:"user"`
}

type CloudObjectStoreDetailsCapsEntry struct {
	Perm string `json:"perm" mapstructure:"perm"`
	Type string `json:"type" mapstructure:"type"`
}

func (x CloudObjectStoreDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Object Store UniqId: %s\n", x.UniqId))
	slice = append(slice, fmt.Sprintf("\tDisplay Name: %s\n", x.DisplayName))
	slice = append(slice, fmt.Sprintf("\tHost: %s\n", x.Host))
	slice = append(slice, fmt.Sprintf("\tUserId: %s\n", x.UserId))
	slice = append(slice, fmt.Sprintf("\tMax Buckets: %d\n", x.MaxBuckets))
	slice = append(slice, fmt.Sprintf("\tCaps:\n"))
	for _, key := range x.Caps {
		slice = append(slice, fmt.Sprintf("\t\tPerm: %s\n", key.Perm))
		slice = append(slice, fmt.Sprintf("\t\tType: %s\n", key.Type))
	}
	slice = append(slice, fmt.Sprintf("\tKeys:\n"))
	for _, key := range x.Keys {
		slice = append(slice, fmt.Sprintf("\t\tAccess Key: %s\n", key.AccessKey))
		slice = append(slice, fmt.Sprintf("\t\tSecret Key: %s\n", key.SecretKey))
		slice = append(slice, fmt.Sprintf("\t\tUser: %s\n", key.User))
	}
	if x.Accnt != 0 {
		slice = append(slice, fmt.Sprintf("\tAccount: %d\n", x.Accnt))
	}
	slice = append(slice, fmt.Sprintf("\tSuspended: %t\n", x.Suspended))

	return strings.Join(slice[:], "")
}

type CloudObjectStoreDelete struct {
	Deleted string `json:"deleted" mapstructure:"deleted"`
}

type CloudObjectStoreDiskSpace struct {
	Buckets []map[string]interface{} `json:"buckets" mapstructure:"buckets"`
	Total   int64                    `json:"total" mapstructure:"total"`
}

type CloudObjectStoreDeleteKey struct {
	Deleted string `json:"deleted" mapstructure:"deleted"`
}
