package ov

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSNMPv3User(t *testing.T) {
	var (
		d      *OVTest
		c      *ov.OVClient
		testId string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_snmp_v3_user")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test snmpv3 ser already exist
		testId = d.Tc.GetTestData(d.Env, "Id").(string)

		testSnmpv3User, err := c.GetSNMPv3UserById(testId)
		assert.NoError(t, err, "CreateSNMPv3User get the SNMPv3User error -> %s", err)

		if testSnmpv3User.URI.IsNil() {
			testSnmpv3User = ov.SNMPv3User{
				UserName:                 "user7",
				SecurityLevel:            "Authentication and privacy",
				AuthenticationProtocol:   "SHA1",
				AuthenticationPassphrase: "authPass",
				PrivacyProtocol:          "AES-128",
				PrivacyPassphrase:        "1234567812345678",
			}

			_, err := c.CreateSNMPv3Users(testSnmpv3User)
			assert.NoError(t, err, "CreateSNMPv3User error -> %s", err)

			_, err = c.CreateSNMPv3Users(testSnmpv3User)
			assert.Error(t, err, "CreateSNMPv3User should error because the SNMPv3User already exists, err-> %s", err)

		} else {
			log.Warnf("The SNMPv3User already exist, so skipping CreateSNMPv3User test for %s", testId)
		}

		// reload the test profile that we just created
		testSnmpv3User, err = c.GetSNMPv3UserById(testId)
		assert.NoError(t, err, "GetSNMPv3User error -> %s", err)
	}

}
func TestGetSNMPv3UserById(t *testing.T) {
	var (
		d      *OVTest
		c      *ov.OVClient
		testId string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_snmp_v3_user")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testId = d.Tc.GetTestData(d.Env, "Id").(string)

		testSnmpv3User, err := c.GetSNMPv3UserById(testId)
		assert.NoError(t, err, "GetSNMPv3UserById threw an error -> %s", err)
		assert.Equal(t, testId, testSnmpv3User.Id)

		testSnmpv3User, err = c.GetSNMPv3UserById("bad")
		assert.NoError(t, err, "GetSNMPv3UserByName with fake name -> %s", err)
		assert.Equal(t, "", testSnmpv3User.Id)

	} else {
		d, c = getTestDriverU("test_snmp_v3_user")
		testId = d.Tc.GetTestData(d.Env, "Id").(string)
		data, err := c.GetSNMPv3UserById(testId)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetSNMPv3Users(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_snmp_v3_user")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		SNMPv3Users, err := c.GetSNMPv3Users("", "", "", "")
		assert.NoError(t, err, "GetSNMPv3Users threw error -> %s, %+v\n", err, SNMPv3Users)

		SNMPv3Users, err = c.GetSNMPv3Users("", "", "", "name:asc")
		assert.NoError(t, err, "GetSNMPv3Users name:asc error -> %s, %+v\n", err, SNMPv3Users)

	} else {
		_, c = getTestDriverU("test_snmp_v3_user")
		data, err := c.GetSNMPv3Users("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteSNMPv3UserNotFound(t *testing.T) {
	var (
		c              *ov.OVClient
		testId         = "fake"
		testSnmpv3User ov.SNMPv3User
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_snmp_v3_user")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteSNMPv3UserById(testId)
		assert.NoError(t, err, "DeleteSNMPv3User err-> %s", err)

		testSnmpv3User, err = c.GetSNMPv3UserById(testId)
		assert.NoError(t, err, "GetSNMPv3UserByName with deleted ethernet network -> %+v", err)
		assert.Equal(t, "", testSnmpv3User.Id, fmt.Sprintf("Problem getting snmp v3 user id, %+v", testSnmpv3User))
	} else {
		_, c = getTestDriverU("test_snmp_v3_user")
		err := c.DeleteSNMPv3UserById(testId)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testSnmpv3User))
	}
}

func TestDeleteSNMPv3UserById(t *testing.T) {
	var (
		d              *OVTest
		c              *ov.OVClient
		testId         string
		testSnmpv3User ov.SNMPv3User
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_snmp_v3_user")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testId = d.Tc.GetTestData(d.Env, "Id").(string)

		err := c.DeleteSNMPv3UserById(testId)
		assert.NoError(t, err, "DeleteSNMPv3User err-> %s", err)

		testSnmpv3User, err = c.GetSNMPv3UserById(testId)
		assert.NoError(t, err, "GetSNMPv3UserById with deleted ethernet network-> %+v", err)
		assert.Equal(t, "", testSnmpv3User.Id, fmt.Sprintf("Problem getting snmp v3 user id, %+v", testSnmpv3User))
	} else {
		_, c = getTestDriverU("test_snmp_v3_user")
		err := c.DeleteSNMPv3UserById("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testSnmpv3User))
	}

}

func TestDeleteSNMPv3UserByName(t *testing.T) {
	var (
		d              *OVTest
		c              *ov.OVClient
		testUserName   string
		testSnmpv3User ov.SNMPv3User
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_snmp_v3_user")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testUserName = d.Tc.GetTestData(d.Env, "UserName").(string)

		err := c.DeleteSNMPv3UserByName(testUserName)
		assert.NoError(t, err, "DeleteSNMPv3User err-> %s", err)
	} else {
		_, c = getTestDriverU("test_snmp_v3_user")
		err := c.DeleteSNMPv3UserByName("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testSnmpv3User))
	}

}
