package ov

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetSwitchTypes(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_switch_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		switchTypes, err := c.GetSwitchTypes("", "")
		assert.NoError(t, err, "GetSwitchTypes threw error -> %s, %+v\n", err, switchTypes)

		switchTypes, err = c.GetSwitchTypes("", "name:asc")
		assert.NoError(t, err, "GetSwitchTypes name:asc error -> %s, %+v\n", err, switchTypes)

	} else {
		_, c = getTestDriverU("test_switch_type")
		data, err := c.GetSwitchTypes("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetSwitchTypeByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_switch_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testSwitchType, err := c.GetSwitchTypeByName(testName)
		assert.NoError(t, err, "GetSwitchTypeByName thew an error -> %s", err)
		assert.Equal(t, testName, testSwitchType.Name)

		testSwitchType, err = c.GetSwitchTypeByName("bad")
		assert.NoError(t, err, "GetSwitchTypeByName with fake name -> %s", err)
		assert.Equal(t, "", testSwitchType.Name)

	} else {
		d, c = getTestDriverU("test_switch_type")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetSwitchTypeByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
