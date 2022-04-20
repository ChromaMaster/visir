package public_ip_ipinfo

import (
	"io"
	"net/http"
)

type PublicIPAddressProvider struct{}

func (p *PublicIPAddressProvider) GetPublicIPAddress() (string, error) {
	response, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func NewPublicIPAddressProvider() *PublicIPAddressProvider {
	return &PublicIPAddressProvider{}
}
