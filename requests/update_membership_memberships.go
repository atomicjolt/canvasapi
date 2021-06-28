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
	"github.com/atomicjolt/string_utils"
)

// UpdateMembershipMemberships Accept a membership request, or add/remove moderator rights.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
// # MembershipID (Required) ID
//
// Form Parameters:
// # WorkflowState (Optional) . Must be one of acceptedCurrently, the only allowed value is "accepted"
// # Moderator (Optional) no description
//
type UpdateMembershipMemberships struct {
	Path struct {
		GroupID      string `json:"group_id" url:"group_id,omitempty"`           //  (Required)
		MembershipID string `json:"membership_id" url:"membership_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` //  (Optional) . Must be one of accepted
		Moderator     string `json:"moderator" url:"moderator,omitempty"`           //  (Optional)
	} `json:"form"`
}

func (t *UpdateMembershipMemberships) GetMethod() string {
	return "PUT"
}

func (t *UpdateMembershipMemberships) GetURLPath() string {
	path := "groups/{group_id}/memberships/{membership_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{membership_id}", fmt.Sprintf("%v", t.Path.MembershipID))
	return path
}

func (t *UpdateMembershipMemberships) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMembershipMemberships) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateMembershipMemberships) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateMembershipMemberships) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.MembershipID == "" {
		errs = append(errs, "'MembershipID' is required")
	}
	if t.Form.WorkflowState != "" && !string_utils.Include([]string{"accepted"}, t.Form.WorkflowState) {
		errs = append(errs, "WorkflowState must be one of accepted")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMembershipMemberships) Do(c *canvasapi.Canvas) (*models.GroupMembership, error) {
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
