package ov
import (
	"os"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// Test SessionLogin
func TestSessionLogin(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.SessionLogin()
		// fmt.Printf("after SessionLogin: %s -> (err) %s", data.ID, err)
		assert.NoError(t, err, "SessionLogin threw error -> %s", err)
		assert.NotEmpty(t, data.ID, fmt.Sprintf("SessionLogin is empty! something went wrong, err -> %s, data -> %+v\n", err, data) )
		assert.Equal(t, "none", c.APIKey)
	} else {
		c = getTestDriverU()
		data, err := c.SessionLogin()
		assert.Error(t,err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n",err, data))
		assert.Equal(t, "none", c.APIKey)
	}
}
