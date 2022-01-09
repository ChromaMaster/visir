package private_ip_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ChromaMaster/visir/pkg/infrastructure/network/private_ip"
	"github.com/ChromaMaster/visir/pkg/system/network"
)

var _ = Describe("Infrastructure / InternalIP", func() {
	It("returns the internal IP address list", func() {
		ipProvider := private_ip.NewPrivateIPAddressProvider()

		ipList, err := ipProvider.GetPrivateIPAddresses()

		Expect(err).ToNot(HaveOccurred())
		Expect(ipList).To(ContainElements(
			network.IpAddress{Iface: "lo", Address: []string{"127.0.0.1/8"}},
			network.IpAddress{Iface: "wlp6s0", Address: []string{"192.168.1.36/24"}},
			network.IpAddress{Iface: "docker0", Address: []string{"172.17.0.1/16"}},
		))
	})
})
