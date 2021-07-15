package storm

import liquidweb "github.com/liquidweb/liquidweb-go"

// ServerBackend is the interface for storm servers.
type ServerBackend interface {
	Create(ServerParams) (*Server, error)
	List(ServerListParams) (*ServerList, error)
	Details(string) (*Server, error)
	Update(ServerParams) (*Server, error)
	Destroy(string) (*ServerDeletion, error)
	Status(string) (*ServerStatus, error)
	Stop(string, ...bool) (*ServerStop, error)
	Reboot(string) (*ServerReboot, error)
	Start(string) (*ServerStart, error)
}

// ServerClient is the API client for storm servers.
type ServerClient struct {
	Backend liquidweb.Backend
}

// List will fetch a list of storm servers.
func (c *ServerClient) List(params ServerListParams) (*ServerList, error) {
	var result ServerList
	err := c.Backend.CallIntoInterface("v1/Storm/Server/list", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Create a new storm server.
func (c *ServerClient) Create(params ServerParams) (*Server, error) {
	var result Server
	err := c.Backend.CallIntoInterface("v1/Storm/Server/create", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Details fetches the details for a storm server.
func (c *ServerClient) Details(id string) (*Server, error) {
	var result Server
	params := ServerParams{UniqID: id}

	err := c.Backend.CallIntoInterface("v1/Storm/Server/details", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update a storm server.
func (c *ServerClient) Update(params ServerParams) (*Server, error) {
	var result Server

	err := c.Backend.CallIntoInterface("v1/Storm/Server/update", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Destroy a storm server.
func (c *ServerClient) Destroy(id string) (*ServerDeletion, error) {
	var result ServerDeletion
	params := ServerParams{UniqID: id}

	err := c.Backend.CallIntoInterface("v1/Storm/Server/destroy", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Status returns the current status of a storm server.
func (c *ServerClient) Status(id string) (*ServerStatus, error) {
	var result ServerStatus
	params := ServerParams{UniqID: id}

	err := c.Backend.CallIntoInterface("v1/Storm/Server/status", params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Stop a storm server.
func (c *ServerClient) Stop(uniqId string, force ...bool) (*ServerStop, error) {
	var result ServerStop
	args := map[string]interface{}{
		"uniq_id": uniqId,
	}

	if len(force) > 0 {
		args["force"] = force[0]
	}
	err := c.Backend.CallIntoInterface("bleed/server/shutdown", args, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Reboot a storm server.
func (c *ServerClient) Reboot(uniqId string) (*ServerReboot, error) {
	var result ServerReboot
	args := map[string]interface{}{
		"uniq_id": uniqId,
	}

	err := c.Backend.CallIntoInterface("bleed/storm/server/reboot", args, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Start a Storm Server.
func (c *ServerClient) Start(uniqId string) (*ServerStart, error) {
	var result ServerStart
	args := map[string]interface{}{
		"uniq_id": uniqId,
	}

	err := c.Backend.CallIntoInterface("bleed/server/start", args, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
