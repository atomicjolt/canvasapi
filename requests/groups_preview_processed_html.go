package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GroupsPreviewProcessedHtml Preview html content processed for this group
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Form Parameters:
// # Html (Optional) The html content to process
//
type GroupsPreviewProcessedHtml struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Html string `json:"html"` //  (Optional)
	} `json:"form"`
}

func (t *GroupsPreviewProcessedHtml) GetMethod() string {
	return "POST"
}

func (t *GroupsPreviewProcessedHtml) GetURLPath() string {
	path := "groups/{group_id}/preview_html"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GroupsPreviewProcessedHtml) GetQuery() (string, error) {
	return "", nil
}

func (t *GroupsPreviewProcessedHtml) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *GroupsPreviewProcessedHtml) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GroupsPreviewProcessedHtml) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
