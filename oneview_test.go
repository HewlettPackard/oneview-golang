package oneview

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/drivers"
	"github.com/docker/machine/drivers/oneview/icsp"
	"github.com/docker/machine/drivers/oneview/ov"
	"github.com/docker/machine/drivers/oneview/rest"
	"github.com/docker/machine/drivers/oneview/testconfig"
	"github.com/stretchr/testify/assert"
)

// import (
// 	"io/ioutil"
// 	"os"
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// )

type OneViewTest struct {
	Tc         *testconfig.TestConfig
	OVClient   *ov.OVClient
	ICSPClient *icsp.ICSPClient
	Env        string
}

// get Environment
func (ot *OneViewTest) GetEnvironment() {
	if os.Getenv("ONEVIEW_TEST_ENV") != "" {
		ot.Env = os.Getenv("ONEVIEW_TEST_ENV")
		return
	}
	ot.Env = "dev"
	return
}

// get a test driver for acceptance testing
func (ot *OneViewTest) GetTestDriverA() (*OneViewTest, *ov.OVClient, *icsp.ICSPClient) {
	// os.Setenv("DEBUG", "true")  // remove comment to debug logs
	var tc *testconfig.TestConfig
	ot = &OneViewTest{Tc: tc.NewTestConfig(), Env: "dev"}
	ot.GetEnvironment()
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
	ot.ICSPClient = &icsp.ICSPClient{
		rest.Client{
			User:     os.Getenv("ONEVIEW_ICSP_USER"),
			Password: os.Getenv("ONEVIEW_ICSP_PASSWORD"),
			Domain:   os.Getenv("ONEVIEW_ICSP_DOMAIN"),
			Endpoint: os.Getenv("ONEVIEW_ICSP_ENDPOINT"),
			// ConfigDir:
			SSLVerify:  false,
			APIVersion: 108,
			APIKey:     "none",
		},
	}

	ot.OVClient = &ov.OVClient{
		rest.Client{
			User:     os.Getenv("ONEVIEW_OV_USER"),
			Password: os.Getenv("ONEVIEW_OV_PASSWORD"),
			Domain:   os.Getenv("ONEVIEW_OV_DOMAIN"),
			Endpoint: os.Getenv("ONEVIEW_OV_ENDPOINT"),
			// ConfigDir:
			SSLVerify:  false,
			APIVersion: 120,
			APIKey:     "none",
		},
	}
	// fmt.Println("Setting up test with getTestDriverA")
	return ot, ot.OVClient, ot.ICSPClient
}

func TestCreateServer(t *testing.T) {
	var (
		driver             Driver
		d                  *OneViewTest
		c                  *ov.OVClient
		ic                 *icsp.ICSPClient
		template, hostname string
	)
	d, c, ic = d.GetTestDriverA()

	template = d.Tc.GetTestData(d.Env, "TemplateProfile").(string)
	hostname = d.Tc.GetTestData(d.Env, "HostName").(string)
	inner := drivers.NewBaseDriver(hostname, "", "", "")
	driver = Driver{
		ClientICSP: ic,
		ClientOV:   c,
		BaseDriver: inner,
	}

	err := c.CreateMachine(hostname, template)
	assert.NoError(t, err, "CreateMachine threw error -> %s\n", err)

	err = driver.getBlade()
	assert.NoError(t, err, "getBlade threw error -> %s\n", err)

	// power on the server, and leave it in that state
	err = driver.Hardware.PowerOn()
	assert.NoError(t, err, "PowerOn threw error -> %s\n", err)

	// power on the server, and leave it in that state
	err = driver.Hardware.PowerOff()
	assert.NoError(t, err, "PowerOff threw error -> %s\n", err)
}

func TestCustomizeServer(t *testing.T) {
	var (
		err                                                 error
		driver                                              Driver
		d                                                   *OneViewTest
		c                                                   *ov.OVClient
		ic                                                  *icsp.ICSPClient
		serialNumber, osbuildplan, hostname, user, pass, ip string
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		user = os.Getenv("ONEVIEW_ILO_USER")
		pass = os.Getenv("ONEVIEW_ILO_PASSWORD")
		d, c, ic = d.GetTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		ip = d.Tc.GetTestData(d.Env, "IloIPAddress").(string)
		// serialNumber := d.Tc.GetTestData(d.Env, "FreeBladeSerialNumber").(string)
		hostname = d.Tc.GetTestData(d.Env, "HostName").(string)
		inner := drivers.NewBaseDriver(hostname, "", "", "")
		driver = Driver{
			ClientICSP: ic,
			ClientOV:   c,
			BaseDriver: inner,
		}
		err = driver.getBlade()
		assert.NoError(t, err, "getBlade threw error -> %s\n", err)
		serialNumber = driver.Profile.SerialNumber.String()

		osbuildplan = d.Tc.GetTestData(d.Env, "OSBuildPlan").(string)

		var sp *icsp.CustomServerAttributes
		sp = sp.New()
		sp.Set("docker_user", "docker")
		// TODO: make a util for this
		if len(os.Getenv("proxy_enable")) > 0 {
			sp.Set("proxy_enable", os.Getenv("proxy_enable"))
		} else {
			sp.Set("proxy_enable", "false")
		}

		strProxy := os.Getenv("proxy_config")
		sp.Set("proxy_config", strProxy)

		sp.Set("docker_hostname", hostname+"-@server_name@")
		// interface
		sp.Set("interface", fmt.Sprintf("eno%d", 50)) // TODO: what argument should we call 50 besides slotid ??

		// check if the server now exist
		cs := icsp.CustomizeServer{
			HostName:         hostname,     // machine-rack-enclosure-bay
			SerialNumber:     serialNumber, // get it
			ILoUser:          user,
			IloPassword:      pass,
			IloIPAddress:     ip,
			IloPort:          443,
			OSBuildPlan:      osbuildplan, // name of the OS build plan
			PublicSlotID:     1,           // this is the slot id of the public interface
			ServerProperties: sp,
		}
		// create d.Server and apply a build plan and configure the custom attributes
		err := ic.CustomizeServer(cs)
		assert.NoError(t, err, "CustomizeServer threw error -> %s\n", err)

	}
}
