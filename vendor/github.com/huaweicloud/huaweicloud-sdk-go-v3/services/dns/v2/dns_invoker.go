package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
)

type CreateCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomLineInvoker) Invoke() (*model.CreateCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomLineResponse), nil
	}
}

type CreateLineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLineGroupInvoker) Invoke() (*model.CreateLineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLineGroupResponse), nil
	}
}

type DeleteCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomLineInvoker) Invoke() (*model.DeleteCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomLineResponse), nil
	}
}

type DeleteLineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLineGroupInvoker) Invoke() (*model.DeleteLineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLineGroupResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ListCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomLineInvoker) Invoke() (*model.ListCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomLineResponse), nil
	}
}

type ListLineGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLineGroupsInvoker) Invoke() (*model.ListLineGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLineGroupsResponse), nil
	}
}

type ListNameServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNameServersInvoker) Invoke() (*model.ListNameServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNameServersResponse), nil
	}
}

type ShowApiInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiInfoInvoker) Invoke() (*model.ShowApiInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiInfoResponse), nil
	}
}

type ShowDomainQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainQuotaInvoker) Invoke() (*model.ShowDomainQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainQuotaResponse), nil
	}
}

type ShowLineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLineGroupInvoker) Invoke() (*model.ShowLineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLineGroupResponse), nil
	}
}

type UpdateCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCustomLineInvoker) Invoke() (*model.UpdateCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCustomLineResponse), nil
	}
}

type UpdateLineGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLineGroupsInvoker) Invoke() (*model.UpdateLineGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLineGroupsResponse), nil
	}
}

type CreateEipRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEipRecordSetInvoker) Invoke() (*model.CreateEipRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEipRecordSetResponse), nil
	}
}

type ListPtrRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPtrRecordsInvoker) Invoke() (*model.ListPtrRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPtrRecordsResponse), nil
	}
}

type RestorePtrRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestorePtrRecordInvoker) Invoke() (*model.RestorePtrRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestorePtrRecordResponse), nil
	}
}

type ShowPtrRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPtrRecordSetInvoker) Invoke() (*model.ShowPtrRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPtrRecordSetResponse), nil
	}
}

type UpdatePtrRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePtrRecordInvoker) Invoke() (*model.UpdatePtrRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePtrRecordResponse), nil
	}
}

type BatchDeleteRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRecordSetWithLineInvoker) Invoke() (*model.BatchDeleteRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRecordSetWithLineResponse), nil
	}
}

type BatchUpdateRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateRecordSetWithLineInvoker) Invoke() (*model.BatchUpdateRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateRecordSetWithLineResponse), nil
	}
}

type CreateRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetInvoker) Invoke() (*model.CreateRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetResponse), nil
	}
}

type CreateRecordSetWithBatchLinesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetWithBatchLinesInvoker) Invoke() (*model.CreateRecordSetWithBatchLinesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetWithBatchLinesResponse), nil
	}
}

type CreateRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetWithLineInvoker) Invoke() (*model.CreateRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetWithLineResponse), nil
	}
}

type DeleteRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordSetInvoker) Invoke() (*model.DeleteRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordSetResponse), nil
	}
}

type DeleteRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordSetsInvoker) Invoke() (*model.DeleteRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordSetsResponse), nil
	}
}

type ListRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsInvoker) Invoke() (*model.ListRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsResponse), nil
	}
}

type ListRecordSetsByZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsByZoneInvoker) Invoke() (*model.ListRecordSetsByZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsByZoneResponse), nil
	}
}

type ListRecordSetsWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsWithLineInvoker) Invoke() (*model.ListRecordSetsWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsWithLineResponse), nil
	}
}

type SetRecordSetsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRecordSetsStatusInvoker) Invoke() (*model.SetRecordSetsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRecordSetsStatusResponse), nil
	}
}

type ShowRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetInvoker) Invoke() (*model.ShowRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetResponse), nil
	}
}

type ShowRecordSetByZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetByZoneInvoker) Invoke() (*model.ShowRecordSetByZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetByZoneResponse), nil
	}
}

type ShowRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetWithLineInvoker) Invoke() (*model.ShowRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetWithLineResponse), nil
	}
}

type UpdateRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordSetInvoker) Invoke() (*model.UpdateRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordSetResponse), nil
	}
}

type UpdateRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordSetsInvoker) Invoke() (*model.UpdateRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordSetsResponse), nil
	}
}

type BatchCreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagInvoker) Invoke() (*model.BatchCreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagInvoker) Invoke() (*model.ListTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ShowResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceTagInvoker) Invoke() (*model.ShowResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceTagResponse), nil
	}
}

type AssociateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRouterInvoker) Invoke() (*model.AssociateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRouterResponse), nil
	}
}

type CreatePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateZoneInvoker) Invoke() (*model.CreatePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateZoneResponse), nil
	}
}

type CreatePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePublicZoneInvoker) Invoke() (*model.CreatePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePublicZoneResponse), nil
	}
}

type DeletePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateZoneInvoker) Invoke() (*model.DeletePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateZoneResponse), nil
	}
}

type DeletePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublicZoneInvoker) Invoke() (*model.DeletePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublicZoneResponse), nil
	}
}

type DisassociateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateRouterInvoker) Invoke() (*model.DisassociateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateRouterResponse), nil
	}
}

type ListPrivateZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateZonesInvoker) Invoke() (*model.ListPrivateZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateZonesResponse), nil
	}
}

type ListPublicZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicZonesInvoker) Invoke() (*model.ListPublicZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicZonesResponse), nil
	}
}

type ShowPrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateZoneInvoker) Invoke() (*model.ShowPrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateZoneResponse), nil
	}
}

type ShowPrivateZoneNameServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateZoneNameServerInvoker) Invoke() (*model.ShowPrivateZoneNameServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateZoneNameServerResponse), nil
	}
}

type ShowPublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicZoneInvoker) Invoke() (*model.ShowPublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicZoneResponse), nil
	}
}

type ShowPublicZoneNameServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicZoneNameServerInvoker) Invoke() (*model.ShowPublicZoneNameServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicZoneNameServerResponse), nil
	}
}

type UpdatePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateZoneInvoker) Invoke() (*model.UpdatePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateZoneResponse), nil
	}
}

type UpdatePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicZoneInvoker) Invoke() (*model.UpdatePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicZoneResponse), nil
	}
}

type UpdatePublicZoneStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicZoneStatusInvoker) Invoke() (*model.UpdatePublicZoneStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicZoneStatusResponse), nil
	}
}
