package ov

import (
	"fmt"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateFCoENetwork(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_fcoe_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		//find out if the test fcoe network already exists
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testFcoeNet, err := c.GetFCoENetworkByName(testName)
		assert.NoError(t, err, "CreateFCoENetwork get the FCoENetwork error --> %s", err)

		if testFcoeNet.URI.IsNil() {
			testFcoeNet = FCoENetwork{
				Name:   testName,
				VlanId: 143,
				Type:   d.Tc.GetTestData(d.Env, "Type").(string),
			}
			err := c.CreateFCoENetwork(testFcoeNet)
			assert.NoError(t, err, "CreateFCoENetwork error --> %s", err)

			err = c.CreateFCoENetwork(testFcoeNet)
			assert.Error(t, err, "CreateFCoENetwork should error because FCoENetwork already exists, err --> %s ", err)

		} else {
			log.Warnf("The fcoeNetwork already exists, so skipping CreateFCoENetwork test for %s", testName)
		}
		//reload the profile we just created
		testFcoeNet, err = c.GetFCoENetworkByName(testName)
		assert.NoError(t, err, "GetFCoENetworks error --> %s", err)
	}
}

func TestGetFCoENetworkByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_fcoe_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testFcoeNet, err := c.GetFCoENetworkByName(testName)
		assert.NoError(t, err, "GetFCoENetworkByName threw an error --> %s", err)
		assert.Equal(t, testName, testFcoeNet.Name)

		testFcoeNet, err = c.GetFCoENetworkByName("bad")
		assert.NoError(t, err, "GetFCoENetworkByName with fake name --> %s", err)
		assert.Equal(t, "", testFcoeNet.Name)
	} else {

		d, c = getTestDriverU("test_fcoe_network")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetFCoENetworkByName(testName)

		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetFCoENetworks(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {

		_, c = getTestDriverA("test_fcoe_network")
		if c == nil {
			t.Fatalf("Failed to execute GetTestDrive() ")
		}
		fcoeNets, err := c.GetFCoENetworks("", "")
		assert.NoError(t, err, "GetFCoENetworks threw error -->  %s, %+v\n", fcoeNets)

		fcoeNets, err = c.GetFCoENetworks("", "name:asc")
		assert.NoError(t, err, "GetFCoENetworks name:asc error --> %s, %+v\n", err, fcoeNets)

	} else {
		_, c = getTestDriverU("test_fcoe_network")
		data, err := c.GetFCoENetworks("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteFCoENetworkNotFound(t *testing.T) {
	var (
		c           *OVClient
		testName    = "fake"
		testFcoeNet FCoENetwork
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_fcoe_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteFCoENetwork(testName)
		assert.NoError(t, err, "DeleteFCoENetwork err-> %s", err)

		testFcoeNet, err = c.GetFCoENetworkByName(testName)
		assert.NoError(t, err, "GetFCoENetworkByName with deleted fcoe network -> %+v", err)
		assert.Equal(t, "", testFcoeNet.Name, fmt.Sprintf("Problem getting fcoe name, %+v", testFcoeNet))
	} else {
		_, c = getTestDriverU("test_fcoe_network")
		err := c.DeleteFCoENetwork(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testFcoeNet))
	}
}

func TestDeleteFCoENetwork(t *testing.T) {
	var (
		d           *OVTest
		c           *OVClient
		testName    string
		testFcoeNet FCoENetwork
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_fcoe_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteFCoENetwork(testName)
		assert.NoError(t, err, "DeleteFCoENetwork err-> %s", err)

		testFcoeNet, err = c.GetFCoENetworkByName(testName)
		assert.NoError(t, err, "GetFCoENetworkByName with deleted fcoe network-> %+v", err)
		assert.Equal(t, "", testFcoeNet.Name, fmt.Sprintf("Problem getting fcoe net name, %+v", testFcoeNet))
	} else {
		_, c = getTestDriverU("test_fcoe_network")
		err := c.DeleteFCoENetwork("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testFcoeNet))
	}

}
