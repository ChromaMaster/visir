package network_test

import (
	"errors"
	"github.com/ChromaMaster/visir/pkg/domain/system/network"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakePrivateIPAddressProvider struct {
	errorMustBeReturned bool
}

func (p *FakePrivateIPAddressProvider) GetPrivateIPAddresses() ([]network.IpAddress, error) {
	if p.errorMustBeReturned {
		return nil, errors.New("")
	}
	return []network.IpAddress{
		{"eth0", []string{"127.0.0.1"}},
		{"eth1", []string{"8.8.8.8"}},
	}, nil
}

func (p *FakePrivateIPAddressProvider) shouldReturnAnError() {
	p.errorMustBeReturned = true
}

var _ = Describe("System Info / IP Address", func() {
	var (
		ipAddressService          *network.IPAddressService
		publicAddressProvider     *FakePublicIPAddressProvider
		internalIPAddressProvider *FakePrivateIPAddressProvider
	)
	BeforeEach(func() {
		publicAddressProvider = &FakePublicIPAddressProvider{}
		internalIPAddressProvider = &FakePrivateIPAddressProvider{}
		ipAddressService = network.NewIPAddressService(publicAddressProvider, internalIPAddressProvider)
	})

	When("someone requests the public IP address", func() {
		It("returns the public IP address", func() {
			ipAddress, err := ipAddressService.GetPublicIPAddress()

			Expect(err).ToNot(HaveOccurred())
			Expect(ipAddress).To(Equal("127.0.0.1"))
		})
	})

	When("the provider is unable to extract the IP", func() {
		It("returns an error", func() {
			publicAddressProvider.shouldReturnError()

			ipAddress, err := ipAddressService.GetPublicIPAddress()

			Expect(err).To(MatchError("unable to extract IP address"))
			Expect(ipAddress).To(BeEmpty())
		})
	})

	When("someone requests the private IP addresses", func() {
		It("returns the private IP addresses", func() {
			ipAddress, err := ipAddressService.GetPrivateIPAddresses()

			Expect(err).ToNot(HaveOccurred())
			Expect(ipAddress).To(ConsistOf([]network.IpAddress{{"eth0", []string{"127.0.0.1"}}, {"eth1", []string{"8.8.8.8"}}}))
		})

		When("the external provider is unable to retrieve the IP addresses", func() {
			It("returns an error", func() {
				internalIPAddressProvider.shouldReturnAnError()

				ipAddress, err := ipAddressService.GetPrivateIPAddresses()

				Expect(err).To(MatchError("unable to extract IP addresses"))
				Expect(ipAddress).To(BeEmpty())
			})
		})
	})
})

type FakePublicIPAddressProvider struct {
	errorMustBeReturned bool
}

func (f *FakePublicIPAddressProvider) GetPublicIPAddress() (string, error) {
	if f.errorMustBeReturned {
		return "", errors.New("unknown error")
	}
	return "127.0.0.1", nil
}

func (f *FakePublicIPAddressProvider) shouldReturnError() {
	f.errorMustBeReturned = true
}
