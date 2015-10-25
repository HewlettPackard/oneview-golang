package ov

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/liboneview"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/docker/machine/libmachine/log"
)

// introduced in v200 for oneview, allows for an easier method
// to build the profiles for servers and associate them.
// we don't operate on a new struct, we simply use the ServerProfile struct

// ProfileTemplatesNotSupported - determine these functions are supported
func (c *OVClient) ProfileTemplatesNotSupported() bool {
	var currentversion liboneview.Version
	var asc liboneview.APISupport
	currentversion = currentversion.CalculateVersion(c.APIVersion, 108) // force icsp to 108 version since icsp version doesn't matter
	asc = asc.NewByName("profile_templates.go")
	if !asc.IsSupported(currentversion) {
		log.Debugf("ProfileTemplates client version not supported: %+v", currentversion)
		return true
	}
	return false
}

// IsProfileTemplates - returns true when we should use GetProfileTemplate...
func (c *OVClient) IsProfileTemplates() bool {
	return !c.ProfileTemplatesNotSupported()
}

// get a server profile template by name
func (c *OVClient) GetProfileTemplateByName(name string) (ServerProfile, error) {
	var (
		profile ServerProfile
	)
	// v2 way to get ServerProfile
	if c.IsProfileTemplates() {
		profiles, err := c.GetProfileTemplates(fmt.Sprintf("name matches '%s'", name), "name:asc")
		if profiles.Total > 0 {
			return profiles.Members[0], err
		} else {
			return profile, err
		}
	} else {

		// v1 way to get a ServerProfile
		profiles, err := c.GetProfiles(fmt.Sprintf("name matches '%s'", name), "name:asc")
		if profiles.Total > 0 {
			return profiles.Members[0], err
		} else {
			return profile, err
		}
	}

}

// get a server profiles
func (c *OVClient) GetProfileTemplates(filter string, sort string) (ServerProfileList, error) {
	var (
		uri      = "/rest/server-profile-templates"
		q        map[string]interface{}
		profiles ServerProfileList
	)
	q = make(map[string]interface{})
	if filter != "" {
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
		return profiles, err
	}

	log.Debugf("GetProfileTemplates %s", data)
	if err := json.Unmarshal([]byte(data), &profiles); err != nil {
		return profiles, err
	}
	return profiles, nil
}
