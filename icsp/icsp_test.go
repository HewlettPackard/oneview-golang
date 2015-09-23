package icsp

import (
	"os"

	"github.com/docker/machine/drivers/oneview/rest"
  "github.com/docker/machine/drivers/oneview/testconfig"
)

//TODO: need to learn a better way of how integration testing works with bats
// and use that instead
// Acceptance test
// 1) craete an environment, /.oneview.env, script to export these values:
// ONEVIEW_ICSP_ENDPOINT=https://blah
// ONEVIEW_ICSP_PASSWORD=pass
// ONEVIEW_ICSP_USER=user
// ONEVIEW_ICSP_DOMAIN=LOCAL
// ONEVIEW_SSLVERIFY=true
// 2) setup gotest container
//    docker run -it --env-file ./.oneview.env -v $(pwd):/go/src/github.com/docker/machine --name testicsp docker-machine
//    exit
//    docker start goaccept
// 3) setup alias:
//   alias testicsp='docker exec goaccept godep go test -test.v=true --short'
// 4) to refresh env options, use
//    docker rm -f testicsp
//    ... repeat steps 2
type ICSPTest struct {
	Tc     *testconfig.TestConfig
	Client *ICSPClient
	Env    string
}
// get Environment
func (ot *ICSPTest) GetEnvironment() {
	if os.Getenv("ONEVIEW_TEST_ENV") != "" {
		ot.Env = os.Getenv("ONEVIEW_TEST_ENV")
		return
	}
	ot.Env = "dev"
	return
}

// get a test driver for acceptance testing
func getTestDriverA() (*ICSPTest, *ICSPClient) {
	// os.Setenv("DEBUG", "true")  // remove comment to debug logs
	var ot *ICSPTest
	var tc *testconfig.TestConfig
	ot = &ICSPTest{Tc: tc.NewTestConfig(), Env: "dev"}
	ot.GetEnvironment()
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
  ot.Client = &ICSPClient{
    rest.Client{
      User:       os.Getenv("ONEVIEW_ICSP_USER"),
      Password:   os.Getenv("ONEVIEW_ICSP_PASSWORD"),
			Domain:     os.Getenv("ONEVIEW_ICSP_DOMAIN"),
      Endpoint:   os.Getenv("ONEVIEW_ICSP_ENDPOINT"),
			// ConfigDir:
      SSLVerify:  false,
      APIVersion: 108,
			APIKey:     "none",
    },
  }
	// fmt.Println("Setting up test with getTestDriverA")
  return ot, ot.Client
}

// Unit test
func getTestDriverU() (*ICSPTest, *ICSPClient) {
	var ot *ICSPTest
	var tc *testconfig.TestConfig
	ot = &ICSPTest{Tc: tc.NewTestConfig(), Env: "dev"}
	ot.GetEnvironment()
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
  ot.Client = &ICSPClient{
    rest.Client{
      User:       "foo",
      Password:   "bar",
			Domain:     "LOCAL",
      Endpoint:   "https://icsptestcase",
      SSLVerify:  false,
      APIVersion: 108,
			APIKey:     "none",
    },
  }
	// fmt.Println("Setting up test with getTestDriverU")
  return ot, ot.Client
}
