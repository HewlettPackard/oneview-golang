package ov

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetOSDeploymentPlanByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_os_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testOsdp, err := c.GetOSDeploymentPlanByName(testName)
		assert.NoError(t, err, "GetOSDeploymentByName thew an error -> %s", err)
		assert.Equal(t, testName, testOsdp.Name)

		testOsdp, err = c.GetOSDeploymentPlanByName("bad")
		assert.NoError(t, err, "GetOSDeploymentPlanByName with fake name -> %s", err)
		assert.Equal(t, "", testOsdp.Name)

	} else {
		d, c = getTestDriverU("test_os_deployment_plan")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetOSDeploymentPlanByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetOSDeploymentPlans(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_os_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		osdps, err := c.GetOSDeploymentPlans("", "")
		assert.NoError(t, err, "GetOSDeploymentPlans threw error -> %s, %+v\n", err, osdps)

		osdps, err = c.GetOSDeploymentPlans("", "name:asc")
		assert.NoError(t, err, "GetOSDeploymentPlans name:asc error -> %s, %+v\n", err, osdps)

	} else {
		_, c = getTestDriverU("test_os_deployment_plan")
		data, err := c.GetOSDeploymentPlans("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
