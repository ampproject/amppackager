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

package fake

import (
	"math"
	"net"
	"sync"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/cidr"
)

const (
	valuePoolMagicID     = types.ID(math.MaxInt64)
	valuePoolResourceKey = "meta"
)

func pool() *valuePool {
	InitDataStore()
	return vp
}

var vp *valuePool

type valuePool struct {
	CurrentID            int64
	CurrentSharedIP      net.IP
	SharedNetMaskLen     int
	SharedDefaultGateway net.IP
	CurrentMACAddress    net.HardwareAddr
	CurrentSubnets       map[int]string
	dataStore            Store
	mu                   sync.Mutex
}

var poolMu sync.Mutex

func initValuePool(s Store) *valuePool {
	poolMu.Lock()
	defer poolMu.Unlock()

	v := s.Get(valuePoolResourceKey, iaas.APIDefaultZone, valuePoolMagicID)
	if v != nil {
		vp = v.(*valuePool)
		vp.dataStore = s
		return vp
	}

	vp = &valuePool{
		CurrentID:            int64(100000000000),
		CurrentSharedIP:      net.IP{192, 0, 2, 2},
		SharedNetMaskLen:     24,
		SharedDefaultGateway: net.IP{192, 0, 2, 1},
		CurrentMACAddress:    net.HardwareAddr{0x00, 0x00, 0x5E, 0x00, 0x53, 0x00},
		CurrentSubnets: map[int]string{
			24: "24.0.0.0/24",
			25: "25.0.0.0/25",
			26: "26.0.0.0/26",
			27: "27.0.0.0/27",
			28: "28.0.0.0/28",
		},
		dataStore: s,
	}
	return vp
}

func (p *valuePool) store() {
	p.dataStore.Put(valuePoolResourceKey, iaas.APIDefaultZone, valuePoolMagicID, p)
}

func (p *valuePool) generateID() types.ID {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.CurrentID++

	p.store()
	return types.ID(p.CurrentID)
}

func (p *valuePool) nextSharedIP() net.IP {
	p.mu.Lock()
	defer p.mu.Unlock()

	ip := p.CurrentSharedIP.To4()
	ip[3]++
	p.CurrentSharedIP = ip
	p.store()

	ret := net.IP{0x00, 0x00, 0x00, 0x00}
	copy(ret, ip)
	return ret
}

func (p *valuePool) nextMACAddress() net.HardwareAddr {
	p.mu.Lock()
	defer p.mu.Unlock()

	mac := []byte(p.CurrentMACAddress)
	mac[5]++
	p.CurrentMACAddress = mac
	p.store()

	ret := net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	copy(ret, mac)
	return ret
}

func (p *valuePool) nextSubnet(maskLen int) *assignedSubnet {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, currentSubnet, _ := net.ParseCIDR(p.CurrentSubnets[maskLen])
	next, _ := cidr.NextSubnet(currentSubnet, maskLen) // ignore result
	p.CurrentSubnets[maskLen] = next.String()

	count := cidr.AddressCount(next)
	current := next.IP
	var defaultGateway, networkAddr string

	var addresses []string
	for i := uint64(0); i < count; i++ {
		// [0]: ネットワークアドレス
		// [1:3]: ルータ自身が利用
		// [len]: ブロードキャスト
		if i < 4 || i == count-1 {
			if i == 0 {
				networkAddr = current.String()
			}
			if i == 1 {
				defaultGateway = current.String()
			}
			current = cidr.Inc(current)
			continue
		}
		addresses = append(addresses, current.String())
		current = cidr.Inc(current)
	}

	p.store()
	return &assignedSubnet{
		defaultRoute:   defaultGateway,
		networkAddress: networkAddr,
		networkMaskLen: maskLen,
		addresses:      addresses,
	}
}

func (p *valuePool) nextSubnetFull(maskLen int, defaultRoute string) *assignedSubnet {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, currentSubnet, _ := net.ParseCIDR(p.CurrentSubnets[maskLen])
	next, _ := cidr.NextSubnet(currentSubnet, maskLen) // ignore result
	p.CurrentSubnets[maskLen] = next.String()

	count := cidr.AddressCount(next)
	current := next.IP
	var networkAddr string

	var addresses []string
	for i := uint64(0); i < count; i++ {
		addresses = append(addresses, current.String())
		current = cidr.Inc(current)
	}

	p.store()
	return &assignedSubnet{
		defaultRoute:   defaultRoute,
		networkAddress: networkAddr,
		networkMaskLen: maskLen,
		addresses:      addresses,
	}
}

type assignedSubnet struct {
	defaultRoute   string
	networkMaskLen int
	networkAddress string
	addresses      []string
}
