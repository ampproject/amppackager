// Code generated by protoc-gen-goext. DO NOT EDIT.

package oauth

func (m *SubjectClaims) SetSub(v string) {
	m.Sub = v
}

func (m *SubjectClaims) SetName(v string) {
	m.Name = v
}

func (m *SubjectClaims) SetGivenName(v string) {
	m.GivenName = v
}

func (m *SubjectClaims) SetFamilyName(v string) {
	m.FamilyName = v
}

func (m *SubjectClaims) SetPreferredUsername(v string) {
	m.PreferredUsername = v
}

func (m *SubjectClaims) SetPicture(v string) {
	m.Picture = v
}

func (m *SubjectClaims) SetEmail(v string) {
	m.Email = v
}

func (m *SubjectClaims) SetZoneinfo(v string) {
	m.Zoneinfo = v
}

func (m *SubjectClaims) SetLocale(v string) {
	m.Locale = v
}

func (m *SubjectClaims) SetPhoneNumber(v string) {
	m.PhoneNumber = v
}

func (m *SubjectClaims) SetSubType(v SubjectType) {
	m.SubType = v
}

func (m *SubjectClaims) SetFederation(v *Federation) {
	m.Federation = v
}

func (m *Federation) SetId(v string) {
	m.Id = v
}

func (m *Federation) SetName(v string) {
	m.Name = v
}
