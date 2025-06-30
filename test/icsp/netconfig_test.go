package icsp

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/HewlettPackard/oneview-golang/icsp"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

var n icsp.NetConfig

// GetNetConfigTestData - loads n with NetConfig
func GetNetConfigTestData() error {
	var (
		d *ICSPTest
	)
	d, _ = getTestDriverU()
	jsonData := d.Tc.GetTestData(d.Env, "NetConfigJSON").(string)
	log.Debugf("jsonData => %s", jsonData)
	err := json.Unmarshal([]byte(jsonData), &n)
	return err
}

// TestNetConfigType - verify we can serialize NetConfig
func TestNetConfigType(t *testing.T) {
	err := GetNetConfigTestData()
	assert.NoError(t, err, "Should Unmarshal NetConfig data -> %s", err)
	assert.True(t, len(n.Interfaces) == 2, "Should have 2 interfaces")
}

// TestNewNetConfig - verify we can create an empty netconfig type
func TestNewNetConfig(t *testing.T) {
	var emptyconfig utils.Nstring
	emptyconfig.Nil()
	// TODO: determine configuration options for network customization
	n = icsp.NewNetConfig(emptyconfig, //s.HostName,
		emptyconfig, // workgroup utils.Nstring,
		emptyconfig, // domain utils.Nstring,
		emptyconfig, // winslist utils.Nstring,
		emptyconfig, // dnsnamelist utils.Nstring,
		emptyconfig) // dnssearchlist utils.Nstring)
	assert.Equal(t, "", n.Hostname, "should have empty hostname")
	assert.Equal(t, "", n.Workgroup, "should have empty workgroup")
	assert.Equal(t, "", n.Domain, "should have empty domain")
	assert.Equal(t, emptyconfig, n.WINSList, "should have empty winslist")
	assert.Equal(t, emptyconfig, n.DNSNameList, "should have empty dnsnamelist")
	assert.Equal(t, emptyconfig, n.DNSSearchList, "should have empty dnssearchlist")
}

// TODO TestNewNetConfigInterface - verify we can create an interface type

// TestNetConfigAddAllDHCP - verify we can convert a servers interfaces to dhcp
func TestNetConfigAddAllDHCP(t *testing.T) {
	var (
		d *ICSPTest
		c *icsp.ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		serialNumber := d.Tc.GetTestData(d.Env, "FreeICSPSerialNumber").(string)
		s, err := c.GetServerBySerialNumber(serialNumber) // fake serial number
		assert.NoError(t, err, "should GetServerBySerialNumber -> %s, %+v", err, s)

		err = GetNetConfigTestData()
		assert.NoError(t, err, "should GetNetConfigTestData -> %s", err)

		var emptyconfig utils.Nstring
		emptyconfig.Nil()
		n.AddAllDHCP(s.Interfaces, false, emptyconfig)
		for _, neti := range n.Interfaces {
			assert.True(t, neti.DHCPv4, fmt.Sprintf("Should have interface, %s, dhcpv4 enabled", neti.MACAddr))
			assert.False(t, neti.IPv6Autoconfig, "Should have ipv6 auto config as false")
		}
	}
}

// TODO TestSetStaticInterface - verify we can set one of the servers interface to static
// TesttoJSON - verify we can return a json string of the NetConfig object
func TestToJSON(t *testing.T) {
	var (
		d *ICSPTest
		c *icsp.ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		serialNumber := d.Tc.GetTestData(d.Env, "FreeICSPSerialNumber").(string)
		s, err := c.GetServerBySerialNumber(serialNumber) // fake serial number
		assert.NoError(t, err, "should GetServerBySerialNumber -> %s, %+v", err, s)

		err = GetNetConfigTestData()
		assert.NoError(t, err, "should GetNetConfigTestData -> %s", err)

		var emptyconfig utils.Nstring
		emptyconfig.Nil()
		n.AddAllDHCP(s.Interfaces, false, emptyconfig)
		data, err := n.ToJSON()
		assert.NoError(t, err, "Should convert object to json, %s", err)
		log.Infof("n JSON -> %s", data)
		assert.True(t, len(data) > 0, "Should have some data")
	}
}

// TestNetConfigSave - save netconfig to hpsa_netconfig
func TestNetConfigSave(t *testing.T) {
	var (
		d *ICSPTest
		c *icsp.ICSPClient
	)
	if os.Getenv("ICSP_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA()
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		serialNumber := d.Tc.GetTestData(d.Env, "FreeICSPSerialNumber").(string)
		s, err := c.GetServerBySerialNumber(serialNumber) // fake serial number
		assert.NoError(t, err, "should GetServerBySerialNumber -> %s, %+v", err, s)

		err = GetNetConfigTestData()
		assert.NoError(t, err, "should GetNetConfigTestData -> %s", err)

		var emptyconfig utils.Nstring
		emptyconfig.Nil()
		n.AddAllDHCP(s.Interfaces, false, emptyconfig)
		s, err = n.Save(s)
		assert.NoError(t, err, "should Save -> %s, %+v", err, s)
		// place those strings into custom attributes
		_, err = c.SaveServer(s)
		assert.NoError(t, err, "should SaveServer -> %s, %+v", err, s)
	}
}
