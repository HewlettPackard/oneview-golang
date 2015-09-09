package ov
import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// get server hardware test
func TestServerHardware(t *testing.T) {
	var (
		c *OVClient
		testData    string = "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57"
		expectsData string = "2M25090RMW"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
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
		c *OVClient
		testHwType_URI  string = "/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C"
		testHWGroup_URI string = "/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAvailableHardware(testHwType_URI, testHWGroup_URI)
		assert.NoError(t, err, "GetAvailableHardware threw error -> %s", err)
		// fmt.Printf("data.Connections -> %+v\n", data)
		assert.NotEqual(t, "", data.Name)

	}
}
