package rest

import "strings"

// Helpers
func (c *Client) Sanatize(s string) (string) {
	if strings.LastIndex(s, "/") > 0 {
		s = strings.Trim(s, "/")
	}
	return s
}
