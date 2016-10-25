package ov

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateLogicalSwitchGroup(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_logical_switch_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test logicalSwitchGroup already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testLogicalSwitchGroup, err := c.GetLogicalSwitchGroupByName(testName)
		assert.NoError(t, err, "CreateLogicalSwitchGroup get the LogicalSwitchError error -> %s", err)

		if testLogicalSwitchGroup.URI.IsNil() {

			switchMapData := d.Tc.GetTestData(d.Env, "SwitchMapTemplate").(map[string]interface{})

			locationEntry := LocationEntry{
				RelativeValue: 1,
				Type:          switchMapData["SwitchMapEntryTemplates"].([]interface{})[0].(map[string]interface{})["LogicalLocation"].(map[string]interface{})["LocationEntries"].([]interface{})[0].(map[string]interface{})["Type"].(string),
			}
			locationEntries := make([]LocationEntry, 1)
			locationEntries[0] = locationEntry
			logicalLocation := LogicalLocation{
				LocationEntries: locationEntries,
			}

			switchMapEntry := SwitchMapEntry{
				PermittedSwitchTypeUri: utils.NewNstring(switchMapData["SwitchMapEntryTemplates"].([]interface{})[0].(map[string]interface{})["PermittedSwitchTypeUri"].(string)),
				LogicalLocation:        logicalLocation,
			}
			switchMapEntries := make([]SwitchMapEntry, 1)
			switchMapEntries[0] = switchMapEntry

			switchMapTemplate := SwitchMapTemplate{
				SwitchMapEntryTemplates: switchMapEntries,
			}
			//log.Warnf("%s", switchMapData["SwitchMapEntryTemplates"].([]interface{})[0].(map[string]interface{})["LogicalLocation"].(map[string]interface{})["LocationEntries"].([]interface{})[0].(map[string]interface{})["RelativeValue"])
			testLogicalSwitchGroup = LogicalSwitchGroup{
				Name:              testName,
				Type:              d.Tc.GetTestData(d.Env, "Type").(string),
				Category:          d.Tc.GetTestData(d.Env, "Category").(string),
				SwitchMapTemplate: switchMapTemplate,
			}
			err := c.CreateLogicalSwitchGroup(testLogicalSwitchGroup)
			assert.NoError(t, err, "CreateLogicalSwitchGroup error -> %s", err)

			err = c.CreateLogicalSwitchGroup(testLogicalSwitchGroup)
			assert.Error(t, err, "CreateLogicalSwitchGroup should error because the LogicalSwitchGroup already exists, err-> %s", err)

		} else {
			log.Warnf("The logicalSwitchGroup already exists, so skipping CreateLogicalSwitchGroup test for %s", testName)
		}

		// reload the test profile that we just created
		testLogicalSwitchGroup, err = c.GetLogicalSwitchGroupByName(testName)
		assert.NoError(t, err, "GetLogicalSwitchGroupByName error -> %s", err)
	}

}

func TestGetLogicalSwitchGroups(t *testing.T) {
	var (
		c *OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_logical_switch_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		logicalSwitchGroups, err := c.GetLogicalSwitchGroups("", "")
		assert.NoError(t, err, "GetLogicalSwitchGroups threw error -> %s, %+v\n", err, logicalSwitchGroups)

		logicalSwitchGroups, err = c.GetLogicalSwitchGroups("", "name:asc")
		assert.NoError(t, err, "GetLogicalSwitchGroups name:asc error -> %s, %+v\n", err, logicalSwitchGroups)

	} else {
		_, c = getTestDriverU("test_logical_switch_group")
		data, err := c.GetLogicalSwitchGroups("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetLogicalSwitchGroupByName(t *testing.T) {
	var (
		d        *OVTest
		c        *OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_logical_switch_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testLogicalSwitchGroup, err := c.GetLogicalSwitchGroupByName(testName)
		assert.NoError(t, err, "GetLogicalSwitchGroupByName thew an error -> %s", err)
		assert.Equal(t, testName, testLogicalSwitchGroup.Name)

		testLogicalSwitchGroup, err = c.GetLogicalSwitchGroupByName("bad")
		assert.NoError(t, err, "GetLogicalSwitchGroupByName with fake name -> %s", err)
		assert.Equal(t, "", testLogicalSwitchGroup.Name)

	} else {
		d, c = getTestDriverU("test_logical_switch_group")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetLogicalSwitchGroupByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteLogicalSwitchGroupNotFound(t *testing.T) {
	var (
		c                      *OVClient
		testName               = "fake"
		testLogicalSwitchGroup LogicalSwitchGroup
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_logical_switch_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteLogicalSwitchGroup(testName)
		assert.NoError(t, err, "DeleteLogicalSwitchGroup err-> %s", err)

		testLogicalSwitchGroup, err = c.GetLogicalSwitchGroupByName(testName)
		assert.NoError(t, err, "GetLogicalSwitchGroupByName with deleted logical switch group -> %+v", err)
		assert.Equal(t, "", testLogicalSwitchGroup.Name, fmt.Sprintf("Problem getting logical switch group name, %+v", testLogicalSwitchGroup))
	} else {
		_, c = getTestDriverU("test_logical_switch_group")
		err := c.DeleteLogicalSwitchGroup(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testLogicalSwitchGroup))
	}
}

func TestDeleteLogicalSwitchGroup(t *testing.T) {
	var (
		d                      *OVTest
		c                      *OVClient
		testName               string
		testLogicalSwitchGroup LogicalSwitchGroup
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_logical_switch_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteLogicalSwitchGroup(testName)
		assert.NoError(t, err, "DeleteLogicalSwitchGroup err-> %s", err)

		testLogicalSwitchGroup, err = c.GetLogicalSwitchGroupByName(testName)
		assert.NoError(t, err, "GetLogicalSwitchGroupByName with deleted logical switch gorup-> %+v", err)
		assert.Equal(t, "", testLogicalSwitchGroup.Name, fmt.Sprintf("Problem getting logicalSwitchGroup name, %+v", testLogicalSwitchGroup))
	} else {
		_, c = getTestDriverU("test_logical_switch_group")
		err := c.DeleteLogicalSwitchGroup("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testLogicalSwitchGroup))
	}

}
