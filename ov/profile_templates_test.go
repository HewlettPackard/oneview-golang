package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// GetProfileTemplateByName get a template profile
func TestGetProfileTemplateByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testname string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		// determine if we should execute
		testname = d.Tc.GetTestData(d.Env, "TemplateProfile").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileTemplateByName(testname)
		assert.NoError(t, err, "GetProfileTemplateByName threw error -> %s", err)
		assert.Equal(t, testname, data.Name)

		data, err = c.GetProfileTemplateByName("foo")
		assert.NoError(t, err, "GetProfileTemplateByName with fake name -> %+v", err)
		assert.Equal(t, "", data.Name)

	} else {
		d, c = getTestDriverU("dev")
		// determine if we should execute
		if c.ProfileTemplatesNotSupported() {
			return
		}

		testname = d.Tc.GetTestData(d.Env, "ServerProfileName").(string)
		data, err := c.GetProfileTemplateByName(testname)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
