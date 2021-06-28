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

// CreateObserverPairingCode If the user is a student, will generate a code to be used with self registration
// or observees APIs to link another user to this student.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # UserID (Required) ID
//
type CreateObserverPairingCode struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *CreateObserverPairingCode) GetMethod() string {
	return "POST"
}

func (t *CreateObserverPairingCode) GetURLPath() string {
	path := "users/{user_id}/observer_pairing_codes"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *CreateObserverPairingCode) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateObserverPairingCode) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CreateObserverPairingCode) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CreateObserverPairingCode) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateObserverPairingCode) Do(c *canvasapi.Canvas) (*models.PairingCode, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PairingCode{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
