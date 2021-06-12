package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// InviteOthersToGroup Sends an invitation to all supplied email addresses which will allow the
// receivers to join the group.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Form Parameters:
// # Invitees (Required) An array of email addresses to be sent invitations.
//
type InviteOthersToGroup struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Invitees []string `json:"invitees"` //  (Required)
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

func (t *InviteOthersToGroup) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *InviteOthersToGroup) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Form.Invitees == nil {
		errs = append(errs, "'Invitees' is required")
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
