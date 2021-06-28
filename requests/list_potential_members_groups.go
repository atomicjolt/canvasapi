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

// ListPotentialMembersGroups A paginated list of the users who can potentially be added to a
// collaboration in the given context.
//
// For courses, this consists of all enrolled users.  For groups, it is comprised of the
// group members plus the admins of the course containing the group.
// https://canvas.instructure.com/doc/api/collaborations.html
//
// Path Parameters:
// # GroupID (Required) ID
//
type ListPotentialMembersGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListPotentialMembersGroups) GetMethod() string {
	return "GET"
}

func (t *ListPotentialMembersGroups) GetURLPath() string {
	path := "groups/{group_id}/potential_collaborators"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListPotentialMembersGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListPotentialMembersGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPotentialMembersGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPotentialMembersGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPotentialMembersGroups) Do(c *canvasapi.Canvas) ([]*models.User, error) {
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
