package ov

import (
	"fmt"
	"os"
	"testing"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateApplianceTimeandLocal(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_appliancetime_and_local")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		applianceTimeandLocal := ov.ApplianceTimeandLocal{
			locale:     "en_US.UTF-8",
			dateTime:   "2014-09-11T12:10:33",
			timezone:   "UTC"
		}
		err := c.CreateApplianceTimeandLocal(applianceTimeandLocal)
		assert.NoError(t, err, "CreateApplianceTimeandLocal error -> %s", err)
 		err = c.CreateApplianceTimeandLocal(applianceTimeandLocal)
		assert.Error(t, err, "CreateApplianceTimeandLocal should error becaue the network already exists, err -> %s", err)
	}
}

func TestGetApplianceTimenandLocals(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_appliancetime_and_local")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		fcNetworks, err := c.GetApplianceTimeandLocals("", "", "", "")
		assert.NoError(t, err, "GetApplianceTimeandLocals threw an error -> %s. %+v\n", err, fcNetworks)

	} else {
		_, c = getTestDriverU("test_appliancetime_and_local")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetApplianceTimeandLocals("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}
}