package i3s

import (
	"os"
	"testing"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/testconfig"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

type I3STest struct {
	Tc     *testconfig.TestConfig
	Client *I3SClient
	Env    string
}

// get Environment
func (ot *I3STest) GetEnvironment(env string) {
	if os.Getenv("ONEVIEW_TEST_ENV") != "" {
		ot.Env = os.Getenv("ONEVIEW_TEST_ENV")
		return
	}
	ot.Env = env
	return
}

// get a test driver for acceptance testing
func getTestDriverA(env string) (*I3STest, *I3SClient) {
	// os.Setenv("DEBUG", "true")  // remove comment to debug logs
	var ot *I3STest
	var tc *testconfig.TestConfig
	ot = &I3STest{Tc: tc.NewTestConfig(), Env: "dev"}
	ot.GetEnvironment(env)
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
	ot.Client = &I3SClient{
		rest.Client{
			User:     os.Getenv("ONEVIEW_I3S_USER"),
			Password: os.Getenv("ONEVIEW_I3S_PASSWORD"),
			Domain:   os.Getenv("ONEVIEW_I3S_DOMAIN"),
			Endpoint: os.Getenv("ONEVIEW_I3S_ENDPOINT"),
			// ConfigDir:
			SSLVerify: false,
			APIKey:    "none",
		},
	}
	// TODO: implement ot.Client.RefreshVersion()
	return ot, ot.Client
}

// Unit test
func getTestDriverU(env string) (*I3STest, *I3SClient) {
	var ot *I3STest
	var tc *testconfig.TestConfig
	ot = &I3STest{Tc: tc.NewTestConfig(), Env: "dev"}
	ot.GetEnvironment(env)
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
	ot.Client = &I3SClient{
		rest.Client{
			User:       "foo",
			Password:   "bar",
			Domain:     "LOCAL",
			Endpoint:   "https://i3stestcase",
			SSLVerify:  false,
			APIVersion: 300,
			APIKey:     "none",
		},
	}
	return ot, ot.Client
}

/*
//Not sure what we are aiming to do with this. Will leave commented for now
// Test Getting New I3SClient
func TestNewI3SClient(t *testing.T) {
	var (
		c *I3SClient
	)
	log.Debug("implements unit test for TestNewI3SClient")
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
	} else {
		_, c = getTestDriverU()
	}
	assert.True(t, (c != nil), "Failed to get proper client")
}*/
