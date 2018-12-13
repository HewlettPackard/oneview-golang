package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Scope struct {
	Description utils.Nstring `json:"description,omitempty"` // "description": "Scope 1",
	//Modified              string        `json:"modified,omitempty"`              // "modified": "20150831T154835.250Z",
	Name string `json:"name,omitempty"` // "name": "Ethernet Network 1",
	//State                 string        `json:"state,omitempty"`                 // "state": "Normal",
	//Status                string        `json:"status,omitempty"`                // "status": "Critical",
	Type string        `json:"type,omitempty"` // "type": "scopesV3",
	URI  utils.Nstring `json:"uri,omitempty"`  // "uri": "/rest/scopess/e2f0031b-52bd-4223-9ac1-d91cb519d548"
}

type ScopeList struct {
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Count       int           `json:"count,omitempty"`       // "count": 1,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []Scope       `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetScopeByName(name string) (Scope, error) {
	var (
		scp Scope
	)
	scps, err := c.GetScopes(fmt.Sprintf("name matches '%s'", name), "name:asc")
	if scps.Total > 0 {
		return scps.Members[0], err
	} else {
		return scp, err
	}
}

func (c *OVClient) GetScopes(query string, sort string) (ScopeList, error) {
	var (
		uri    = "/rest/scopes"
		q      map[string]interface{}
		Scopes ScopeList
	)
	q = make(map[string]interface{})
	if len(query) > 0 {
		q["query"] = query
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
		return Scopes, err
	}

	log.Debugf("GetScopes %s", data)
	if err := json.Unmarshal([]byte(data), &Scopes); err != nil {
		return Scopes, err
	}
	return Scopes, nil
}

func (c *OVClient) CreateScope(scp Scope) error {
	log.Infof("Initializing creation of scope for %s.", scp.Name)
	var (
		uri = "/rest/scopes"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, scp)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, scp)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new scopes request: %s", err)
		return err
	}

	log.Debugf("Response New Scope %s", data)
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

func (c *OVClient) DeleteScope(name string) error {
	var (
		scp Scope
		err error
		t   *Task
		uri string
	)

	scp, err = c.GetScopeByName(name)
	if err != nil {
		return err
	}
	if scp.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", scp.URI, scp)
		log.Debugf("task -> %+v", t)
		uri = scp.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete scope request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete scope %s", data)
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
		log.Infof("Scope could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateScope(scp Scope) error {
	log.Infof("Initializing update of scope for %s.", scp.Name)
	var (
		uri = scp.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, scp)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, scp)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update scope request: %s", err)
		return err
	}

	log.Debugf("Response update Scope %s", data)
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
