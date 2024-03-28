package bunny

import "context"

// VideoLibraries represents the response of the List Video Library API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_index
type VideoLibraries PaginationReply[VideoLibrary]

// VideoLibraryListOpts represents both PaginationOptions and the other optional
// query parameters of the List endpoint.
type VideoLibraryListOpts struct {
	VideoLibraryGetOpts
	PaginationOptions
}

// List retrieves the Video Libraries.
// If opts is nil, DefaultPaginationPerPage and DefaultPaginationPage will be used.
// if opts.Page or or opts.PerPage is < 1, the related DefaultPagination values are used.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_index
func (s *VideoLibraryService) List(
	ctx context.Context,
	opts *VideoLibraryListOpts,
) (*VideoLibraries, error) {
	const path = "/videolibrary"
	var res VideoLibraries

	// NOTE: The resourceList function is not used for the purpose of
	// providing the extra query param options in VideoLibraryGetOpts. In the future
	// hopefully it can be removed for a better solution. See the following discussion:
	// https://github.com/simplesurance/bunny-go/pull/27#discussion_r1021270152

	// Ensure that opts.Page is >=1, if it isn't bunny.net will send a
	// different response JSON object, that contains only a single Object,
	// without items and paginations fields. Enforcing opts.page =>1 ensures
	// that we always unmarshall into the same struct.
	if opts == nil {
		opts = &VideoLibraryListOpts{
			PaginationOptions: PaginationOptions{
				Page:    DefaultPaginationPage,
				PerPage: DefaultPaginationPerPage,
			},
		}
	} else {
		opts.ensureConstraints()
	}

	req, err := s.client.newGetRequest(path, opts)
	if err != nil {
		return nil, err
	}

	if err := s.client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
