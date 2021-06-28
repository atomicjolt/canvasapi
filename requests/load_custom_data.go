package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// LoadCustomData Load custom user data.
//
// Arbitrary JSON data can be stored for a User.  This API call
// retrieves that data for a (optional) given scope.
// See {api:UsersController#set_custom_data Store Custom Data} for details and
// examples.
//
// On success, this endpoint returns an object containing the data that was requested.
//
// Responds with status code 400 if the namespace parameter, +ns+, is missing or invalid,
// or if the specified scope does not contain any data.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # UserID (Required) ID
//
// Query Parameters:
// # Ns (Required) The namespace from which to retrieve the data.  This should be something other
//    Canvas API apps aren't likely to use, such as a reverse DNS for your organization.
//
type LoadCustomData struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Ns string `json:"ns" url:"ns,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *LoadCustomData) GetMethod() string {
	return "GET"
}

func (t *LoadCustomData) GetURLPath() string {
	path := "users/{user_id}/custom_data"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *LoadCustomData) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *LoadCustomData) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *LoadCustomData) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *LoadCustomData) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Query.Ns == "" {
		errs = append(errs, "'Ns' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *LoadCustomData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
