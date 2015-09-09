package ov
import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// get server hardware test
func TestServerHardware(t *testing.T) {
	var (
		d *OVTest
		c *OVClient
		testData string
		expectsData string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		testData    = d.Tc.GetTestData(d.Env, "ServerHardwareURI").(string)
		expectsData = d.Tc.GetExpectsData(d.Env, "SerialNumber").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetServerHardware(testData)
		assert.NoError(t, err, "GetServerHardware threw error -> %s", err)
		// fmt.Printf("data.Connections -> %+v\n", data)
		assert.Equal(t, expectsData, data.SerialNumber)

	}
}

// get server hardware test
func TestGetAvailableHardware(t *testing.T) {
	var (
		d *OVTest
		c *OVClient
		testHwType_URI  string
		testHWGroup_URI string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		testHwType_URI  = d.Tc.GetTestData(d.Env, "HardwareTypeURI").(string)
		testHWGroup_URI = d.Tc.GetTestData(d.Env, "GroupURI").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAvailableHardware(testHwType_URI, testHWGroup_URI)
		assert.NoError(t, err, "GetAvailableHardware threw error -> %s", err)
		// fmt.Printf("data.Connections -> %+v\n", data)
		assert.NotEqual(t, "", data.Name)

	}
}
