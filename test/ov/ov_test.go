package ov

import (
	"os"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/testconfig"
	"github.com/docker/machine/libmachine/log"

	"github.com/HewlettPackard/oneview-golang/ov"
)

//TODO: need to learn a better way of how integration testing works with bats
// and use that instead
// Acceptance test
// 1) craete an environment, /.oneview.env, script to export these values:
// ONEVIEW_OV_ENDPOINT=https://blah
// ONEVIEW_OV_PASSWORD=pass
// ONEVIEW_OV_USER=user
// ONEVIEW_OV_DOMAIN=LOCAL
// ONEVIEW_SSLVERIFY=true
// 2) setup gotest container
//    docker run -it --env-file ./.oneview.env -v $(pwd):/go/src/github.com/docker/machine --name goaccept docker-machine
//    exit
//    docker start goaccept
// 3) setup alias:
//   alias goaccept='docker exec goaccept godep go test -test.v=true --short'
// 4) to refresh env options, use
//    docker rm -f goaccept
//    ... repeat steps 2
type OVTest struct {
	Tc     *testconfig.TestConfig
	Client *ov.OVClient
	Env    string
}

// get Environment
func (ot *OVTest) GetEnvironment(env string) {
	if os.Getenv("ONEVIEW_TEST_ENV") != "" {
		ot.Env = os.Getenv("ONEVIEW_TEST_ENV")
		return
	}
	ot.Env = env
	return
}

// get a test driver for acceptance testing
func getTestDriverA(env string) (*OVTest, *ov.OVClient) {
	// os.Setenv("DEBUG", "true")  // remove comment to debug logs
	var ot *OVTest
	var tc *testconfig.TestConfig
	ot = &OVTest{Tc: tc.NewTestConfig(), Env: env}
	ot.GetEnvironment(env)
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
	ot.Client = &ov.OVClient{
		rest.Client{
			User:     os.Getenv("ONEVIEW_OV_USER"),
			Password: os.Getenv("ONEVIEW_OV_PASSWORD"),
			Domain:   os.Getenv("ONEVIEW_OV_DOMAIN"),
			Endpoint: os.Getenv("ONEVIEW_OV_ENDPOINT"),
			// ConfigDir:
			SSLVerify: false,
			APIKey:    "none",
		},
	}
	err := ot.Client.RefreshVersion()
	if err != nil {
		log.Errorf("Problem with getting api version refreshed : %+v", err)
	}
	// fmt.Println("Setting up test with getTestDriverA")
	return ot, ot.Client
}

// Unit test
func getTestDriverU(env string) (*OVTest, *ov.OVClient) {
	var ot *OVTest
	var tc *testconfig.TestConfig
	ot = &OVTest{Tc: tc.NewTestConfig(), Env: env}
	ot.GetEnvironment(env)
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
	ot.Client = &ov.OVClient{
		rest.Client{
			User:       "foo",
			Password:   "bar",
			Domain:     "LOCAL",
			Endpoint:   "https://ovtestcase",
			SSLVerify:  false,
			APIVersion: 120,
			APIKey:     "none",
		},
	}
	// fmt.Println("Setting up test with getTestDriverU")
	return ot, ot.Client
}
