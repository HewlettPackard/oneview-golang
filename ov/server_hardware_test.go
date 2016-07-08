package ov

import (
	"os"
	"testing"

	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// get server hardware test
func TestServerHardware(t *testing.T) {
	var (
		d           *OVTest
		c           *OVClient
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
		data, err := c.GetServerHardware(testData)
		log.Debugf("%+v", data)
		assert.NoError(t, err, "GetServerHardware threw error -> %s", err)
		// fmt.Printf("data.Connections -> %+v\n", data)
		assert.Equal(t, expectsData, data.SerialNumber.String())

	}
}

// get server hardware test
func TestGetAvailableHardware(t *testing.T) {
	var (
		d               *OVTest
		c               *OVClient
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
		c           *OVClient
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
		s, err := c.GetServerHardware(testData)
		ip := s.GetIloIPAddress()
		log.Debugf("server -> %+v", s)
		log.Debugf("ip -> %+v", ip)
		assert.NoError(t, err, "GetServerHardware threw error -> %s", err)
		assert.Equal(t, expectsData, ip)

	}

}
