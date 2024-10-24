package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
)

type DnsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDnsClient(hcClient *httpclient.HcHttpClient) *DnsClient {
	return &DnsClient{HcClient: hcClient}
}

func DnsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateCustomLine 创建单个自定义线路
//
// 创建单个自定义线路
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateCustomLine(request *model.CreateCustomLineRequest) (*model.CreateCustomLineResponse, error) {
	requestDef := GenReqDefForCreateCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomLineResponse), nil
	}
}

// CreateCustomLineInvoker 创建单个自定义线路
func (c *DnsClient) CreateCustomLineInvoker(request *model.CreateCustomLineRequest) *CreateCustomLineInvoker {
	requestDef := GenReqDefForCreateCustomLine()
	return &CreateCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLineGroup 创建线路分组
//
// 创建一个线路分组。 该接口部分区域未上线、如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateLineGroup(request *model.CreateLineGroupRequest) (*model.CreateLineGroupResponse, error) {
	requestDef := GenReqDefForCreateLineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLineGroupResponse), nil
	}
}

// CreateLineGroupInvoker 创建线路分组
func (c *DnsClient) CreateLineGroupInvoker(request *model.CreateLineGroupRequest) *CreateLineGroupInvoker {
	requestDef := GenReqDefForCreateLineGroup()
	return &CreateLineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomLine 删除单个自定义线路
//
// 删除单个自定义线路
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteCustomLine(request *model.DeleteCustomLineRequest) (*model.DeleteCustomLineResponse, error) {
	requestDef := GenReqDefForDeleteCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomLineResponse), nil
	}
}

// DeleteCustomLineInvoker 删除单个自定义线路
func (c *DnsClient) DeleteCustomLineInvoker(request *model.DeleteCustomLineRequest) *DeleteCustomLineInvoker {
	requestDef := GenReqDefForDeleteCustomLine()
	return &DeleteCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLineGroup 删除线路分组
//
// 删除单个线路分组。该接口部分区域未上线、如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteLineGroup(request *model.DeleteLineGroupRequest) (*model.DeleteLineGroupResponse, error) {
	requestDef := GenReqDefForDeleteLineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLineGroupResponse), nil
	}
}

// DeleteLineGroupInvoker 删除线路分组
func (c *DnsClient) DeleteLineGroupInvoker(request *model.DeleteLineGroupRequest) *DeleteLineGroupInvoker {
	requestDef := GenReqDefForDeleteLineGroup()
	return &DeleteLineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询所有的云解析服务API版本号
//
// 查询所有的云解析服务API版本号列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询所有的云解析服务API版本号
func (c *DnsClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomLine 查询自定义线路
//
// 查询自定义线路
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListCustomLine(request *model.ListCustomLineRequest) (*model.ListCustomLineResponse, error) {
	requestDef := GenReqDefForListCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomLineResponse), nil
	}
}

// ListCustomLineInvoker 查询自定义线路
func (c *DnsClient) ListCustomLineInvoker(request *model.ListCustomLineRequest) *ListCustomLineInvoker {
	requestDef := GenReqDefForListCustomLine()
	return &ListCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLineGroups 查询线路分组列表
//
// 查询线路分组列表。该接口部分区域未上线、如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListLineGroups(request *model.ListLineGroupsRequest) (*model.ListLineGroupsResponse, error) {
	requestDef := GenReqDefForListLineGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLineGroupsResponse), nil
	}
}

// ListLineGroupsInvoker 查询线路分组列表
func (c *DnsClient) ListLineGroupsInvoker(request *model.ListLineGroupsRequest) *ListLineGroupsInvoker {
	requestDef := GenReqDefForListLineGroups()
	return &ListLineGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNameServers 查询名称服务器列表
//
// 查询名称服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListNameServers(request *model.ListNameServersRequest) (*model.ListNameServersResponse, error) {
	requestDef := GenReqDefForListNameServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNameServersResponse), nil
	}
}

// ListNameServersInvoker 查询名称服务器列表
func (c *DnsClient) ListNameServersInvoker(request *model.ListNameServersRequest) *ListNameServersInvoker {
	requestDef := GenReqDefForListNameServers()
	return &ListNameServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiInfo 查询指定的云解析服务API版本号
//
// 查询指定的云解析服务API版本号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowApiInfo(request *model.ShowApiInfoRequest) (*model.ShowApiInfoResponse, error) {
	requestDef := GenReqDefForShowApiInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiInfoResponse), nil
	}
}

// ShowApiInfoInvoker 查询指定的云解析服务API版本号
func (c *DnsClient) ShowApiInfoInvoker(request *model.ShowApiInfoRequest) *ShowApiInfoInvoker {
	requestDef := GenReqDefForShowApiInfo()
	return &ShowApiInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainQuota 查询租户配额
//
// 查询单租户在DNS服务下的资源配额，包括公网zone配额、内网zone配额、Record Set配额、PTR Record配额、入站终端节点配额、出站终端节点配额、自定义线路配额、线路分组配额等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowDomainQuota(request *model.ShowDomainQuotaRequest) (*model.ShowDomainQuotaResponse, error) {
	requestDef := GenReqDefForShowDomainQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainQuotaResponse), nil
	}
}

// ShowDomainQuotaInvoker 查询租户配额
func (c *DnsClient) ShowDomainQuotaInvoker(request *model.ShowDomainQuotaRequest) *ShowDomainQuotaInvoker {
	requestDef := GenReqDefForShowDomainQuota()
	return &ShowDomainQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLineGroup 查询线路分组
//
// 查询线路分组。该接口部分区域未上线、如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowLineGroup(request *model.ShowLineGroupRequest) (*model.ShowLineGroupResponse, error) {
	requestDef := GenReqDefForShowLineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLineGroupResponse), nil
	}
}

// ShowLineGroupInvoker 查询线路分组
func (c *DnsClient) ShowLineGroupInvoker(request *model.ShowLineGroupRequest) *ShowLineGroupInvoker {
	requestDef := GenReqDefForShowLineGroup()
	return &ShowLineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCustomLine 更新单个自定义线路
//
// 更新单个自定义线路
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateCustomLine(request *model.UpdateCustomLineRequest) (*model.UpdateCustomLineResponse, error) {
	requestDef := GenReqDefForUpdateCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomLineResponse), nil
	}
}

// UpdateCustomLineInvoker 更新单个自定义线路
func (c *DnsClient) UpdateCustomLineInvoker(request *model.UpdateCustomLineRequest) *UpdateCustomLineInvoker {
	requestDef := GenReqDefForUpdateCustomLine()
	return &UpdateCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLineGroups 更新线路分组
//
// 更新单个线路分组。该接口部分区域未上线、如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateLineGroups(request *model.UpdateLineGroupsRequest) (*model.UpdateLineGroupsResponse, error) {
	requestDef := GenReqDefForUpdateLineGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLineGroupsResponse), nil
	}
}

// UpdateLineGroupsInvoker 更新线路分组
func (c *DnsClient) UpdateLineGroupsInvoker(request *model.UpdateLineGroupsRequest) *UpdateLineGroupsInvoker {
	requestDef := GenReqDefForUpdateLineGroups()
	return &UpdateLineGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEipRecordSet 设置弹性IP的PTR记录
//
// 设置弹性IP的PTR记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateEipRecordSet(request *model.CreateEipRecordSetRequest) (*model.CreateEipRecordSetResponse, error) {
	requestDef := GenReqDefForCreateEipRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEipRecordSetResponse), nil
	}
}

// CreateEipRecordSetInvoker 设置弹性IP的PTR记录
func (c *DnsClient) CreateEipRecordSetInvoker(request *model.CreateEipRecordSetRequest) *CreateEipRecordSetInvoker {
	requestDef := GenReqDefForCreateEipRecordSet()
	return &CreateEipRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPtrRecords 查询租户弹性IP的PTR记录列表
//
// 查询租户弹性IP的PTR记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPtrRecords(request *model.ListPtrRecordsRequest) (*model.ListPtrRecordsResponse, error) {
	requestDef := GenReqDefForListPtrRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPtrRecordsResponse), nil
	}
}

// ListPtrRecordsInvoker 查询租户弹性IP的PTR记录列表
func (c *DnsClient) ListPtrRecordsInvoker(request *model.ListPtrRecordsRequest) *ListPtrRecordsInvoker {
	requestDef := GenReqDefForListPtrRecords()
	return &ListPtrRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestorePtrRecord 将弹性IP的PTR记录恢复为默认值
//
// 将弹性IP的PTR记录恢复为默认值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) RestorePtrRecord(request *model.RestorePtrRecordRequest) (*model.RestorePtrRecordResponse, error) {
	requestDef := GenReqDefForRestorePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestorePtrRecordResponse), nil
	}
}

// RestorePtrRecordInvoker 将弹性IP的PTR记录恢复为默认值
func (c *DnsClient) RestorePtrRecordInvoker(request *model.RestorePtrRecordRequest) *RestorePtrRecordInvoker {
	requestDef := GenReqDefForRestorePtrRecord()
	return &RestorePtrRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPtrRecordSet 查询单个弹性IP的PTR记录
//
// 查询单个弹性IP的PTR记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPtrRecordSet(request *model.ShowPtrRecordSetRequest) (*model.ShowPtrRecordSetResponse, error) {
	requestDef := GenReqDefForShowPtrRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPtrRecordSetResponse), nil
	}
}

// ShowPtrRecordSetInvoker 查询单个弹性IP的PTR记录
func (c *DnsClient) ShowPtrRecordSetInvoker(request *model.ShowPtrRecordSetRequest) *ShowPtrRecordSetInvoker {
	requestDef := GenReqDefForShowPtrRecordSet()
	return &ShowPtrRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePtrRecord 修改弹性IP的PTR记录
//
// 修改弹性IP的PTR记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePtrRecord(request *model.UpdatePtrRecordRequest) (*model.UpdatePtrRecordResponse, error) {
	requestDef := GenReqDefForUpdatePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePtrRecordResponse), nil
	}
}

// UpdatePtrRecordInvoker 修改弹性IP的PTR记录
func (c *DnsClient) UpdatePtrRecordInvoker(request *model.UpdatePtrRecordRequest) *UpdatePtrRecordInvoker {
	requestDef := GenReqDefForUpdatePtrRecord()
	return &UpdatePtrRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRecordSetWithLine 批量删除某个Zone下的Record Set资源
//
// 批量删除某个Zone下的Record Set资源，当删除的资源不存在时，则默认删除成功。
// 响应结果中只包含本次实际删除的资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchDeleteRecordSetWithLine(request *model.BatchDeleteRecordSetWithLineRequest) (*model.BatchDeleteRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForBatchDeleteRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRecordSetWithLineResponse), nil
	}
}

// BatchDeleteRecordSetWithLineInvoker 批量删除某个Zone下的Record Set资源
func (c *DnsClient) BatchDeleteRecordSetWithLineInvoker(request *model.BatchDeleteRecordSetWithLineRequest) *BatchDeleteRecordSetWithLineInvoker {
	requestDef := GenReqDefForBatchDeleteRecordSetWithLine()
	return &BatchDeleteRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateRecordSetWithLine 批量修改RecordSet
//
// 批量修改RecordSet。属于原子性操作，请求Record Set将全部完成修改，或不做任何修改。
// 仅公网Zone支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchUpdateRecordSetWithLine(request *model.BatchUpdateRecordSetWithLineRequest) (*model.BatchUpdateRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForBatchUpdateRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateRecordSetWithLineResponse), nil
	}
}

// BatchUpdateRecordSetWithLineInvoker 批量修改RecordSet
func (c *DnsClient) BatchUpdateRecordSetWithLineInvoker(request *model.BatchUpdateRecordSetWithLineRequest) *BatchUpdateRecordSetWithLineInvoker {
	requestDef := GenReqDefForBatchUpdateRecordSetWithLine()
	return &BatchUpdateRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordSet 创建单个Record Set
//
// 创建单个Record Set
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateRecordSet(request *model.CreateRecordSetRequest) (*model.CreateRecordSetResponse, error) {
	requestDef := GenReqDefForCreateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetResponse), nil
	}
}

// CreateRecordSetInvoker 创建单个Record Set
func (c *DnsClient) CreateRecordSetInvoker(request *model.CreateRecordSetRequest) *CreateRecordSetInvoker {
	requestDef := GenReqDefForCreateRecordSet()
	return &CreateRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordSetWithBatchLines 批量线路创建RecordSet
//
// 批量线路创建RecordSet。属于原子性操作，如果存在一个参数校验不通过，则创建失败。仅公网Zone支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateRecordSetWithBatchLines(request *model.CreateRecordSetWithBatchLinesRequest) (*model.CreateRecordSetWithBatchLinesResponse, error) {
	requestDef := GenReqDefForCreateRecordSetWithBatchLines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetWithBatchLinesResponse), nil
	}
}

// CreateRecordSetWithBatchLinesInvoker 批量线路创建RecordSet
func (c *DnsClient) CreateRecordSetWithBatchLinesInvoker(request *model.CreateRecordSetWithBatchLinesRequest) *CreateRecordSetWithBatchLinesInvoker {
	requestDef := GenReqDefForCreateRecordSetWithBatchLines()
	return &CreateRecordSetWithBatchLinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordSetWithLine 创建单个Record Set
//
// 创建单个Record Set，仅适用于公网DNS
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateRecordSetWithLine(request *model.CreateRecordSetWithLineRequest) (*model.CreateRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForCreateRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetWithLineResponse), nil
	}
}

// CreateRecordSetWithLineInvoker 创建单个Record Set
func (c *DnsClient) CreateRecordSetWithLineInvoker(request *model.CreateRecordSetWithLineRequest) *CreateRecordSetWithLineInvoker {
	requestDef := GenReqDefForCreateRecordSetWithLine()
	return &CreateRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecordSet 删除单个Record Set
//
// 删除单个Record Set。删除有添加智能解析的记录集时，需要用Record Set多线路管理模块中删除接口进行删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteRecordSet(request *model.DeleteRecordSetRequest) (*model.DeleteRecordSetResponse, error) {
	requestDef := GenReqDefForDeleteRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetResponse), nil
	}
}

// DeleteRecordSetInvoker 删除单个Record Set
func (c *DnsClient) DeleteRecordSetInvoker(request *model.DeleteRecordSetRequest) *DeleteRecordSetInvoker {
	requestDef := GenReqDefForDeleteRecordSet()
	return &DeleteRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecordSets 删除单个Record Set
//
// 删除单个Record Set
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteRecordSets(request *model.DeleteRecordSetsRequest) (*model.DeleteRecordSetsResponse, error) {
	requestDef := GenReqDefForDeleteRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetsResponse), nil
	}
}

// DeleteRecordSetsInvoker 删除单个Record Set
func (c *DnsClient) DeleteRecordSetsInvoker(request *model.DeleteRecordSetsRequest) *DeleteRecordSetsInvoker {
	requestDef := GenReqDefForDeleteRecordSets()
	return &DeleteRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordSets 查询租户Record Set资源列表
//
// 查询租户Record Set资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListRecordSets(request *model.ListRecordSetsRequest) (*model.ListRecordSetsResponse, error) {
	requestDef := GenReqDefForListRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsResponse), nil
	}
}

// ListRecordSetsInvoker 查询租户Record Set资源列表
func (c *DnsClient) ListRecordSetsInvoker(request *model.ListRecordSetsRequest) *ListRecordSetsInvoker {
	requestDef := GenReqDefForListRecordSets()
	return &ListRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordSetsByZone 查询单个Zone下Record Set列表
//
// 查询单个Zone下Record Set列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListRecordSetsByZone(request *model.ListRecordSetsByZoneRequest) (*model.ListRecordSetsByZoneResponse, error) {
	requestDef := GenReqDefForListRecordSetsByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsByZoneResponse), nil
	}
}

// ListRecordSetsByZoneInvoker 查询单个Zone下Record Set列表
func (c *DnsClient) ListRecordSetsByZoneInvoker(request *model.ListRecordSetsByZoneRequest) *ListRecordSetsByZoneInvoker {
	requestDef := GenReqDefForListRecordSetsByZone()
	return &ListRecordSetsByZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordSetsWithLine 查询租户Record Set资源列表
//
// 查询租户Record Set资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListRecordSetsWithLine(request *model.ListRecordSetsWithLineRequest) (*model.ListRecordSetsWithLineResponse, error) {
	requestDef := GenReqDefForListRecordSetsWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsWithLineResponse), nil
	}
}

// ListRecordSetsWithLineInvoker 查询租户Record Set资源列表
func (c *DnsClient) ListRecordSetsWithLineInvoker(request *model.ListRecordSetsWithLineRequest) *ListRecordSetsWithLineInvoker {
	requestDef := GenReqDefForListRecordSetsWithLine()
	return &ListRecordSetsWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRecordSetsStatus 设置Record Set状态
//
// 设置Record Set状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) SetRecordSetsStatus(request *model.SetRecordSetsStatusRequest) (*model.SetRecordSetsStatusResponse, error) {
	requestDef := GenReqDefForSetRecordSetsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecordSetsStatusResponse), nil
	}
}

// SetRecordSetsStatusInvoker 设置Record Set状态
func (c *DnsClient) SetRecordSetsStatusInvoker(request *model.SetRecordSetsStatusRequest) *SetRecordSetsStatusInvoker {
	requestDef := GenReqDefForSetRecordSetsStatus()
	return &SetRecordSetsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordSet 查询单个Record Set
//
// 查询单个Record Set。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowRecordSet(request *model.ShowRecordSetRequest) (*model.ShowRecordSetResponse, error) {
	requestDef := GenReqDefForShowRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetResponse), nil
	}
}

// ShowRecordSetInvoker 查询单个Record Set
func (c *DnsClient) ShowRecordSetInvoker(request *model.ShowRecordSetRequest) *ShowRecordSetInvoker {
	requestDef := GenReqDefForShowRecordSet()
	return &ShowRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordSetByZone 查询单个Zone下Record Set列表
//
// 查询单个Zone下Record Set列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowRecordSetByZone(request *model.ShowRecordSetByZoneRequest) (*model.ShowRecordSetByZoneResponse, error) {
	requestDef := GenReqDefForShowRecordSetByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetByZoneResponse), nil
	}
}

// ShowRecordSetByZoneInvoker 查询单个Zone下Record Set列表
func (c *DnsClient) ShowRecordSetByZoneInvoker(request *model.ShowRecordSetByZoneRequest) *ShowRecordSetByZoneInvoker {
	requestDef := GenReqDefForShowRecordSetByZone()
	return &ShowRecordSetByZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordSetWithLine 查询单个Record Set
//
// 查询单个Record Set，仅适用于公网DNS
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowRecordSetWithLine(request *model.ShowRecordSetWithLineRequest) (*model.ShowRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForShowRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetWithLineResponse), nil
	}
}

// ShowRecordSetWithLineInvoker 查询单个Record Set
func (c *DnsClient) ShowRecordSetWithLineInvoker(request *model.ShowRecordSetWithLineRequest) *ShowRecordSetWithLineInvoker {
	requestDef := GenReqDefForShowRecordSetWithLine()
	return &ShowRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecordSet 修改单个Record Set
//
// 修改单个Record Set
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateRecordSet(request *model.UpdateRecordSetRequest) (*model.UpdateRecordSetResponse, error) {
	requestDef := GenReqDefForUpdateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetResponse), nil
	}
}

// UpdateRecordSetInvoker 修改单个Record Set
func (c *DnsClient) UpdateRecordSetInvoker(request *model.UpdateRecordSetRequest) *UpdateRecordSetInvoker {
	requestDef := GenReqDefForUpdateRecordSet()
	return &UpdateRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecordSets 修改单个Record Set
//
// 修改单个Record Set
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateRecordSets(request *model.UpdateRecordSetsRequest) (*model.UpdateRecordSetsResponse, error) {
	requestDef := GenReqDefForUpdateRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetsResponse), nil
	}
}

// UpdateRecordSetsInvoker 修改单个Record Set
func (c *DnsClient) UpdateRecordSetsInvoker(request *model.UpdateRecordSetsRequest) *UpdateRecordSetsInvoker {
	requestDef := GenReqDefForUpdateRecordSets()
	return &UpdateRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateTag 为指定实例批量添加或删除标签
//
// 为指定实例批量添加或删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchCreateTag(request *model.BatchCreateTagRequest) (*model.BatchCreateTagResponse, error) {
	requestDef := GenReqDefForBatchCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagResponse), nil
	}
}

// BatchCreateTagInvoker 为指定实例批量添加或删除标签
func (c *DnsClient) BatchCreateTagInvoker(request *model.BatchCreateTagRequest) *BatchCreateTagInvoker {
	requestDef := GenReqDefForBatchCreateTag()
	return &BatchCreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 为指定实例添加标签
//
// 为指定实例添加标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 为指定实例添加标签
func (c *DnsClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除资源标签
//
// 删除资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除资源标签
func (c *DnsClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTag 使用标签查询资源实例
//
// 使用标签查询资源实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListTag(request *model.ListTagRequest) (*model.ListTagResponse, error) {
	requestDef := GenReqDefForListTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagResponse), nil
	}
}

// ListTagInvoker 使用标签查询资源实例
func (c *DnsClient) ListTagInvoker(request *model.ListTagRequest) *ListTagInvoker {
	requestDef := GenReqDefForListTag()
	return &ListTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询指定实例类型的所有标签集合
//
// 查询指定实例类型的所有标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询指定实例类型的所有标签集合
func (c *DnsClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTag 查询指定实例的标签信息
//
// 查询指定实例的标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowResourceTag(request *model.ShowResourceTagRequest) (*model.ShowResourceTagResponse, error) {
	requestDef := GenReqDefForShowResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagResponse), nil
	}
}

// ShowResourceTagInvoker 查询指定实例的标签信息
func (c *DnsClient) ShowResourceTagInvoker(request *model.ShowResourceTagRequest) *ShowResourceTagInvoker {
	requestDef := GenReqDefForShowResourceTag()
	return &ShowResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRouter 在内网Zone上关联VPC
//
// 在内网Zone上关联VPC
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) AssociateRouter(request *model.AssociateRouterRequest) (*model.AssociateRouterResponse, error) {
	requestDef := GenReqDefForAssociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRouterResponse), nil
	}
}

// AssociateRouterInvoker 在内网Zone上关联VPC
func (c *DnsClient) AssociateRouterInvoker(request *model.AssociateRouterRequest) *AssociateRouterInvoker {
	requestDef := GenReqDefForAssociateRouter()
	return &AssociateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivateZone 创建单个内网Zone
//
// 创建单个内网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreatePrivateZone(request *model.CreatePrivateZoneRequest) (*model.CreatePrivateZoneResponse, error) {
	requestDef := GenReqDefForCreatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateZoneResponse), nil
	}
}

// CreatePrivateZoneInvoker 创建单个内网Zone
func (c *DnsClient) CreatePrivateZoneInvoker(request *model.CreatePrivateZoneRequest) *CreatePrivateZoneInvoker {
	requestDef := GenReqDefForCreatePrivateZone()
	return &CreatePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePublicZone 创建单个公网Zone
//
// 创建单个公网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreatePublicZone(request *model.CreatePublicZoneRequest) (*model.CreatePublicZoneResponse, error) {
	requestDef := GenReqDefForCreatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicZoneResponse), nil
	}
}

// CreatePublicZoneInvoker 创建单个公网Zone
func (c *DnsClient) CreatePublicZoneInvoker(request *model.CreatePublicZoneRequest) *CreatePublicZoneInvoker {
	requestDef := GenReqDefForCreatePublicZone()
	return &CreatePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePrivateZone 删除单个内网Zone
//
// 删除单个内网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeletePrivateZone(request *model.DeletePrivateZoneRequest) (*model.DeletePrivateZoneResponse, error) {
	requestDef := GenReqDefForDeletePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateZoneResponse), nil
	}
}

// DeletePrivateZoneInvoker 删除单个内网Zone
func (c *DnsClient) DeletePrivateZoneInvoker(request *model.DeletePrivateZoneRequest) *DeletePrivateZoneInvoker {
	requestDef := GenReqDefForDeletePrivateZone()
	return &DeletePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePublicZone 删除单个公网Zone
//
// 删除单个公网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeletePublicZone(request *model.DeletePublicZoneRequest) (*model.DeletePublicZoneResponse, error) {
	requestDef := GenReqDefForDeletePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicZoneResponse), nil
	}
}

// DeletePublicZoneInvoker 删除单个公网Zone
func (c *DnsClient) DeletePublicZoneInvoker(request *model.DeletePublicZoneRequest) *DeletePublicZoneInvoker {
	requestDef := GenReqDefForDeletePublicZone()
	return &DeletePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateRouter 在内网Zone上解关联VPC
//
// 在内网Zone上解关联VPC
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DisassociateRouter(request *model.DisassociateRouterRequest) (*model.DisassociateRouterResponse, error) {
	requestDef := GenReqDefForDisassociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRouterResponse), nil
	}
}

// DisassociateRouterInvoker 在内网Zone上解关联VPC
func (c *DnsClient) DisassociateRouterInvoker(request *model.DisassociateRouterRequest) *DisassociateRouterInvoker {
	requestDef := GenReqDefForDisassociateRouter()
	return &DisassociateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivateZones 查询内网Zone列表
//
// 查询内网Zone列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPrivateZones(request *model.ListPrivateZonesRequest) (*model.ListPrivateZonesResponse, error) {
	requestDef := GenReqDefForListPrivateZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateZonesResponse), nil
	}
}

// ListPrivateZonesInvoker 查询内网Zone列表
func (c *DnsClient) ListPrivateZonesInvoker(request *model.ListPrivateZonesRequest) *ListPrivateZonesInvoker {
	requestDef := GenReqDefForListPrivateZones()
	return &ListPrivateZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicZones 查询公网Zone列表
//
// 查询公网Zone列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPublicZones(request *model.ListPublicZonesRequest) (*model.ListPublicZonesResponse, error) {
	requestDef := GenReqDefForListPublicZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicZonesResponse), nil
	}
}

// ListPublicZonesInvoker 查询公网Zone列表
func (c *DnsClient) ListPublicZonesInvoker(request *model.ListPublicZonesRequest) *ListPublicZonesInvoker {
	requestDef := GenReqDefForListPublicZones()
	return &ListPublicZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateZone 查询单个内网Zone
//
// 查询单个内网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPrivateZone(request *model.ShowPrivateZoneRequest) (*model.ShowPrivateZoneResponse, error) {
	requestDef := GenReqDefForShowPrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneResponse), nil
	}
}

// ShowPrivateZoneInvoker 查询单个内网Zone
func (c *DnsClient) ShowPrivateZoneInvoker(request *model.ShowPrivateZoneRequest) *ShowPrivateZoneInvoker {
	requestDef := GenReqDefForShowPrivateZone()
	return &ShowPrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateZoneNameServer 查询单个内网Zone的名称服务器
//
// 查询单个内网Zone的名称服务器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPrivateZoneNameServer(request *model.ShowPrivateZoneNameServerRequest) (*model.ShowPrivateZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPrivateZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneNameServerResponse), nil
	}
}

// ShowPrivateZoneNameServerInvoker 查询单个内网Zone的名称服务器
func (c *DnsClient) ShowPrivateZoneNameServerInvoker(request *model.ShowPrivateZoneNameServerRequest) *ShowPrivateZoneNameServerInvoker {
	requestDef := GenReqDefForShowPrivateZoneNameServer()
	return &ShowPrivateZoneNameServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicZone 查询单个公网Zone
//
// 查询单个公网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPublicZone(request *model.ShowPublicZoneRequest) (*model.ShowPublicZoneResponse, error) {
	requestDef := GenReqDefForShowPublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneResponse), nil
	}
}

// ShowPublicZoneInvoker 查询单个公网Zone
func (c *DnsClient) ShowPublicZoneInvoker(request *model.ShowPublicZoneRequest) *ShowPublicZoneInvoker {
	requestDef := GenReqDefForShowPublicZone()
	return &ShowPublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicZoneNameServer 查询单个公网Zone的名称服务器
//
// 查询单个公网Zone的名称服务器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPublicZoneNameServer(request *model.ShowPublicZoneNameServerRequest) (*model.ShowPublicZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPublicZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneNameServerResponse), nil
	}
}

// ShowPublicZoneNameServerInvoker 查询单个公网Zone的名称服务器
func (c *DnsClient) ShowPublicZoneNameServerInvoker(request *model.ShowPublicZoneNameServerRequest) *ShowPublicZoneNameServerInvoker {
	requestDef := GenReqDefForShowPublicZoneNameServer()
	return &ShowPublicZoneNameServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrivateZone 修改单个内网Zone
//
// 修改单个内网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePrivateZone(request *model.UpdatePrivateZoneRequest) (*model.UpdatePrivateZoneResponse, error) {
	requestDef := GenReqDefForUpdatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateZoneResponse), nil
	}
}

// UpdatePrivateZoneInvoker 修改单个内网Zone
func (c *DnsClient) UpdatePrivateZoneInvoker(request *model.UpdatePrivateZoneRequest) *UpdatePrivateZoneInvoker {
	requestDef := GenReqDefForUpdatePrivateZone()
	return &UpdatePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublicZone 修改单个公网Zone
//
// 修改单个公网Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePublicZone(request *model.UpdatePublicZoneRequest) (*model.UpdatePublicZoneResponse, error) {
	requestDef := GenReqDefForUpdatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicZoneResponse), nil
	}
}

// UpdatePublicZoneInvoker 修改单个公网Zone
func (c *DnsClient) UpdatePublicZoneInvoker(request *model.UpdatePublicZoneRequest) *UpdatePublicZoneInvoker {
	requestDef := GenReqDefForUpdatePublicZone()
	return &UpdatePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublicZoneStatus 设置单个公网Zone状态
//
// 设置单个公网Zone状态，支持暂停、启用Zone
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePublicZoneStatus(request *model.UpdatePublicZoneStatusRequest) (*model.UpdatePublicZoneStatusResponse, error) {
	requestDef := GenReqDefForUpdatePublicZoneStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicZoneStatusResponse), nil
	}
}

// UpdatePublicZoneStatusInvoker 设置单个公网Zone状态
func (c *DnsClient) UpdatePublicZoneStatusInvoker(request *model.UpdatePublicZoneStatusRequest) *UpdatePublicZoneStatusInvoker {
	requestDef := GenReqDefForUpdatePublicZoneStatus()
	return &UpdatePublicZoneStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
