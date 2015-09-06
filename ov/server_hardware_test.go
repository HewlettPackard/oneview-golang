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
