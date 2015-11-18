package icsp

import (
	"encoding/json"
	"testing"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// TestNetConfigType - verify we can serialize NetConfig
func TestNetConfigType(t *testing.T) {
	var (
		d *ICSPTest
		n NetConfig
	)
	d, _ = getTestDriverU()
	jsonData := d.Tc.GetTestData(d.Env, "NetConfigJSON").(string)
	log.Debugf("jsonData => %s", jsonData)
	err := json.Unmarshal([]byte(jsonData), &n)
	assert.NoError(t, err, "Unmarshal NetConfig data -> %s, %+v\n", err, jsonData)
	assert.True(t, len(n.Interfaces) == 2, "Should have 2 interfaces")
}

// TODO TestNewNetConfig - verify we can create an empty netconfig type
// TODO TestNewNetConfigInterface - verify we can create an interface type
// TODO TestAddAllDHCP - verify we can convert a servers interfaces to dhcp
// TODO TestSetStaticInterface - verify we can set one of the servers interface to static
// TODO TesttoJSON - verify we can return a json string of the NetConfig object
// TODO TestSave - save netconfig to hpsa_netconfig
