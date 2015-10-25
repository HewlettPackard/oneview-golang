package icsp

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// TestGetJobs
func TestGetJobs(t *testing.T) {
	var (
		// d *OVTest
		c *ICSPClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetJobs()
		assert.NoError(t, err, "GetJobs threw error -> %s, %+v\n", err, data)
		if data.Total > 0 {
			log.Debugf("data -> %+v", data.Members[0])
			assert.Condition(t, func() bool { return len(data.Members[0].URI) > 0 }, "has no uri content")
			data, err := c.GetJob(ODSUri{URI: data.Members[0].URI})
			assert.NoError(t, err, "GetJob threw error -> %s, %+v\n", err, data)
		}
	} else {
		_, c = getTestDriverU()
		data, err := c.GetJobs()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
