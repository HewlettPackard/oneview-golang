package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type UplinkSet struct {
	Name                           string           `json:"name"`                     // "name": "Uplink77",
	LogicalInterconnectURI         utils.Nstring    `json:"logicalInterconnectUri"`   // "logicalInterconnectUri": "/rest/logical-interconnects/7769cae0-b680-435b-9b87-9b864c81657f",
	NetworkURIs                    []utils.Nstring  `json:"networkUris"`              // "networkUris": "/rest/uplink-sets/e2f0031b-52bd-4223-9ac1-d91cb519d548",
	FcNetworkURIs                  []utils.Nstring  `json:"[]"`                       // "fcNetworkUris": "[]",
	FcoeNetworkURIs                []utils.Nstring  `json:"[]"`                       // "fcoeNetworkUris": "[]",
	PortConfigInfos                string           `json:"[]"`                       // "portConfigInfos": "[]",
	ConnectionMode                 string           `json:"connectionMode"`           // "connectionMode":"Auto",
	NetworkType                    string           `json:"networkType"`              // "networkType":"Ethernet",
	ManualLoginRedistributionState string           `json:"manualLoginRedistributionState"` //"manualLoginRedistributionState":"NotSupported"
}

type UplinkSetList struct {
	Total             int                 `json:"total,omitempty"`             // "total": 1,
	Count             int                 `json:"count,omitempty"`             // "count": 1,
	Start             int                 `json:"start,omitempty"`             // "start": 0,
	PrevPageURI       utils.Nstring       `json:"prevPageUri,omitempty"`       // "prevPageUri": null,
	NextPageURI       utils.Nstring       `json:"nextPageUri,omitempty"`       // "nextPageUri": null,
	URI               utils.Nstring       `json:"uri,omitempty"`               // "uri": "/rest/uplink-sets?start=0&count=10"
	Members           []UplinkSet         `json:"members,omitempty"`           // "members":[]
	Type              string              `json:"type,omitempty"`              // "type": "UplinkSetCollectionV4",
}

func (c *OVClient) GetUplinkSetByName(name string) (UplinkSet, error) {
	var (
		upSet UplinkSet
	)
	upSets, err := c.GetUplinkSets(fmt.Sprintf("name matches '%s'", name), "name:asc")
	if upSets.Total > 0 {
		return upSets.Members[0], err
	} else {
		return upSet,err
	}
}

func (c *OVClient) GetUplinkSets(filter string, sort string) (UplinkSetList, error) {
	var (
		uri              = "/rest/uplink-sets"
		q                map[string]interface{}
		uplinkSets UplinkSetList
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
		return uplinkSets, err
	}

	log.Debugf("GetUplinkSets %s", data)
	if err := json.Unmarshal([]byte(data), &uplinkSets); err != nil {
		return uplinkSets, err
	}
	return uplinkSets, nil
}


func (c *OVClient) GetUplinkSetById(id string) ([]string, error) {
	var (
		uri          = "/rest/uplink-sets/"
		uplinkSetId = new([]string)
	)
	uri = uri + id
	fmt.Println(uri)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return *uplinkSetId, err
	}
	log.Infof("GetUplinkSetId %s", data)
	if err := json.Unmarshal([]byte(data), uplinkSetId); err != nil {
		return *uplinkSetId, err
	}
	return *uplinkSetId, nil
}

func (c *OVClient) CreateUplinkSet(eNet UplinkSet) error {
	log.Infoof("Initializing creation of uplink-set for %s.",upSet.Name)
	var (
		uri = "/rest/uplink-sets"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, upSet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, upSet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new Uplink Set request: %s", err)
		return err
	}

	log.Debugf("Response New Uplink Set %s", data)
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


func (c *OVClient) DeleteUplinkSet(name string) error {
	var (
		eNet UplinkSet 
		err  error
		t    *Task
		uri  string
	)

	upSet, err = c.GetUplinkSetByName(name)
	if err != nil {
		return err
	}
	if upSet.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", upSet.URI, eNet)
		log.Debugf("task -> %+v", t)
		uri = upSet.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete uplink-set  request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete uplink-set %s", data)
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
		log.Infof("uplink-set could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateUplinkSet(upSet UplinkSet) error {
	log.Infof("Initializing update of uplink-set for %s.", upSet.Name)
	var (
		uri = eNet.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, upSet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, upSet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update uplink-set request: %s", err)
		return err
	}

	log.Debugf("Response update Uplink-set %s", data)
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
