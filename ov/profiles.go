package ov

import (
	"encoding/json"
	"fmt"
	"github.com/docker/machine/log"
	"github.com/docker/machine/drivers/oneview/rest"
)

// ServerProfile , server profile object for ov
type ServerProfile struct {
	Connections            []Connection `json:"connections,omitempty"`
	Type                   string `json:"type,omitempty"`                  // "type": "ServerProfileV4",
	URI                    string `json:"uri,omitempty"`                   // "uri": "/rest/server-profiles/9979b3a4-646a-4c3e-bca6-80ca0b403a93",
	Name                   string `json:"name,omitempty"`                  // "name": "Server_Profile_scs79",
	Description            string `json:"description,omitempty"`           // "description": "Docker Machine Bay 16",
	SerialNumber           string `json:"serialNumber,omitempty"`          // "serialNumber": "2M25090RMW",
	UUID                   string `json:"uuid,omitempty"`                  // "uuid": "30373237-3132-4D32-3235-303930524D57",
	ServerHardwareURI      string `json:"serverHardwareUri,omitempty"`     // "serverHardwareUri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57",
	ServerHardwareTypeURI  string `json:"serverHardwareTypeUri,omitempty"` // "serverHardwareTypeUri": "/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	EnclosureGroupURI      string `json:"enclosureGroupUri,omitempty"`     // "enclosureGroupUri": "/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039",
	EnclosureURI           string `json:"enclosureUri,omitempty"`          // "enclosureUri": "/rest/enclosures/092SN51207RR",
	EnclosureBay           int    `json:"enclosureBay,omitempty"`          // "enclosureBay": 16,
	Affinity               string `json:"affinity,omitempty"`              // "affinity": "Bay",
	Created                string `json:"created,omitempty"`               // "created": "20150831T154835.250Z",
	Modified               string `json:"modified,omitempty"`              // "modified": "20150902T175611.657Z",
	Status                 string `json:"status,omitempty"`                // "status": "Critical",
	State                  string `json:"state,omitempty"`                 // "state": "Normal",
	InProgress             bool   `json:"inProgress,omitempty"`            // "inProgress": false,
	TaskURI                string `json:"taskUri,omitempty"`               // "taskUri": "/rest/tasks/6F0DF438-7D30-41A2-A36D-62AB866BC7E8",
	ETAG                   string `json:"eTag,omitempty"`        	         // "eTag": "1441036118675/8"

}

// ServerProfileList a list of ServerProfile objects
type ServerProfileList struct {
	Total        int     `json:"total,omitempty"`        // "total": 1,
	Count        int     `json:"count,omitempty"`        // "count": 1,
	Start        int     `json:"start,omitempty"`        // "start": 0,
	PrevPageURI  string  `json:"prevPageUri,omitempty"`  // "prevPageUri": null,
	NextPageURI  string  `json:"nextPageUri,omitempty"`  //"nextPageUri": null,
	URI          string  `json:"uri,omitempty"`          // "uri": "/rest/server-profiles?filter=serialNumber%20matches%20%272M25090RMW%27&sort=name:asc"
	Members      []ServerProfile `json:"members,omitempty"`   //"members":[]
}

// get a server profile by name
func (c *OVClient) GetProfileByName(name string)(ServerProfile, error) {
	var (
		profile ServerProfile
	)
	profiles, err := c.GetProfiles( fmt.Sprintf("name matches '%s'",name), "name:asc" )
	if profiles.Total > 0 {
		return profiles.Members[0], err
	} else {
		return profile, err
	}
}

// GetProfileBySN  accepts serial number
func (c *OVClient) GetProfileBySN(serialnum string)(ServerProfile, error) {
	var (
		profile ServerProfile
	)
	profiles, err := c.GetProfiles( fmt.Sprintf("serialNumber matches '%s'",serialnum), "name:asc" )
	if profiles.Total > 0 {
		return profiles.Members[0], err
	} else {
		return profile, err
	}
}

// get a server profiles
func (c *OVClient) GetProfiles(filter string, sort string)(ServerProfileList, error) {
	var (
		uri    = "/rest/server-profiles"
		q      = map[string]string{}
		profiles ServerProfileList
	)

	if filter != "" {
		q["filter"] = filter
	}

  if sort != "" {
		q["sort"] = sort
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions( c.GetAuthHeaderMap() )
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri , nil)
	if err != nil {
		return profiles, err
	}

	log.Debugf("GetProfiles %s", data)
	if err := json.Unmarshal([]byte(data), &profiles); err != nil {
		return profiles, err
	}
	return profiles, nil
}
