package icsp

import (
	"github.com/docker/machine/drivers/oneview/rest"
)
// ICSPClient - wrapper class for icsp api's
type ICSPClient struct {
	rest.Client
}

// new Client
func (c *ICSPClient) NewICSPClient(user string, password string, domain string, endpoint string, sslverify bool, apiversion int) (*ICSPClient) {
	return &ICSPClient{
		rest.Client{
			User:       user,
			Password:   password,
			Domain:     domain,
			Endpoint:   endpoint,
			SSLVerify:  sslverify,
			APIVersion: apiversion,
			APIKey:     "none",
		},
	}
}
