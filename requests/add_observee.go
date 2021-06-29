package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// AddObservee Registers a user as being observed by the given user.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.ObserveeID (Required) ID
//
// Form Parameters:
// # Form.RootAccountID (Optional) The ID for the root account to associate with the observation link.
//    If not specified, a link will be created for each root account associated
//    to both the observer and observee.
//
type AddObservee struct {
	Path struct {
		UserID     string `json:"user_id" url:"user_id,omitempty"`         //  (Required)
		ObserveeID string `json:"observee_id" url:"observee_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		RootAccountID int64 `json:"root_account_id" url:"root_account_id,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *AddObservee) GetMethod() string {
	return "PUT"
}

func (t *AddObservee) GetURLPath() string {
	path := "users/{user_id}/observees/{observee_id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{observee_id}", fmt.Sprintf("%v", t.Path.ObserveeID))
	return path
}

func (t *AddObservee) GetQuery() (string, error) {
	return "", nil
}

func (t *AddObservee) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AddObservee) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AddObservee) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.ObserveeID == "" {
		errs = append(errs, "'Path.ObserveeID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddObservee) Do(c *canvasapi.Canvas) (*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
