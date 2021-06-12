package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleGroupMembershipUsers Returns the group membership with the given membership id or user id.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
// # UserID (Required) ID
//
type GetSingleGroupMembershipUsers struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
		UserID  string `json:"user_id"`  //  (Required)
	} `json:"path"`
}

func (t *GetSingleGroupMembershipUsers) GetMethod() string {
	return "GET"
}

func (t *GetSingleGroupMembershipUsers) GetURLPath() string {
	path := "groups/{group_id}/users/{user_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *GetSingleGroupMembershipUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleGroupMembershipUsers) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleGroupMembershipUsers) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleGroupMembershipUsers) Do(c *canvasapi.Canvas) (*models.GroupMembership, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.GroupMembership{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}