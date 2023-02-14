package sbpool

import "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"

const Schema = "http://schemas.ultradns.com/SBPool.jsonschema"

type Profile struct {
	Context          string               `json:"@context,omitempty"`
	Description      string               `json:"description"`
	Order            string               `json:"order,omitempty"`
	Status           string               `json:"status,omitempty"`
	RunProbes        bool                 `json:"runProbes"`
	ActOnProbes      bool                 `json:"actOnProbes"`
	MaxActive        int                  `json:"maxActive,omitempty"`
	MaxServed        int                  `json:"maxServed,omitempty"`
	FailureThreshold int                  `json:"failureThreshold,omitempty"`
	RDataInfo        []*pool.RDataInfo    `json:"rdataInfo,omitempty"`
	BackupRecords    []*pool.BackupRecord `json:"backupRecords,omitempty"`
}

func (profile *Profile) SetContext() {
	profile.Context = Schema
}

func (profile *Profile) GetContext() string {
	return profile.Context
}
