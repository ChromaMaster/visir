package public_ip_ipinfo_test

import (
	"io"
	"net/http"
	"github.com/ChromaMaster/visir/pkg/infrastructure/network/public_ip_ipinfo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Infrastructure / IPInfo", func() {
	It("returns the public IP address", func() {
		ipinfoProvider := public_ip_ipinfo.NewPublicIPAddressProvider()

		ipAddress, err := ipinfoProvider.GetPublicIPAddress()

		Expect(err).ToNot(HaveOccurred())
		Expect(ipAddress).To(Equal(actualPublicIPAddress()))
	})
})

func actualPublicIPAddress() string {
	response, err := http.Get("https://ifconfig.me")
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	return string(bytes)
}
