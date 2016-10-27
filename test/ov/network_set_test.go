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

func TestCreateNetworkSet(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_network_set")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test network set already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testNetSet, err := c.GetNetworkSetByName(testName)
		assert.NoError(t, err, "CreateNetworkSet get the NetworkSet error -> %s", err)

		networkUris := make([]utils.Nstring, 0)

		if testNetSet.URI.IsNil() {
			testNetSet = ov.NetworkSet{
				Name:        testName,
				Type:        d.Tc.GetTestData(d.Env, "Type").(string),
				NetworkUris: networkUris,
			}
			err := c.CreateNetworkSet(testNetSet)
			assert.NoError(t, err, "CreateNetworkSet error -> %s", err)

			err = c.CreateNetworkSet(testNetSet)
			assert.Error(t, err, "CreateNetworkSet should error because the Network Set already exists, err-> %s", err)

		} else {
			log.Warnf("The networkSet already exist, so skipping CreateNetworkSet test for %s", testName)
		}

		// reload the test profile that we just created
		testNetSet, err = c.GetNetworkSetByName(testName)
		assert.NoError(t, err, "CreateNetworkSet get the server profile error -> %s", err)
	}
}

func TestGetNetworkSetByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_network_set")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testNetSet, err := c.GetNetworkSetByName(testName)
		assert.NoError(t, err, "GetNetworkSetByName thew an error -> %s", err)
		assert.Equal(t, testName, testNetSet.Name)

		testNetSet, err = c.GetNetworkSetByName("bad")
		assert.NoError(t, err, "GetNetworkSetByName with fake name -> %s", err)
		assert.Equal(t, "", testNetSet.Name)

	} else {
		d, c = getTestDriverU("test_network_set")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetNetworkSetByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}

}

func TestGetNetworkSets(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_network_set")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		networkSets, err := c.GetNetworkSets("", "")
		assert.NoError(t, err, "GetNetworkSets threw error -> %s, %+v\n", err, networkSets)

		networkSets, err = c.GetNetworkSets("", "name:asc")
		assert.NoError(t, err, "GetNetworkSets name:asc error -> %s, %+v\n", err, networkSets)

	} else {
		_, c = getTestDriverU("test_network_set")
		data, err := c.GetNetworkSets("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteNetworkSetNotFound(t *testing.T) {
	var (
		c          *ov.OVClient
		testName   = "fake"
		testNetSet ov.NetworkSet
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_network_set")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteNetworkSet(testName)
		assert.NoError(t, err, "DeleteNetworkSet err-> %s", err)

		testNetSet, err = c.GetNetworkSetByName(testName)
		assert.NoError(t, err, "GetNetworkSetByName with deleted network set -> %+v", err)
		assert.Equal(t, "", testNetSet.Name, fmt.Sprintf("Problem getting network set wtih name, %+v", testNetSet))
	} else {
		_, c = getTestDriverU("test_network_set")
		err := c.DeleteNetworkSet(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testNetSet))
	}
}

func TestDeleteNetworkSet(t *testing.T) {
	var (
		d          *OVTest
		c          *ov.OVClient
		testName   string
		testNetSet ov.NetworkSet
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_network_set")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteNetworkSet(testName)
		assert.NoError(t, err, "DeleteNetworkSet err-> %s", err)

		testNetSet, err = c.GetNetworkSetByName(testName)
		assert.NoError(t, err, "GetNetworkSetByName with deleted network set-> %+v", err)
		assert.Equal(t, "", testNetSet.Name, fmt.Sprintf("Problem getting netset name, %+v", testNetSet))
	} else {
		_, c = getTestDriverU("test_network_set")
		err := c.DeleteNetworkSet("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testNetSet))
	}

}
