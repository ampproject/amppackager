package bunny

import "context"

const (
	// DefaultPaginationPage is the default value that is used for
	// PaginationOptions.Page if it is unset.
	DefaultPaginationPage = 1
	// DefaultPaginationPerPage is the default value that is used for
	// PaginationOptions.PerPage if it is unset.
	DefaultPaginationPerPage = 1000
)

// PaginationOptions specifies optional parameters for List APIs.
type PaginationOptions struct {
	// Page the page to return
	Page int32 `url:"page,omitempty"`
	// PerPage how many entries to return per page
	PerPage int32 `url:"per_page,omitempty"`
}

// PaginationReply represents the pagination information contained in a
// List API endpoint response.
//
// Ex. Bunny.net API docs:
// - https://docs.bunny.net/reference/pullzonepublic_index
// - https://docs.bunny.net/reference/storagezonepublic_index
type PaginationReply[Item any] struct {
	Items        []*Item `json:"Items,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage"`
	TotalItems   *int32  `json:"TotalItems"`
	HasMoreItems *bool   `json:"HasMoreItems"`
}

func (p *PaginationOptions) ensureConstraints() {
	if p.Page < 1 {
		p.Page = DefaultPaginationPage
	}

	if p.PerPage < 1 {
		p.PerPage = DefaultPaginationPerPage
	}
}

func resourceList[Resp any](
	ctx context.Context,
	client *Client,
	path string,
	opts *PaginationOptions,
) (*Resp, error) {
	var res Resp

	// Ensure that opts.Page is >=1, if it isn't bunny.net will send a
	// different response JSON object, that contains only a single Object,
	// without items and paginations fields. Enforcing opts.page =>1 ensures
	// that we always unmarshall into the same struct.
	if opts == nil {
		opts = &PaginationOptions{
			Page:    DefaultPaginationPage,
			PerPage: DefaultPaginationPerPage,
		}
	} else {
		opts.ensureConstraints()
	}

	req, err := client.newGetRequest(path, opts)
	if err != nil {
		return nil, err
	}

	if err := client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
