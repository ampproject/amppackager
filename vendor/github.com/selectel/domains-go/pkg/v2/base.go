package v2

import (
	"context"
	"io"
	"net/http"
)

const (
	rootPath                 = "/zones"
	zonePath                 = "/zones/%v"
	zonePathUpdateState      = "/zones/%v/state"
	zonePathUpdateProtection = "/zones/%v/protection"
	rrsetPath                = "/zones/%v/rrset"
	singleRRSetPath          = "/zones/%v/rrset/%v"
)

type (
	DNSManager[Z any, S any] interface {
		ZoneManager[Z]
		RRSetManager[S]
	}
	DNSClient[Z any, S any] interface {
		DNSManager[Z, S]
		WithHeaders(headers http.Header) DNSClient[Z, S]
	}

	ZoneManager[Z any] interface {
		GetZone(ctx context.Context, zoneID string, options *map[string]string) (*Z, error)
		ListZones(ctx context.Context, options *map[string]string) (Listable[Z], error)
		CreateZone(ctx context.Context, zone Creatable) (*Z, error)
		DeleteZone(ctx context.Context, zoneID string) error
		UpdateZoneState(ctx context.Context, zoneID string, disabled bool) error
		UpdateZoneComment(ctx context.Context, zoneID string, comment string) error
		UpdateProtectionState(ctx context.Context, zoneID string, protected bool) error
	}

	RRSetManager[S any] interface {
		CreateRRSet(ctx context.Context, zoneID string, rrset Creatable) (*S, error)
		GetRRSet(ctx context.Context, zoneID, rrsetID string) (*S, error)
		ListRRSets(ctx context.Context, zoneID string, options *map[string]string) (Listable[S], error)
		UpdateRRSet(ctx context.Context, zoneID, rrsetID string, rrset Updatable) error
		DeleteRRSet(ctx context.Context, zoneID, rrsetID string) error
	}

	Listable[T any] interface {
		GetCount() int
		GetNextOffset() int
		GetItems() []*T
	}

	Creatable interface {
		CreationForm() (io.Reader, error)
	}

	Updatable interface {
		UpdateForm() (io.Reader, error)
	}
)
