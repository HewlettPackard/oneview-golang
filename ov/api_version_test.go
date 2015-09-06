package ov
import (
	"os"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// Test GetAPIVersion
func TestGetAPIVersion(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAPIVersion()
		// fmt.Printf("after GetAPIVersion: %s -> (err) %s", data.CurrentVersion, err)
		// assert.Error(t,err, fmt.Sprintf("Error caught as expected: %s",err))
		assert.NoError(t, err, "GetAPIVersion threw error -> %s", err)
		assert.Equal(t, 120, data.CurrentVersion)
		assert.Equal(t, 1, data.MinimumVersion)
	} else {
		c = getTestDriverU()
		data, err := c.GetAPIVersion()
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
	}

}
