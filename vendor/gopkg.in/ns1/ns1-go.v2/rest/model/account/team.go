package account

// Team wraps an NS1 /accounts/teams resource
type Team struct {
	ID          string         `json:"id,omitempty"`
	Name        string         `json:"name"`
	Permissions PermissionsMap `json:"permissions"`
	IPWhitelist []IPWhitelist  `json:"ip_whitelist"`
}

// IPWhitelist wraps the IP whitelist for Teams.
type IPWhitelist struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
