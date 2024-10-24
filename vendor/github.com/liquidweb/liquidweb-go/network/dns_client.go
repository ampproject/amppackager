package network

import (
	"fmt"

	liquidweb "github.com/liquidweb/liquidweb-go"
)

type DNSZoneBackend interface {
	List(liquidweb.ListMeta) (DNSZoneList, error)
	ListAll() (DNSZoneList, error)
	Create(DNSZoneCreateParams) (DNSZone, error)
	Details(string) (DNSZone, error)
	DeleteByName(string) error
}

type DNSZoneClient struct {
	Backend liquidweb.Backend
}

func (c *DNSZoneClient) List(params liquidweb.ListMeta) (result DNSZoneList, err error) {
	err = c.Backend.CallIntoInterface("bleed/Network/DNS/Zone/list", params, &result)
	return
}

func (c *DNSZoneClient) ListAll() (result DNSZoneList, err error) {
	var reqParams liquidweb.ListMeta
	var incrementalResult DNSZoneList
	incrementalResult.PageTotal = 2

	for incrementalResult.PageTotal > incrementalResult.PageNum {
		incrementalResult, err = c.List(reqParams)
		if err != nil {
			return DNSZoneList{}, err
		}
		result.Items = append(result.Items, incrementalResult.Items...)
		reqParams.PageNum = incrementalResult.PageNum + 1
	}
	return
}

func (c *DNSZoneClient) Create(params DNSZoneCreateParams) (result DNSZone, err error) {
	err = c.Backend.CallIntoInterface("bleed/Network/DNS/Zone/create", params, &result)
	return
}

func (c *DNSZoneClient) DeleteByName(zoneName string) (err error) {
	result := struct {
		DeletedDomain string `json:"name"`
	}{}
	err = c.Backend.CallIntoInterface("bleed/Network/DNS/Zone/delete", struct {
		Name string `json:"name"`
	}{
		Name: zoneName,
	}, &result)
	if err != nil {
		return err
	}
	if result.DeletedDomain != zoneName {
		return fmt.Errorf("requested delete of %s deleted %s", zoneName, result.DeletedDomain)
	}
	return nil
}

func (c *DNSZoneClient) Details(zoneName string) (zone DNSZone, err error) {
	var result DNSZone
	err = c.Backend.CallIntoInterface("bleed/Network/DNS/Zone/details", struct {
		Name string `json:"name"`
	}{
		Name: zoneName,
	}, &result)
	return result, err
}

// DNSBackend describes the interface for interactions with the API.
type DNSBackend interface {
	Create(*DNSRecordParams) (*DNSRecord, error)
	Details(int) (*DNSRecord, error)
	List(*DNSRecordParams) (*DNSRecordList, error)
	ListAll(string) (DNSRecordList, error)
	Update(*DNSRecordParams) (*DNSRecord, error)
	Delete(*DNSRecordParams) (*DNSRecordDeletion, error)
}

// DNSClient is the backend implementation for interacting with DNS Records.
type DNSClient struct {
	Backend liquidweb.Backend
}

// Create creates a new DNS Record.
func (c *DNSClient) Create(params *DNSRecordParams) (*DNSRecord, error) {
	var result DNSRecord
	err := c.Backend.CallIntoInterface("v1/Network/DNS/Record/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details returns details about a DNS Record.
func (c *DNSClient) Details(id int) (*DNSRecord, error) {
	var result DNSRecord
	params := DNSRecordParams{ID: id}

	err := c.Backend.CallIntoInterface("v1/Network/DNS/Record/details", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a list of DNS Records.
func (c *DNSClient) List(params *DNSRecordParams) (*DNSRecordList, error) {
	list := &DNSRecordList{}

	err := c.Backend.CallIntoInterface("v1/Network/DNS/Record/list", params, list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// List returns a list of DNS Records.
func (c *DNSClient) ListAll(zone string) (result DNSRecordList, err error) {
	var reqParams DNSRecordParams
	incrementalResult := &DNSRecordList{}
	reqParams.PageNum = 1
	reqParams.Zone = zone
	incrementalResult.PageTotal = 2

	for incrementalResult.PageTotal > incrementalResult.PageNum {
		incrementalResult, err = c.List(&reqParams)
		if err != nil {
			return DNSRecordList{}, err
		}
		result.Items = append(result.Items, incrementalResult.Items...)
		reqParams.PageNum++
	}
	return
}

// Update will update a DNS Record.
func (c *DNSClient) Update(params *DNSRecordParams) (*DNSRecord, error) {
	var result DNSRecord
	err := c.Backend.CallIntoInterface("v1/Network/DNS/Record/update", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete will delete a DNS Record.
func (c *DNSClient) Delete(params *DNSRecordParams) (*DNSRecordDeletion, error) {
	var result DNSRecordDeletion
	err := c.Backend.CallIntoInterface("v1/Network/DNS/Record/delete", params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
