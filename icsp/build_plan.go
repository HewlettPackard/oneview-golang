package icsp

import (
	"encoding/json"

	"github.com/docker/machine/drivers/oneview/rest"
	"github.com/docker/machine/drivers/oneview/utils"
	"github.com/docker/machine/log"
)

// URLEndPoint(s) export this constant
const (
	URLEndPointBuildPlan = "/rest/os-deployment-build-plans"
)

// BuildPlanItem
type BuildPlanItem struct {
	CfgFileDownload  bool          `json:"cfgFileDownload,omitempty"`  // cfgFileDownload - Boolean that indicates whether the current step is used for downloading configuration file or uploading it
	CfgFileOverwrite bool          `json:"cfgFileOverwrite,omitempty"` // cfgFileOverwrite - Flag that indicates whether or not to overwrite the file on the target server if the step type is 'Config File' and it is a download
	CodeType         string        `json:"codeType,omitempty"`         // codeType -  Supported types for scripts: OGFS, Python, Unix, Windows .BAT, Windows VBScript. Supported types for packages: Install ZIP. Supported type for configuration file: Config File
	ID               string        `json:"id:omitempty"`               // id - System-assigned id of the Step
	Name             string        `json:"name,omitempty"`             // name  - name of step
	Parameters       string        `json:"parameters,omitempty"`       // parameters -  Additional parameters that affect the operations of the Step
	Type             string        `json:"type,omitempty"`             // type - TYpe of the step
	URI              utils.Nstring `json:"uri,omitempty"`              // uri - The canonical URI of the Step
}

// BuildPlanHistory
type BuildPlanHistory struct {
	Summary string `json:"summary,omitempty"` // summary - A time ordered array of change log entries. An empty array is returned if no entries were found
	User    string `json:"user,omitempty"`    // user - User to whom log entries belong
	Time    string `json:"time,omitempty"`    // time - Time window for log entries. Default is 90 days
}

// BuildPlanCustAttrs
type BuildPlanCustAttrs struct {
	Attribute string `json:"attribute,omitempty"` // Attribute - Name of the name/value custom attribute pair associated with this OS Build Plan
	Value     string `json:"value,omitempty"`     // Value - Value of the name/value custom attribute pair associated with this OS Build Plan
}

// OSDBuildPlan struct
type OSDBuildPlan struct {
	Arch              string             `json:"arch,omitempty"`
	BuildPlanHistory  []BuildPlanHistory `json:"buildPlanHistory,omitempty"` // buildPlanHistory array
	BuildPlanStepType string             `json:"buildPlanStepType,omitempty"`
}

// OSBuildPlan
type OSBuildPlan struct {
	Category    string         `json:"category,omitempty"`    //Category - Resource category used for authorizations and resource type groupings
	Count       int            `json:"count,omiitempty"`      // Count - The actual number of resources returned in the specified page
	Created     string         `json:"created,omitempty"`     // Created - Date and time when the resource was created
	ETag        string         `json:"eTag,omitempty"`        // ETag - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource
	Members     []OSDBuildPlan `json:"members,omitempty"`     // Members - array of BuildPlans
	Modified    string         `json:"modified,omitempty"`    // Modified -
	NextPageURI utils.Nstring  `json:"nextPageUri,omitempty"` // NextPageURI - URI pointing to the page of resources following the list of resources contained in the specified collection
	PrevPageURI utils.Nstring  `json:"prevPageURI,omitempty"` // PrevPageURI - URI pointing to the page of resources preceding the list of resources contained in the specified collection
	Start       int            `json:"start,omitempty"`       // Start - The row or record number of the first resource returned in the specified page
	Total       int            `json:"total,omitempty"`       // Total -  The total number of resources that would be returned from the query (including any filters), without pagination or enforced resource limits
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
func (c *ICSPClient) GetBuildPlanByName(osdbuildPan string) (int, error) {
	//TODO
	return 0, nil
}
