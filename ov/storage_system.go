package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type StorageSystemV4 struct {
	Hostname	       string                 `json:"hostname,omitempty"`
	Credentials            Credentials            `json:"credentials,omitempty"`
	Category               string                 `json:"category,omitempty"`
	ETAG                   string                 `json:"eTag,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	State                  string                 `json:"state,omitempty"`
	Status                 string                 `json:"status,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	URI                    utils.Nstring          `json:"uri,omitempty"`
	Family		       string                 `json:"family,omitempty"`
	StoragePoolsUri        utils.Nstring          `json:"storagePoolsUri,omitempty"`
	TotalCapacity	       utils.Nstring          `json:"TotalCapacity,omitempty"`
}

type Credentials struct {
	Password         string        `json:"password,omitempty"`
	Username	 string        `json:"username,omitempty"`
}

type StorageSystemsListV4 struct {
	Total       int               `json:"total,omitempty"`       // "total": 1,
	Count       int               `json:"count,omitempty"`       // "count": 1,
	Start       int               `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring     `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring     `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring     `json:"uri,omitempty"`         // "uri": "/rest/storage-systems"
	Members     []StorageSystemV4 `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetStorageSystemByName(name string) (StorageSystemV4, error) {
	var (
		sSystem StorageSystemV4
	)
	sSystems, err := c.GetStorageSystems(fmt.Sprintf("name matches '%s'", name), "name:asc")
	if sSystems.Total > 0 {
		return sSystems.Members[0], err
	} else {
		return sSystem, err
	}
}

func (c *OVClient) GetStorageSystems(filter string, sort string) (StorageSystemsListV4, error) {
	var (
		uri   = "/rest/storage-systems"
		q     map[string]interface{}
		sSystem StorageSystemsListV4
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
		return sSystem, err
	}

	log.Debugf("GetStorageSystems %s", data)
	if err := json.Unmarshal([]byte(data), &sSystem); err != nil {
		return sSystem, err
	}
	return sSystem, nil
}

func (c *OVClient) CreateStorageSystem(sSystem StorageSystemV4) error {
	log.Infof("Initializing creation of storage volume for %s.", sSystem.Name)
	var (
		uri = "/rest/storage-systems"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, sSystem)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, sSystem)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new storage volume request: %s", err)
		return err
	}

	log.Debugf("Response New StorageSystem %s", data)
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

func (c *OVClient) DeleteStorageSystem(name string) error {
	var (
		sSystem StorageSystemV4
		err  error
		t    *Task
		uri  string
	)

	sSystem, err = c.GetStorageSystemByName(name)
	if err != nil {
		return err
	}
	if sSystem.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", sSystem.URI, sSystem)
		log.Debugf("task -> %+v", t)
		uri = sSystem.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete storage volume request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete storage volume %s", data)
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
		log.Infof("StorageSystem could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateStorageSystem(sSystem StorageSystemV4) error {
	log.Infof("Initializing update of storage volume for %s.", sSystem.Name)
	var (
		uri = sSystem.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, sSystem)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, sSystem)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update StorageSystem request: %s", err)
		return err
	}

	log.Debugf("Response update StorageSystem %s", data)
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
