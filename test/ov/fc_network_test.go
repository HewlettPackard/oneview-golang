package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateLogicalFCNetwork(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		fcNetwork, err := c.GetFCNetworkByName(testName)
		assert.NoError(t, err, "getFcNetworkByName error -> %s", err)

		if fcNetwork.URI.IsNil() {
			falseVar := false
			fcNetwork := FCNetwork{
				AutoLoginRedistribution: falseVar,
				Description:             "Test FC Network",
				LinkStabilityTime:       30,
				FabricType:              "fabrictype",
				Name:                    testName,
				ConnectionTemplateUri: "http://auri.com/template",
			}
			err := c.CreateFCNetwork(fcNetwork)
			assert.NoError(t, err, "CreateFCNetwork error -> %s", err)

			err = c.CreateFCNetwork(fcNetwork)
			assert.Error(t, err, "CreateFCNetwork should error becaue the network already exists, err -> %s", err)
		} else {
			log.Warnf("The FCNetwork already exists so skipping CreateFCNetwork test for %s", testName)
		}

		//Reload the the test profile
		fcNetwork, err = c.GetFCNetworkByName(testName)
		assert.NoError(t, err, "GetFCNetwork error -> %s", err)
	}
}

func TestGetFCNetworkByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)

	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		testFcNetwork, err := c.GetFCNetworkByName(testName)
		assert.NoError(t, err, "GetFCNetworkByName threw error -> %s, %+v\n", err, testFcNetwork)
		assert.Equal(t, testName, testFcNetwork.Name)

		testFcNetwork, err = c.GetFCNetworkByName("bad")
		assert.NoError(t, err, "GetFCNetworkByName with fake name -> %s", err)
		assert.Equal(t, "", testFcNetwork.Name)

	} else {
		d, c = getTestDriverU("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetFCNetworkByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetFCNetworks(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		fcNetworks, err := c.GetFCNetworks("", "")
		assert.NoError(t, err, "GetFCNetworks threw an error -> %s. %+v\n", err, fcNetworks)

		fcNetworks, err = c.GetFCNetworks("", "name:asc")
		assert.NoError(t, err, "GetFCNetworks name:asc error -> %s. %+v\n", err, fcNetworks)

	} else {
		_, c = getTestDriverU("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetFCNetworks("", "")
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	_, c = getTestDriverU("test_fc_network")
	data, err := c.GetProfiles("", "")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func TestDeleteFCNetwork(t *testing.T) {
	var (
		c           *OVClient
		testProfile ServerProfile
	)
	_, c = getTestDriverU("test_fc_network")
	err := c.DeleteProfile("footest")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testProfile))
}
