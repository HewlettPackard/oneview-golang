package icsp

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// TestGetProfiles
func TestGetAllBuildPlans(t *testing.T) {
	var (
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAllBuildPlans()
		assert.NoError(t, err, "GetServers threw error -> %s, %+v\n", err, data)

	} else {
		_, c = getTestDriverU()
		data, err := c.GetServers()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

// TestGetProfiles
func TestGetBuildPlanByName(t *testing.T) {
	var (
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetBuildPlanByName("ProLiant OS - RHEL 7.0 x64 Scripted Install")
		assert.NoError(t, err, "GetServers threw error -> %s, %+v\n", err, data)
		if assert.NotNil(t, data.URI) {
			log.Debug(fmt.Sprintf("plan uri: %v", data.URI))
		}

	} else {
		_, c = getTestDriverU()
		data, err := c.GetServers()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))

	}
}
