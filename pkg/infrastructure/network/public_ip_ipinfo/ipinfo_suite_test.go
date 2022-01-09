package public_ip_ipinfo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIpinfo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ipinfo Suite")
}
