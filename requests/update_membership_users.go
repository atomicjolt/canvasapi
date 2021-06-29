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

// UpdateMembershipUsers Accept a membership request, or add/remove moderator rights.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.UserID (Required) ID
//
// Form Parameters:
// # Form.WorkflowState (Optional) . Must be one of acceptedCurrently, the only allowed value is "accepted"
// # Form.Moderator (Optional) no description
//
type UpdateMembershipUsers struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		UserID  string `json:"user_id" url:"user_id,omitempty"`   //  (Required)
	} `json:"path"`

	Form struct {
		WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` //  (Optional) . Must be one of accepted
		Moderator     string `json:"moderator" url:"moderator,omitempty"`           //  (Optional)
	} `json:"form"`
}

func (t *UpdateMembershipUsers) GetMethod() string {
	return "PUT"
}

func (t *UpdateMembershipUsers) GetURLPath() string {
	path := "groups/{group_id}/users/{user_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *UpdateMembershipUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMembershipUsers) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateMembershipUsers) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateMembershipUsers) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Form.WorkflowState != "" && !string_utils.Include([]string{"accepted"}, t.Form.WorkflowState) {
		errs = append(errs, "WorkflowState must be one of accepted")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMembershipUsers) Do(c *canvasapi.Canvas) (*models.GroupMembership, error) {
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
