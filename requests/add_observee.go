package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// AddObservee Registers a user as being observed by the given user.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # UserID (Required) ID
// # ObserveeID (Required) ID
//
// Form Parameters:
// # RootAccountID (Optional) The ID for the root account to associate with the observation link.
//    If not specified, a link will be created for each root account associated
//    to both the observer and observee.
//
type AddObservee struct {
	Path struct {
		UserID     string `json:"user_id"`     //  (Required)
		ObserveeID string `json:"observee_id"` //  (Required)
	} `json:"path"`

	Form struct {
		RootAccountID int64 `json:"root_account_id"` //  (Optional)
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

func (t *AddObservee) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *AddObservee) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ObserveeID == "" {
		errs = append(errs, "'ObserveeID' is required")
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
