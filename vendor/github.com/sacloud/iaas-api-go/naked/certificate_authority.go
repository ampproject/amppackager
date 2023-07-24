// Copyright 2022-2023 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package naked

import (
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// CertificateAuthority プライベートCA
type CertificateAuthority struct {
	ID           types.ID                      `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string                        `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string                        `yaml:"description"`
	Tags         types.Tags                    `yaml:"tags"`
	Icon         *Icon                         `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time                    `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time                    `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability           `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass string                        `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Provider     *Provider                     `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *CertificateAuthoritySettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                        `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *CertificateAuthorityStatus   `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// CertificateAuthorityStatus CertificateAuthorityステータス
type CertificateAuthorityStatus struct {
	Country          string    `json:"country,omitempty" yaml:"country,omitempty" structs:",omitempty"`
	Organization     string    `json:"organization,omitempty" yaml:"organization,omitempty" structs:",omitempty"`
	OrganizationUnit []string  `json:"organization_unit,omitempty" yaml:"organization_unit,omitempty" structs:",omitempty"`
	CommonName       string    `json:"common_name,omitempty" yaml:"common_name,omitempty" structs:",omitempty"`
	NotAfter         time.Time `json:"not_after,omitempty" yaml:"not_after,omitempty" structs:",omitempty"`
	Subject          string    `json:"subject,omitempty" yaml:"subject,omitempty" structs:",omitempty"`
}

// CertificateAuthoritySettings CertificateAuthorityセッティング
type CertificateAuthoritySettings struct {
	// 現在は常に空となる。実際の設定は以下APIからCA/クライアント/サーバ別で取得する
	//
	// CA: GET /commonserviceitem/:id/certificateauthority
	// サーバ証明書: GET /commonserviceitem/:id/certificateauthority/servers
	// クライアント証明書: GET /commonserviceitem/:id/certificateauthority/clients
}

// CertificateAuthorityDetail CAの詳細情報
//
// GET /commonserviceitem/:id/certificateauthorityの戻り値
type CertificateAuthorityDetail struct {
	Subject         string           `json:"subject,omitempty" yaml:"subject,omitempty" structs:",omitempty"`
	CertificateData *CertificateData `json:"certificate_data,omitempty" yaml:"certificate_data,omitempty" structs:",omitempty"`
}

// CertificateAuthorityServerDetail サーバ証明書の詳細情報
//
// GET /commonserviceitem/:id/certificateauthority/serversの戻り値を構成する
// (実際にはFind系のラッパーがある)
type CertificateAuthorityServerDetail struct {
	ID              string           `json:"id,omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Subject         string           `json:"subject,omitempty" yaml:"subject,omitempty" structs:",omitempty"`
	SANs            []string         `json:"sans,omitempty" yaml:"sans,omitempty" structs:",omitempty"`
	EMail           string           `json:"email,omitempty" yaml:"email,omitempty" structs:",omitempty"`
	IssueState      string           `json:"issue_state,omitempty" yaml:"issue_state,omitempty" structs:",omitempty"`
	CertificateData *CertificateData `json:"certificate_data,omitempty" yaml:"certificate_data,omitempty" structs:",omitempty"`
	URL             string           `json:"url,omitempty" yaml:"url,omitempty" structs:",omitempty"` // 常に空のはず
}

// CertificateAuthorityClientDetail クライアント証明書の詳細情報
//
// GET /commonserviceitem/:id/certificateauthority/clientsの戻り値を構成する
// (実際にはFind系のラッパーがある)
type CertificateAuthorityClientDetail struct {
	ID              string                                    `json:"id,omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Subject         string                                    `json:"subject,omitempty" yaml:"subject,omitempty" structs:",omitempty"`
	EMail           string                                    `json:"email,omitempty" yaml:"email,omitempty" structs:",omitempty"`
	IssuanceMethod  types.ECertificateAuthorityIssuanceMethod `json:"issuance_method,omitempty" yaml:"issuance_method,omitempty" structs:",omitempty"`
	IssueState      string                                    `json:"issue_state,omitempty" yaml:"issue_state,omitempty" structs:",omitempty"`
	CertificateData *CertificateData                          `json:"certificate_data,omitempty" yaml:"certificate_data,omitempty" structs:",omitempty"`
	URL             string                                    `json:"url,omitempty" yaml:"url,omitempty" structs:",omitempty"`
}

// CertificateData CA/クライアント/サーバの各証明書の情報
type CertificateData struct {
	CertificatePEM string    `json:"certificate_pem,omitempty" yaml:"certificate_pem,omitempty" structs:",omitempty"`
	Subject        string    `json:"subject,omitempty" yaml:"subject,omitempty" structs:",omitempty"`
	SerialNumber   string    `json:"serial_number,omitempty" yaml:"serial_number,omitempty" structs:",omitempty"`
	NotBefore      time.Time `json:"not_before,omitempty" yaml:"not_before,omitempty" structs:",omitempty"`
	NotAfter       time.Time `json:"not_after,omitempty" yaml:"not_after,omitempty" structs:",omitempty"`
}

type CertificateAuthorityAddClientParameter struct {
	Status *CertificateAuthorityAddClientParameterBody `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

type CertificateAuthorityAddClientParameterBody struct {
	Country                   string                                    `json:"country,omitempty" yaml:"country,omitempty" structs:",omitempty"`
	Organization              string                                    `json:"organization,omitempty" yaml:"organization,omitempty" structs:",omitempty"`
	OrganizationUnit          []string                                  `json:"organization_unit,omitempty" yaml:"organization_unit,omitempty" structs:",omitempty"`
	CommonName                string                                    `json:"common_name,omitempty" yaml:"common_name,omitempty" structs:",omitempty"`
	NotAfter                  time.Time                                 `json:"not_after,omitempty" yaml:"not_after,omitempty" structs:",omitempty"`
	EMail                     string                                    `json:"email,omitempty" yaml:"email,omitempty" structs:",omitempty"`
	IssuanceMethod            types.ECertificateAuthorityIssuanceMethod `json:"issuance_method,omitempty" yaml:"issuance_method,omitempty" structs:",omitempty"`
	CertificateSigningRequest string                                    `json:"certificate_signing_request,omitempty" yaml:"certificate_signing_request,omitempty" structs:",omitempty"`
	PublicKey                 string                                    `json:"public_key,omitempty" yaml:"public_key,omitempty" structs:",omitempty"`
}

type CertificateAuthorityAddServerParameter struct {
	Status *CertificateAuthorityAddServerParameterBody `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

type CertificateAuthorityAddServerParameterBody struct {
	Country                   string    `json:"country,omitempty" yaml:"country,omitempty" structs:",omitempty"`
	Organization              string    `json:"organization,omitempty" yaml:"organization,omitempty" structs:",omitempty"`
	OrganizationUnit          []string  `json:"organization_unit,omitempty" yaml:"organization_unit,omitempty" structs:",omitempty"`
	CommonName                string    `json:"common_name,omitempty" yaml:"common_name,omitempty" structs:",omitempty"`
	NotAfter                  time.Time `json:"not_after,omitempty" yaml:"not_after,omitempty" structs:",omitempty"`
	SANs                      []string  `json:"sans,omitempty" yaml:"sans,omitempty" structs:",omitempty"`
	CertificateSigningRequest string    `json:"certificate_signing_request,omitempty" yaml:"certificate_signing_request,omitempty" structs:",omitempty"`
	PublicKey                 string    `json:"public_key,omitempty" yaml:"public_key,omitempty" structs:",omitempty"`
}

type CertificateAuthorityAddClientOrServerResult struct {
	ID string `json:"id,omitempty" yaml:"id,omitempty" structs:",omitempty"`
}
