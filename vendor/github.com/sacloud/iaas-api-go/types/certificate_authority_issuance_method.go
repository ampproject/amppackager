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

package types

// ECertificateAuthorityIssuanceMethod マネージドPKI(CA)での証明書発行方法
type ECertificateAuthorityIssuanceMethod string

// String ECertificateAuthorityIssuanceMethodの文字列表現
func (p ECertificateAuthorityIssuanceMethod) String() string {
	return string(p)
}

// CertificateAuthorityIssuanceMethods ECertificateAuthorityIssuanceMethodがとりうる値
var CertificateAuthorityIssuanceMethods = struct {
	URL       ECertificateAuthorityIssuanceMethod
	EMail     ECertificateAuthorityIssuanceMethod
	PublicKey ECertificateAuthorityIssuanceMethod
	CSR       ECertificateAuthorityIssuanceMethod
}{
	URL:       ECertificateAuthorityIssuanceMethod("url"),
	EMail:     ECertificateAuthorityIssuanceMethod("email"),
	PublicKey: ECertificateAuthorityIssuanceMethod("public_key"),
	CSR:       ECertificateAuthorityIssuanceMethod("csr"),
}

// CertificateAuthorityIssuanceMethodStrings x
var CertificateAuthorityIssuanceMethodStrings = []string{
	CertificateAuthorityIssuanceMethods.URL.String(),
	CertificateAuthorityIssuanceMethods.EMail.String(),
	CertificateAuthorityIssuanceMethods.PublicKey.String(),
	CertificateAuthorityIssuanceMethods.CSR.String(),
}
