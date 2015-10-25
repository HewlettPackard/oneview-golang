package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/docker/machine/libmachine/log"
)

// APIVersion struct
type APIVersion struct {
	CurrentVersion int `json:"currentVersion,omitempty"`
	MinimumVersion int `json:"minimumVersion,omitempty"`
}

// GetAPIVersion - returns the api version for OneView server
// returns structure APIVersion
func (c *OVClient) GetAPIVersion() (APIVersion, error) {
	var (
		uri        = "/rest/version"
		apiversion APIVersion
	)

	//c.AuthHeaders := map[string]interface{}{"auth": []interface{}{auth}}
	c.SetAuthHeaderOptions(c.GetAuthHeaderMapNoVer())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return apiversion, err
	}

	log.Debugf("GetAPIVersion %s", data)
	if err := json.Unmarshal([]byte(data), &apiversion); err != nil {
		return apiversion, err
	}
	return apiversion, err
}

// RefreshVersion - refresh the max api Version for the client
func (c *OVClient) RefreshVersion() error {
	var v APIVersion
	v, err := c.GetAPIVersion()
	if err != nil {
		return err
	}
	c.APIVersion = v.CurrentVersion
	return nil
}
