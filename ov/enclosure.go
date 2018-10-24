package ov

import (
	// "fmt"
	"encoding/json"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Enclosure struct {
	ActiveOaPreferredIP					string               `json:"activeOaPreferredIP,omitempty"`		// "activeOaPreferredIP": "16.124.135.110",
	AssetTag							string               `json:"assetTag,omitempty"`				// "assetTag": "",
	Category                            string               `json:"category,omitempty"`                // "category": "enclosures",
	Created                             string               `json:"created,omitempty"`                 // "created": "20150831T154835.250Z",
	Description                         string               `json:"description,omitempty"`             // "description": "Enclosure Group 1",
	DeviceBayCount						int 				 `json:"deviceBayCount,omitempty"`			// "deviceBayCount": 16,
	DeviceBays 							[]DeviceBayMap		 `json:"deviceBays,omitempty`				// "deviceBays": [],
	ETAG                                string               `json:"eTag,omitempty"`                    // "eTag": "1441036118675/8",
	EnclosureGroupUri                   utils.Nstring        `json:"enclosureGroupUri,omitempty"`       // "enclosureGroupUri": "/rest/enclosure-groups/293e8efe-c6b1-4783-bf88-2d35a8e49071",
	EnclosureType                    	string        		 `json:"enclosureType,omitempty"`           // "enclosureType": "BladeSystem c7000 Enclosure",
	FwBaselineName         				string               `json:"fwBaselineName,omitempty"`  		// "fwBaselineName": null,
	FwBaselineUri             			utils.Nstring        `json:"fwBaselineUri,omitempty"`           // "fwBaselineUri": null,
	InterconnectBayCount                int      			 `json:"interconnectBayCount,omitempty"`	// "interconnectBayCount": 8,
	InterconnectBays 					[]InterconnectBayMap `json:"interconnectBayMappings"` 			// "interconnectBays": [],
	IsFwManaged							bool				 `json:"isFwManaged"`						// "isFwManaged": false,
	LicensingIntent						string               `json:"licensingIntent,omitempty"`			// "licensingIntent": "OneView",
	Modified                            string               `json:"modified,omitempty"`         		// "modified": "20150831T154835.250Z",
	Name                                string               `json:"name,omitempty"`             		// "name": "e10",
	OA 									[]OAMap				 `json:"oa,omitempty"`						// "oa": [],
	OaBayCount							int 				 `json:"oaBayCount,omitempty"`				// "oaBayCount": 2,
	PartNumber                    		string               `json:"partNumber,omitempty"` 				// "partNumber": "403320-B21",
	RackName                        	string               `json:"rackName,omitempty"`     			// "rackName": "Rack-Renamed",
	RefreshState						string 				 `json:"refreshState,omitempty"`			// "refreshState": "NotRefreshing",
	SerialNumber                        string               `json:"serialNumber,omitempty"`     		// "serialNumber": "USE62519EE",
	StandbyOaPreferredIP                string               `json:"standbyOaPreferredIP,omitempty"`	// "standbyOaPreferredIP": "",
	State                               string               `json:"state,omitempty"`            		// "state": "Configured",
	StateReason							string 				 `json:"StateReason"`						// "stateReason": "None",
	Status                              string               `json:"status,omitempty"`           		// "status": "Critical",
	Type                                string               `json:"type,omitempty"`             		// "type": "Enclosure",
	URI                                 utils.Nstring        `json:"uri,omitempty"`              		// "uri": "/rest/enclosures/09USE62519EE",
	UUID								string               `json:"uuid,omitempty"`					// "uuid": "09USE62519EE",
	VcmDomainId							string               `json:"vcmDomainId,omitempty"`				// "vcmDomainId": "@914ae756bdbce70cf7cbce65d34a23",
	VcmDomainName						string               `json:"vcmDomainName,omitempty"`			// "vcmDomainName": "OneViewDomain",
	VcmMode								bool                 `json:"vcmMode,omitempty"`					// "vcmMode": true,
	VcmUrl								string        		 `json:"vcmUrl,omitempty"`					// "vcmUrl": "https://16.124.128.80"
}

type EnclosureList struct {
	Total       int              `json:"total,omitempty"`       // "total": 1,
	Count       int              `json:"count,omitempty"`       // "count": 1,
	Start       int              `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring    `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring    `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring    `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []Enclosure 	 `json:"members,omitempty"`     // "members":[]
}

type DeviceBayMap struct {
	AvailableForFullHeightProfile 	bool			`json:"availableForFullHeightProfile"`	// "availableForFullHeightProfile": false,
    AvailableForHalfHeightProfile   bool 			`json:"availableForHalfHeightProfile"`	// "availableForHalfHeightProfile": true,
    BayNumber 						int 			`json:"bayNumber"`						// "bayNumber": 1,
    Category 						string 			`json:"category,omitempty"`				// "category": "device-bays",
    CoveredByDevice 				utils.Nstring 	`json:"coveredByDevice,omitempty"` 		// "coveredByDevice": "/rest/server-hardware/30373237-3132-4D32-3236-303730344E54",
    CoveredByProfile 				string 			`json:"coveredByProfile,omitempty"`		// "coveredByProfile": null,
    Created 						string 			`json:"created,omitempty"`				// "created": null,
    DevicePresence 					string 			`json:"devicePresence,omitempty"`		// "devicePresence": "Present",
    DeviceUri 						utils.Nstring	`json:"deviceUri,omitempty"`			// "deviceUri": "/rest/server-hardware/30373237-3132-4D32-3236-303730344E54",
    EnclosureUri 					utils.Nstring 	`json:"enclosureUri,omitempty"`			// "enclosureUri": null,
    ETAG 							string 			`json:"eTag,omitempty"`					// "eTag": null,
    Model 							string 			`json:"model,omitempty"`				// "model": null,
    Modified 						string 			`json:"modified,omitempty"`				// "modified": null,
    ProfileUri 						utils.Nstring 	`json:"profileUri,omitempty"`			// "profileUri": null,
    Type 							string 			`json:"type,omitempty"`					// "type": "DeviceBay",
    URI 							utils.Nstring 	`json:"uri,omitempty"`					// "uri": "/rest/enclosures/09USE62519EE/device-bays/1"
}
// type InterconnectBayMap struct {
// 	InterconnectBay             int           `json:"interconnectBay,omitempty"`             // "interconnectBay": 0,
// 	InterconnectUri  			utils.Nstring `json:"interconnectUri,omitempty"`			 // "interconnectUri": "/rest/interconnects/d8aecda2-8bb8-4198-bf84-790bb7b72a06",
// 	LogicalInterconnectGroupUri utils.Nstring `json:"logicalInterconnectGroupUri,omitempty"` // "logicalInterconnectGroupUri": ""
// }

type OAMap struct {
    BayNumber 		int 				`json:"bayNumber"`					// "bayNumber": 1,
    DhcpEnable 		bool 				`json:"dhcpEnable"`					// "dhcpEnable": false,
    DhcpIpv6Enable 	bool 				`json:"dhcpIpv6Enable"`				// "dhcpIpv6Enable": false,
    FqdnHostName 	string 				`json:"fqdnHostName,omitempty"`		// "fqdnHostName": "e10-oa.vse.rdlabs.hpecorp.net",
    FwBuildDate 	string 				`json:"fwBuildDate,omitempty"`		// "fwBuildDate": "Jun 17 2016",
    FwVersion  		string 				`json:"fwVersion,omitempty"`		// "fwVersion": "4.60",
    IpAddress 		string 				`json:"ipAddress,omitempty"`		// "ipAddress": "16.124.135.110",
    Ipv6Addresses	[]Ipv6Addresses		`json:"ipv6Addresses,omitempty"`	// "ipv6Addresses": []
    Role 			string 				`json:"role,omitempty"`				// "role": "Active",
    State 			string 				`json:"state,omitempty"`			// "state": null

}

type Ipv6Addresses struct {
	Address 	string 		`json:"address,omitempty"`	// "address": "",
    Type        string 		`json:"type,omitempty"`		// "type": "NotSet"
}

// func (c *OVClient) GetEnclosureGroupByName(name string) (EnclosureGroup, error) {
// 	var (
// 		enclosureGroup EnclosureGroup
// 	)
// 	enclosureGroups, err := c.GetEnclosureGroups(fmt.Sprintf("name matches '%s'", name), "name:asc")
// 	if enclosureGroups.Total > 0 {
// 		return enclosureGroups.Members[0], err
// 	} else {
// 		return enclosureGroup, err
// 	}
// }

// func (c *OVClient) GetEnclosureGroupByUri(uri utils.Nstring) (EnclosureGroup, error) {
// 	var (
// 		enclosureGroup EnclosureGroup
// 	)
// 	// refresh login
// 	c.RefreshLogin()
// 	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
// 	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
// 	if err != nil {
// 		return enclosureGroup, err
// 	}
// 	log.Debugf("GetEnclosureGroup %s", data)
// 	if err := json.Unmarshal([]byte(data), &enclosureGroup); err != nil {
// 		return enclosureGroup, err
// 	}
// 	return enclosureGroup, nil
// }

func (c *OVClient) GetEnclosures(filter string, sort string) (EnclosureList, error) {
	var (
		uri             = "/rest/enclosures"
		q               map[string]interface{}
		enclosures 		EnclosureList
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
		return enclosures, err
	}
	log.Debugf("GetEnclosures %s", data)
	if err := json.Unmarshal([]byte(data), &enclosures); err != nil {
		return enclosures, err
	}
	return enclosures, nil
}

// func (c *OVClient) CreateEnclosureGroup(eGroup EnclosureGroup) error {
// 	log.Infof("Initializing creation of enclosure group for %s.", eGroup.Name)
// 	var (
// 		uri = "/rest/enclosure-groups"
// 		t   *Task
// 	)

// 	// refresh login
// 	c.RefreshLogin()
// 	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

// 	t = t.NewProfileTask(c)
// 	t.ResetTask()
// 	data, err := c.RestAPICall(rest.POST, uri, eGroup)
// 	if err != nil {
// 		log.Errorf("Error submitting new enclosure group request: %s", err)
// 		return err
// 	}

// 	log.Debugf("Response New EnclosureGroup %s", data)
// 	if err := json.Unmarshal([]byte(data), &t); err != nil {
// 		t.TaskIsDone = true
// 		log.Errorf("Error with task un-marshal: %s", err)
// 		return err
// 	}

// 	return nil
// }

// func (c *OVClient) DeleteEnclosureGroup(name string) error {
// 	var (
// 		enclosureGroup EnclosureGroup
// 		err            error
// 		t              *Task
// 		uri            string
// 	)

// 	enclosureGroup, err = c.GetEnclosureGroupByName(name)
// 	if err != nil {
// 		return err
// 	}
// 	if enclosureGroup.Name != "" {
// 		t = t.NewProfileTask(c)
// 		t.ResetTask()
// 		log.Debugf("REST : %s \n %+v\n", enclosureGroup.URI, enclosureGroup)
// 		log.Debugf("task -> %+v", t)
// 		uri = enclosureGroup.URI.String()
// 		if uri == "" {
// 			log.Warn("Unable to post delete, no uri found.")
// 			t.TaskIsDone = true
// 			return err
// 		}
// 		_, err := c.RestAPICall(rest.DELETE, uri, nil)
// 		if err != nil {
// 			log.Errorf("Error submitting delete enclosure group request: %s", err)
// 			t.TaskIsDone = true
// 			return err
// 		}

// 		return nil
// 	} else {
// 		log.Infof("EnclosureGroup could not be found to delete, %s, skipping delete ...", name)
// 	}
// 	return nil
// }

// func (c *OVClient) UpdateEnclosureGroup(enclosureGroup EnclosureGroup) error {
// 	log.Infof("Initializing update of enclosure group for %s.", enclosureGroup.Name)
// 	var (
// 		uri = enclosureGroup.URI.String()
// 		t   *Task
// 	)
// 	// refresh login
// 	c.RefreshLogin()
// 	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

// 	t = t.NewProfileTask(c)
// 	t.ResetTask()
// 	log.Debugf("REST : %s \n %+v\n", uri, enclosureGroup)
// 	log.Debugf("task -> %+v", t)
// 	data, err := c.RestAPICall(rest.PUT, uri, enclosureGroup)
// 	if err != nil {
// 		t.TaskIsDone = true
// 		log.Errorf("Error submitting update enclosure group request: %s", err)
// 		return err
// 	}

// 	log.Debugf("Response update EnclosureGroup %s", data)
// 	if err := json.Unmarshal([]byte(data), &t); err != nil {
// 		t.TaskIsDone = true
// 		log.Errorf("Error with task un-marshal: %s", err)
// 		return err
// 	}

// 	return nil
// }
