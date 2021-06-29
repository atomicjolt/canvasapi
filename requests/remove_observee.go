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

// RemoveObservee Unregisters a user as being observed by the given user.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.ObserveeID (Required) ID
//
// Query Parameters:
// # Query.RootAccountID (Optional) If specified, only removes the link for the given root account
//
type RemoveObservee struct {
	Path struct {
		UserID     string `json:"user_id" url:"user_id,omitempty"`         //  (Required)
		ObserveeID string `json:"observee_id" url:"observee_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		RootAccountID int64 `json:"root_account_id" url:"root_account_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *RemoveObservee) GetMethod() string {
	return "DELETE"
}

func (t *RemoveObservee) GetURLPath() string {
	path := "users/{user_id}/observees/{observee_id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{observee_id}", fmt.Sprintf("%v", t.Path.ObserveeID))
	return path
}

func (t *RemoveObservee) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RemoveObservee) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveObservee) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveObservee) HasErrors() error {
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

func (t *RemoveObservee) Do(c *canvasapi.Canvas) (*models.User, error) {
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
