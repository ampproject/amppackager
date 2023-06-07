package pool

const (
	RD  = "RD_POOL"
	SF  = "SF_POOL"
	SLB = "SLB_POOL"
	SB  = "SB_POOL"
	TC  = "TC_POOL"
	DIR = "DIR_POOL"
)

// Monitor structure for SF and SLB Pool.
type Monitor struct {
	Method          string `json:"method,omitempty"`
	URL             string `json:"url,omitempty"`
	TransmittedData string `json:"transmittedData"`
	SearchString    string `json:"searchString"`
}

// BackupRecord structure for SB and TC pool.
type BackupRecord struct {
	RData            string `json:"rdata,omitempty"`
	FailOverDelay    int    `json:"failoverDelay"`
	AvailableToServe bool   `json:"availableToServe"`
}

// RDataInfo structure for SB and TC pool.
type RDataInfo struct {
	State            string `json:"state,omitempty"`
	Status           string `json:"status,omitempty"`
	RunProbes        bool   `json:"runProbes"`
	AvailableToServe bool   `json:"availableToServe"`
	Priority         int    `json:"priority"`
	FailoverDelay    int    `json:"failoverDelay"`
	Threshold        int    `json:"threshold"`
	Weight           int    `json:"weight,omitempty"`
}
