package ov

import (
	"fmt"
	"os"
	"testing"
	//"time"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// Test SessionLogin
func TestSessionLogin(t *testing.T) {
	var (
		// d *OVTest
		c *OVClient
		// env = os.Getenv("ONEVIEW_TEST_ENV") || "dev"
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.SessionLogin()
		log.Debugf("after SessionLogin: %s -> (err) %s", data.ID, err)

		assert.NoError(t, err, "SessionLogin threw error -> %s", err)
		assert.NotEmpty(t, data.ID, fmt.Sprintf("SessionLogin is empty! something went wrong, err -> %s, data -> %+v\n", err, data))
		assert.Equal(t, "none", c.APIKey)

		c.APIKey = data.ID
		err = c.SessionLogout()
		assert.NoError(t, err, "SessionLogout threw error -> %s", err)
		data, err = c.SessionLogin()
		assert.NoError(t, err, "SessionLogin threw error -> %s", err)
	} else {
		_, c = getTestDriverU("dev")
		data, err := c.SessionLogin()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
		assert.Equal(t, "none", c.APIKey)
	}
}

// Test SessionLogout
func TestSessionLogout(t *testing.T) {
	var (
		//d *OVTest
		c *OVClient
		//testSerial string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		//data, err := c.SessionLogin()
		// this is needed so we can "copy" the session id to the ov client
		err := c.RefreshLogin()
		//log.Debugf(" login key -> %s, session data -> %+v", c.APIKey, data)
		log.Debugf(" login key -> %s", c.APIKey)
		assert.NoError(t, err, "SessionLogin threw error -> %s", err)
		//assert.NotEmpty(t, data.ID, fmt.Sprintf("SessionLogin is empty! something went wrong, err -> %s, data -> %+v\n", err, data))
		//assert.Equal(t, "none", c.APIKey)
		err = c.SessionLogout()
		assert.NoError(t, err, "SessionLogout threw error -> %s", err)
		// test if we can perform an op after logout
		//_, err = c.GetProfileBySN(testSerial)
		//assert.Error(t, err, "SessionLogin threw error -> %s", err)
	} else {
		/*_, c = getTestDriverU("dev")
		data, err := c.SessionLogin()
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
		assert.Equal(t, "none", c.APIKey)
		*/
	}
}

// Get the current idle timeout from the logged in session
func TestGetIdleTimeout(t *testing.T) {
	var (
		c *OVClient
		// d *OVTest
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		err := c.RefreshLogin()
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
		c        *OVClient
		testtime int64
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
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
