package bunny

import "context"

// PullZoneAddOptions are the request parameters for the Get Pull Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_add
type PullZoneAddOptions struct {
	// The name of the pull zone.
	Name string `json:"Name,omitempty"`
	// The origin URL of the pull zone where the files are fetched from.
	OriginURL string `json:"OriginUrl,omitempty"`

	// The ID of the storage zone that the pull zone is linked to. (Optional)
	StorageZoneID *int64 `json:"StorageZoneId,omitempty"`
	// The type of the pull zone. Standard = 0, Volume = 1. (Optional)
	Type int `json:"Type,omitempty"`
}

// Add creates a new Pull Zone.
// opts and the non-optional parameters in the struct must be specified for a successful request.
// On success the created PullZone is returned.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_add
func (s *PullZoneService) Add(ctx context.Context, opts *PullZoneAddOptions) (*PullZone, error) {
	return resourcePostWithResponse[PullZone](
		ctx,
		s.client,
		"/pullzone",
		opts,
	)
}
