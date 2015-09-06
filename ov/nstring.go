package ov

import (
	"encoding/json"
)

type Nstring string

func (n *Nstring) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, (*string)(n))
}
