package icsp

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetAPIVersion
func TestGetAPIVersion(t *testing.T) {
	var (
		d *OVTest
		c *OVClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAPIVersion()
		// fmt.Printf("after GetAPIVersion: %s -> (err) %s", data.CurrentVersion, err)
		// assert.Error(t,err, fmt.Sprintf("Error caught as expected: %s",err))
		assert.NoError(t, err, "GetAPIVersion threw error -> %s", err)
		assert.Equal(t, true, d.Tc.EqualFaceI(d.Tc.GetExpectsData(d.Env, "CurrentVersion"), data.CurrentVersion))
		assert.Equal(t, true, d.Tc.EqualFaceI(d.Tc.GetExpectsData(d.Env, "MinimumVersion"), data.MinimumVersion))
	} else {
		_, c = getTestDriverU()
		data, err := c.GetAPIVersion()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}

}
