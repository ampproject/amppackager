package types

import (
	"encoding/json"
	"net"
)

func ParseIPNet(str string) (*IPNet, error) {
	_, ipnet, err := net.ParseCIDR(str)
	if err != nil {
		return nil, err
	}
	return &IPNet{*ipnet}, nil
}

// +k8s:deepcopy-gen=false
type IPNet struct {
	net.IPNet
}

func (i IPNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *IPNet) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "<nil>" {
		i.IPNet = net.IPNet{}
		return nil
	}

	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}
	i.IPNet = *ipnet
	return nil
}

func (i *IPNet) DeepCopyInto(in *IPNet) {
	i.IPNet.IP = make(net.IP, len(in.IPNet.IP))
	i.IPNet.Mask = make(net.IPMask, len(in.IPNet.Mask))
	copy(i.IPNet.IP, in.IPNet.IP)
	copy(i.IPNet.Mask, in.IPNet.Mask)
}

func (i *IPNet) DeepCopy() *IPNet {
	res := &IPNet{}
	res.DeepCopyInto(i)
	return res
}
