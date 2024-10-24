package bunny

import "context"

func resourceGet[Resp any](ctx context.Context, client *Client, path string, params interface{}) (*Resp, error) {
	req, err := client.newGetRequest(path, params)
	if err != nil {
		return nil, err
	}

	var res Resp
	if err = client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
