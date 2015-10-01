package icsp

import (
	"github.com/docker/machine/drivers/oneview/rest"
	"github.com/docker/machine/log"
)

// ICSPClient - wrapper class for icsp api's
type ICSPClient struct {
	rest.Client
}

// new Client
func (c *ICSPClient) NewICSPClient(user string, password string, domain string, endpoint string, sslverify bool, apiversion int) *ICSPClient {
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

// CustomServerAttributes setup custom attributes to apply to a server
type CustomServerAttributes struct {
	Values map[string]string // Hash of strings for custom attributes
}

func (cs *CustomServerAttributes) New() *CustomServerAttributes {
	return &CustomServerAttributes{
		Values: make(map[string]string),
	}
}

func (cs *CustomServerAttributes) Set(key string, value string) {
	cs.Values[key] = value
}

func (cs *CustomServerAttributes) Get(key string) string {
	return cs.Values[key]
}

// CustomizeServer - use customizeserver when working with creating a new server
// server create if it's missing
// server apply deployment job
type CustomizeServer struct {
	HostName         string                  // provide a hostname to set
	SerialNumber     string                  // should be the serial number for the server
	ILoUser          string                  // should be the user name for ilo administration
	IloPassword      string                  // should be the ilo password to use
	IloIPAddress     string                  // PXE ip address for ilo
	IloPort          int                     // port number for ilo server
	OSBuildPlan      string                  // name of the OS build plan
	ServerProperties *CustomServerAttributes // name value pairs for server custom attributes
}

// Customize Server
func (c *ICSPClient) CustomizeServer(cs CustomizeServer) error {
	s, err := c.GetServerBySerialNumber(cs.SerialNumber)
	if err != nil {
		return err
	}
	if !s.URI.IsNil() {
		if err := c.CreateServer(cs.ILoUser, cs.IloPassword, cs.IloIPAddress, cs.IloPort); err != nil {
			return err
		}
	} else {
		log.Infof("ICSP server was already created, %s, skipping", cs.HostName)
	}

	// save the server attributes to the server
	for k, v := range cs.ServerProperties.Values {
		s.SetCustomAttribute(k, "server", v)
	}

	// save it
	new_server, err := c.SaveServer(s)
	if err != nil {
		return err
	}

	// apply the build Plan
	if err := c.ApplyDeploymentJobs(cs.OSBuildPlan, new_server); err != nil {
		return err
	}
	return nil
}
