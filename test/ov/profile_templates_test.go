package ov

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateProfileTemplate(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_server_profile_template")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test Server Profile Template already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testServerProfileTemplate, err := c.GetProfileTemplateByName(testName)
		assert.NoError(t, err, "CreateServerProfileTemplate get the ServerProfileTemplate error -> %s", err)

		if testServerProfileTemplate.URI.IsNil() {
			testServerProfileTemplate = ov.ServerProfile{
				Name: testName,
				Type: d.Tc.GetTestData(d.Env, "Type").(string),
				ServerHardwareTypeURI: utils.NewNstring(d.Tc.GetTestData(d.Env, "ServerHardwareTypeUri").(string)),
				EnclosureGroupURI:     utils.NewNstring(d.Tc.GetTestData(d.Env, "EnclosureGroupUri").(string)),
			}

			err := c.CreateProfileTemplate(testServerProfileTemplate)
			assert.NoError(t, err, "CreateServerProfileTemplate error -> %s", err)

			//err = c.CreateProfileTemplate(testServerProfileTemplate)
			//assert.Error(t, err, "CreateServerProfileTemplate should error because the Server Profile Template already exists, err-> %s", err)

		} else {
			log.Warnf("The serverProfileTemplate already exist, so skipping CreateServerProfileTemplate test for %s", testName)
		}

		// reload the test profile that we just created
		testServerProfileTemplate, err = c.GetProfileTemplateByName(testName)
		assert.NoError(t, err, "GetServerProfileTemplate error -> %s", err)
	}
}

// GetProfileTemplateByName get a template profile
func TestGetProfileTemplateByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testname string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_server_profile_template")
		// determine if we should execute
		testname = d.Tc.GetTestData(d.Env, "Name").(string)
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
		d, c = getTestDriverU("test_server_profile_template")
		// determine if we should execute
		if c.ProfileTemplatesNotSupported() {
			return
		}

		testname = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetProfileTemplateByName(testname)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteProfileTemplateNotFound(t *testing.T) {
	var (
		c                         *ov.OVClient
		testName                  = "fake"
		testServerProfileTemplate ov.ServerProfile
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_server_profile_template")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteProfileTemplate(testName)
		assert.NoError(t, err, "DeleteServerProfileTemplate err-> %s", err)

		testServerProfileTemplate, err = c.GetProfileTemplateByName(testName)
		assert.NoError(t, err, "GetServerProfileTemplateByName with deleted serverProfileTemplate -> %+v", err)
		assert.Equal(t, "", testServerProfileTemplate.Name, fmt.Sprintf("Problem getting serverProfileTemplate name, %+v", testServerProfileTemplate))
	} else {
		_, c = getTestDriverU("test_server_profile_template")
		err := c.DeleteProfileTemplate(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testServerProfileTemplate))
	}
}

func TestDeleteProfileTemplate(t *testing.T) {
	var (
		d                         *OVTest
		c                         *ov.OVClient
		testName                  string
		testServerProfileTemplate ov.ServerProfile
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_server_profile_template")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteProfileTemplate(testName)
		assert.NoError(t, err, "DeleteServerProfileTemplate err-> %s", err)

		testServerProfileTemplate, err = c.GetProfileTemplateByName(testName)
		assert.NoError(t, err, "GetServerProfileTemplateByName with deleted serverProfileTemplate-> %+v", err)
		assert.Equal(t, "", testServerProfileTemplate.Name, fmt.Sprintf("Problem getting serverProfileTemplate name, %+v", testServerProfileTemplate))
	} else {
		_, c = getTestDriverU("test_server_profile_template")
		err := c.DeleteProfileTemplate("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testServerProfileTemplate))
	}

}
