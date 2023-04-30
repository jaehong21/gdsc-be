package entity

import (
	"fmt"
	"net"
)

type Building struct {
	ID             string `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	Location       string `db:"location" json:"location"`
	NetworkAddress string `db:"network_address" json:"network_address,omitempty"`
	SubnetMask     string `db:"subnet_mask" json:"subnet_mask,omitempty"`
}

func (b *Building) VerifyRequestIP(requestIP string) error {
	requestNetowrkAddress, err := GetNetworkAddress(requestIP, b.SubnetMask)
	if err != nil {
		return err
	}

	if b.NetworkAddress != requestNetowrkAddress {
		return fmt.Errorf("invalid request IP address: %s", requestIP)
	}

	return nil
}

func GetNetworkAddress(requestIP string, subnetMask string) (string, error) {
	ip := net.ParseIP(requestIP).To4()
	if ip == nil {
		return "", fmt.Errorf("invalid IP address: %s", requestIP)
	}

	mask := net.IPMask(net.ParseIP(subnetMask).To4())
	if mask == nil {
		return "", fmt.Errorf("invalid subnet mask: %s", subnetMask)
	}

	network := net.IP(make([]byte, len(ip)))
	for i := 0; i < len(ip); i++ {
		network[i] = ip[i] & mask[i]
	}

	return network.String(), nil
}
