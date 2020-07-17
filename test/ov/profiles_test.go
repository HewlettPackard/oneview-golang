package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// CreateProfileFromTemplate(name string, new_template ServerProfile, blade ServerHardware)
// test create profile
func TestCreateProfileFromTemplate(t *testing.T) {
	var (
		d                *OVTest
		c                *ov.OVClient
		testHostName     string
		testBladeSerial  string
		testTemplateName string
		testBlades       ov.ServerHardwareList
		testTemplate     ov.ServerProfile
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testHostName = d.Tc.GetTestData(d.Env, "HostName").(string)
		testBladeSerial = d.Tc.GetTestData(d.Env, "FreeBladeSerialNumber").(string)
		testTemplateName = d.Tc.GetTestData(d.Env, "TemplateProfile").(string)
		if os.Getenv("ONEVIEW_TEST_PROVISION") != "true" {
			log.Info("env ONEVIEW_TEST_PROVISION != true for TestCreateProfileFromTemplate")
			log.Infof("Skipping, create profile for : %s, %s, %s", testHostName, testBladeSerial, testTemplateName)
			return
		}

		testBlades, _ = c.GetServerHardwareList([]string{fmt.Sprintf("serialNumber matches '%s'", testBladeSerial)}, "name:asc", "", "", "")
		assert.True(t, (len(testBlades.Members) > 0), "Did not get any blades from server hardware list")

		testTemplate, _ = c.GetProfileTemplateByName(testTemplateName)
		assert.Equal(t, testTemplateName, testTemplate.Name, fmt.Sprintf("Problem getting template name, %+v", testTemplate))

		// find out if the test profile already exist
		testProfile, err := c.GetProfileByName(testHostName)
		assert.NoError(t, err, "CreateProfileFromTemplate get the server profile error -> %s", err)

		if len(testBlades.Members) > 0 && testProfile.URI.IsNil() {
			err := c.CreateProfileFromTemplate(testHostName, testTemplate, testBlades.Members[0])
			assert.NoError(t, err, "CreateProfileFromTemplate error -> %s", err)

			err = c.CreateProfileFromTemplate(testHostName, testTemplate, testBlades.Members[0])
			assert.Error(t, err, "CreateProfileFromTemplate should error because a template already exist, err-> %s", err)

		} else {
			log.Warnf("The testHostName already exist, so skipping CreateProfileTemplate test for %s", testHostName)
		}

		// reload the test profile that we just created
		testProfile, err = c.GetProfileByName(testHostName)
		assert.NoError(t, err, "CreateProfileFromTemplate get the server profile error -> %s", err)

		// get the server hardware associated with that test profile
		log.Debugf("testProfile -> %+v", testProfile)
		testBlade, err := c.GetServerHardwareByUri(testProfile.ServerHardwareURI)
		assert.NoError(t, err, "CreateProfileFromTemplate call to GetServerHardwareByUri got error -> %s", err)

		// power on the server, and leave it in that state
		var pt *ov.PowerTask
		log.Debugf("testBlade -> %+v", testBlade)
		pt = pt.NewPowerTask(testBlade)
		pt.Timeout = 46 // timeout is 20 sec
		log.Info("------- Setting Power to On")
		err = pt.PowerExecutor(ov.P_ON)
		assert.NoError(t, err, "PowerExecutor threw no errors -> %s", err)
	}

}

// TestSubmitNewProfile functionality
func TestSubmitNewProfile(t *testing.T) {
	var (
		d            *OVTest
		c            *ov.OVClient
		testHostName string
		testProfile  ov.ServerProfile
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		testHostName = d.Tc.GetTestData(d.Env, "ServerProfileName").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		isAvailable, err := c.GetAvailableServers(testProfile.ServerHardwareURI.String())
		assert.NoError(t, err, "GetAvailableServers get the server hardware error -> %s", err)
		assert.Equal(t, "", isAvailable, fmt.Sprintf("Is given hardware available: %+v", isAvailable))

		testServerHardware, err := c.GetServerHardwareByUri(testProfile.ServerHardwareURI)
		assert.NoError(t, err, "SubmitNewProfile call to GetServerHardwareByUri got error -> %s", err)

		testProfile, err := c.GetProfileByName(testHostName)
		assert.NoError(t, err, "GetProfileByName with created profile -> %+v", err)
		assert.Equal(t, "", testProfile.Name, fmt.Sprintf("Problem getting profile name, %+v", testHostName))

		var pt *ov.PowerTask
		pt = pt.NewPowerTask(testServerHardware)
		pt.Timeout = 46 // timeout is 20 sec
		log.Info("------- Setting Power to off")
		err = pt.PowerExecutor(ov.P_OFF)
		assert.NoError(t, err, "PowerExecutor threw no errors -> %s", err)

	} else {
		_, c = getTestDriverU("dev")
		err := c.DeleteProfile("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testProfile))
	}

}

// find Server_Profile_scs79
func TestGetProfileByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testname string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		testname = d.Tc.GetTestData(d.Env, "ServerProfileName").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileByName(testname)
		assert.NoError(t, err, "GetProfileByName threw error -> %s", err)
		assert.Equal(t, testname, data.Name)

		data, err = c.GetProfileByName("foo")
		assert.NoError(t, err, "GetProfileByName with fake name -> %+v", err)
		assert.Equal(t, "", data.Name)

	} else {
		d, c = getTestDriverU("dev")
		testname = d.Tc.GetTestData(d.Env, "ServerProfileName").(string)
		data, err := c.GetProfileByName(testname)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

// TestGetProfileConnectionByName verify functionality
func TestGetConnectionByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testname string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		testname = d.Tc.GetTestData(d.Env, "ServerProfileName").(string)
		pubcname := d.Tc.GetTestData(d.Env, "FreePublicConnection").(string)
		expectsmac := d.Tc.GetExpectsData(d.Env, "MACAddress").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		profile, err := c.GetProfileByName(testname)
		assert.NoError(t, err, "GetProfileByName threw error -> %s", err)

		for _, c := range profile.Connections {
			log.Debugf("connection -> %d %s %s %s", c.ID, c.Name, c.MAC, c.MacType)
		}

		connection, err := profile.GetConnectionByName(pubcname)
		log.Debugf("Got connection -> %+v", connection)
		assert.NoError(t, err, "GetConnectionByName threw error -> %s", err)
		assert.Equal(t, expectsmac, connection.MAC.String(), "GetConnectionByName failed to get connection %+v", err, connection)
	}
}

// TestGetProfileNameBySN
// Acceptance test ->
// /rest/server-profiles
// ?filter=serialNumber matches '2M25090RMW'&sort=name:asc
func TestGetProfileBySN(t *testing.T) {
	var (
		d          *OVTest
		c          *ov.OVClient
		testSerial string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		testSerial = d.Tc.GetTestData(d.Env, "SerialNumber").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfileBySN(testSerial)
		assert.NoError(t, err, "GetProfileBySN threw error -> %s", err)
		assert.Equal(t, testSerial, data.SerialNumber.String())

		data, err = c.GetProfileBySN("foo")
		assert.NoError(t, err, "GetProfileBySN with fake serial number -> %+v", err)
		assert.Equal(t, "null", data.SerialNumber.String())

	} else {
		d, c = getTestDriverU("dev")
		testSerial = d.Tc.GetTestData(d.Env, "SerialNumber").(string)
		data, err := c.GetProfileBySN(testSerial)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

// TestGetProfiles
func TestGetProfiles(t *testing.T) {
	var (
		// d *OVTest
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetProfiles("", "", "", "", "")
		assert.NoError(t, err, "GetProfiles threw error -> %s, %+v\n", err, data)

		data, err = c.GetProfiles("", "", "", "name:asc", "")
		assert.NoError(t, err, "GetProfiles name:asc error -> %s, %+v", err, data)

	} else {
		_, c = getTestDriverU("dev")
		data, err := c.GetProfiles("", "", "", "", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

// test for not found profile
// should not delete a profile that doesn't exist
func TestDeleteProfileNotFound(t *testing.T) {
	var (
		c               *ov.OVClient
		testProfileName = "fake_profile_doesnt_exist"
		testProfile     ov.ServerProfile
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteProfile(testProfileName)
		assert.NoError(t, err, "DeleteProfile err-> %s", err)

		testProfile, err = c.GetProfileByName(testProfileName)
		assert.NoError(t, err, "GetProfileByName with deleted profile -> %+v", err)
		assert.Equal(t, "", testProfile.Name, fmt.Sprintf("Problem getting template name, %+v", testProfile))
	} else {
		_, c = getTestDriverU("dev")
		err := c.DeleteProfile(testProfileName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testProfile))
	}
}

// test DeleteProfile
func TestDeleteProfile(t *testing.T) {
	var (
		d               *OVTest
		c               *ov.OVClient
		testProfileName string
		testProfile     ov.ServerProfile
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testProfileName = d.Tc.GetTestData(d.Env, "HostName").(string)

		err := c.DeleteProfile(testProfileName)
		assert.NoError(t, err, "DeleteProfile err-> %s", err)

		testProfile, err = c.GetProfileByName(testProfileName)
		assert.NoError(t, err, "GetProfileByName with deleted profile -> %+v", err)
		assert.Equal(t, "", testProfile.Name, fmt.Sprintf("Problem getting template name, %+v", testProfile))
	} else {
		_, c = getTestDriverU("dev")
		err := c.DeleteProfile("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testProfile))
	}

}
