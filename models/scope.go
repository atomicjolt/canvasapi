package models

type Scope struct {
	Resource     string `json:"resource"`      // The resource the scope is associated with.Example: courses
	ResourceName string `json:"resource_name"` // The localized resource name.Example: Courses
	Controller   string `json:"controller"`    // The controller the scope is associated to.Example: courses
	Action       string `json:"action"`        // The controller action the scope is associated to.Example: index
	Verb         string `json:"verb"`          // The HTTP verb for the scope.Example: GET
	Scope        string `json:"scope"`         // The identifier for the scope.Example: url:GET|/api/v1/courses
}

func (t *Scope) HasError() error {
	return nil
}
