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

// GetSingleGroupMembershipMemberships Returns the group membership with the given membership id or user id.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
// # MembershipID (Required) ID
//
type GetSingleGroupMembershipMemberships struct {
	Path struct {
		GroupID      string `json:"group_id" url:"group_id,omitempty"`           //  (Required)
		MembershipID string `json:"membership_id" url:"membership_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleGroupMembershipMemberships) GetMethod() string {
	return "GET"
}

func (t *GetSingleGroupMembershipMemberships) GetURLPath() string {
	path := "groups/{group_id}/memberships/{membership_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{membership_id}", fmt.Sprintf("%v", t.Path.MembershipID))
	return path
}

func (t *GetSingleGroupMembershipMemberships) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleGroupMembershipMemberships) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleGroupMembershipMemberships) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleGroupMembershipMemberships) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.MembershipID == "" {
		errs = append(errs, "'MembershipID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleGroupMembershipMemberships) Do(c *canvasapi.Canvas) (*models.GroupMembership, error) {
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
