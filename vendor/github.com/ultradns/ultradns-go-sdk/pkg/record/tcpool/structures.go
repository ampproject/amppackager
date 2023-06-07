package tcpool

import "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"

const Schema = "http://schemas.ultradns.com/TCPool.jsonschema"

type Profile struct {
	Context          string             `json:"@context,omitempty"`
	Description      string             `json:"description"`
	Status           string             `json:"status,omitempty"`
	RunProbes        bool               `json:"runProbes"`
	ActOnProbes      bool               `json:"actOnProbes"`
	MaxToLB          int                `json:"maxToLB"`
	FailureThreshold int                `json:"failureThreshold"`
	RDataInfo        []*pool.RDataInfo  `json:"rdataInfo,omitempty"`
	BackupRecord     *pool.BackupRecord `json:"backupRecord,omitempty"`
}

func (profile *Profile) SetContext() {
	profile.Context = Schema
}

func (profile *Profile) GetContext() string {
	return profile.Context
}
