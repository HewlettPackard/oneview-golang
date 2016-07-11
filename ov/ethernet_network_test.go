package ov

import (
	"fmt"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateEthernetNetwork(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ethernet_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test ethernet network already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testEthNet, err := c.GetEthernetNetworkByName(testName)
		assert.NoError(t, err, "CreateEthernetNetwork get the EthernetNetwork error -> %s", err)

		if testEthNet.URI.IsNil() {
			testEthNet = EthernetNetwork{
				Name:                testName,
				VlanId:              7,
				Purpose:             d.Tc.GetTestData(d.Env, "Purpose").(string),
				SmartLink:           d.Tc.GetTestData(d.Env, "SmartLink").(bool),
				PrivateNetwork:      d.Tc.GetTestData(d.Env, "PrivateNetwork").(bool),
				EthernetNetworkType: d.Tc.GetTestData(d.Env, "EthernetNetworkType").(string),
				Type:                d.Tc.GetTestData(d.Env, "Type").(string),
			}
			err := c.CreateEthernetNetwork(testEthNet)
			assert.NoError(t, err, "CreateEthernetNetwork error -> %s", err)

			err = c.CreateEthernetNetwork(testEthNet)
			assert.Error(t, err, "CreateEthernetNetwork should error because the EthernetNetwork already exists, err-> %s", err)

		} else {
			log.Warnf("The ethernetNetwork already exist, so skipping CreateEthernetNetwork test for %s", testName)
		}

		// reload the test profile that we just created
		testEthNet, err = c.GetEthernetNetworkByName(testName)
		assert.NoError(t, err, "GetEthernetNetwork error -> %s", err)
	}

}

func TestGetEthernetNetworkByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ethernet_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testEthNet, err := c.GetEthernetNetworkByName(testName)
		assert.NoError(t, err, "GetEthernetNetworkByName thew an error -> %s", err)
		assert.Equal(t, testName, testEthNet.Name)

		testEthNet, err = c.GetEthernetNetworkByName("bad")
		assert.NoError(t, err, "GetEthernetNetworkByName with fake name -> %s", err)
		assert.Equal(t, "", testEthNet.Name)

	} else {
		d, c = getTestDriverU("test_ethernet_network")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetEthernetNetworkByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetEthernetNetworks(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_ethernet_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		ethernetNetworks, err := c.GetEthernetNetworks("", "")
		assert.NoError(t, err, "GetEthernetNetworks threw error -> %s, %+v\n", err, ethernetNetworks)

		ethernetNetworks, err = c.GetEthernetNetworks("", "name:asc")
		assert.NoError(t, err, "GetEthernetNetworks name:asc error -> %s, %+v\n", err, ethernetNetworks)

	} else {
		_, c = getTestDriverU("test_ethernet_network")
		data, err := c.GetEthernetNetworks("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteEthernetNetworkNotFound(t *testing.T) {
	var (
		c          *OVClient
		testName   = "fake"
		testEthNet EthernetNetwork
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_ethernet_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteEthernetNetwork(testName)
		assert.NoError(t, err, "DeleteEthernetNetwork err-> %s", err)

		testEthNet, err = c.GetEthernetNetworkByName(testName)
		assert.NoError(t, err, "GetEthernetNetworkByName with deleted ethernet network -> %+v", err)
		assert.Equal(t, "", testEthNet.Name, fmt.Sprintf("Problem getting ethernet name, %+v", testEthNet))
	} else {
		_, c = getTestDriverU("test_ethernet_network")
		err := c.DeleteEthernetNetwork(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testEthNet))
	}
}

func TestDeleteEthernetNetwork(t *testing.T) {
	var (
		d          *OVTest
		c          *OVClient
		testName   string
		testEthNet EthernetNetwork
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ethernet_network")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteEthernetNetwork(testName)
		assert.NoError(t, err, "DeleteEthernetNetwork err-> %s", err)

		testEthNet, err = c.GetEthernetNetworkByName(testName)
		assert.NoError(t, err, "GetEthernetNetworkByName with deleted ethernet network-> %+v", err)
		assert.Equal(t, "", testEthNet.Name, fmt.Sprintf("Problem getting ethnet name, %+v", testEthNet))
	} else {
		_, c = getTestDriverU("test_ethernet_network")
		err := c.DeleteEthernetNetwork("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testEthNet))
	}

}
