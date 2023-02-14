package rdpool

const Schema = "http://schemas.ultradns.com/RDPool.jsonschema"

type Profile struct {
	Context     string `json:"@context,omitempty"`
	Order       string `json:"order,omitempty"`
	Description string `json:"description,omitempty"`
}

func (profile *Profile) SetContext() {
	profile.Context = Schema
}

func (profile *Profile) GetContext() string {
	return profile.Context
}
