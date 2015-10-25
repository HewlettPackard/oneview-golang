package icsp

import (
	"encoding/json"
	"testing"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// TestDeploymentJobs deployment job struct testing
func TestDeploymentJobs(t *testing.T) {
	var (
		d  *ICSPTest
		dj *DeploymentJobs
	)
	d, _ = getTestDriverU()
	jsonData := d.Tc.GetTestData(d.Env, "DeploymentJobsJSON").(string)
	log.Debugf("jsonData => %s", jsonData)
	err := json.Unmarshal([]byte(jsonData), &dj)
	assert.NoError(t, err, "Unmarshal DeploymentJobs threw error -> %s, %+v\n", err, jsonData)
}

// TestPersonalizeServerData deployment job struct testing
func TestPersonalizeServerData(t *testing.T) {
	var (
		d *ICSPTest
		o *OSDPersonalizeServerDataV2
	)
	d, _ = getTestDriverU()
	jsonData := d.Tc.GetTestData(d.Env, "OSDPersonalizeServerDataV2JSON").(string)
	log.Debugf("jsonData => %s", jsonData)
	err := json.Unmarshal([]byte(jsonData), &o)
	assert.NoError(t, err, "Unmarshal DeploymentJobs threw error -> %s, %+v\n", err, jsonData)
	//OSDPersonalizeServerDataV2JSON2
	jsonData = d.Tc.GetTestData(d.Env, "OSDPersonalizeServerDataV2JSON2").(string)
	log.Debugf("jsonData => %s", jsonData)
	err = json.Unmarshal([]byte(jsonData), &o)
	assert.NoError(t, err, "Unmarshal OSDPersonalizeServerDataV2JSON2 threw error -> %s, %+v\n", err, jsonData)
}

// TestOSDPersonalityDataV2 test struct
func TestOSDPersonalityDataV2(t *testing.T) {
	var (
		d *ICSPTest
		o *OSDNicDataV2
	)
	d, _ = getTestDriverU()
	jsonData := d.Tc.GetTestData(d.Env, "OSDPersonalityDataV2JSON").(string)
	log.Debugf("jsonData => %s", jsonData)
	err := json.Unmarshal([]byte(jsonData), &o)
	assert.NoError(t, err, "Unmarshal OSDPersonalityDataV2 threw error -> %s, %+v\n", err, jsonData)
}

// TestOSDNicDataV2 test struct
func TestOSDNicDataV2(t *testing.T) {
	var (
		d *ICSPTest
		o *OSDNicDataV2
	)
	d, _ = getTestDriverU()
	jsonData := d.Tc.GetTestData(d.Env, "InterfaceJSON").(string)
	log.Debugf("jsonData => %s", jsonData)
	err := json.Unmarshal([]byte(jsonData), &o)
	assert.NoError(t, err, "Unmarshal InterfaceJSON threw error -> %s, %+v\n", err, jsonData)
}
