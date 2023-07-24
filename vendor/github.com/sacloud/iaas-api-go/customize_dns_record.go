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
	"fmt"
	"strings"

	"github.com/sacloud/iaas-api-go/types"
)

// Equal 名前/タイプ/値が同じレコードの場合trueを返す
func (o *DNSRecord) Equal(r *DNSRecord) bool {
	return o.Name == r.Name && o.Type == r.Type && o.RData == r.RData
}

// NewDNSRecord レコードを生成して返す
func NewDNSRecord(t types.EDNSRecordType, name, rdata string, ttl int) *DNSRecord {
	switch t {
	case
		types.DNSRecordTypes.NS,
		types.DNSRecordTypes.CNAME,
		types.DNSRecordTypes.MX,
		types.DNSRecordTypes.ALIAS,
		types.DNSRecordTypes.PTR:
		if rdata != "" && !strings.HasSuffix(rdata, ".") {
			rdata += "."
		}
	}

	return &DNSRecord{
		Name:  name,
		Type:  t,
		RData: rdata,
		TTL:   ttl,
	}
}

// MXRecord MXレコード型
type MXRecord struct {
	Name     string
	RData    string
	TTL      int
	Priority int
}

// Type レコードタイプ
func (r *MXRecord) Type() types.EDNSRecordType {
	return types.DNSRecordTypes.MX
}

// ToRecord *DNSRecord型へ変換
func (r *MXRecord) ToRecord() *DNSRecord {
	rdata := r.RData
	if rdata != "" && !strings.HasSuffix(rdata, ".") {
		rdata += "."
	}
	return &DNSRecord{
		Name:  r.Name,
		Type:  r.Type(),
		RData: fmt.Sprintf("%d %s", r.Priority, rdata),
		TTL:   r.TTL,
	}
}

// NewMXRecord MXレコードを生成して返す
func NewMXRecord(name, rdata string, ttl, priority int) *DNSRecord {
	return (&MXRecord{
		Name:     name,
		RData:    rdata,
		Priority: priority,
		TTL:      ttl,
	}).ToRecord()
}

// SRVRecord SRVレコード型
type SRVRecord struct {
	Name     string
	RData    string
	TTL      int
	Priority int
	Weight   int
	Port     int
}

// Type レコードタイプ
func (r *SRVRecord) Type() types.EDNSRecordType {
	return types.DNSRecordTypes.SRV
}

// ToRecord *DNSRecordに変換
func (r *SRVRecord) ToRecord() *DNSRecord {
	return &DNSRecord{
		Name:  r.Name,
		Type:  r.Type(),
		RData: fmt.Sprintf("%d %d %d %s", r.Priority, r.Weight, r.Port, r.RData),
		TTL:   r.TTL,
	}
}

// NewSRVRecord SRVレコードを生成して返す
func NewSRVRecord(name, rdata string, ttl, priority, weight, port int) *DNSRecord {
	return (&SRVRecord{
		Name:     name,
		RData:    rdata,
		TTL:      ttl,
		Priority: priority,
		Weight:   weight,
		Port:     port,
	}).ToRecord()
}
