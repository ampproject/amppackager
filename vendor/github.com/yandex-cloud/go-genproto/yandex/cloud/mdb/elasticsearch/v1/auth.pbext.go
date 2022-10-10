// Code generated by protoc-gen-goext. DO NOT EDIT.

package elasticsearch

func (m *AuthProviders) SetProviders(v []*AuthProvider) {
	m.Providers = v
}

type AuthProvider_Settings = isAuthProvider_Settings

func (m *AuthProvider) SetSettings(v AuthProvider_Settings) {
	m.Settings = v
}

func (m *AuthProvider) SetType(v AuthProvider_Type) {
	m.Type = v
}

func (m *AuthProvider) SetName(v string) {
	m.Name = v
}

func (m *AuthProvider) SetOrder(v int64) {
	m.Order = v
}

func (m *AuthProvider) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *AuthProvider) SetHidden(v bool) {
	m.Hidden = v
}

func (m *AuthProvider) SetDescription(v string) {
	m.Description = v
}

func (m *AuthProvider) SetHint(v string) {
	m.Hint = v
}

func (m *AuthProvider) SetIcon(v string) {
	m.Icon = v
}

func (m *AuthProvider) SetSaml(v *SamlSettings) {
	m.Settings = &AuthProvider_Saml{
		Saml: v,
	}
}

func (m *SamlSettings) SetIdpEntityId(v string) {
	m.IdpEntityId = v
}

func (m *SamlSettings) SetIdpMetadataFile(v []byte) {
	m.IdpMetadataFile = v
}

func (m *SamlSettings) SetSpEntityId(v string) {
	m.SpEntityId = v
}

func (m *SamlSettings) SetKibanaUrl(v string) {
	m.KibanaUrl = v
}

func (m *SamlSettings) SetAttributePrincipal(v string) {
	m.AttributePrincipal = v
}

func (m *SamlSettings) SetAttributeGroups(v string) {
	m.AttributeGroups = v
}

func (m *SamlSettings) SetAttributeName(v string) {
	m.AttributeName = v
}

func (m *SamlSettings) SetAttributeEmail(v string) {
	m.AttributeEmail = v
}

func (m *SamlSettings) SetAttributeDn(v string) {
	m.AttributeDn = v
}
