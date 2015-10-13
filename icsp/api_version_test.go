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
		d *ICSPTest
		c *ICSPClient
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
		assert.True(t, d.Tc.EqualFaceI(d.Tc.GetExpectsData(d.Env+"_icsp", "CurrentVersion"), data.CurrentVersion))
		assert.True(t, d.Tc.EqualFaceI(d.Tc.GetExpectsData(d.Env+"_icsp", "MinimumVersion"), data.MinimumVersion))
	} else {
		_, c = getTestDriverU()
		data, err := c.GetAPIVersion()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}

}

// TestRefreshVersion get the latest version being used from api
func TestRefreshVersion(t *testing.T) {
	var (
		d *ICSPTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		c.APIVersion = -1
		assert.Equal(t, -1, c.APIVersion)
		err := c.RefreshVersion()
		assert.NoError(t, err, "RefreshVersion threw error -> %s", err)
		assert.True(t, d.Tc.EqualFaceI(d.Tc.GetExpectsData(d.Env+"_icsp", "CurrentVersion"), c.APIVersion))
	} else {
		_, c = getTestDriverU()
		err := c.RefreshVersion()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s", err))
	}

}
