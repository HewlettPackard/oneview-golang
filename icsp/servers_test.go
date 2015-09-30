package icsp

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/machine/log"
	"github.com/stretchr/testify/assert"
)

//TODO: implement create server unt test
func TestCreateServer(t *testing.T) {
	var (
		d              *ICSPTest
		c              *ICSPClient
		user, pass, ip string
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		user = os.Getenv("ONEVIEW_ILO_USER")
		pass = os.Getenv("ONEVIEW_ILO_PASSWORD")
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		ip = d.Tc.GetTestData(d.Env, "IloIPAddress").(string)
		serialNumber := d.Tc.GetTestData(d.Env, "FreeBladeSerialNumber").(string)
		log.Debug("implements acceptance test for TestCreateServer")
		s, err := c.GetServerBySerialNumber(serialNumber) // fake serial number
		assert.NoError(t, err, "GetServerBySerialNumber fake threw error -> %s, %+v\n", err, s)
		if s.URI.String() != "null" {
			// create the server
			err := c.CreateServer(user, pass, ip, 443)
			assert.NoError(t, err, "CreateServer threw error -> %s\n", err)
		} else {
			// create the server
			err := c.CreateServer(user, pass, ip, 443)
			assert.Error(t, err, "CreateServer should throw conflict error  -> %s\n", err)
		}

		// check if the server now exist
	} else {
		log.Debug("implements unit test for TestCreateServer")
	}
}

// TestSaveServer implement save server
func TestSaveServer(t *testing.T) {
	var (
		d *ICSPTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		log.Debug("implements acceptance test for TestCreateServer")
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// get a Server
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
	} else {
		log.Debug("implements unit test for TestCreateServer")
	}
}

// TestGetProfiles
func TestGetServers(t *testing.T) {
	var (
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetServers()
		assert.NoError(t, err, "GetServers threw error -> %s, %+v\n", err, data)

	} else {
		_, c = getTestDriverU()
		data, err := c.GetServers()
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetServerByName(t *testing.T) {
	var (
		d *ICSPTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		IcspName := d.Tc.GetTestData(d.Env, "IcspName").(string)
		data, err := c.GetServerByName(IcspName)
		assert.NoError(t, err, "GetServerByName threw error -> %s, %+v\n", err, data)
	}
}

func TestGetServerByHostName(t *testing.T) {
	var (
		d *ICSPTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testHostName := d.Tc.GetTestData(d.Env, "FreeHostName").(string)
		data, err := c.GetServerByHostName(testHostName)
		assert.NoError(t, err, "GetServerByName threw error -> %s, %+v\n", err, data)
	}
}

func TestGetServerBySerialNumber(t *testing.T) {
	var (
		d *ICSPTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		serialNumber := d.Tc.GetTestData(d.Env, "FreeBladeSerialNumber").(string)
		data, err := c.GetServerBySerialNumber(serialNumber)
		assert.NoError(t, err, "GetServerBySerialNumber threw error -> %s, %+v\n", err, data)

		// negative test
		data, err = c.GetServerBySerialNumber("SXXXX33333") // fake serial number
		assert.NoError(t, err, "GetServerBySerialNumber fake threw error -> %s, %+v\n", err, data)
		assert.Equal(t, data.URI.String(), "null", "GetServerBySerialNumber on fake should be nil")
	}
}

//TODO: implement test for delete
func TestDeleteServer(t *testing.T) {
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		log.Debug("implements acceptance test for TestDeleteServer")
		// check if the server exist
		// delete the server
	} else {
		log.Debug("implements unit test for TestDeleteServer")
	}
}

func TestIsServerManaged(t *testing.T) {
	var (
		//d *ICSPTest
		c *ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.IsServerManaged("2M251204DZ")
		assert.NoError(t, err, "IsServerManaged -> %s, %+v\n", err, data)
		assert.True(t, data)
	}
}
