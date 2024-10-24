package bunny

import "context"

type loadFreeCertificateQueryParams struct {
	Hostname string `url:"hostname,omitempty"`
}

// LoadFreeCertificate represents the Load Free Certificate API Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_loadfreecertificate
func (s *PullZoneService) LoadFreeCertificate(ctx context.Context, hostname string) error {
	params := loadFreeCertificateQueryParams{Hostname: hostname}

	req, err := s.client.newGetRequest("/pullzone/loadFreeCertificate", &params)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
