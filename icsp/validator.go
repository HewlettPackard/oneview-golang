package icsp

// URLEndPoint export this constant
const URLEndPointValidator = "/rest/authz/validator"

// Authz struct ...
type Authz struct {
	authorized string `json:"authorized,omitempty"`
}
