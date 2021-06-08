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

func TestCreateIpv4Range(t *testing.T) {
	var (
		d  *OVTest
		c  *ov.OVClient
		id string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		id = d.Tc.GetTestData(d.Env, "Id").(string)

		ipv4Range, err := c.GetIPv4RangebyId("", id)
		assert.NoError(t, err, "GetIPv4RangebyId error -> %s", err)
		fragments := make([]ov.StartStopFragments, 1)
		fragments[0] = ov.StartStopFragments{StartAddress: "10.16.0.10", EndAddress: "10.16.0.100"}
		if ipv4Range.URI.IsNil() {
			ipv4Range := ov.CreateIpv4Range{
				SubnetUri:          utils.Nstring(d.Tc.GetTestData(d.Env, "SubnetUri").(string)),
				StartStopFragments: fragments,
				Name:               "testName",
				Type:               d.Tc.GetTestData(d.Env, "Type").(string),
			}
			_, err := c.CreateIPv4Range(ipv4Range)
			assert.NoError(t, err, "CreateIPv4Range error -> %s", err)

			_, err = c.CreateIPv4Range(ipv4Range)
			assert.Error(t, err, "CreateIPv4Range should error becaue the range already exists, err -> %s", err)
		} else {
			log.Warnf("The Ipv4Range already exists so skipping CreateIPv4Range test for %s", id)
		}

		ipv4Range, err = c.GetIPv4RangebyId("", id)
		assert.NoError(t, err, "GetIPv4RangebyId error -> %s", err)
	}
}

func GetIPv4RangebyId(t *testing.T) {
	var (
		d      *OVTest
		c      *ov.OVClient
		testId string
	)

	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testId = d.Tc.GetTestData(d.Env, "Id").(string)
		testIpv4Range, err := c.GetIPv4RangebyId("", testId)
		assert.NoError(t, err, "GetIPv4RangebyId threw error -> %s, %+v\n", err, testIpv4Range)

		testIpv4Range, err = c.GetIPv4RangebyId("", "bad")
		assert.NoError(t, err, "GetIPv4RangebyId with fake id -> %s", err)
		assert.Equal(t, "", testIpv4Range.Name)

	} else {
		d, c = getTestDriverU("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testId = d.Tc.GetTestData(d.Env, "Id").(string)
		data, err := c.GetIPv4RangebyId("", testId)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetAllocatedFragments(t *testing.T) {
	var (
		d      *OVTest
		c      *ov.OVClient
		testId string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testId = d.Tc.GetTestData(d.Env, "Id").(string)
		allocatedFragments, err := c.GetAllocatedFragments("", "", testId)
		assert.NoError(t, err, "GetAllocatedFragments threw an error -> %s. %+v\n", err, allocatedFragments)

		allocatedFragments, err = c.GetAllocatedFragments("", "", testId)
		assert.NoError(t, err, "GetAllocatedFragments name:asc error -> %s. %+v\n", err, allocatedFragments)

	} else {
		d, c = getTestDriverU("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetAllocatedFragments("", "", testId)
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	d, c = getTestDriverU("test_ipv4_range")
	testId = d.Tc.GetTestData(d.Env, "Id").(string)
	data, err := c.GetAllocatedFragments("", "", testId)
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func TestGetFreeFragments(t *testing.T) {
	var (
		d      *OVTest
		c      *ov.OVClient
		testId string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testId = d.Tc.GetTestData(d.Env, "Id").(string)
		freeFragments, err := c.GetFreeFragments("", "", testId)
		assert.NoError(t, err, "TestGetFreeFragments threw an error -> %s. %+v\n", err, freeFragments)

		freeFragments, err = c.GetFreeFragments("", "", testId)
		assert.NoError(t, err, "TestGetFreeFragments name:asc error -> %s. %+v\n", err, freeFragments)

	} else {
		d, c = getTestDriverU("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetFreeFragments("", "", testId)
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	d, c = getTestDriverU("test_ipv4_range")
	testId = d.Tc.GetTestData(d.Env, "Id").(string)
	data, err := c.GetFreeFragments("", "", testId)
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func TestDeleteIpv4Range(t *testing.T) {
	var (
		d             *OVTest
		c             *ov.OVClient
		id            string
		testIpv4Range ov.Ipv4Range
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_ipv4_range")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		id = d.Tc.GetTestData(d.Env, "Id").(string)

		err := c.DeleteIpv4Range(id)
		assert.NoError(t, err, "DeleteIpv4Range err-> %s", err)

		testIpv4Range, err = c.GetIPv4RangebyId("", id)
		assert.NoError(t, err, "GetIPv4RangebyId with deleted ipv4 Range-> %+v", err)
		assert.Equal(t, "", testIpv4Range.Name, fmt.Sprintf("Problem getting ipv4Range name, %+v", testIpv4Range))
	} else {
		_, c = getTestDriverU("test_ipv4_range")
		err := c.DeleteIpv4Range(id)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testIpv4Range))
	}

}
