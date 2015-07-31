package ov
import (
	"os"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
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
func getTestDriverA() (*OVClient) {
  client := &OVClient{
    Client{
      User:       os.Getenv("ONEVIEW_OV_USER"),
      Password:   os.Getenv("ONEVIEW_OV_PASSWORD"),
			Domain:     os.Getenv("ONEVIEW_OV_DOMAIN"),
      Endpoint:   os.Getenv("ONEVIEW_OV_ENDPOINT"),
      SSLVerify:  false,
      APIVersion: 120,
			APIKey:     "none",
    },
  }
	// fmt.Println("Setting up test with getTestDriverA")
  return client
}

// Unit test
func getTestDriverU() (*OVClient) {
  client := &OVClient{
    Client{
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
	return client
}

// Test GetAPIVersion
func TestGetAPIVersion(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAPIVersion()
		// fmt.Printf("after GetAPIVersion: %s -> (err) %s", data.CurrentVersion, err)
		// assert.Error(t,err, fmt.Sprintf("Error caught as expected: %s",err))
		assert.NoError(t, err, "GetAPIVersion threw error -> %s", err)
		assert.Equal(t, 120, data.CurrentVersion)
		assert.Equal(t, 1, data.MinimumVersion)
	} else {
		c = getTestDriverU()
		data, err := c.GetAPIVersion()
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
	}

}

// Test SessionLogin
func TestSessionLogin(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.SessionLogin()
		// fmt.Printf("after SessionLogin: %s -> (err) %s", data.ID, err)
		assert.NoError(t, err, "SessionLogin threw error -> %s", err)
		assert.NotEmpty(t, data.ID, fmt.Sprintf("SessionLogin is empty! something went wrong, err -> %s, data -> %+v\n", err, data) )
		assert.Equal(t, "none", c.APIKey)
	} else {
		c = getTestDriverU()
		data, err := c.SessionLogin()
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
		assert.Equal(t, "none", c.APIKey)
	}
}

// TestGetProfileNameBySN
// Acceptance test ->
// /rest/server-profiles
// ?filter=serialNumber matches '2M25090RMW'&sort=name:asc
func TestGetProfileNameBySN(t *testing.T) {
	var (
		c *OVClient
		testSerial string = "2M25090RMW"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileNameBySN(testSerial)
		// fmt.Printf("after GetProfileNameBySN: %s -> (err) %s", data, err)
		assert.NoError(t, err, "GetProfileNameBySN threw error -> %s", err)
		assert.Equal(t, testSerial, data.SerialNumber)

		data, err = c.GetProfileNameBySN("foo")
		assert.NoError(t, err, "GetProfileNameBySN with fake serial number -> %+v", err)
		assert.Equal(t, "", data.SerialNumber)

	} else {
		c = getTestDriverU()
		data, err := c.GetProfileNameBySN(testSerial)
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
	}
}
