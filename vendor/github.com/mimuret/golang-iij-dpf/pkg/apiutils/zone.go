package apiutils

import (
	"context"
	"fmt"

	"github.com/miekg/dns"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
)

var (
	ErrZoneNotFound   = fmt.Errorf("zone not found")
	ErrRecordNotFound = fmt.Errorf("record not found")
)

func getZoneFromSearchKeyWords(ctx context.Context, cl api.ClientInterface, keywords *core.ZoneListSearchKeywords) (*core.Zone, error) {
	zoneList := &core.ZoneList{}
	if _, err := cl.ListAll(ctx, zoneList, keywords); err != nil {
		return nil, fmt.Errorf("failed to search zone: %w", err)
	}
	for _, zone := range zoneList.Items {
		if len(keywords.Name) > 0 && keywords.Name[0] == zone.Name {
			return &zone, nil
		}
		if len(keywords.ServiceCode) > 0 && keywords.ServiceCode[0] == zone.ServiceCode {
			return &zone, nil
		}
	}
	return nil, ErrZoneNotFound
}

func GetZoneIdFromServiceCode(ctx context.Context, cl api.ClientInterface, serviceCode string) (string, error) {
	z, err := GetZoneFromServiceCode(ctx, cl, serviceCode)
	if err != nil {
		return "", err
	}
	return z.ID, nil
}

func GetZoneFromServiceCode(ctx context.Context, cl api.ClientInterface, serviceCode string) (*core.Zone, error) {
	return getZoneFromSearchKeyWords(ctx, cl, &core.ZoneListSearchKeywords{
		ServiceCode: api.KeywordsString{serviceCode},
	})
}

func GetZoneIDFromZonename(ctx context.Context, cl api.ClientInterface, zonename string) (string, error) {
	z, err := GetZoneFromZonename(ctx, cl, zonename)
	if err != nil {
		return "", err
	}
	return z.ID, nil
}

func GetZoneFromZonename(ctx context.Context, cl api.ClientInterface, zonename string) (*core.Zone, error) {
	return getZoneFromSearchKeyWords(ctx, cl, &core.ZoneListSearchKeywords{
		Name: api.KeywordsString{zonename},
	})
}

func GetRecordFromZoneName(ctx context.Context, cl api.ClientInterface, zonename string, recordName string, rrtype zones.Type) (*zones.Record, error) {
	z, err := GetZoneFromZonename(ctx, cl, zonename)
	if err != nil {
		return nil, err
	}
	return GetRecordFromZoneID(ctx, cl, z.ID, recordName, rrtype)
}

func GetRecordFromZoneID(ctx context.Context, cl api.ClientInterface, zoneID string, recordName string, rrtype zones.Type) (*zones.Record, error) {
	recordName = dns.CanonicalName(recordName)
	keywords := &zones.RecordListSearchKeywords{
		Name: api.KeywordsString{recordName},
	}
	currentList := &zones.CurrentRecordList{
		AttributeMeta: zones.AttributeMeta{
			ZoneID: zoneID,
		},
	}
	if _, err := cl.ListAll(ctx, currentList, keywords); err != nil {
		return nil, fmt.Errorf("failed to search records: %w", err)
	}
	for _, record := range currentList.Items {
		if recordName == record.Name && record.RRType == rrtype {
			return &record, nil
		}
	}
	return nil, ErrRecordNotFound
}
