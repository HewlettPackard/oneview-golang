package icsp

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/log"
	"github.com/stretchr/testify/assert"
)

// TestGetProfiles
func TestGetServers(t *testing.T) {
	var (
		// d *OVTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
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

//TODO: implement test for delete

// TestSetCustomAttribute
func TestSetCustomAttribute(t *testing.T) {
	var (
		d *ICSPTest
		// c *ICSPClient
		ca CustomAttribute
	)
	// FOR NOW acceptance and unit test on this one is the same.

	// if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
	//TODO: think about acceptance test
	// log.Debug("no acceptance test yet")
	// } else {
	// unit test case
	d, _ = getTestDriverU()
	jsonStringData := d.Tc.GetTestData(d.Env, "CustomAttributeJSONString").(string)
	testKey1 := d.Tc.GetTestData(d.Env, "KeyTest1").(string)
	testScope1 := d.Tc.GetTestData(d.Env, "ScopeTest1").(string)
	expectsValue1 := d.Tc.GetExpectsData(d.Env, "ValueTest1").(string)
	log.Infof("jsonStringData => %s", jsonStringData)
	log.Infof("testKey1 => %s, testScope1 => %s == %s", testKey1, testScope1, expectsValue1)

	err := json.Unmarshal([]byte(jsonStringData), &ca)
	assert.NoError(t, err, "Unmarshal jsonStringData threw error -> %s, %+v\n", err, jsonStringData)
	// TODO: complete test case 1, simple read of custom attributes
	// assert.Equal(t, expectsValue1, )
	// TODO: test case 2 , setting a value on existing attribute
	// TODO: test case 3, appending a new attribute
	// }

}
