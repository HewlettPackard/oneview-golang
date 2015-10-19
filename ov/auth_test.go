package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/log"
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
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.SessionLogin()
		// fmt.Printf("after SessionLogin: %s -> (err) %s", data.ID, err)
		assert.NoError(t, err, "SessionLogin threw error -> %s", err)
		assert.NotEmpty(t, data.ID, fmt.Sprintf("SessionLogin is empty! something went wrong, err -> %s, data -> %+v\n", err, data))
		assert.Equal(t, "none", c.APIKey)
	} else {
		_, c = getTestDriverU()
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
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		//data, err := c.SessionLogin()
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
		/*_, c = getTestDriverU()
		data, err := c.SessionLogin()
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
		assert.Equal(t, "none", c.APIKey)
		*/
	}
}
