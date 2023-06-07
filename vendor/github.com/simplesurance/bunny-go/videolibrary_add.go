package bunny

import "context"

// VideoLibraryAddOptions are the request parameters for the Get Video Library API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_add
type VideoLibraryAddOptions struct {
	// The name of the Video Library.
	Name *string `json:"Name,omitempty"`

	// The geo-replication regions of the underlying storage zone (Optional)
	ReplicationRegions []string `json:"ReplicationRegions,omitempty"`
}

// Add creates a new Video Library.
// opts and the non-optional parameters in the struct must be specified for a successful request.
// On success the created VideoLibrary is returned.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_add
func (s *VideoLibraryService) Add(ctx context.Context, opts *VideoLibraryAddOptions) (*VideoLibrary, error) {
	return resourcePostWithResponse[VideoLibrary](
		ctx,
		s.client,
		"/videolibrary",
		opts,
	)
}
