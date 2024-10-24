// Code generated by protoc-gen-goext. DO NOT EDIT.

package opensearch

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1/config"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *GetClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClustersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListClustersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClustersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClustersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClustersResponse) SetClusters(v []*Cluster) {
	m.Clusters = v
}

func (m *ListClustersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *CreateClusterRequest) SetConfigSpec(v *ConfigCreateSpec) {
	m.ConfigSpec = v
}

func (m *CreateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *CreateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *CreateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateClusterRequest) SetConfigSpec(v *ConfigUpdateSpec) {
	m.ConfigSpec = v
}

func (m *UpdateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *UpdateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *UpdateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *UpdateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *ListClusterLogsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *ListClusterLogsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *ListClusterLogsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterLogsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterLogsRequest) SetAlwaysNextPageToken(v bool) {
	m.AlwaysNextPageToken = v
}

func (m *ListClusterLogsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterLogsRequest) SetServiceType(v ListClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *LogRecord) SetTimestamp(v *timestamppb.Timestamp) {
	m.Timestamp = v
}

func (m *LogRecord) SetMessage(v map[string]string) {
	m.Message = v
}

func (m *ListClusterLogsResponse) SetLogs(v []*LogRecord) {
	m.Logs = v
}

func (m *ListClusterLogsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *StreamLogRecord) SetRecord(v *LogRecord) {
	m.Record = v
}

func (m *StreamLogRecord) SetNextRecordToken(v string) {
	m.NextRecordToken = v
}

func (m *StreamClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StreamClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *StreamClusterLogsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *StreamClusterLogsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *StreamClusterLogsRequest) SetRecordToken(v string) {
	m.RecordToken = v
}

func (m *StreamClusterLogsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *StreamClusterLogsRequest) SetServiceType(v StreamClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *ListClusterOperationsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListClusterOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterHostsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterHostsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterHostsResponse) SetHosts(v []*Host) {
	m.Hosts = v
}

func (m *ListClusterHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *MoveClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterMetadata) SetSourceFolderId(v string) {
	m.SourceFolderId = v
}

func (m *MoveClusterMetadata) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *StartClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ConfigCreateSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigCreateSpec) SetAdminPassword(v string) {
	m.AdminPassword = v
}

func (m *ConfigCreateSpec) SetOpensearchSpec(v *OpenSearchCreateSpec) {
	m.OpensearchSpec = v
}

func (m *ConfigCreateSpec) SetDashboardsSpec(v *DashboardsCreateSpec) {
	m.DashboardsSpec = v
}

func (m *ConfigCreateSpec) SetAccess(v *Access) {
	m.Access = v
}

func (m *KeystoreSetting) SetName(v string) {
	m.Name = v
}

func (m *KeystoreSetting) SetValue(v string) {
	m.Value = v
}

type OpenSearchCreateSpec_Config = isOpenSearchCreateSpec_Config

func (m *OpenSearchCreateSpec) SetConfig(v OpenSearchCreateSpec_Config) {
	m.Config = v
}

func (m *OpenSearchCreateSpec) SetPlugins(v []string) {
	m.Plugins = v
}

func (m *OpenSearchCreateSpec) SetNodeGroups(v []*OpenSearchCreateSpec_NodeGroup) {
	m.NodeGroups = v
}

func (m *OpenSearchCreateSpec) SetOpensearchConfig_2(v *config.OpenSearchConfig2) {
	m.Config = &OpenSearchCreateSpec_OpensearchConfig_2{
		OpensearchConfig_2: v,
	}
}

func (m *OpenSearchCreateSpec) SetKeystoreSettings(v []*KeystoreSetting) {
	m.KeystoreSettings = v
}

func (m *OpenSearchCreateSpec_NodeGroup) SetName(v string) {
	m.Name = v
}

func (m *OpenSearchCreateSpec_NodeGroup) SetResources(v *Resources) {
	m.Resources = v
}

func (m *OpenSearchCreateSpec_NodeGroup) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *OpenSearchCreateSpec_NodeGroup) SetZoneIds(v []string) {
	m.ZoneIds = v
}

func (m *OpenSearchCreateSpec_NodeGroup) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *OpenSearchCreateSpec_NodeGroup) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *OpenSearchCreateSpec_NodeGroup) SetRoles(v []OpenSearch_GroupRole) {
	m.Roles = v
}

func (m *DashboardsCreateSpec) SetNodeGroups(v []*DashboardsCreateSpec_NodeGroup) {
	m.NodeGroups = v
}

func (m *DashboardsCreateSpec_NodeGroup) SetName(v string) {
	m.Name = v
}

func (m *DashboardsCreateSpec_NodeGroup) SetResources(v *Resources) {
	m.Resources = v
}

func (m *DashboardsCreateSpec_NodeGroup) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *DashboardsCreateSpec_NodeGroup) SetZoneIds(v []string) {
	m.ZoneIds = v
}

func (m *DashboardsCreateSpec_NodeGroup) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *DashboardsCreateSpec_NodeGroup) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *ConfigUpdateSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigUpdateSpec) SetAdminPassword(v string) {
	m.AdminPassword = v
}

func (m *ConfigUpdateSpec) SetOpensearchSpec(v *OpenSearchClusterUpdateSpec) {
	m.OpensearchSpec = v
}

func (m *ConfigUpdateSpec) SetDashboardsSpec(v *DashboardsClusterUpdateSpec) {
	m.DashboardsSpec = v
}

func (m *ConfigUpdateSpec) SetAccess(v *Access) {
	m.Access = v
}

type OpenSearchClusterUpdateSpec_Config = isOpenSearchClusterUpdateSpec_Config

func (m *OpenSearchClusterUpdateSpec) SetConfig(v OpenSearchClusterUpdateSpec_Config) {
	m.Config = v
}

func (m *OpenSearchClusterUpdateSpec) SetPlugins(v []string) {
	m.Plugins = v
}

func (m *OpenSearchClusterUpdateSpec) SetOpensearchConfig_2(v *config.OpenSearchConfig2) {
	m.Config = &OpenSearchClusterUpdateSpec_OpensearchConfig_2{
		OpensearchConfig_2: v,
	}
}

func (m *OpenSearchClusterUpdateSpec) SetSetKeystoreSettings(v []*KeystoreSetting) {
	m.SetKeystoreSettings = v
}

func (m *OpenSearchClusterUpdateSpec) SetRemoveKeystoreSettings(v []string) {
	m.RemoveKeystoreSettings = v
}

func (m *BackupClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BackupClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterRequest) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RestoreClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *RestoreClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *RestoreClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *RestoreClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *RestoreClusterRequest) SetConfigSpec(v *ConfigCreateSpec) {
	m.ConfigSpec = v
}

func (m *RestoreClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *RestoreClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *RestoreClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *RestoreClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *RestoreClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *RestoreClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *RestoreClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterMetadata) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RescheduleMaintenanceRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceRequest) SetRescheduleType(v RescheduleMaintenanceRequest_RescheduleType) {
	m.RescheduleType = v
}

func (m *RescheduleMaintenanceRequest) SetDelayedUntil(v *timestamppb.Timestamp) {
	m.DelayedUntil = v
}

func (m *RescheduleMaintenanceMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceMetadata) SetDelayedUntil(v *timestamppb.Timestamp) {
	m.DelayedUntil = v
}

func (m *ListClusterBackupsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterBackupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterBackupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterBackupsResponse) SetBackups(v []*Backup) {
	m.Backups = v
}

func (m *ListClusterBackupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *DeleteOpenSearchNodeGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteOpenSearchNodeGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateOpenSearchNodeGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateOpenSearchNodeGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateOpenSearchNodeGroupRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateOpenSearchNodeGroupRequest) SetNodeGroupSpec(v *OpenSearchNodeGroupUpdateSpec) {
	m.NodeGroupSpec = v
}

func (m *OpenSearchNodeGroupUpdateSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *OpenSearchNodeGroupUpdateSpec) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *OpenSearchNodeGroupUpdateSpec) SetRoles(v []OpenSearch_GroupRole) {
	m.Roles = v
}

func (m *OpenSearchNodeGroupUpdateSpec) SetZoneIds(v []string) {
	m.ZoneIds = v
}

func (m *OpenSearchNodeGroupUpdateSpec) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *OpenSearchNodeGroupUpdateSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *AddOpenSearchNodeGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddOpenSearchNodeGroupRequest) SetNodeGroupSpec(v *OpenSearchCreateSpec_NodeGroup) {
	m.NodeGroupSpec = v
}

func (m *DeleteDashboardsNodeGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteDashboardsNodeGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateDashboardsNodeGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateDashboardsNodeGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateDashboardsNodeGroupRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateDashboardsNodeGroupRequest) SetNodeGroupSpec(v *DashboardsNodeGroupUpdateSpec) {
	m.NodeGroupSpec = v
}

func (m *DashboardsNodeGroupUpdateSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *DashboardsNodeGroupUpdateSpec) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *DashboardsNodeGroupUpdateSpec) SetZoneIds(v []string) {
	m.ZoneIds = v
}

func (m *DashboardsNodeGroupUpdateSpec) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *DashboardsNodeGroupUpdateSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *AddDashboardsNodeGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddDashboardsNodeGroupRequest) SetNodeGroupSpec(v *DashboardsCreateSpec_NodeGroup) {
	m.NodeGroupSpec = v
}

func (m *AddNodeGroupMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddNodeGroupMetadata) SetName(v string) {
	m.Name = v
}

func (m *UpdateNodeGroupMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateNodeGroupMetadata) SetName(v string) {
	m.Name = v
}

func (m *DeleteNodeGroupMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteNodeGroupMetadata) SetName(v string) {
	m.Name = v
}

func (m *GetAuthSettingsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateAuthSettingsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateAuthSettingsRequest) SetSettings(v *AuthSettings) {
	m.Settings = v
}

func (m *UpdateAuthSettingsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}
