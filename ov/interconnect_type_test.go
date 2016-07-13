package ov

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetInterconnectTypes(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_interconnect_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		interconnectTypes, err := c.GetInterconnectTypes("", "")
		assert.NoError(t, err, "GetInterconnectTypes threw error -> %s, %+v\n", err, interconnectTypes)

		interconnectTypes, err = c.GetInterconnectTypes("", "name:asc")
		assert.NoError(t, err, "GetInterconnectTypes name:asc error -> %s, %+v\n", err, interconnectTypes)

	} else {
		_, c = getTestDriverU("test_interconnect_type")
		data, err := c.GetInterconnectTypes("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetInterconnectTypeByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_interconnect_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testInterconnectType, err := c.GetInterconnectTypeByName(testName)
		assert.NoError(t, err, "GetInterconnectTypeByName thew an error -> %s", err)
		assert.Equal(t, testName, testInterconnectType.Name)

		testInterconnectType, err = c.GetInterconnectTypeByName("bad")
		assert.NoError(t, err, "GetInterconnectTypeByName with fake name -> %s", err)
		assert.Equal(t, "", testInterconnectType.Name)

	} else {
		d, c = getTestDriverU("test_interconnect_type")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetInterconnectTypeByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
