package usecase

import (
	"github.com/ChromaMaster/visir/pkg/domain/system/network"
	"github.com/ChromaMaster/visir/pkg/infrastructure/network/private_ip"
	"github.com/ChromaMaster/visir/pkg/infrastructure/network/public_ip_ipinfo"
)

type UseCase interface {
	Execute(text string) string
}

type Factory interface {
	NewEchoUseCase() UseCase
	NewPublicIpUseCase() UseCase
}

type factory struct {
}

func NewFactory() Factory {
	return &factory{}
}

type EchoUseCase struct{}

func (f *factory) NewEchoUseCase() UseCase {
	return &EchoUseCase{}
}

func (e *EchoUseCase) Execute(text string) string {
	return text
}

type PublicIpUseCase struct{}

func (f *factory) NewPublicIpUseCase() UseCase {
	return &PublicIpUseCase{}
}

func (p PublicIpUseCase) Execute(text string) string {
	n := network.NewIPAddressService(public_ip_ipinfo.NewPublicIPAddressProvider(), private_ip.NewPrivateIPAddressProvider())
	ip, err := n.GetPublicIPAddress()
	if err != nil {
		return "error"
	}
	return ip
}
