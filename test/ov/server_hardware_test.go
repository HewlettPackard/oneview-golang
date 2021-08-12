package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// get server hardware test
func TestServerHardware(t *testing.T) {
	var (
		d           *OVTest
		c           *ov.OVClient
		testData    utils.Nstring
		expectsData string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		testData = utils.Nstring(d.Tc.GetTestData(d.Env, "ServerHardwareURI").(string))
		expectsData = d.Tc.GetExpectsData(d.Env, "SerialNumber").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetServerHardwareByUri(testData)
		log.Debugf("%+v", data)
		assert.NoError(t, err, "GetServerHardwareByUri threw error -> %s", err)
		// fmt.Printf("data.Connections -> %+v\n", data)
		assert.Equal(t, expectsData, data.SerialNumber.String())

	}
}

func TestGetServerHardwareByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_server_hardware")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testServerHardware, err := c.GetServerHardwareByName(testName)
		assert.NoError(t, err, "GetServerHardwareByName thew an error -> %s", err)
		assert.Equal(t, testName, testServerHardware.Name)

		testServerHardware, err = c.GetServerHardwareByName("bad")
		assert.NoError(t, err, "GetServerHardwareByName with fake name -> %s", err)
		assert.Equal(t, "", testServerHardware.Name)

	} else {
		d, c = getTestDriverU("test_server_hardware")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetServerHardwareByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

// get server hardware test
func TestGetAvailableHardware(t *testing.T) {
	var (
		d               *OVTest
		c               *ov.OVClient
		testHwType_URI  utils.Nstring
		testHWGroup_URI utils.Nstring
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		testHwType_URI = utils.Nstring(d.Tc.GetTestData(d.Env, "HardwareTypeURI").(string))
		testHWGroup_URI = utils.Nstring(d.Tc.GetTestData(d.Env, "GroupURI").(string))
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAvailableHardware(testHwType_URI, testHWGroup_URI)
		assert.NoError(t, err, "GetAvailableHardware threw error -> %s", err)
		// fmt.Printf("data.Connections -> %+v\n", data)
		log.Debugf("Abailable server -> %+v", data)
		log.Infof("Server Name -> %+v", data.Name)
		assert.NotEqual(t, "", data.Name)

	}
}

// TestGetIloIPAddress verify get ip address for hardware
func TestGetIloIPAddress(t *testing.T) {

	var (
		d           *OVTest
		c           *ov.OVClient
		testData    utils.Nstring
		expectsData string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("dev")
		testData = utils.Nstring(d.Tc.GetTestData(d.Env, "ServerHardwareURI").(string))
		expectsData = d.Tc.GetExpectsData(d.Env, "IloIPAddress").(string)
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		s, err := c.GetServerHardwareByUri(testData)
		assert.NoError(t, err, "GetServerHardwareByUri threw error -> %s", err)
		ip := s.GetIloIPAddress()
		log.Debugf("server -> %+v", s)
		log.Debugf("ip -> %+v", ip)
		assert.Equal(t, expectsData, ip)

	}

}

// TestAddRackServer verify adds hardware to the appliance
func TestAddRackServers(t *testing.T) {

	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_server_hardware")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test server hardware already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testSh, err := c.GetServerHardwareByName(testName)
		assert.NoError(t, err, "Create Server Hardware error -> %s", err)

		if testSh.URI.IsNil() {
			testSh = ov.ServerHardware{
				Hostname:           d.Tc.GetTestData(d.Env, "Hostname").(string),
				Username:           d.Tc.GetTestData(d.Env, "Username").(string),
				Password:           d.Tc.GetTestData(d.Env, "Password").(string),
				ConfigurationState: d.Tc.GetTestData(d.Env, "ConfigurationState").(string),
				LicensingIntent:    d.Tc.GetTestData(d.Env, "LicensingIntent").(string),
			}
			err := c.AddRackServer(testSh)
			assert.NoError(t, err, "Add server hardware error -> %s", err)

			err = c.AddRackServer(testSh)
			assert.Error(t, err, "AddRackServer should give error because hardware already exists, err-> %s", err)

		} else {
			log.Warnf("already exist, so skipping AddRackServer test for %s", testName)
		}

		// reload the test profile that we just created
		testSh, err = c.GetServerHardwareByName(testName)
		assert.NoError(t, err, "GetServerHardware error -> %s", err)
	}

}

// TestAddMultipleRackServers verify adds multiple hardwares
func TestAddMultipleRackServers(t *testing.T) {

	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_server_hardware")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test server hardware already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testSh := ov.ServerHardware{
			MpHostsAndRanges:   d.Tc.GetTestData(d.Env, "MpHostsAndRanges"),
			Username:           d.Tc.GetTestData(d.Env, "Username").(string),
			Password:           d.Tc.GetTestData(d.Env, "Password").(string),
			ConfigurationState: d.Tc.GetTestData(d.Env, "ConfigurationState").(string),
			LicensingIntent:    d.Tc.GetTestData(d.Env, "LicensingIntent").(string),
		}
		err := c.AddMultipleRackServers(testSh)
		assert.NoError(t, err, "Add server hardware error -> %s", err)

		err = c.AddMultipleRackServers(testSh)
		assert.Error(t, err, "AddRackServer should give error because hardware already exists, err-> %s", err)

		// reload the test profile that we just created
		testSh, err = c.GetServerHardwareByName(testName)
		assert.NoError(t, err, "GetServerHardware error -> %s", err)
	}

}
