package goacmedns

// Account is a struct that holds the registration response from an ACME-DNS
// server. It represents an API username/key that can be used to update TXT
// records for the account's subdomain.
type Account struct {
	FullDomain string `json:"fulldomain"`
	SubDomain  string `json:"subdomain"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	// ServerURL contains the URL of the acme-dns server the Account was registered with
	// (may be empty for Account instances registered before this field was added).
	ServerURL string `json:"server_url"`
}
