package ov

import (
	"fmt"
	"os"
	"testing"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateLogicalFCNetwork(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
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
			fcNetwork := ov.FCNetwork{
				AutoLoginRedistribution: falseVar,
				Description:             "Test FC Network",
				LinkStabilityTime:       30,
				FabricType:              d.Tc.GetTestData(d.Env, "FabricType").(string),
				Name:                    testName,
				Type:                    d.Tc.GetTestData(d.Env, "Type").(string),
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
		c        *ov.OVClient
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
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		fcNetworks, err := c.GetFCNetworks("", "","","")
		assert.NoError(t, err, "GetFCNetworks threw an error -> %s. %+v\n", err, fcNetworks)

		fcNetworks, err = c.GetFCNetworks("", "name:asc", "", "")
		assert.NoError(t, err, "GetFCNetworks name:asc error -> %s. %+v\n", err, fcNetworks)

	} else {
		_, c = getTestDriverU("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetFCNetworks("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	_, c = getTestDriverU("test_fc_network")
	data, err := c.GetFCNetworks("", "", "", "")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func TestDeleteFCNetwork(t *testing.T) {
	var (
		d         *OVTest
		c         *ov.OVClient
		testName  string
		testFcNet ov.FCNetwork
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_fc_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteFCNetwork(testName)
		assert.NoError(t, err, "DeleteFcNetwork err-> %s", err)

		testFcNet, err = c.GetFCNetworkByName(testName)
		assert.NoError(t, err, "GetFcNetworkByName with deleted fc network-> %+v", err)
		assert.Equal(t, "", testFcNet.Name, fmt.Sprintf("Problem getting fcnet name, %+v", testFcNet))
	} else {
		_, c = getTestDriverU("test_fc_network")
		err := c.DeleteFCNetwork("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testFcNet))
	}

}
