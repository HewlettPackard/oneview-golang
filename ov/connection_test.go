package ov
import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test working with connections
// Acceptance test ->
// /rest/server-profiles
// ?filter=serialNumber matches '2M25090RMW'&sort=name:asc
func TestConnections(t *testing.T) {
	var (
		c *OVClient
		testSerial string = "2M25090RMW"
		getsmac    string = "34:64:A9:BB:E6:98"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileBySN(testSerial)
		assert.NoError(t, err, "GetProfileBySN threw error -> %s", err)
		// fmt.Printf("data.Connections -> %+v\n", data)
		assert.Equal(t, getsmac, data.Connections[0].MAC)

	}
}
