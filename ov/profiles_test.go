package ov
import (
	"os"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// TestGetProfileNameBySN
// Acceptance test ->
// /rest/server-profiles
// ?filter=serialNumber matches '2M25090RMW'&sort=name:asc
func TestGetProfileNameBySN(t *testing.T) {
	var (
		c *OVClient
		testSerial string = "2M25090RMW"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileNameBySN(testSerial)
		// fmt.Printf("after GetProfileNameBySN: %s -> (err) %s", data, err)
		assert.NoError(t, err, "GetProfileNameBySN threw error -> %s", err)
		assert.Equal(t, testSerial, data.SerialNumber)

		data, err = c.GetProfileNameBySN("foo")
		assert.NoError(t, err, "GetProfileNameBySN with fake serial number -> %+v", err)
		assert.Equal(t, "", data.SerialNumber)

	} else {
		c = getTestDriverU()
		data, err := c.GetProfileNameBySN(testSerial)
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
	}
}
