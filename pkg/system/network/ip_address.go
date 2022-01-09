package network

import (
	"errors"
)

type IpAddress struct {
	Iface   string
	Address []string
}

type IPAddressService struct {
	publicIPAddressProvider  PublicIPAddressProvider
	privateIPAddressProvider PrivateIPAddressProvider
}

func (s *IPAddressService) GetPublicIPAddress() (string, error) {
	address, err := s.publicIPAddressProvider.GetPublicIPAddress()
	if err != nil {
		return "", errors.New("unable to extract IP address")
	}
	return address, nil
}

func (s *IPAddressService) GetPrivateIPAddresses() ([]IpAddress, error) {
	addresses, err := s.privateIPAddressProvider.GetPrivateIPAddresses()
	if err != nil {
		return nil, errors.New("unable to extract IP addresses")
	}
	return addresses, nil
}

type PublicIPAddressProvider interface {
	GetPublicIPAddress() (string, error)
}

type PrivateIPAddressProvider interface {
	GetPrivateIPAddresses() ([]IpAddress, error)
}

func NewIPAddressService(publicIPAddressProvider PublicIPAddressProvider, privateIPAddressProvider PrivateIPAddressProvider) *IPAddressService {
	return &IPAddressService{
		publicIPAddressProvider:  publicIPAddressProvider,
		privateIPAddressProvider: privateIPAddressProvider,
	}
}
