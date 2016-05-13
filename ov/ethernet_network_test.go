package ov

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEthernetNetworkByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testname string
	)
	d, c = getTestDriverU("test_ethernet_network")
	testname = d.Tc.GetTestData(d.Env, "Name").(string)
	data, err := c.GetProfileByName(testname)
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func TestGetEthernetNetworks(t *testing.T) {
	var (
		c *OVClient
	)
	_, c = getTestDriverU("test_ethernet_network")
	data, err := c.GetProfiles("", "")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func TestDeleteEthernetNetwork(t *testing.T) {
	var (
		c           *OVClient
		testProfile ServerProfile
	)
	_, c = getTestDriverU("test_ethernet_network")
	err := c.DeleteProfile("footest")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testProfile))
}
