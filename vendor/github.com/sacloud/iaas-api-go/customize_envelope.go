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

package iaas

import (
	"encoding/json"

	"github.com/sacloud/iaas-api-go/naked"
	"github.com/sacloud/iaas-api-go/search"
)

// Note: sacloud/配下でのUnmarshalJSONの実装
//
// v2/internal/dslではAPIからの戻り値が以下のようにラップされていることを期待している
// {
//    "<API名>": {},
//    "is_ok": true
// }
//
// この形式に沿っていないAPIについてはレスポンスのエンベロープ(struct xxxResponseEnvelope)でUnmarshalJSONを実装する必要がある

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (a *authStatusReadResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias authStatusReadResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedAuthStatus naked.AuthStatus
	if err := json.Unmarshal(data, &nakedAuthStatus); err != nil {
		return err
	}
	tmp.AuthStatus = &nakedAuthStatus

	*a = authStatusReadResponseEnvelope(tmp)
	return nil
}

func (b *billDetailsCSVResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias billDetailsCSVResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedBillDetailCSV naked.BillDetailCSV
	if err := json.Unmarshal(data, &nakedBillDetailCSV); err != nil {
		return err
	}
	tmp.CSV = &nakedBillDetailCSV

	*b = billDetailsCSVResponseEnvelope(tmp)
	return nil
}

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (a *certificateAuthorityAddClientResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias certificateAuthorityAddClientResponseEnvelope

	// is_okなどの共通的な項目
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	// ラッパーが省略されている、レスポンスボディの直下にある項目
	var result naked.CertificateAuthorityAddClientOrServerResult
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	tmp.CertificateAuthority = &result

	*a = certificateAuthorityAddClientResponseEnvelope(tmp)
	return nil
}

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (a *certificateAuthorityAddServerResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias certificateAuthorityAddServerResponseEnvelope

	// is_okなどの共通的な項目
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	// ラッパーが省略されている、レスポンスボディの直下にある項目
	var result naked.CertificateAuthorityAddClientOrServerResult
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	tmp.CertificateAuthority = &result

	*a = certificateAuthorityAddServerResponseEnvelope(tmp)
	return nil
}

func (m *mobileGatewaySetSIMRoutesRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias struct {
		SIMRoutes []*naked.MobileGatewaySIMRoute `json:"sim_routes"`
	}
	tmp := &alias{
		SIMRoutes: m.SIMRoutes,
	}
	if len(tmp.SIMRoutes) == 0 {
		tmp.SIMRoutes = make([]*naked.MobileGatewaySIMRoute, 0)
	}
	return json.Marshal(tmp)
}

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (s *serverGetVNCProxyResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias serverGetVNCProxyResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedVNCProxy naked.VNCProxyInfo
	if err := json.Unmarshal(data, &nakedVNCProxy); err != nil {
		return err
	}
	tmp.VNCProxyInfo = &nakedVNCProxy

	*s = serverGetVNCProxyResponseEnvelope(tmp)
	return nil
}

/*
 * 検索時に固定パラメータを設定するための実装
 */

func (s autoBackupFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias autoBackupFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "autobackup"
	return json.Marshal(tmp)
}

func (s autoScaleFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias autoScaleFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "autoscale"
	return json.Marshal(tmp)
}

func (s certificateAuthorityFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias certificateAuthorityFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "certificateauthority"
	return json.Marshal(tmp)
}

func (s containerRegistryFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias containerRegistryFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "containerregistry"
	return json.Marshal(tmp)
}

func (s eSMEFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias eSMEFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "esme"
	return json.Marshal(tmp)
}

func (s dNSFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias dNSFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "dns"
	return json.Marshal(tmp)
}

func (s simpleMonitorFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias simpleMonitorFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "simplemon"
	return json.Marshal(tmp)
}

func (s gSLBFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias gSLBFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "gslb"
	return json.Marshal(tmp)
}

func (s proxyLBFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias proxyLBFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "proxylb"
	return json.Marshal(tmp)
}

func (s sIMFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias sIMFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "sim"
	return json.Marshal(tmp)
}

func (s localRouterFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias localRouterFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "localrouter"
	return json.Marshal(tmp)
}

func (s enhancedDBFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias enhancedDBFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Provider.Class")] = "enhanceddb"
	return json.Marshal(tmp)
}

func (s databaseFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias databaseFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "database"
	return json.Marshal(tmp)
}

func (s loadBalancerFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias loadBalancerFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "loadbalancer"
	return json.Marshal(tmp)
}

func (s vPCRouterFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias vPCRouterFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "vpcrouter"
	return json.Marshal(tmp)
}

func (s nFSFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias nFSFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "nfs"
	return json.Marshal(tmp)
}

func (s mobileGatewayFindRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias mobileGatewayFindRequestEnvelope
	tmp := alias(s)
	if tmp.Filter == nil {
		tmp.Filter = search.Filter{}
	}
	tmp.Filter[search.Key("Class")] = "mobilegateway"
	return json.Marshal(tmp)
}

/*
 * for Shared Archive
 */

func (s archiveShareRequestEnvelope) MarshalJSON() ([]byte, error) {
	type alias archiveShareRequestEnvelope
	tmp := alias(s)
	tmp.Shared = true
	return json.Marshal(tmp)
}

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (s *archiveShareResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias archiveShareResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedData naked.ArchiveShareInfo
	if err := json.Unmarshal(data, &nakedData); err != nil {
		return err
	}
	tmp.ArchiveShareInfo = &nakedData

	*s = archiveShareResponseEnvelope(tmp)
	return nil
}

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (a *vPCRouterLogsResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias vPCRouterLogsResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedLogs naked.VPCRouterLog
	if err := json.Unmarshal(data, &nakedLogs); err != nil {
		return err
	}
	tmp.VPCRouter = &nakedLogs

	*a = vPCRouterLogsResponseEnvelope(tmp)
	return nil
}

// UnmarshalJSON APIからの戻り値でレスポンスボディ直下にデータを持つことへの対応
func (a *vPCRouterPingResponseEnvelope) UnmarshalJSON(data []byte) error {
	type alias vPCRouterPingResponseEnvelope

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var nakedResult naked.VPCRouterPingResult
	if err := json.Unmarshal(data, &nakedResult); err != nil {
		return err
	}
	tmp.VPCRouter = &nakedResult

	*a = vPCRouterPingResponseEnvelope(tmp)
	return nil
}
