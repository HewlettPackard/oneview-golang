package ov

import (
	"encoding/json"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type TimeConfig struct {
	DisplayName utils.Nstring `json:"displayName,omitempty"`
	Locale      string `json:"locale,omitempty"`
	LocaleName  string `json:"localName,omitempty"`
}

type TimeConfigList struct {
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Count       int           `json:"count,omitempty"`       // "count": 1,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	ETag        utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring `json:"uri,omitempty"`         // "uri": "",
	Members     []TimeConfig  `json:"members,omitempty"`     // "members":[]
	Created     string        `json:"status,omitempty"`
	Category    string        `json:"status,omitempty"`
	Type        string        `json:"status,omitempty"`
	Modified    string        `json:"status,omitempty"`
}

func (c *OVClient) GetTimeConfigs() (TimeConfigList, error) {
	var (
		uri         = "/rest/appliance/configuration/timeconfig/locales"
		timeConfigs TimeConfigList
	)


	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return timeConfigs, err
	}
	log.Debugf("GetTimeConfig List %s", data)
	if err := json.Unmarshal(data, &timeConfigs); err != nil {
		return timeConfigs, err
	}
	return timeConfigs, nil
}
