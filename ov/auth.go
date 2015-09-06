package ov

import (
	"strconv"
	"encoding/json"
	"github.com/docker/machine/log"
	"github.com/docker/machine/drivers/oneview/rest"
)

// Marshel a json into a auth header
type AuthHeader struct {
	ContentType   string `json:"Content-Type,omitempty"`
	XAPIVersion   int    `json:"X-API-Version,omitempty"`
	auth          string `json:"auth,omitempty"`
}
// Generate an auth Header map
func (c *OVClient) GetAuthHeaderMap() map[string]string {
	return map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"X-API-Version": strconv.Itoa(c.APIVersion),
		"auth": c.APIKey,
	}
}

// Session struct
type Session struct {
	ID string `json:"sessionID,omitempty"`
}

// auth structure
type Auth struct {
	UserName  string `json:"userName,omitempty"`
	Password  string `json:"password,omitempty"`
	Domain    string `json:"authLoginDomain,omitempty"`
}

// Refresh login authkey
// Should make sure we have a valid APIKey
func (c *OVClient) RefreshLogin() (error) {
	if c.APIKey == "" || c.APIKey == "none"	{
		log.Debugf("Getting new session id")
	  s, err := c.SessionLogin()
		if err != nil {
			return err
		}
	  c.APIKey = s.ID
	}
	// TODO: validate that session id is valid, if not refresh it
	return nil
}

// Login to OneView and get a session ID
// returns Session structure
func (c *OVClient) SessionLogin() (Session, error) {
	var (
		uri    = "/rest/login-sessions"
		body   = Auth{UserName: c.User, Password: c.Password, Domain: c.Domain}
		session Session
	)

	c.SetAuthHeaderOptions( c.GetAuthHeaderMap() )
	data, err := c.RestAPICall(rest.POST, uri , body)
	if err != nil {
		return session, err
	}

	log.Debugf("SessionLogin %s", data)
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return session, err
	}
	// Update APIKey
  return session, err
}
