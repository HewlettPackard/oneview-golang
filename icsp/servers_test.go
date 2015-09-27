package icsp

import (
	"fmt"
	"os"
	"testing"

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
