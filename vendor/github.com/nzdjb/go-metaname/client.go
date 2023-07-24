// Package metaname provides a client for the Metaname API.
package metaname

import (
	"context"
	"encoding/json"

	"github.com/AdamSLevy/jsonrpc2/v14"
)

type iJsonRpc2Client interface {
	Request(context context.Context, host string, method string, params interface{}, result interface{}) error
}

// A client for the Metaname API.
type MetanameClient struct {
	RpcClient        iJsonRpc2Client
	Host             string
	AccountReference string
	APIKey           string
}

// A ResourceRecord is a representation of a DNS record.
//
// Aux should be nil for records other than MX and SRV records, where it represents the priority.
// Reference should be nil when supplying a ResourceRecord, but will be populated when retrieving a record.
//
// https://metaname.net/api/1.1/doc#Resource_record_details
type ResourceRecord struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Aux       *int    `json:"aux"`
	Ttl       int     `json:"ttl"`
	Data      string  `json:"data"`
	Reference *string `json:"reference,omitempty"`
}

// Create a new MetanameClient with some default values.
func NewMetanameClient(accountReference string, apiKey string) *MetanameClient {
	return &MetanameClient{
		RpcClient:        &jsonrpc2.Client{},
		Host:             "https://metaname.net/api/1.1",
		AccountReference: accountReference,
		APIKey:           apiKey,
	}
}

// Creates a DNS record in the zone for the given domain and returns a reference that can be used for updating and deleting it.
//
// https://metaname.net/api/1.1/doc#create_dns_record
func (c *MetanameClient) CreateDnsRecord(ctx context.Context, domainName string, record ResourceRecord) (string, error) {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, record}
	var result string
	err := c.RpcClient.Request(ctx, c.Host, "create_dns_record", params, &result)
	return result, err
}

// Updates the details of a DNS record in a zone.
//
// https://metaname.net/api/1.1/doc#update_dns_record
func (c *MetanameClient) UpdateDnsRecord(ctx context.Context, domainName string, reference string, record ResourceRecord) error {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, reference, record}
	err := c.RpcClient.Request(ctx, c.Host, "update_dns_record", params, nil)
	return ignoreNullResultError(err)
}

// Delete a DNS record from a zone.
//
// https://metaname.net/api/1.1/doc#delete_dns_record
func (c *MetanameClient) DeleteDnsRecord(ctx context.Context, domainName string, reference string) error {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, reference}
	err := c.RpcClient.Request(ctx, c.Host, "delete_dns_record", params, nil)
	return ignoreNullResultError(err)
}

// Retrieve all the DNS records in a zone.
//
// https://metaname.net/api/1.1/doc#dns_zone
func (c *MetanameClient) DnsZone(ctx context.Context, domainName string) ([]ResourceRecord, error) {
	params := []interface{}{c.AccountReference, c.APIKey, domainName}
	var result []ResourceRecord
	err := c.RpcClient.Request(ctx, c.Host, "dns_zone", params, &result)
	return result, err
}

// Create or update a zone.
//
// https://metaname.net/api/1.1/doc#configure_zone
func (c *MetanameClient) ConfigureZone(ctx context.Context, zoneName string, records []ResourceRecord, options interface{}) error {
	params := []interface{}{c.AccountReference, c.APIKey, zoneName, records, options}
	err := c.RpcClient.Request(ctx, c.Host, "configure_zone", params, nil)
	return ignoreNullResultError(err)
}

// Workaround until https://github.com/AdamSLevy/jsonrpc2/issues/11 is fixed.
type nullSafeResponse struct {
	Result interface{} `json:"result"`
}

func ignoreNullResultError(err error) error {
	if unexerr, ok := err.(jsonrpc2.ErrorUnexpectedHTTPResponse); ok {
		var res nullSafeResponse
		unmerr := json.Unmarshal(unexerr.Body, &res)
		if unmerr != nil {
			return err
		} else if res.Result == nil {
			return nil
		}
	}
	return err
}
