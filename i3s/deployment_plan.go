package i3s

import (
	"encoding/json"
	"github.com/HewlettPackard/oneview-golang/utils"
)

type DeploymentPlan struct {
	Category         string            `json:"category,omitempty"`         // "category": "oe-deployment-plan",
	Created          string            `json:"created,omitempty"`          // "created": "20150831T154835.250Z",
	CustomAttributes []CustomAttribute `json:"customAttributes,omitempty"` // "customAttributes": [],
	Description      string            `json:"description,omitempty"`      // "description": "Deployment Plan 1",
	ETAG             string            `json:"eTag,omitempty"`             // "eTag": "1441036118675/8",
	GoldenImageUri   utils.Nstring     `json:"goldenImageUri,omitempty"`   // "goldenImageUri": "",
	HPProvided       bool              `json:"hpProvided,omitempty"`       // "hpProvided": false,
	ID               string            `json:"id,omitempty"`               // "id": "1",
	Modified         string            `json:"modified,omitempty"`         // "modified": "20150831T154835.250Z",
	Name             string            `json:"name,omitempty"`             // "name": "Deployment Plan 1",
	OEBuildPlanURI   utils.Nstring     `json:"oeBuildPlanUri,omitempty"`   // "oeBuildPlanUri": "/rest/build-plans/4f907ab5-7133-4081-bb5a-4a6dea398046",
	State            string            `json:"state,omitempty"`            // "state": "Normal",
	Status           string            `json:"status,omitempty"`           // "status": "Critical",
	Type             string            `json:"type,omitempty"`             // "type": "OEDeploymentPlan",
	URI              utils.Nstring     `json:"uri,omitempty"`              // "uri": "/rest/deployment-plans/31e5dcba-b8ac-4f64-bbaa-7a4474f11994"
}

type CustomAttribute struct {
	Constraints string `json:"constraints,omitempty"` // "constraints": "{\"unit\":\"KB\"}",
	Description string `json:"description,omitempty"` // "description": "custom attribute 1",
	Editable    bool   `json:"editable,omitempty"`    // "editable": false,
	ID          string `json:"id,omitempty"`          // "id": "1",
	Name        string `json:"name,omitempty"`        // "name": "volumesize",
	Type        string `json:"type,omitempty"`        // "type": "number",
	Value       string `json:"value,omitempty"`       // "value": "1045898",
	Visible     bool   `json:"visible,omitempty"`     // "visible": true,
}

type DeploymentPlanList struct {
	Total       int              `json:"total,omitempty"`       // "total": 1,
	Count       int              `json:"count,omitempty"`       // "count": 1,
	Start       int              `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring    `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring    `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring    `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []DeploymentPlan `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetDeploymentPlanByName(name string) (DeploymentPlan, error) {
	var (
		depPlan DeploymentPlan
	)
	depPlans, err := c.GetDeploymentPlans(fmt.Sprintf("name matches '%s'", name), "name:asc")
	if err != nil {
		return depPlan, err
	}
	if depPlans.Total > 0 {
		return depPlans.Members[0], err
	} else {
		return depPlan, err
	}
}

func (c *OVClient) GetDeploymentPlans(filter string, sort string) (DeploymentPlanList, error) {
	var (
		uri             = "/rest/deployment-plans"
		q               map[string]interface{}
		deploymentPlans DeploymentPlanList
	)
	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return deploymentPlans, err
	}

	log.Debugf("GetDeploymentPlans %s", data)
	if err := json.Unmarshal([]byte(data), &deploymentPlans); err != nil {
		return deploymentPlans, err
	}
	return deploymentPlans, nil
}

func (c *OVClient) CreateDeploymentPlan(deploymentPlan DeploymentPlan) error {
	log.Infof("Initializing creation of deploymentPlan for %s.", deploymentPlan.Name)
	var (
		uri = "/rest/deployment-plans"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, deploymentPlan)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, deploymentPlan)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new deployment plan request: %s", err)
		return err
	}

	log.Debugf("Response New DeploymentPlan %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) DeleteDeploymentPlan(name string) error {
	var (
		deploymentPlan DeploymentPlan
		err            error
		t              *Task
		uri            string
	)

	deploymentPlan, err = c.GetDeploymentPlanByName(name)
	if err != nil {
		return err
	}
	if deploymentPlan.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", deploymentPlan.URI, deploymentPlan)
		log.Debugf("task -> %+v", t)
		uri = deploymentPlan.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete deployment plan request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete deployment plan %s", data)
		if err := json.Unmarshal([]byte(data), &t); err != nil {
			t.TaskIsDone = true
			log.Errorf("Error with task un-marshal: %s", err)
			return err
		}
		err = t.Wait()
		if err != nil {
			return err
		}
		return nil
	} else {
		log.Infof("DeploymentPlan could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateDeploymentPlan(deploymentPlan DeploymentPlan) error {
	log.Infof("Initializing update of deployment plan for %s.", deploymentPlan.Name)
	var (
		uri = deploymentPlan.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, deploymentPlan)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, deploymentPlan)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update deployment plan request: %s", err)
		return err
	}

	log.Debugf("Response update DeploymentPlan %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}
