package v2

import (
	"context"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// SnapshotExport represents exported Snapshot information.
type SnapshotExport struct {
	MD5sum       *string
	PresignedURL *string
}

// Snapshot represents a Snapshot.
type Snapshot struct {
	CreatedAt  *time.Time
	ID         *string
	InstanceID *string
	Name       *string
	State      *string

	c    *Client
	zone string
}

func snapshotFromAPI(client *Client, zone string, s *papi.Snapshot) *Snapshot {
	return &Snapshot{
		CreatedAt:  s.CreatedAt,
		ID:         s.Id,
		InstanceID: s.Instance.Id,
		Name:       s.Name,
		State:      (*string)(s.State),

		c:    client,
		zone: zone,
	}
}

func (s Snapshot) get(ctx context.Context, client *Client, zone, id string) (interface{}, error) {
	return client.GetSnapshot(ctx, zone, id)
}

// Export exports the Snapshot and returns the exported Snapshot information.
func (s *Snapshot) Export(ctx context.Context) (*SnapshotExport, error) {
	resp, err := s.c.ExportSnapshotWithResponse(apiv2.WithZone(ctx, s.zone), *s.ID)
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(s.c.timeout).
		WithInterval(s.c.pollInterval).
		Poll(ctx, s.c.OperationPoller(s.zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	expSnapshot, err := s.c.GetSnapshotWithResponse(apiv2.WithZone(ctx, s.zone), *res.(*papi.Reference).Id)
	if err != nil {
		return nil, err
	}

	return &SnapshotExport{
		MD5sum:       expSnapshot.JSON200.Export.Md5sum,
		PresignedURL: expSnapshot.JSON200.Export.PresignedUrl,
	}, nil
}

// ListSnapshots returns the list of existing Snapshots in the specified zone.
func (c *Client) ListSnapshots(ctx context.Context, zone string) ([]*Snapshot, error) {
	list := make([]*Snapshot, 0)

	resp, err := c.ListSnapshotsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Snapshots != nil {
		for i := range *resp.JSON200.Snapshots {
			list = append(list, snapshotFromAPI(c, zone, &(*resp.JSON200.Snapshots)[i]))
		}
	}

	return list, nil
}

// GetSnapshot returns the Snapshot corresponding to the specified ID in the specified zone.
func (c *Client) GetSnapshot(ctx context.Context, zone, id string) (*Snapshot, error) {
	resp, err := c.GetSnapshotWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return snapshotFromAPI(c, zone, resp.JSON200), nil
}

// DeleteSnapshot deletes the specified Snapshot in the specified zone.
func (c *Client) DeleteSnapshot(ctx context.Context, zone, id string) error {
	resp, err := c.DeleteSnapshotWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
