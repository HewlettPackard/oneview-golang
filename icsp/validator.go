package icsp

// URLEndPoint export this constant
const URLEndPointValidator = "/rest/authz/validator"

// Authz struct ...
type Authz struct {
	authorized string `json:"authorized,omitempty"`
}

type CategoryAction struct {
	actionDto   string `json:"actionDto,omitempty"`   // actionDto - the action name
	categoryDto string `json:"categoryDto,omitempty"` // categoryDto - the category name
}
