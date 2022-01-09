package private_ip_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrivateIp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrivateIp Suite")
}
