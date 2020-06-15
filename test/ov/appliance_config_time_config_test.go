package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

func TestGetTimeConfigs(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_appliance_time_config")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		timeConfigs, err := c.GetTimeConfigs()
		assert.NoError(t, err, "GetTimeConfigs threw an error -> %s. %+v\n", err, timeConfigs)

		timeConfigs, err = c.GetTimeConfigs()
		assert.NoError(t, err, "GetTimeConfigs name:asc error -> %s. %+v\n", err, timeConfigs)

	} else {
		_, c = getTestDriverU("test_appliance_time_config")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetTimeConfigs("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	_, c = getTestDriverU("test_appliance_time_config")
	data, err := c.GetTimeConfigs()
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}
