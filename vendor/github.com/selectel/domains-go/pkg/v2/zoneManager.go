package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type (
	// Zone represents an unmarshalled zone body from API response.
	Zone struct {
		ID        string    `json:"id"`
		ProjectID string    `json:"project_id"`
		Name      string    `json:"name"`
		Comment   string    `json:"comment"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Disabled  bool      `json:"disabled"`
		delegationInfo
	}

	delegationInfo struct {
		DelegationCheckedAt time.Time `json:"delegation_checked_at"`
		LastDelegatedAt     time.Time `json:"last_delegated_at"`
		LastCheckStatus     bool      `json:"last_check_status"`
	}

	zoneCreateForm struct {
		Name string `json:"name"`
	}

	zoneUpdateComment struct {
		Comment string `json:"comment"`
	}

	zoneUpdateState struct {
		Disabled bool `json:"disabled"`
	}
	zoneProtectionState struct {
		Protected bool `json:"protected"`
	}
)

func (z *Zone) CreationForm() (io.Reader, error) {
	form := zoneCreateForm{Name: z.Name}
	body, err := json.Marshal(form)

	return bytes.NewReader(body), err
}

// GetZone returns a single zone by its id.
func (c *Client) GetZone(ctx context.Context, zoneID string, _ *map[string]string) (*Zone, error) {
	r, e := c.prepareRequest(
		ctx, http.MethodGet, fmt.Sprintf(zonePath, zoneID), nil, nil, nil,
	)

	return processRequest[Zone](c.httpClient, r, e)
}

// ListZones returns a list of zones by options.
func (c *Client) ListZones(ctx context.Context, options *map[string]string) (Listable[Zone], error) {
	r, e := c.prepareRequest(
		ctx, http.MethodGet, rootPath, nil, options, nil,
	)

	return processRequest[List[Zone]](c.httpClient, r, e)
}

// CreateZone request to create of a new zone.
func (c *Client) CreateZone(ctx context.Context, zone Creatable) (*Zone, error) {
	body, err := zone.CreationForm()
	if err != nil {
		return nil, fmt.Errorf("create zone: %w", err)
	}
	r, e := c.prepareRequest(
		ctx, http.MethodPost, rootPath, body, nil, nil,
	)

	return processRequest[Zone](c.httpClient, r, e)
}

// DeleteZone request to delete of the zone by id.
func (c *Client) DeleteZone(ctx context.Context, zoneID string) error {
	r, e := c.prepareRequest(
		ctx, http.MethodDelete, fmt.Sprintf(zonePath, zoneID), nil, nil, nil,
	)
	_, err := processRequest[Zone](c.httpClient, r, e)

	return err
}

// UpdateZoneComment request to update the comment for zone by zoneID.
func (c *Client) UpdateZoneComment(ctx context.Context, zoneID string, comment string) error {
	updateComment, err := json.Marshal(zoneUpdateComment{
		Comment: comment,
	})
	if err != nil {
		return fmt.Errorf("zone marshal: %w", err)
	}
	form := bytes.NewReader(updateComment)
	r, e := c.prepareRequest(
		ctx, http.MethodPatch, fmt.Sprintf(zonePath, zoneID), form, nil, nil,
	)
	_, err = processRequest[Zone](c.httpClient, r, e)

	return err
}

// UpdateZoneState request to enable/disable service for zone by zoneID.
func (c *Client) UpdateZoneState(ctx context.Context, zoneID string, disabled bool) error {
	updateState, err := json.Marshal(zoneUpdateState{
		Disabled: disabled,
	})
	if err != nil {
		return fmt.Errorf("zone marshal: %w", err)
	}
	form := bytes.NewReader(updateState)
	r, e := c.prepareRequest(
		ctx, http.MethodPatch, fmt.Sprintf(zonePathUpdateState, zoneID), form, nil, nil,
	)
	_, err = processRequest[Zone](c.httpClient, r, e)

	return err
}

// UpdateProtectionState request to enable/disable zone protection from delete operation.
func (c *Client) UpdateProtectionState(ctx context.Context, zoneID string, protected bool) error {
	updateState, err := json.Marshal(zoneProtectionState{
		Protected: protected,
	})
	if err != nil {
		return fmt.Errorf("zone marshal: %w", err)
	}
	form := bytes.NewReader(updateState)
	r, e := c.prepareRequest(
		ctx, http.MethodPatch, fmt.Sprintf(zonePathUpdateProtection, zoneID), form, nil, nil,
	)
	_, err = processRequest[Zone](c.httpClient, r, e)

	return err
}
