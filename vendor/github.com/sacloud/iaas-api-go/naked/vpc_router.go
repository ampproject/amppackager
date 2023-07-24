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
	"encoding/json"
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// VPCRouter VPCルータ
type VPCRouter struct {
	ID           types.ID            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Class        string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Plan         *AppliancePlan      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Settings     *VPCRouterSettings  `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Instance     *Instance           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Icon         *Icon               `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Switch       *Switch             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Interfaces   Interfaces          `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Tags         types.Tags          `yaml:"tags"`
}

// VPCRouterSettingsUpdate VPCルータ
type VPCRouterSettingsUpdate struct {
	Settings     *VPCRouterSettings `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SettingsHash string             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterSettings VPCルータ 設定
type VPCRouterSettings struct {
	Router *VPCRouterSetting `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterSetting VPCルータ 設定
type VPCRouterSetting struct {
	InternetConnection   *VPCRouterInternetConnection   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Interfaces           VPCRouterInterfaces            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	VRID                 int                            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	StaticNAT            *VPCRouterStaticNAT            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PortForwarding       *VPCRouterPortForwarding       `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Firewall             *VPCRouterFirewall             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DHCPServer           *VPCRouterDHCPServer           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DHCPStaticMapping    *VPCRouterDHCPStaticMappings   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DNSForwarding        *VPCRouterDNSForwarding        `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PPTPServer           *VPCRouterPPTPServer           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	L2TPIPsecServer      *VPCRouterL2TPIPsecServer      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	WireGuard            *VPCRouterWireGuard            `json:"WireGuardServer,omitempty" yaml:",omitempty" structs:",omitempty"`
	RemoteAccessUsers    *VPCRouterRemoteAccessUsers    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SiteToSiteIPsecVPN   *VPCRouterSiteToSiteIPsecVPN   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	StaticRoutes         *VPCRouterStaticRoutes         `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SyslogHost           string                         `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ScheduledMaintenance *VPCRouterScheduledMaintenance `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterInternetConnection インターフェース
type VPCRouterInternetConnection struct {
	Enabled types.StringFlag `yaml:"enabled"`
}

// VPCRouterInterface インターフェース
type VPCRouterInterface struct {
	IPAddress        []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	VirtualIPAddress string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	IPAliases        []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NetworkMaskLen   int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	// Index 仮想フィールド、VPCルータなどでInterfaces(実体は[]*Interface)を扱う場合にUnmarshalJSONの中で設定される
	//
	// Findした際のAPIからの応答にも同名のフィールドが含まれるが無関係。
	Index int `json:"-"`
}

// VPCRouterInterfaces Interface配列
//
// 配列中にnullが返ってくる(VPCルータなど)への対応のためのtype
type VPCRouterInterfaces []*VPCRouterInterface

// UnmarshalJSON 配列中にnullが返ってくる(VPCルータなど)への対応
func (i *VPCRouterInterfaces) UnmarshalJSON(b []byte) error {
	type alias VPCRouterInterfaces
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	var dest []*VPCRouterInterface
	for i, v := range a {
		if v != nil {
			if v.Index == 0 {
				v.Index = i
			}
			dest = append(dest, v)
		}
	}

	*i = VPCRouterInterfaces(dest)
	return nil
}

// MarshalJSON 配列中にnullが入る場合(VPCルータなど)への対応
func (i *VPCRouterInterfaces) MarshalJSON() ([]byte, error) {
	max := 0
	for _, iface := range *i {
		if max < iface.Index {
			max = iface.Index
		}
	}

	var dest = make([]*VPCRouterInterface, max+1)
	for _, iface := range *i {
		dest[iface.Index] = iface
	}

	return json.Marshal(dest)
}

// MarshalJSON JSON
func (i *VPCRouterInterface) MarshalJSON() ([]byte, error) {
	type alias struct {
		IPAddress        []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
		VirtualIPAddress string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
		IPAliases        []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
		NetworkMaskLen   int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	}

	tmp := alias{
		IPAddress:        i.IPAddress,
		VirtualIPAddress: i.VirtualIPAddress,
		IPAliases:        i.IPAliases,
		NetworkMaskLen:   i.NetworkMaskLen,
	}
	return json.Marshal(tmp)
}

// VPCRouterStaticNAT スタティックNAT
type VPCRouterStaticNAT struct {
	Config  []*VPCRouterStaticNATConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag            `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterStaticNAT) MarshalJSON() ([]byte, error) {
	if f == nil || f.Config == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterStaticNAT
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterStaticNATConfig スタティックNAT
type VPCRouterStaticNATConfig struct {
	GlobalAddress  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PrivateAddress string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Description    string `yaml:"description"`
}

// VPCRouterPortForwarding ポートフォワーディング設定
type VPCRouterPortForwarding struct {
	Config  []*VPCRouterPortForwardingConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                 `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterPortForwarding) MarshalJSON() ([]byte, error) {
	if f == nil || f.Config == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterPortForwarding
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterPortForwardingConfig ポートフォワーディング設定
type VPCRouterPortForwardingConfig struct {
	Protocol       types.EVPCRouterPortForwardingProtocol `json:",omitempty"` // プロトコル(tcp/udp)
	GlobalPort     types.StringNumber                     `json:",omitempty"` // グローバル側ポート
	PrivateAddress string                                 `json:",omitempty"` // プライベートIPアドレス
	PrivatePort    types.StringNumber                     `json:",omitempty"` // プライベート側ポート
	Description    string                                 `json:",omitempty"` // 説明
}

// VPCRouterFirewall ファイアウォール
type VPCRouterFirewall struct {
	Config  VPCRouterFirewallConfigs `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag         `yaml:"enabled"`
}

// MarshalJSON 常にEnabledをtrueに設定する
func (f *VPCRouterFirewall) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	f.Enabled = types.StringTrue
	type alias VPCRouterFirewall
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterFirewallConfigs VPCルータのファイアウォール設定
//
// 配列のインデックスで対象インターフェースを表す
type VPCRouterFirewallConfigs [8]*VPCRouterFirewallConfig

// UnmarshalJSON 配列中にnullが返ってくる(VPCルータなど)への対応
func (i *VPCRouterFirewallConfigs) UnmarshalJSON(b []byte) error {
	type alias VPCRouterFirewallConfigs
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	var dest [8]*VPCRouterFirewallConfig
	for i, v := range a {
		if v != nil {
			if v.Index == 0 {
				v.Index = i
			}
			dest[v.Index] = v
		}
	}

	*i = VPCRouterFirewallConfigs(dest)
	return nil
}

// MarshalJSON 配列中にnullが入る場合(VPCルータなど)への対応
func (i *VPCRouterFirewallConfigs) MarshalJSON() ([]byte, error) {
	var dest [8]*VPCRouterFirewallConfig
	for _, iface := range *i {
		if iface != nil {
			if iface.Receive == nil {
				iface.Receive = make([]*VPCRouterFirewallRule, 0)
			}
			if iface.Send == nil {
				iface.Send = make([]*VPCRouterFirewallRule, 0)
			}
			dest[iface.Index] = iface
		}
	}

	for i, v := range dest {
		if v == nil {
			dest[i] = &VPCRouterFirewallConfig{
				Receive: make([]*VPCRouterFirewallRule, 0),
				Send:    make([]*VPCRouterFirewallRule, 0),
				Index:   i,
			}
		}
	}

	return json.Marshal(dest)
}

// VPCRouterFirewallConfig ファイアウォール
type VPCRouterFirewallConfig struct {
	Receive []*VPCRouterFirewallRule `yaml:"receive"`
	Send    []*VPCRouterFirewallRule `yaml:"send"`

	// Index 仮想フィールド UnmarshalJSONの中で設定される
	Index int `json:"-"`
}

// VPCRouterFirewallRule ファイアウォール ルール
type VPCRouterFirewallRule struct {
	Protocol           types.Protocol           `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SourceNetwork      types.VPCFirewallNetwork `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SourcePort         types.VPCFirewallPort    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DestinationNetwork types.VPCFirewallNetwork `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DestinationPort    types.VPCFirewallPort    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Action             types.Action             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Logging            types.StringFlag         `yaml:"enabled"`
	Description        string                   `yaml:"description"`
}

// VPCRouterDHCPServer DHCPサーバ
type VPCRouterDHCPServer struct {
	Config  []*VPCRouterDHCPServerConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag             `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterDHCPServer) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterDHCPServer
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterDHCPServerConfig DHCPサーバ
type VPCRouterDHCPServerConfig struct {
	Interface  string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStop  string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStart string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DNSServers []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterDHCPStaticMappings DHCPスタティックマッピング
type VPCRouterDHCPStaticMappings struct {
	Config  []*VPCRouterDHCPStaticMappingConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                    `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterDHCPStaticMappings) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterDHCPStaticMappings
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterDHCPStaticMappingConfig DHCPスタティックマッピング
type VPCRouterDHCPStaticMappingConfig struct {
	MACAddress string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	IPAddress  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterDNSForwarding DNSフォワーディング
type VPCRouterDNSForwarding struct {
	Interface  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DNSServers []string
	Enabled    types.StringFlag `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterDNSForwarding) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if f.Interface != "" || len(f.DNSServers) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterDNSForwarding
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterPPTPServer PPTP
type VPCRouterPPTPServer struct {
	Config  *VPCRouterPPTPServerConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag           `yaml:"enabled"`
}

// VPCRouterPPTPServerConfig PPTP
type VPCRouterPPTPServerConfig struct {
	RangeStart string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStop  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterL2TPIPsecServer L2TP
type VPCRouterL2TPIPsecServer struct {
	Config  *VPCRouterL2TPIPsecServerConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                `yaml:"enabled"`
}

// VPCRouterL2TPIPsecServerConfig L2TP
type VPCRouterL2TPIPsecServerConfig struct {
	RangeStart      string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RangeStop       string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PreSharedSecret string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterWireGuard WireGuardサーバ
type VPCRouterWireGuard struct {
	Config  *VPCRouterWireGuardConfig `yaml:"config"`
	Enabled types.StringFlag          `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterWireGuard) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if f.Config == nil {
		f.Config = &VPCRouterWireGuardConfig{}
	}
	if f.Config.IPAddress != "" {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterWireGuard
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterWireGuardConfig WireGuardサーバコンフィグ
type VPCRouterWireGuardConfig struct {
	IPAddress string                    `yaml:"ip_address"` // xxx.xxx.xxx.xxx/nn
	Peers     []*VPCRouterWireGuardPeer `yaml:"peers"`
}

// MarshalJSON Peersがnullを許容しないため代わりに[]を設定する
func (f *VPCRouterWireGuardConfig) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if f.Peers == nil {
		f.Peers = []*VPCRouterWireGuardPeer{}
	}
	type alias VPCRouterWireGuardConfig
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterWireGuardPeer ピアの設定
type VPCRouterWireGuardPeer struct {
	Name      string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PublicKey string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterRemoteAccessUserConfig リモートアクセスユーザー
type VPCRouterRemoteAccessUserConfig struct {
	UserName string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Password string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterRemoteAccessUsers リモートアクセスユーザー
type VPCRouterRemoteAccessUsers struct {
	Config  []*VPCRouterRemoteAccessUserConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag                   `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterRemoteAccessUsers) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterRemoteAccessUsers
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterSiteToSiteIPsecVPN サイト間VPN
type VPCRouterSiteToSiteIPsecVPN struct {
	Config         []*VPCRouterSiteToSiteIPsecVPNConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	IKE            *VPCRouterSiteToSiteIPsecVPNIKE      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ESP            *VPCRouterSiteToSiteIPsecVPNESP      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	EncryptionAlgo string                               `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // aes128 or aes256
	HashAlgo       string                               `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // sha1 or sha256
	DHGroup        string                               `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // modp1024 or modp2048 or modp3072 or modp4096
	Enabled        types.StringFlag                     `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterSiteToSiteIPsecVPN) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterSiteToSiteIPsecVPN
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterSiteToSiteIPsecVPNConfig サイト間VPN
type VPCRouterSiteToSiteIPsecVPNConfig struct {
	Peer            string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PreSharedSecret string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	RemoteID        string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Routes          []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	LocalPrefix     []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterSiteToSiteIPsecVPNIKE サイト間VPNのIKE設定値
type VPCRouterSiteToSiteIPsecVPNIKE struct {
	Lifetime int // ISAKMP SAの寿命
	DPD      VPCRouterSiteToSiteIPsecVPNIKEDPD
}

// VPCRouterSiteToSiteIPsecVPNIKEDPD サイト間VPNのIKE設定値 - DPD
type VPCRouterSiteToSiteIPsecVPNIKEDPD struct {
	Interval int // IKEキープアライブ(DPD) インターバル
	Timeout  int // IKEキープアライブ(DPD) タイムアウト
}

// VPCRouterSiteToSiteIPsecVPNESP サイト間VPNのESP設定値
type VPCRouterSiteToSiteIPsecVPNESP struct {
	Lifetime int // IPsec SAの寿命
}

// VPCRouterStaticRoutes スタティックルート
type VPCRouterStaticRoutes struct {
	Config  []*VPCRouterStaticRouteConfig `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Enabled types.StringFlag              `yaml:"enabled"`
}

// MarshalJSON Configが一つ以上ある場合にEnabledをtrueに設定する
func (f *VPCRouterStaticRoutes) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	if len(f.Config) > 0 {
		f.Enabled = types.StringTrue
	}
	type alias VPCRouterStaticRoutes
	a := alias(*f)
	return json.Marshal(&a)
}

// VPCRouterStaticRouteConfig スタティックルート
type VPCRouterStaticRouteConfig struct {
	Prefix  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NextHop string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// VPCRouterScheduledMaintenance
type VPCRouterScheduledMaintenance struct {
	DayOfWeek int // 0が日曜日
	Hour      int // 0-23時
}

// VPCRouterStatus ステータス
type VPCRouterStatus struct {
	FirewallReceiveLogs    []string             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	FirewallSendLogs       []string             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	VPNLogs                []string             `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SessionCount           int                  `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PercentageOfMemoryFree []types.StringNumber `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DHCPServerLeases       []struct {
		IPAddress  string
		MACAddress string
	} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	L2TPIPsecServerSessions []struct {
		User      string
		IPAddress string
		TimeSec   int
	} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PPTPServerSessions []struct {
		User      string
		IPAddress string
		TimeSec   int
	} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	WireGuard *struct {
		PublicKey string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SiteToSiteIPsecVPNPeers []struct {
		Status string
		Peer   string
	} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	SessionAnalysis *struct {
		SourceAndDestination []*VPCRouterStatisticsValue `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
		DestinationAddress   []*VPCRouterStatisticsValue `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
		DestinationPort      []*VPCRouterStatisticsValue `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
		SourceAddress        []*VPCRouterStatisticsValue `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

type VPCRouterLog struct {
	Log string
}

type VPCRouterPingResult struct {
	Result []string
}

// VPCRouterStatisticsValue VPCルータのセッション統計情報
type VPCRouterStatisticsValue struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Count int    `json:"count,omitempty" yaml:"count,omitempty" structs:",omitempty"`
}
