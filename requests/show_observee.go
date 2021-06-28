package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowObservee Gets information about an observed user.
//
// *Note:* all users are allowed to view their own observees.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # UserID (Required) ID
// # ObserveeID (Required) ID
//
type ShowObservee struct {
	Path struct {
		UserID     string `json:"user_id" url:"user_id,omitempty"`         //  (Required)
		ObserveeID string `json:"observee_id" url:"observee_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ShowObservee) GetMethod() string {
	return "GET"
}

func (t *ShowObservee) GetURLPath() string {
	path := "users/{user_id}/observees/{observee_id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{observee_id}", fmt.Sprintf("%v", t.Path.ObserveeID))
	return path
}

func (t *ShowObservee) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowObservee) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowObservee) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowObservee) HasErrors() error {
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

func (t *ShowObservee) Do(c *canvasapi.Canvas) (*models.User, error) {
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
