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

// GetUserProfile Returns user profile data, including user id, name, and profile pic.
//
// When requesting the profile for the user accessing the API, the user's
// calendar feed URL and LTI user id will be returned as well.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # UserID (Required) ID
//
type GetUserProfile struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetUserProfile) GetMethod() string {
	return "GET"
}

func (t *GetUserProfile) GetURLPath() string {
	path := "users/{user_id}/profile"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *GetUserProfile) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUserProfile) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetUserProfile) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetUserProfile) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUserProfile) Do(c *canvasapi.Canvas) (*models.Profile, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Profile{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
