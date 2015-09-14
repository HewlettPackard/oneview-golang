package icsp

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/docker/machine/drivers/oneview/rest"
	"github.com/docker/machine/log"
)

// URLEndPoint export this constant
const URLEndPointSession = "/rest/login-sessions"

// AuthHeader Marshal a json into a auth header
/*type AuthHeader struct {
	ContentType string `json:"Content-Type,omitempty"`
	XAPIVersion int    `json:"X-API-Version,omitempty"`
	Auth        string `json:"auth,omitempty"`
}
*/

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
