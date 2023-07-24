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

// ESimpleMonitorProtocol シンプル監視 プロトコル
type ESimpleMonitorProtocol string

// String ESimpleMonitorProtocolの文字列表現
func (p ESimpleMonitorProtocol) String() string {
	return string(p)
}

// SimpleMonitorProtocols シンプル監視 プロトコル
var SimpleMonitorProtocols = struct {
	// HTTP http
	HTTP ESimpleMonitorProtocol
	// HTTPS https
	HTTPS ESimpleMonitorProtocol
	// Ping ping
	Ping ESimpleMonitorProtocol
	// TCP tcp
	TCP ESimpleMonitorProtocol
	// DNS dns
	DNS ESimpleMonitorProtocol
	// SSH ssh
	SSH ESimpleMonitorProtocol
	// SMTP smtp
	SMTP ESimpleMonitorProtocol
	// POP3 pop3
	POP3 ESimpleMonitorProtocol
	// SNMP snmp
	SNMP ESimpleMonitorProtocol
	// SSLCertificate sslcertificate
	SSLCertificate ESimpleMonitorProtocol
	// FTP ftp
	FTP ESimpleMonitorProtocol
}{
	HTTP:           ESimpleMonitorProtocol("http"),
	HTTPS:          ESimpleMonitorProtocol("https"),
	Ping:           ESimpleMonitorProtocol("ping"),
	TCP:            ESimpleMonitorProtocol("tcp"),
	DNS:            ESimpleMonitorProtocol("dns"),
	SSH:            ESimpleMonitorProtocol("ssh"),
	SMTP:           ESimpleMonitorProtocol("smtp"),
	POP3:           ESimpleMonitorProtocol("pop3"),
	SNMP:           ESimpleMonitorProtocol("snmp"),
	SSLCertificate: ESimpleMonitorProtocol("sslcertificate"),
	FTP:            ESimpleMonitorProtocol("ftp"),
}

// SimpleMonitorProtocolStrings シンプル監視プロトコルの文字列リスト
var SimpleMonitorProtocolStrings = []string{
	SimpleMonitorProtocols.HTTP.String(),
	SimpleMonitorProtocols.HTTPS.String(),
	SimpleMonitorProtocols.Ping.String(),
	SimpleMonitorProtocols.TCP.String(),
	SimpleMonitorProtocols.DNS.String(),
	SimpleMonitorProtocols.SSH.String(),
	SimpleMonitorProtocols.SMTP.String(),
	SimpleMonitorProtocols.POP3.String(),
	SimpleMonitorProtocols.SNMP.String(),
	SimpleMonitorProtocols.SSLCertificate.String(),
	SimpleMonitorProtocols.FTP.String(),
}
