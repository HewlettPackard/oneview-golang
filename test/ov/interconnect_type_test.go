package ov

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetInterconnectTypes(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_interconnect_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		interconnectTypes, err := c.GetInterconnectTypes("", "", "", "")
		assert.NoError(t, err, "GetInterconnectTypes threw error -> %s, %+v\n", err, interconnectTypes)

		interconnectTypes, err = c.GetInterconnectTypes("", "", "", "name:asc")
		assert.NoError(t, err, "GetInterconnectTypes name:asc error -> %s, %+v\n", err, interconnectTypes)

	} else {
		_, c = getTestDriverU("test_interconnect_type")
		data, err := c.GetInterconnectTypes("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

/*
func TestGetInterconnectTypeByURI(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testURI utils.Nstring
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_interconnect_type")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testURI = utils.NewNstring(d.Tc.GetTestData(d.Env, "URI").(string))

		testInterconnectType, err := c.GetInterconnectTypeByUri(testURI)
		assert.Error(t, err, "GetInterconnectTypeByName thew an error -> %+v", testInterconnectType)
		assert.Equal(t, testURI, testInterconnectType.URI)

		testInterconnectType, err = c.GetInterconnectTypeByName("bad")
		assert.NoError(t, err, "GetInterconnectTypeByURI with fake name -> %s", err)
		assert.Equal(t, "", testInterconnectType.Name)

	} else {
		d, c = getTestDriverU("test_interconnect_type")
		testURI = utils.NewNstring(d.Tc.GetTestData(d.Env, "URI").(string))
		data, err := c.GetInterconnectTypeByUri(testURI)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
*/
