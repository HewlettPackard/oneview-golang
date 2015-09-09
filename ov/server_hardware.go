package ov

import (
	"encoding/json"
	"strings"
	"errors"
	"net/url"
	"github.com/docker/machine/log"
	"github.com/docker/machine/drivers/oneview/rest"
)

// HardwareState
type HardwareState int

const (
	H_UNKNOWN    HardwareState = 1 + iota
	H_ADDING
	H_NOPROFILE_APPLIED
	H_MONITORED
	H_UNMANAGED
	H_REMOVING
	H_REMOVE_FAILED
	H_REMOVED
	H_APPLYING_PROFILE
	H_PROFILE_APPLIED
	H_REMOVING_PROFILE
	H_PROFILE_ERROR
	H_UNSUPPORTED
	H_UPATING_FIRMWARE
)

var hardwarestates = [...]string {
	"Unknown",          // not initialized
	"Adding",           // server being added
	"NoProfileApplied", // server successfully added
	"Monitored",        // server being monitored
	"Unmanaged",        // discovered a supported server
	"Removing",         // server being removed
	"RemoveFailed",     // unsuccessful server removal
	"Removed",          // server successfully removed
	"ApplyingProfile",  // profile being applied to server
	"ProfileApplied",   // profile successfully applied
	"RemovingProfile",  // profile being removed
	"ProfileError",     // unsuccessful profile apply or removal
	"Unsupported",      // server model or version not currently supported by the appliance
	"UpdatingFirmware", // server firmware update in progress
}

func (h HardwareState) String() string { return hardwarestates[h-1] }
func (h HardwareState) Equal(s string) (bool) {return (strings.ToUpper(s) == strings.ToUpper(h.String()))}

// get server hardware from ov
type ServerHardware struct {
	Type                   string  `json:"type,omitempty"`                  // "type": "server-hardware-3",
	Name                   string  `json:"name,omitempty"`                  // "name": "se05, bay 16",
	State                  string  `json:"state,omitempty"`                 // "state": "ProfileApplied",
	StateReason            string  `json:"stateReason,omitempty"`           // "stateReason": "NotApplicable",
	AssetTag               string  `json:"assetTag,omitempty"`              // "assetTag": "[Unknown]",
	Category               string  `json:"category,omitempty"`              // "category": "server-hardware",
	Created                string  `json:"created,omitempty"`               // "created": "2015-08-14T21:02:01.537Z",
	Description            Nstring `json:"description,omitempty"`           // "description": null,
	ETAG                   string  `json:"eTag,omitempty"`                  // "eTag": "1441147370086",
	FormFactor             string  `json:"formFactor,omitempty"`            // "formFactor": "HalfHeight",
	LicensingIntent        string  `json:"licensingIntent,omitempty"`       // "licensingIntent": "OneView",
	LocationURI            string  `json:"locationUri,omitempty"`           // "locationUri": "/rest/enclosures/092SN51207RR",
	MemoryMb               int     `json:"memoryMb,omitempty"`              // "memoryMb": 262144,
	Model                  string  `json:"model,omitempty"`                 // "model": "ProLiant BL460c Gen9",
	Modified               string  `json:"modified,omitempty"`              // "modified": "2015-09-01T22:42:50.086Z",
	MpDnsName              string  `json:"mpDnsName,omitempty"`             // "mpDnsName": "ILO2M25090RMW",
	MpFirwareVersion       string  `json:"mpFirmwareVersion,omitempty"`     // "mpFirmwareVersion": "2.03 Nov 07 2014",
	MpIpAddress            string  `json:"mpIpAddress,omitempty"`           // "mpIpAddress": "172.28.3.136",
	MpModel                string  `json:"mpModel,omitempty"`               // "mpModel": "iLO4",
	PartNumber             string  `json:"partNumber,omitempty"`            // "partNumber": "727021-B21",
	Position               int     `json:"position,omitempty"`              // "position": 16,
	PowerLock              bool    `json:"powerLock,omitempty"`             // "powerLock": false,
	PowerState             string  `json:"powerState,omitempty"`            // "powerState": "Off",
	ProcessorCoreCount     int     `json:"processorCoreCount,omitempty"`    // "processorCoreCount": 14,
	ProcessorCount         int     `json:"processorCount,omitempty"`        // "processorCount": 2,
	ProcessorSpeedMhz      int     `json:"processorSpeedMhz,omitempty"`     // "processorSpeedMhz": 2300,
	ProcessorType          string  `json:"processorType,omitempty"`         // "processorType": "Intel(R) Xeon(R) CPU E5-2695 v3 @ 2.30GHz",
	RefreshState           string  `json:"refreshState,omitempty"`          // "refreshState": "NotRefreshing",
	RomVersion             string  `json:"romVersion,omitempty"`            // "romVersion": "I36 11/03/2014",
	SerialNumber           string  `json:"serialNumber,omitempty"`          // "serialNumber": "2M25090RMW",
	ServerGroupURI         string  `json:"serverGroupUri,omitempty"`        // "serverGroupUri": "/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039",
	ServerHardwareTypeURI  string  `json:"serverHardwareTypeUri,omitempty"` // "serverHardwareTypeUri": "/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	ServerProfileURI       string  `json:"serverProfileUri,omitempty"`      // "serverProfileUri": "/rest/server-profiles/9979b3a4-646a-4c3e-bca6-80ca0b403a93",
	ShortModel             string  `json:"shortModel,omitempty"`            // "shortModel": "BL460c Gen9",
	Status                 string  `json:"status,omitempty"`                // "status": "Warning",
	URI                    string  `json:"uri,omitempty"`                   // "uri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57",
	UUID                   string  `json:"uuid,omitempty"`                  // "uuid": "30373237-3132-4D32-3235-303930524D57",
	VirtualSerialNumber    string  `json:"VirtualSerialNumber,omitempty"`   // "virtualSerialNumber": "",
	VirtualUUID            string  `json:"virtualUuid,omitempty"`           // "virtualUuid": "00000000-0000-0000-0000-000000000000"
	Client                 *OVClient
}

// server hardware list, simillar to ServerProfileList with a TODO
type ServerHardwareList struct {
	Type         string           `json:"type,omitempty"`        // "type": "server-hardware-list-3",
	Category     string           `json:"category,omitempty"`    // "category": "server-hardware",
	Count        int              `json:"count,omitempty"`       // "count": 15,
	Created      string           `json:"created,omitempty"`     // "created": "2015-09-08T04:58:21.489Z",
	ETAG         string           `json:"eTag,omitempty"`        // "eTag": "1441688301489",
	Modified     string           `json:"modified,omitempty"`    // "modified": "2015-09-08T04:58:21.489Z",
	NextPageURI  Nstring          `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI  Nstring          `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start        int              `json:"start,omitempty"`       // "start": 0,
	Total        int              `json:"total,omitempty"`       // "total": 15,
	URI          string           `json:"uri,omitempty"`         // "uri": "/rest/server-hardware?sort=name:asc&filter=serverHardwareTypeUri=%27/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C%27&filter=serverGroupUri=%27/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039%27&start=0&count=100"
	Members      []ServerHardware `json:"members,omitempty"`     //"members":[]
}

// get a server hardware with uri
func (c *OVClient) GetServerHardware(uri string)(ServerHardware, error) {

	var hardware ServerHardware
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions( c.GetAuthHeaderMap() )

	// rest call
	data, err := c.RestAPICall(rest.GET, uri , nil)
	if err != nil {
		return hardware, err
	}

	log.Debugf("GetServerHardware %s", data)
	if err := json.Unmarshal([]byte(data), &hardware); err != nil {
		return hardware, err
	}
	hardware.Client = c
	return hardware, nil
}

// get a server hardware with filters
func (c *OVClient) GetServerHardwareList(filters []string, sort string)(ServerHardwareList, error) {
	var (
		uri    = "/rest/server-hardware"
		q      map[string]interface{}
		serverlist ServerHardwareList
	)
  q = make(map[string]interface{})
	if len(filters) > 0 {
		q["filter"] = filters
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

	var Url *url.URL
	Url, err = url.Parse(c.Sanatize(c.Endpoint))
	Url.Path += uri
	c.GetQueryString(Url)

	log.Infof("uri -> %s qs -> %s , data -> %+v, error %s", uri, Url.RawQuery, data, err)
	if err != nil {
		return serverlist, err
	}

	log.Debugf("GetServerHardwareList %s", data)
	if err := json.Unmarshal([]byte(data), &serverlist); err != nil {
		return serverlist, err
	}
	return serverlist, nil
}

// get available server
// blades = rest_api(:oneview, :get, "/rest/server-hardware?sort=name:asc&filter=serverHardwareTypeUri='#{server_hardware_type_uri}'&filter=serverGroupUri='#{enclosure_group_uri}'")
func (c *OVClient) GetAvailableHardware(hardwaretype_uri string, servergroup_uri string) (hw ServerHardware, err error) {
	var (
		hwlist ServerHardwareList
		f      = []string{	hardwaretype_uri, servergroup_uri }
	)
	if hwlist, err = c.GetServerHardwareList( f, "name:asc"); err != nil {
		return hw, err
	}
	if ! (len(hwlist.Members) > 0) {
		return hw, errors.New("Error! No available blades that are compatible with the server profile!")
	}

	// pick an available blade
	for _, blade := range hwlist.Members {
		if !H_PROFILE_APPLIED.Equal(blade.State) && !H_APPLYING_PROFILE.Equal(blade.State) {
			hw = blade
			break
		}
	}
	if hw.Name == "" {
		return hw, errors.New("No more blades are available for provisioning!")
	}
	return hw, nil
}
