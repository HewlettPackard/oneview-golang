package icsp

// URLEndPoint export this constant
const URLEndPointValidator = "/rest/authz/validator"

// Authz struct ...
type Authz struct {
	Authorized string `json:"authorized,omitempty"`
}

type CategoryAction struct {
	ActionDto   string `json:"actionDto,omitempty"`   // actionDto - the action name
	CategoryDto string `json:"categoryDto,omitempty"` // categoryDto - the category name
}
