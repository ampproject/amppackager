package bunny

import (
	"context"
	"fmt"
)

// RemoveCertificateOptions represents the request parameters for the Remove
// Certificate API Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_removecertificate
type RemoveCertificateOptions struct {
	Hostname *string `json:"Hostname,omitempty"`
}

// RemoveCertificate represents the Remove Certificate API Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_removecertificate
func (s *PullZoneService) RemoveCertificate(ctx context.Context, pullZoneID int64, opts *RemoveCertificateOptions) error {
	path := fmt.Sprintf("/pullzone/%d/removeCertificate", pullZoneID)
	return resourceDelete(ctx, s.client, path, opts)
}
