package cloud

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/windmilleng/tilt/internal/token"
)

// an address like cloud.tilt.dev or localhost:10450
type Address string

const addressEnvName = "TILT_CLOUD_ADDRESS"

func ProvideAddress() Address {
	address := os.Getenv(addressEnvName)
	if address == "" {
		address = "alerts.tilt.dev"
	}

	return Address(address)
}

func RegisterTokenURL(cloudAddress string, t token.Token) string {
	return fmt.Sprintf("https://%s/register_token?token=%s", cloudAddress, t)
}

func URL(cloudAddress string) *url.URL {
	var u url.URL
	u.Host = cloudAddress
	u.Scheme = "https"
	if strings.Split(cloudAddress, ":")[0] == "localhost" {
		u.Scheme = "http"
	}
	return &u
}