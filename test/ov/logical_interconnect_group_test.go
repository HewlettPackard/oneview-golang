package ov

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateLogicalInterconnectGroup(t *testing.T) {
	var (
		d                *OVTest
		interconnectData *OVTest
		c                *ov.OVClient
		testName         string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_logical_interconnect_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		interconnectData, c = getTestDriverA("test_interconnect_type")
		if c == nil || interconnectData == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		testLogicalInterconnectGroup, err := c.GetLogicalInterconnectGroupByName(testName)
		assert.NoError(t, err, "CreateLogicalInterconnectGroup get the LogicalInterconnectError error -> %s", err)

		if testLogicalInterconnectGroup.URI.IsNil() {
			/*
				interconnectMapEntryTemplates := make([]ov.InterconnectMapEntryTemplate, 8)
				for i := 0; i < 8; i++ {
					locationEntry1 := ov.LocationEntry{
						RelativeValue: i + 1,
						Type:          "Bay",
					}
					locationEntry2 := ov.LocationEntry{
						RelativeValue: i + 2,
						Type:          "Enclosure",
					}
					locationEntries := make([]ov.LocationEntry, 2)
					locationEntries[0] = locationEntry1
					locationEntries[1] = locationEntry2
					logicalLocation := ov.LogicalLocation{
						LocationEntries: locationEntries,
					}

					interconnectMapEntryTemplate := ov.InterconnectMapEntryTemplate{
						LogicalLocation:              logicalLocation,
						PermittedInterconnectTypeUri: utils.NewNstring(interconnectData.Tc.GetTestData(interconnectData.Env, "URI").(string)),
					}
					interconnectMapEntryTemplates[i] = interconnectMapEntryTemplate
				}
				interconnectMapTemplate := ov.InterconnectMapTemplate{
					InterconnectMapEntryTemplates: interconnectMapEntryTemplates,
				}

				f := false
				tr := true
				ethernetSettings := ov.EthernetSettings{
					Type: "EthernetInterconnectSettingsV3",
					EnableFastMacCacheFailover: &f,
					EnableIgmpSnooping:         &tr,
					EnableRichTLV:              &tr,
					MacRefreshInterval:         10,
					IgmpIdleTimeoutInterval:    250,
				}

				telemetryConfig := ov.TelemetryConfiguration{
					Type:            "telemetry-configuration",
					EnableTelemetry: &tr,
					SampleCount:     12,
					SampleInterval:  150,
				}

				snmpConfiguration := ov.SnmpConfiguration{
					Type:          "snmp-configuration",
					Enabled:       &f,
					ReadCommunity: "test",
					SystemContact: "sys contact",
					SnmpAccess:    []string{"192.168.1.0/24"},
				}

				qosTrafficClassifiers := make([]ov.QosTrafficClassifier, 4)
				qosTrafficClass := ov.QosTrafficClass{
					BandwidthShare:   "10",
					EgressDot1pValue: 5,
					MaxBandwidth:     10,
					RealTime:         &tr,
					ClassName:        "RealTime",
					Enabled:          &tr,
				}
				qosClassificationMap1 := ov.QosClassificationMap{
					Dot1pClassMapping: []int{5, 6, 7},
					DscpClassMapping:  []string{"DSCP 46, EF", "DSCP 40, CS5", "DSCP 48, CS6", "DSCP 56, CS7"},
				}
				qosTrafficClassifier := ov.QosTrafficClassifier{
					QosTrafficClass:          qosTrafficClass,
					QosClassificationMapping: &qosClassificationMap1,
				}
				qosTrafficClassifiers[0] = qosTrafficClassifier

				qosTrafficClass = ov.QosTrafficClass{
					BandwidthShare:   "fcoe",
					EgressDot1pValue: 3,
					MaxBandwidth:     100,
					RealTime:         &f,
					ClassName:        "FCoE lossless",
					Enabled:          &tr,
				}
				qosClassificationMap2 := ov.QosClassificationMap{
					Dot1pClassMapping: []int{3},
					DscpClassMapping:  []string{},
				}
				qosTrafficClassifier = ov.QosTrafficClassifier{
					QosTrafficClass:          qosTrafficClass,
					QosClassificationMapping: &qosClassificationMap2,
				}
				qosTrafficClassifiers[1] = qosTrafficClassifier

				qosTrafficClass = ov.QosTrafficClass{
					BandwidthShare:   "65",
					EgressDot1pValue: 0,
					MaxBandwidth:     100,
					RealTime:         &f,
					ClassName:        "Best effort",
					Enabled:          &tr,
				}
				qosClassificationMap3 := ov.QosClassificationMap{
					Dot1pClassMapping: []int{1, 0},
					DscpClassMapping:  []string{"DSCP 10, AF11", "DSCP 12, AF12", "DSCP 14, AF13", "DSCP 8, CS1", "DSCP 0, CS0"},
				}
				qosTrafficClassifier = ov.QosTrafficClassifier{
					QosTrafficClass:          qosTrafficClass,
					QosClassificationMapping: &qosClassificationMap3,
				}
				qosTrafficClassifiers[2] = qosTrafficClassifier

				qosTrafficClass = ov.QosTrafficClass{
					BandwidthShare:   "25",
					EgressDot1pValue: 2,
					MaxBandwidth:     100,
					RealTime:         &f,
					ClassName:        "Medium",
					Enabled:          &tr,
				}
				qosClassificationMap4 := ov.QosClassificationMap{
					Dot1pClassMapping: []int{4, 3, 2},
					DscpClassMapping: []string{"DSCP 18, AF21",
						"DSCP 20, AF22",
						"DSCP 22, AF23",
						"DSCP 26, AF31",
						"DSCP 28, AF32",
						"DSCP 30, AF33",
						"DSCP 34, AF41",
						"DSCP 36, AF42",
						"DSCP 38, AF43",
						"DSCP 16, CS2",
						"DSCP 24, CS3",
						"DSCP 32, CS4"},
				}
				qosTrafficClassifier = ov.QosTrafficClassifier{
					QosTrafficClass:          qosTrafficClass,
					QosClassificationMapping: &qosClassificationMap4,
				}
				qosTrafficClassifiers[3] = qosTrafficClassifier

				activeQosConfig := ov.ActiveQosConfig{
					Type:                       "QosConfiguration",
					ConfigType:                 "CustomWithFCoE",
					QosTrafficClassifiers:      qosTrafficClassifiers,
					UplinkClassificationType:   "DOT1P",
					DownlinkClassificationType: "DOT1P_AND_DSCP",
				}

				qosConfiguration := ov.QosConfiguration{
					Type:            "qos-aggregated-configuration",
					ActiveQosConfig: activeQosConfig,
				}*/

			testLogicalInterconnectGroup := ov.LogicalInterconnectGroup{
				Name:          testName,
				Type:          d.Tc.GetTestData(d.Env, "Type").(string),
				EnclosureType: d.Tc.GetTestData(d.Env, "EnclosureType").(string),
				//InterconnectMapTemplate: &interconnectMapTemplate,
				//EthernetSettings:        &ethernetSettings,
				//TelemetryConfiguration:  &telemetryConfig,
				//SnmpConfiguration:       &snmpConfiguration,
				//QosConfiguration:        &qosConfiguration,
			}

			err := c.CreateLogicalInterconnectGroup(testLogicalInterconnectGroup)
			assert.NoError(t, err, "CreateLogicalInterconnectGroup error -> %s", err)

			err = c.CreateLogicalInterconnectGroup(testLogicalInterconnectGroup)
			assert.Error(t, err, "CreateLogicalInterconnectGroup should error because the LogicalInterconnectGroup already exists, err-")
		} else {
			log.Warnf("The logicalInterconnectGroup already exists, so skipping CreateLogicalInterconnectGroup test for %s", testName)
		}
	}

}

func TestGetLogicalInterconnectGroups(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_logical_interconnect_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		logicalInterconnectGroups, err := c.GetLogicalInterconnectGroups("", "")
		assert.NoError(t, err, "GetLogicalInterconnectGroup threw error -> %s, %+v\n", err, logicalInterconnectGroups)

		logicalInterconnectGroups, err = c.GetLogicalInterconnectGroups("", "name:asc")
		assert.NoError(t, err, "GetLogicalInterconnectGroup name:asc error -> %s, %+v\n", err, logicalInterconnectGroups)

	} else {
		_, c = getTestDriverU("test_logical_interconnect_group")
		data, err := c.GetLogicalInterconnectGroups("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetLogicalInterconnectGroupByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_logical_interconnect_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testLogicalInterconnectGroup, err := c.GetLogicalInterconnectGroupByName(testName)
		assert.NoError(t, err, "GetLogicalInterconnectGroupByName thew an error -> %s", err)
		assert.Equal(t, testName, testLogicalInterconnectGroup.Name)

		testLogicalInterconnectGroup, err = c.GetLogicalInterconnectGroupByName("bad")
		assert.NoError(t, err, "GetLogicalInterconnectGroupByName with fake name -> %s", err)
		assert.Equal(t, "", testLogicalInterconnectGroup.Name)

	} else {
		d, c = getTestDriverU("test_logical_interconnect_group")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetLogicalInterconnectGroupByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteLogicalInterconnectGroupNotFound(t *testing.T) {
	var (
		c                            *ov.OVClient
		testName                     = "fake"
		testLogicalInterconnectGroup ov.LogicalInterconnectGroup
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_logical_interconnect_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteLogicalInterconnectGroup(testName)
		assert.NoError(t, err, "DeleteLogicaInterconnectGroup err-> %s", err)

		testLogicalInterconnectGroup, err = c.GetLogicalInterconnectGroupByName(testName)
		assert.NoError(t, err, "GetLogicalInterconnectGroupByName with deleted logical interconnect group -> %+v", err)
		assert.Equal(t, "", testLogicalInterconnectGroup.Name, fmt.Sprintf("Problem getting logical interconnect group name, %+v", testLogicalInterconnectGroup))
	} else {
		_, c = getTestDriverU("test_logical_interconnect_group")
		err := c.DeleteLogicalInterconnectGroup(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testLogicalInterconnectGroup))
	}
}

func TestDeleteLogicalInterconnectGroup(t *testing.T) {
	var (
		d                            *OVTest
		c                            *ov.OVClient
		testName                     string
		testLogicalInterconnectGroup ov.LogicalInterconnectGroup
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_logical_interconnect_group")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteLogicalInterconnectGroup(testName)
		assert.NoError(t, err, "DeleteLogicalInterconnectGroup err-> %s", err)

		testLogicalInterconnectGroup, err = c.GetLogicalInterconnectGroupByName(testName)
		assert.NoError(t, err, "GetLogicalInterconnectGroupByName with deleted logical interconnect gorup-> %+v", err)
		assert.Equal(t, "", testLogicalInterconnectGroup.Name, fmt.Sprintf("Problem getting logicalInterconnectGroup name, %+v", testLogicalInterconnectGroup))
	} else {
		_, c = getTestDriverU("test_logical_interconnect_group")
		err := c.DeleteLogicalInterconnectGroup("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testLogicalInterconnectGroup))
	}

}
