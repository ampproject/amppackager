package bunny

// PullZoneService communicates with the /pullzone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pull-zone
type PullZoneService struct {
	client *Client
}
