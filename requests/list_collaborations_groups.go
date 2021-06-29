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

// ListCollaborationsGroups A paginated list of collaborations the current user has access to in the
// context of the course provided in the url. NOTE: this only returns
// ExternalToolCollaboration type collaborations.
//
//   curl https://<canvas>/api/v1/courses/1/collaborations/
// https://canvas.instructure.com/doc/api/collaborations.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type ListCollaborationsGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListCollaborationsGroups) GetMethod() string {
	return "GET"
}

func (t *ListCollaborationsGroups) GetURLPath() string {
	path := "groups/{group_id}/collaborations"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListCollaborationsGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListCollaborationsGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCollaborationsGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCollaborationsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCollaborationsGroups) Do(c *canvasapi.Canvas) ([]*models.Collaboration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Collaboration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
