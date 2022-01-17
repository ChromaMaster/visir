package private_ip

import (
	"github.com/ChromaMaster/visir/pkg/domain/system/network"
	"net"
)

type PrivateIPAddressProvider struct{}

func NewPrivateIPAddressProvider() *PrivateIPAddressProvider {
	return &PrivateIPAddressProvider{}
}

func (p *PrivateIPAddressProvider) GetPrivateIPAddresses() ([]network.IpAddress, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	result := make([]network.IpAddress, 0, len(ifaces))
	for _, iface := range ifaces {
		addrs, err := extractIPV4FromInterface(iface)
		if err != nil {
			return nil, err
		}

		result = append(result, network.IpAddress{Iface: iface.Name, Address: addrs})
	}

	return result, nil
}

func extractIPV4FromInterface(iface net.Interface) ([]string, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	addresses := make([]string, 0, len(addrs))
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			addresses = append(addresses, addr.String())
		}
	}
	return addresses, nil
}
