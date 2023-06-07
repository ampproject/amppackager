package sfpool

import "github.com/ultradns/ultradns-go-sdk/pkg/record/pool"

const Schema = "http://schemas.ultradns.com/SFPool.jsonschema"

type Profile struct {
	Context                  string        `json:"@context,omitempty"`
	PoolDescription          string        `json:"poolDescription"`
	LiveRecordDescription    string        `json:"liveRecordDescription"`
	LiveRecordState          string        `json:"liveRecordState,omitempty"`
	RegionFailureSensitivity string        `json:"regionFailureSensitivity,omitempty"`
	Status                   string        `json:"status,omitempty"`
	BackupRecord             *BackupRecord `json:"backupRecord,omitempty"`
	Monitor                  *pool.Monitor `json:"monitor,omitempty"`
}

type BackupRecord struct {
	RData       string `json:"rdata,omitempty"`
	Description string `json:"description"`
}

func (profile *Profile) SetContext() {
	profile.Context = Schema
}

func (profile *Profile) GetContext() string {
	return profile.Context
}
