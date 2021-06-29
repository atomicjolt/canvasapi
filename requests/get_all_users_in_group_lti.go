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

// GetAllUsersInGroupLti Get all Canvas users in a group. Tool providers may only access
// groups that belong to the context the tool is installed in.
// https://canvas.instructure.com/doc/api/plagiarism_detection_platform_users.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type GetAllUsersInGroupLti struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetAllUsersInGroupLti) GetMethod() string {
	return "GET"
}

func (t *GetAllUsersInGroupLti) GetURLPath() string {
	path := "/lti/groups/{group_id}/users"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GetAllUsersInGroupLti) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAllUsersInGroupLti) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllUsersInGroupLti) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllUsersInGroupLti) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllUsersInGroupLti) Do(c *canvasapi.Canvas) ([]*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
