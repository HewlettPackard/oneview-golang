package icsp

import (
	"encoding/json"
	"testing"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// TestGetProfiles
func TestODSUri(t *testing.T) {
	var (
		d *ICSPTest
		u ODSUri
	)
	d, _ = getTestDriverU()
	jsonJobURI := d.Tc.GetTestData(d.Env, "JobURIJSONString").(string)
	log.Debugf("jsonJobURI => %s", jsonJobURI)
	err := json.Unmarshal([]byte(jsonJobURI), &u)
	assert.NoError(t, err, "Unmarshal ODSUri for Job threw error -> %s, %+v\n", err, jsonJobURI)
}
