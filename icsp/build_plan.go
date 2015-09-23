package icsp

import (
	"encoding/json"

	"github.com/docker/machine/drivers/oneview/rest"
	"github.com/docker/machine/log"
)

// URLEndPoint(s) export this constant
const (
	URLEndPointBuildPlan = "/rest/os-deployment-build-plans"
)

// BuildPlan struct
type BuildPlan struct {
	// TODO define this
}

// GetAllBuildPlans - returns all OS build plans
// returns BuildPlan
func (c *ICSPClient) GetAllBuildPlans() (APIVersion, error) {
	var (
		uri        = URLEndPointBuildPlan
		apiversion APIVersion
	)

	//c.AuthHeaders := map[string]interface{}{"auth": []interface{}{auth}}
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return apiversion, err
	}

	log.Debugf("GetAllBuildPlans %s", data)
	if err := json.Unmarshal([]byte(data), &apiversion); err != nil {
		return apiversion, err
	}
	return apiversion, err
}

// GetBuildPlan -  returns a build plan
//
// func (c *ICSPClient) GetBuildPlan() (APIVersion, build_id, error) {
func (c *ICSPClient) GetBuildPlan() (int, error) {
	//TODO
 return 0, nil
}
