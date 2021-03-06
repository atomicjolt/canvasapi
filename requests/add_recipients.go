package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// AddRecipients Add recipients to an existing group conversation. Response is similar to
// the GET/show action, except that only includes the
// latest message (e.g. "joe was added to the conversation by bob")
// https://canvas.instructure.com/doc/api/conversations.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Recipients (Required) An array of recipient ids. These may be user ids or course/group ids
//    prefixed with "course_" or "group_" respectively, e.g.
//    recipients[]=1&recipients[]=2&recipients[]=course_3
//
type AddRecipients struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Recipients []string `json:"recipients" url:"recipients,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *AddRecipients) GetMethod() string {
	return "POST"
}

func (t *AddRecipients) GetURLPath() string {
	path := "conversations/{id}/add_recipients"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AddRecipients) GetQuery() (string, error) {
	return "", nil
}

func (t *AddRecipients) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AddRecipients) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AddRecipients) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.Recipients == nil {
		errs = append(errs, "'Form.Recipients' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddRecipients) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
