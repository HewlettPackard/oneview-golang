package icsp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

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
	PublicSlotID     int                     // the public interface that will be used to get public ipaddress
}

// PreApplyDeploymentJobs - will attempt to identify the public interface for this job
//  for now we simply look for interfaces on ethx and save those into a custom attribute called
//  PublicInterface, this can be controlled by providing the slot id.   Example:
//  slotid = x , will map to ethx
//  slotid = 1 , will map to eth1
//  slotid = 2 , will map to eth2
//  public_interface can only be called when the server is in maintenance mode, all others
//   simply fall out
//TODO: a workaround to figuring out how to bubble up public ip address information from the os to icsp after os build plan provisioning
func (c *ICSPClient) PreApplyDeploymentJobs(s Server, slotid int) error {
	var publicinterface Interface
	if PRE_UNPROVISIONED.Equal(s.OpswLifecycle) && S_MAINTENANCE.Equal(s.State) {
		log.Debugf("Applying pre deployment job settings")
		inets := s.GetInterfaces()
		for _, inet := range inets {
			if id, err := strconv.Atoi(inet.Slot[len(inet.Slot)-1:]); id == slotid {
				if err != nil {
					return err
				}
				publicinterface = inet
				break
			}
		}
		// json version of the publicinterface
		publicinterfacejson, err := json.Marshal(publicinterface)
		if err != nil {
			return err
		}
		// save the publicinterface into a custom attribute called public_interface
		s.SetCustomAttribute("public_interface", "server", fmt.Sprintf("%s", bytes.NewBuffer(publicinterfacejson)))

		// save it
		_, err = c.SaveServer(s)
		if err != nil {
			return err
		}
	} else {
		log.Debugf("Skippling the pre-apply deployment jobs settings")
	}
	return nil
}

// Customize Server
func (c *ICSPClient) CustomizeServer(cs CustomizeServer) error {
	s, err := c.GetServerBySerialNumber(cs.SerialNumber)
	if err != nil {
		return err
	}
	if s.SerialNumber != cs.SerialNumber {
		if err := c.CreateServer(cs.ILoUser, cs.IloPassword, cs.IloIPAddress, cs.IloPort); err != nil {
			return err
		}
	} else {
		log.Infof("ICSP server was already created, %s, skipping", cs.HostName)
	}

	// reload that server
	s, err = c.GetServerBySerialNumber(cs.SerialNumber)
	if err != nil {
		return err
	}

	// verify that the server actually has a URI
	if s.URI.IsNil() {
		return fmt.Errorf("Server customization failure, unable to get a valid server from icsp for serial number: %s", cs.SerialNumber)
	}

	// verify that the server retrieved matches it's serial number
	if s.SerialNumber != cs.SerialNumber {
		return fmt.Errorf("Server customization failure, server serial numbers mismatch for %s.", cs.SerialNumber)
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

	// call to capture the public_interface attribute
	if err := c.PreApplyDeploymentJobs(new_server, cs.PublicSlotID); err != nil {
		return err
	}

	// apply the build Plan
	if err := c.ApplyDeploymentJobs(cs.OSBuildPlan, new_server); err != nil {
		return err
	}
	return nil
}
