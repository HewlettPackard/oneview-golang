package ov
import (
	"os"

	"github.com/docker/machine/drivers/oneview/rest"
  "github.com/docker/machine/drivers/oneview/testconfig"
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
	Tc testconfig.TestConfig
	Client *OVClient
}
func getTestDriverA() (*OVTest, *OVClient) {
	// os.Setenv("DEBUG", "true")  // remove comment to debug logs
	var ot *OVTest
	ot.Tc.GetTestingConfiguration(os.Getenv("ONEVIEW_TEST_DATA"))
  client := &OVClient{
    rest.Client{
      User:       os.Getenv("ONEVIEW_OV_USER"),
      Password:   os.Getenv("ONEVIEW_OV_PASSWORD"),
			Domain:     os.Getenv("ONEVIEW_OV_DOMAIN"),
      Endpoint:   os.Getenv("ONEVIEW_OV_ENDPOINT"),
			// ConfigDir:
      SSLVerify:  false,
      APIVersion: 120,
			APIKey:     "none",
    },
  }
	// fmt.Println("Setting up test with getTestDriverA")
  ot.Client = &client
  return &ot, &client
}

// Unit test
func getTestDriverU() (*OVTest, *OVClient) {
	var ot *OVTest
  client := &OVClient{
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
  ot.Client = &client
  return &ot, &client
}
