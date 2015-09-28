package icsp

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetProfiles
func TestGetServers(t *testing.T) {
	var (
		// d *OVTest
		c *ICSPClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
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

func TestGetServerByName(t *testing.T) {
	var (
		c *ICSPClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetServerByName("sRack03-se05-16")
		assert.NoError(t, err, "GetServerByName threw error -> %s, %+v\n", err, data)
	}
}

func TestGetServerByHostName(t *testing.T) {
	var (
		c *ICSPClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetServerByHostName("ezsetupsystem3464a9bbe698")
		assert.NoError(t, err, "GetServerByName threw error -> %s, %+v\n", err, data)
	}
}

func TestGetServerBySerialNumber(t *testing.T) {
	var (
		c *ICSPClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetServerBySerialNumber("2M251204DF")
		assert.NoError(t, err, "GetServerByName threw error -> %s, %+v\n", err, data)
	}
}

//TODO: implement test for delete
