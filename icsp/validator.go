package icsp

import "strconv"

// URLEndPoint ...
type URLEndPoint string

// URLEndPoint export this constant
const URLEndPoint = "/rest/authz/validator"

// AuthHeader Marshal a json into a auth header
type AuthHeader struct {
	ContentType string `json:"Content-Type,omitempty"`
	XAPIVersion int    `json:"X-API-Version,omitempty"`
	Auth        string `json:"auth,omitempty"`
}

// GetAuthHeaderMap Generate an auth Header map ...
func (c *OVClient) GetAuthHeaderMap() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json; charset=utf-8",
		"X-API-Version": strconv.Itoa(c.APIVersion),
		"auth":          c.APIKey,
	}
}

// Authz struct ...
type Authz struct {
	authorized string `json:"authorized,omitempty"`
}
