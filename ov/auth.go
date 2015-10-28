/*
(c) Copyright [2015] Hewlett Packard Enterprise Development LP

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package ov for working with HP OneView
package ov

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/docker/machine/libmachine/log"
)

// AuthHeader Marshal a json into a auth header
type AuthHeader struct {
	ContentType string `json:"Content-Type,omitempty"`
	XAPIVersion int    `json:"X-API-Version,omitempty"`
	Auth        string `json:"auth,omitempty"`
}

// GetAuthHeaderMap Generate an auth Header map
func (c *OVClient) GetAuthHeaderMap() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json; charset=utf-8",
		"X-API-Version": strconv.Itoa(c.APIVersion),
		"auth":          c.APIKey,
	}
}

// GetAuthHeaderMapNoVer generate header without version
func (c *OVClient) GetAuthHeaderMapNoVer() map[string]string {
	return map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"auth":         c.APIKey,
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
func (c *OVClient) RefreshLogin() error {
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

// SessionLogin Login to OneView and get a session ID
// returns Session structure
func (c *OVClient) SessionLogin() (Session, error) {
	var (
		uri     = "/rest/login-sessions"
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
func (c *OVClient) SessionLogout() error {
	var (
		uri = "/rest/login-sessions"
	)
	log.Debugf("Calling logout for header -> %+v", c.GetAuthHeaderMap())
	if c.APIKey == "none" {
		log.Debugf("already logged out")
		return nil
	}
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	_, err := c.RestAPICall(rest.DELETE, uri, nil)
	if err != nil {
		log.Debugf("Error from %s :-> %+v", uri, err)
		return err
	}
	c.APIKey = "none"
	return nil
}
