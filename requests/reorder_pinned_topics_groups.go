package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ReorderPinnedTopicsGroups Puts the pinned discussion topics in the specified order.
// All pinned topics should be included.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Form Parameters:
// # Order (Required) The ids of the pinned discussion topics in the desired order.
//    (For example, "order=104,102,103".)
//
type ReorderPinnedTopicsGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Order []int64 `json:"order"` //  (Required)
	} `json:"form"`
}

func (t *ReorderPinnedTopicsGroups) GetMethod() string {
	return "POST"
}

func (t *ReorderPinnedTopicsGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/reorder"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ReorderPinnedTopicsGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ReorderPinnedTopicsGroups) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ReorderPinnedTopicsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Form.Order == nil {
		errs = append(errs, "'Order' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReorderPinnedTopicsGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
