package bunny

import (
	"context"
	"fmt"
)

// SetForceSSLOptions represents the message is to the the Set Force SSL Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_setforcessl
type SetForceSSLOptions struct {
	Hostname *string `json:"Hostname,omitempty"`
	ForceSSL *bool   `json:"ForceSSL,omitempty"`
}

// SetForceSSL enables or disables the force SSL option for a hostname of a Pull Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_setforcessl
func (s *PullZoneService) SetForceSSL(ctx context.Context, pullzoneID int64, opts *SetForceSSLOptions) error {
	path := fmt.Sprintf("pullzone/%d/setForceSSL", pullzoneID)
	return resourcePost(ctx, s.client, path, opts)
}
