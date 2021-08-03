package client

import (
	lwApi "github.com/liquidweb/go-lwApi"

	"github.com/liquidweb/liquidweb-go/asset"
	network "github.com/liquidweb/liquidweb-go/network"
	"github.com/liquidweb/liquidweb-go/storage"
	"github.com/liquidweb/liquidweb-go/storm"
)

// API is the structure that houses all of our various API clients that interact with various Storm resources.
type API struct {
	NetworkDNS          network.DNSBackend
	NetworkLoadBalancer network.LoadBalancerBackend
	NetworkVIP          network.VIPBackend
	NetworkZone         network.ZoneBackend
	StorageBlockVolume  storage.BlockVolumeBackend

	StormConfig storm.ConfigBackend
	StormServer storm.ServerBackend
	Asset       asset.AssetBackend
}

// NewAPI is the API client for interacting with Storm.
func NewAPI(username string, password string, url string, timeout int) (*API, error) {
	// TODO support auth token. go-lwApi already supports this.
	clientArgs := lwApi.LWAPIConfig{
		Username: &username,
		Password: &password,
		Url:      url,
		Timeout:  uint(timeout),
		Insecure: false, // disable HTTPS validation?
	}
	client, err := NewClient(&clientArgs)
	if err != nil {
		return nil, err
	}

	api := &API{
		NetworkDNS:          &network.DNSClient{Backend: client.httpClient},
		NetworkLoadBalancer: &network.LoadBalancerClient{Backend: client.httpClient},
		NetworkVIP:          &network.VIPClient{Backend: client.httpClient},
		NetworkZone:         &network.ZoneClient{Backend: client.httpClient},
		StorageBlockVolume:  &storage.BlockVolumeClient{Backend: client.httpClient},
		StormConfig:         &storm.ConfigClient{Backend: client.httpClient},
		StormServer:         &storm.ServerClient{Backend: client.httpClient},
		Asset:               &asset.AssetClient{Backend: client.httpClient},
	}

	return api, nil
}
