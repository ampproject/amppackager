package bunny

import "context"

func resourceDelete(
	ctx context.Context,
	client *Client,
	path string,
	requestBody any,
) error {
	req, err := client.newDeleteRequest(path, requestBody)
	if err != nil {
		return err
	}

	return client.sendRequest(ctx, req, nil)
}
