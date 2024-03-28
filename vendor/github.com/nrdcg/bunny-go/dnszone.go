package bunny

// DNSZoneService communicates with the /dnszone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_index
type DNSZoneService struct {
	client *Client
}
