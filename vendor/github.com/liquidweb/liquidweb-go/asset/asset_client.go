package asset

import (
	"github.com/liquidweb/liquidweb-cli/types/api"
	liquidweb "github.com/liquidweb/liquidweb-go"
)

// AssetBackend is the interface for assets.
type AssetBackend interface {
	Details(string) (*apiTypes.Subaccnt, error)
}

// AssetClient is the API client for storm servers.
type AssetClient struct {
	Backend liquidweb.Backend
}

// Details fetches the details of an asset.
func (c *AssetClient) Details(id string) (*apiTypes.Subaccnt, error) {
	var result apiTypes.Subaccnt
	args := map[string]interface{}{
		"uniq_id": id,
	}

	err := c.Backend.CallIntoInterface("bleed/asset/details", args, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
