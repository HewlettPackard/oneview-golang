package icsp

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/docker/machine/libmachine/log"
)

// URLEndPoint export this constant
const URLEndPointSession = "/rest/login-sessions"

// GetAuthHeaderMap Generate an auth Header map ...
// some api endpoints are hiddent, remove api version to get to them
func (c *ICSPClient) GetAuthHeaderMapNoVer() map[string]string {
	return map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"auth":         c.APIKey,
	}
}

// GetAuthHeaderMap Generate an auth Header map ...
func (c *ICSPClient) GetAuthHeaderMap() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json; charset=utf-8",
		"X-API-Version": strconv.Itoa(c.APIVersion),
		"auth":          c.APIKey,
	}
}

// Session struct
type Session struct {
	ID string `json:"sessionID,omitempty"`
}

// Auth structure
type Auth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
	Domain   string `json:"authLoginDomain,omitempty"`
}

// RefreshLogin Refresh login authkey
// Should make sure we have a valid APIKey
func (c *ICSPClient) RefreshLogin() error {
	if c.APIKey == "" || len(strings.TrimSpace(c.APIKey)) == 0 || c.APIKey == "none" {
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

// SessionLogin  to OneView and get a session ID
// returns Session structure
func (c *ICSPClient) SessionLogin() (Session, error) {
	var (
		uri     = URLEndPointSession
		body    = Auth{UserName: c.User, Password: c.Password, Domain: c.Domain}
		session Session
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.POST, uri, body)
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

// SessionLogout Logout to OneView and get a session ID
// returns Session structure
func (c *ICSPClient) SessionLogout() error {
	var (
		uri = "/rest/login-sessions"
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	_, err := c.RestAPICall(rest.DELETE, uri, nil)
	if err != nil {
		return err
	}
	c.APIKey = ""
	// successful logout HTTP status 204 (no content)
	return nil
	/*if err := json.Unmarshal([]byte(data), &session); err != nil {
		return session, err
	}
	*/
}
