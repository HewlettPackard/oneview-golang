package ov

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetServerHardwareTypeByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_server_hardware_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testServerHardwareType, err := c.GetServerHardwareTypeByName(testName)
		assert.NoError(t, err, "GetServerHardwareTypeByName thew an error -> %s", err)
		assert.Equal(t, testName, testServerHardwareType.Name)

		testServerHardwareType, err = c.GetServerHardwareTypeByName("bad")
		assert.Error(t, err, "GetServerHardwareTypeByName with fake name -> %s", err)
		assert.Equal(t, "", testServerHardwareType.Name)

	} else {
		d, c = getTestDriverU("test_server_hardware_type")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetServerHardwareTypeByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetServerHardwareTypes(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_server_hardware_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		serverHardwareTypes, err := c.GetServerHardwareTypes(0,0,"", "")
		assert.NoError(t, err, "GetServerHardwareTypes threw error -> %s, %+v\n", err, serverHardwareTypes)

		serverHardwareTypes, err = c.GetServerHardwareTypes(0,0,"", "name:asc")
		assert.NoError(t, err, "GetServerHardwareTypes name:asc error -> %s, %+v\n", err, serverHardwareTypes)

	} else {
		_, c = getTestDriverU("test_server_hardware_type")
		data, err := c.GetServerHardwareTypes(0,0,"", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
