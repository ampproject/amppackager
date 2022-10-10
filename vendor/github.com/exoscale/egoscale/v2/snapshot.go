package v2

import (
	"context"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// SnapshotExport represents exported Snapshot information.
type SnapshotExport struct {
	MD5sum       *string
	PresignedURL *string
}

// Snapshot represents a Snapshot.
type Snapshot struct {
	CreatedAt  *time.Time
	ID         *string `req-for:"update,delete"`
	InstanceID *string
	Name       *string
	Size       *int64
	State      *string
	Zone       *string
}

func snapshotFromAPI(s *oapi.Snapshot, zone string) *Snapshot {
	return &Snapshot{
		CreatedAt:  s.CreatedAt,
		ID:         s.Id,
		InstanceID: s.Instance.Id,
		Name:       s.Name,
		Size:       s.Size,
		State:      (*string)(s.State),
		Zone:       &zone,
	}
}

// DeleteSnapshot deletes a Snapshot.
func (c *Client) DeleteSnapshot(ctx context.Context, zone string, snapshot *Snapshot) error {
	if err := validateOperationParams(snapshot, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteSnapshotWithResponse(apiv2.WithZone(ctx, zone), *snapshot.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ExportSnapshot exports a Snapshot and returns the exported Snapshot information.
func (c *Client) ExportSnapshot(ctx context.Context, zone string, snapshot *Snapshot) (*SnapshotExport, error) {
	if err := validateOperationParams(snapshot, "update"); err != nil {
		return nil, err
	}

	resp, err := c.ExportSnapshotWithResponse(apiv2.WithZone(ctx, zone), *snapshot.ID)
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	expSnapshot, err := c.GetSnapshotWithResponse(apiv2.WithZone(ctx, zone), *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
	if err != nil {
		return nil, err
	}

	return &SnapshotExport{
		MD5sum:       expSnapshot.JSON200.Export.Md5sum,
		PresignedURL: expSnapshot.JSON200.Export.PresignedUrl,
	}, nil
}

// GetSnapshot returns the Snapshot corresponding to the specified ID.
func (c *Client) GetSnapshot(ctx context.Context, zone, id string) (*Snapshot, error) {
	resp, err := c.GetSnapshotWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return snapshotFromAPI(resp.JSON200, zone), nil
}

// ListSnapshots returns the list of existing Snapshots.
func (c *Client) ListSnapshots(ctx context.Context, zone string) ([]*Snapshot, error) {
	list := make([]*Snapshot, 0)

	resp, err := c.ListSnapshotsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Snapshots != nil {
		for i := range *resp.JSON200.Snapshots {
			list = append(list, snapshotFromAPI(&(*resp.JSON200.Snapshots)[i], zone))
		}
	}

	return list, nil
}
