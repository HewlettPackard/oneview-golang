package ov

import (
_	"github.com/HewlettPackard/oneview-golang/utils"
)

type Configuration struct {
	UserName   string        `json:"username"`
	Password   string        `json:"password"`
	Endpoint   string `json:"endpoint"`
	Domain     string        `json:"domain"`
	ApiVersion int             `json:"apiversion"`
}
