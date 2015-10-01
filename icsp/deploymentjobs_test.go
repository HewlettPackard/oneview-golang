package icsp

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/docker/machine/log"
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

// TODO: ApplyDeploymentJobs
// TestSaveServer implement save server
func TestApplyDeploymentJobs(t *testing.T) {
	var (
		d *ICSPTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		log.Debug("implements acceptance test for ApplyDeploymentJobs")
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// get a Server
		osBuildPlan := d.Tc.GetTestData(d.Env, "OSBuildPlan").(string)
		serialNumber := d.Tc.GetTestData(d.Env, "FreeBladeSerialNumber").(string)
		s, err := c.GetServerBySerialNumber(serialNumber)
		assert.NoError(t, err, "GetServerBySerialNumber threw error -> %s, %+v\n", err, s)
		// set a custom attribute
		s.SetCustomAttribute("docker_user", "server", "docker")
		// use test keys like from https://github.com/mitchellh/vagrant/tree/master/keys
		// private key from https://raw.githubusercontent.com/mitchellh/vagrant/master/keys/vagrant
		s.SetCustomAttribute("public_key", "server", "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEA6NF8iallvQVp22WDkTkyrtvp9eWW6A8YVr+kz4TjGYe7gHzIw+niNltGEFHzD8+v1I2YJ6oXevct1YeS0o9HZyN1Q9qgCgzUFtdOKLv6IedplqoPkcmF0aYet2PkEDo3MlTBckFXPITAMzF8dJSIFo9D8HfdOV0IAdx4O7PtixWKn5y2hMNG0zQPyUecp4pzC6kivAIhyfHilFR61RGL+GPXQ2MWZWFYbAGjyiYJnAmCP3NOTd0jMZEnDkbUvxhMmBYSdETk1rRgm+R4LOzFUGaHqHDLKLX+FIPKcF96hrucXzcWyLbIbEgE98OHlnVYCzRdK8jlqm8tehUc9c9WhQ==")

		// save a server
		news, err := c.SaveServer(s)
		assert.NoError(t, err, "SaveServer threw error -> %s, %+v\n", err, news)
		assert.Equal(t, s.UUID, news.UUID, "Should return a server with the same UUID")

		// verify that the server attribute was saved by getting the server again and checking the value
		_, testValue2 := s.GetValueItem("docker_user", "server")
		assert.Equal(t, "docker", testValue2.Value, "Should return the saved custom attribute")

		err = c.ApplyDeploymentJobs(osBuildPlan, s)
		assert.NoError(t, err, "ApplyDeploymentJobs threw error -> %s, %+v\n", err, news)
	} else {
		var s Server
		log.Debug("implements unit test for ApplyDeploymentJobs")
		err := c.ApplyDeploymentJobs("testbuildplan", s)
		assert.Error(t, err, "ApplyDeploymentJobs threw error -> %s, %+v\n", err, s)
	}
}
