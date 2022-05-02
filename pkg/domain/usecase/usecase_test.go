package usecase_test

import (
	"github.com/ChromaMaster/visir/pkg/domain/system/network"
	"github.com/ChromaMaster/visir/pkg/domain/usecase"
	"github.com/ChromaMaster/visir/pkg/infrastructure/network/private_ip"
	"github.com/ChromaMaster/visir/pkg/infrastructure/network/public_ip_ipinfo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Echo Use Case", func() {
	It("returns the same text that was provided", func() {
		f := usecase.NewFactory()
		echoUsecase := f.NewEchoUseCase()
		result := echoUsecase.Execute("text")
		Expect(result).To(Equal("text"))
	})
})

var _ = Describe("PublicIP Use Case", func() {
	It("returns the public ip", func() {
		f := usecase.NewFactory()
		publicIpUsecase := f.NewPublicIpUseCase()
		result := publicIpUsecase.Execute("")

		net := network.NewIPAddressService(public_ip_ipinfo.NewPublicIPAddressProvider(), private_ip.NewPrivateIPAddressProvider())
		publicIp, _ := net.GetPublicIPAddress()

		Expect(result).To(Equal(publicIp))
	})
})
