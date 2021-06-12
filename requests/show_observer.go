package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowObserver Gets information about an observed user.
//
// *Note:* all users are allowed to view their own observers.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # UserID (Required) ID
// # ObserverID (Required) ID
//
type ShowObserver struct {
	Path struct {
		UserID     string `json:"user_id"`     //  (Required)
		ObserverID string `json:"observer_id"` //  (Required)
	} `json:"path"`
}

func (t *ShowObserver) GetMethod() string {
	return "GET"
}

func (t *ShowObserver) GetURLPath() string {
	path := "users/{user_id}/observers/{observer_id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{observer_id}", fmt.Sprintf("%v", t.Path.ObserverID))
	return path
}

func (t *ShowObserver) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowObserver) GetBody() (string, error) {
	return "", nil
}

func (t *ShowObserver) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ObserverID == "" {
		errs = append(errs, "'ObserverID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowObserver) Do(c *canvasapi.Canvas) (*models.User, error) {
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
