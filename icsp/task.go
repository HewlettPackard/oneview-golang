package icsp

import "github.com/docker/machine/drivers/oneview/utils"

// ODSUri  returned from create server for job uri task
type ODSUri struct {
	URI utils.Nstring `json:"uri,omitempty"` // uri of job
}
