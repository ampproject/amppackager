package bunny

import "context"

func resourcePutWithResponse[Resp any](
	ctx context.Context,
	client *Client,
	path string,
	requestBody any,
) (*Resp, error) {
	var res Resp

	req, err := client.newPutRequest(path, requestBody)
	if err != nil {
		return nil, err
	}

	if err := client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
