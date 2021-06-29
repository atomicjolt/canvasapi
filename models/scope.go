package models

type Scope struct {
	Resource     string `json:"resource" url:"resource,omitempty"`           // The resource the scope is associated with.Example: courses
	ResourceName string `json:"resource_name" url:"resource_name,omitempty"` // The localized resource name.Example: Courses
	Controller   string `json:"controller" url:"controller,omitempty"`       // The controller the scope is associated to.Example: courses
	Action       string `json:"action" url:"action,omitempty"`               // The controller action the scope is associated to.Example: index
	Verb         string `json:"verb" url:"verb,omitempty"`                   // The HTTP verb for the scope.Example: GET
	Scope        string `json:"scope" url:"scope,omitempty"`                 // The identifier for the scope.Example: url:GET|/api/v1/courses
}

func (t *Scope) HasErrors() error {
	return nil
}
