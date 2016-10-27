package i3s

import (
	"os"
	"testing"
	//"time"

	"github.com/HewlettPackard/oneview-golang/i3s"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// Get the current idle timeout from the logged in session
func TestGetIdleTimeout(t *testing.T) {
	var (
		c *i3s.I3SClient
		// d *OVTest
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		log.Debugf(" login key -> %s", c.APIKey)

		timeout, err := c.GetIdleTimeout()
		assert.Error(t, err, "Timeout: %s", err)
		assert.NoError(t, err, "GetIdleTimeout threw error -> %s", err)
		log.Debugf(" idle timeout -> %d", timeout)
	}
}

// Set idle timeout
func TestSetIdleTimeout(t *testing.T) {
	var (
		c        *i3s.I3SClient
		testtime int64
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		testtime = 25000
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.SetIdleTimeout(testtime)
		assert.NoError(t, err, "SetIdleTimeout threw error -> %s", err)

		timeout, err := c.GetIdleTimeout()
		assert.NoError(t, err, "GetIdleTimeout threw error -> %s", err)
		assert.Equal(t, testtime, timeout, "Should get timeout equal, %s", timeout)
		log.Debugf(" idle timeout -> %d", timeout)
	}
}
