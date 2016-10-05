package i3s

import (
	"fmt"
	"os"
	"testing"
	//"time"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// Get the current idle timeout from the logged in session
func TestGetIdleTimeout(t *testing.T) {
	var (
		c *I3SClient
		// d *OVTest
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		log.Debugf(" login key -> %s", c.APIKey)
		assert.NoError(t, err, "RefreshLogin threw error -> %s", err)

		timeout, err := c.GetIdleTimeout()
		assert.NoError(t, err, "GetIdleTimeout threw error -> %s", err)
		log.Debugf(" idle timeout -> %d", timeout)
	}

}

// Set idle timeout
func TestSetIdleTimeout(t *testing.T) {
	var (
		c        *I3SClient
		testtime int64
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		testtime = 25000
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		err := c.RefreshLogin()
		log.Debugf(" login key -> %s", c.APIKey)
		assert.NoError(t, err, "RefreshLogin threw error -> %s", err)

		err = c.SetIdleTimeout(testtime)
		assert.NoError(t, err, "SetIdleTimeout threw error -> %s", err)

		timeout, err := c.GetIdleTimeout()
		assert.NoError(t, err, "GetIdleTimeout threw error -> %s", err)
		assert.Equal(t, testtime, timeout, "Should get timeout equal, %s", timeout)
		log.Debugf(" idle timeout -> %d", timeout)
	}

}

/*
// Test for expired key and see if RefreshLogin can restore the key if we have a bad key
func TestSessionExpiredKey(t *testing.T) {
	var (
		c *OVClient
		d *OVTest
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ethernet_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		err := c.RefreshLogin()
		log.Debugf(" login key -> %s", c.APIKey)
		assert.NoError(t, err, "RefreshLogin threw error -> %s", err)
		// force key to timeout
		err = c.SetIdleTimeout(800)
		assert.NoError(t, err, "SetIdleTimeout threw error -> %s", err)

		// 1 millisecond and we should be timeout with the current key
		time.Sleep(1 * time.Millisecond)

		// verify we are timed out
		_, err = c.GetIdleTimeout()
		assert.Error(t, err, "should be 404 not found -> %s ", err)

		// verify that we can access something from icps with this client
		// This should not fail because it uses RefreshLogin to get a new login session and avoid timeout
		ethNetName := d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetEthernetNetworkByName(ethNetName)
		assert.NoError(t, err, "GetEthNetByName threw error -> %s", err)
		assert.Equal(t, ethNetName, data.Name)

	}

}
*/
