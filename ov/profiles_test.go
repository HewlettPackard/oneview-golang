package ov
import (
	"os"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// find Server_Profile_scs79
func TestGetProfileByName(t *testing.T) {
	var (
		c *OVClient
		testname string = "Server_Profile_scs79"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileByName(testname)
		assert.NoError(t, err, "GetProfileByName threw error -> %s", err)
		assert.Equal(t, testname, data.Name)

		data, err = c.GetProfileByName("foo")
		assert.NoError(t, err, "GetProfileByName with fake name -> %+v", err)
		assert.Equal(t, "", data.Name)

	} else {
		c = getTestDriverU()
		data, err := c.GetProfileByName(testname)
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
	}
}


// TestGetProfileNameBySN
// Acceptance test ->
// /rest/server-profiles
// ?filter=serialNumber matches '2M25090RMW'&sort=name:asc
func TestGetProfileBySN(t *testing.T) {
	var (
		c *OVClient
		testSerial string = "2M25090RMW"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileBySN(testSerial)
		assert.NoError(t, err, "GetProfileBySN threw error -> %s", err)
		assert.Equal(t, testSerial, data.SerialNumber)

		data, err = c.GetProfileBySN("foo")
		assert.NoError(t, err, "GetProfileBySN with fake serial number -> %+v", err)
		assert.Equal(t, "", data.SerialNumber)

	} else {
		c = getTestDriverU()
		data, err := c.GetProfileBySN(testSerial)
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
	}
}

// TestGetProfiles
func TestGetProfiles(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfiles("","")
		assert.NoError(t, err, "GetProfiles threw error -> %s, %+v\n", err, data)

		data, err = c.GetProfiles("", "name:asc")
		assert.NoError(t, err, "GetProfiles name:asc error -> %s, %+v", err, data)

	} else {
		c = getTestDriverU()
		data, err := c.GetProfiles("","")
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
