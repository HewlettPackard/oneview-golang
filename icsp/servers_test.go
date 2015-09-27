package icsp

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/log"
	"github.com/stretchr/testify/assert"
)

//TODO: implement create server
func TestCreateServer(t *testing.T) {
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		log.Debug("implements acceptance test for TestCreateServer")
		// check that the server doesn't exist
		// create the server
		// check if the server now exist
	} else {
		log.Debug("implements unit test for TestCreateServer")
	}
}

//TODO: implement save server
func TestSaveServer(t *testing.T) {
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		log.Debug("implements acceptance test for TestCreateServer")
		// get a Server
		// set a custom attribute
		// save a server
		// verify that the server attribute was saved by getting the server again and checking the value
	} else {
		log.Debug("implements unit test for TestCreateServer")
	}
}

// TestGetProfiles
func TestGetServers(t *testing.T) {
	var (
		// d *OVTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetServers()
		assert.NoError(t, err, "GetServers threw error -> %s, %+v\n", err, data)

	} else {
		_, c = getTestDriverU()
		data, err := c.GetServers()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

//TODO: implement test for delete
func TestDeleteServer(t *testing.T) {
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		log.Debug("implements acceptance test for TestDeleteServer")
		// check if the server exist
		// delete the server
	} else {
		log.Debug("implements unit test for TestDeleteServer")
	}

}
