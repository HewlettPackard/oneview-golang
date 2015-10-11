package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/drivers/oneview/liboneview"
	"github.com/docker/machine/log"
	"github.com/stretchr/testify/assert"
)

// UnlessNoProfileTemplate - determine if current test version is supported
func UnlessNoProfileTemplate(ovversion int) bool {
	var currentversion liboneview.Version
	var asc liboneview.APISupport
	currentversion = currentversion.CalculateVersion(ovversion, 108) // hard coded icsp version for testing
	asc = asc.NewByName("profile_templates.go")
	if !asc.IsSupported(currentversion) {
		log.Infof("skipping testing for GetProfileTemplateByName, client version not supported: %+v", currentversion)
		return true
	}
	return false
}

// GetProfileTemplateByName get a template profile
func TestGetProfileTemplateByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testname string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		// determine if we should execute
		if UnlessNoProfileTemplate(c.APIVersion) {
			return
		}

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
		d, c = getTestDriverU()
		// determine if we should execute
		if UnlessNoProfileTemplate(c.APIVersion) {
			return
		}

		testname = d.Tc.GetTestData(d.Env, "ServerProfileName").(string)
		data, err := c.GetProfileTemplateByName(testname)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
