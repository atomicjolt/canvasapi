package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// InviteOthersToGroup Sends an invitation to all supplied email addresses which will allow the
// receivers to join the group.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
// Form Parameters:
// # Form.Invitees (Required) An array of email addresses to be sent invitations.
//
type InviteOthersToGroup struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Invitees []string `json:"invitees" url:"invitees,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *InviteOthersToGroup) GetMethod() string {
	return "POST"
}

func (t *InviteOthersToGroup) GetURLPath() string {
	path := "groups/{group_id}/invite"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *InviteOthersToGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *InviteOthersToGroup) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *InviteOthersToGroup) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *InviteOthersToGroup) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Form.Invitees == nil {
		errs = append(errs, "'Form.Invitees' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *InviteOthersToGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
