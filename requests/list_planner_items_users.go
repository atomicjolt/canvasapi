package requests

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListPlannerItemsUsers Retrieve the paginated list of objects to be shown on the planner for the
// current user with the associated planner override to override an item's
// visibility if set.
//
// Planner items for a student may also be retrieved by a linked observer. Use
// the path that accepts a user_id and supply the student's id.
// https://canvas.instructure.com/doc/api/planner.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
// Query Parameters:
// # Query.StartDate (Optional) Only return items starting from the given date.
//    The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # Query.EndDate (Optional) Only return items up to the given date.
//    The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # Query.ContextCodes (Optional) List of context codes of courses and/or groups whose items you want to see.
//    If not specified, defaults to all contexts associated to the current user.
//    Note that concluded courses will be ignored unless specified in the includes[]
//    parameter. The format of this field is the context type, followed by an underscore,
//    followed by the context id. For example: course_42, group_123
// # Query.Filter (Optional) . Must be one of new_activityOnly return items that have new or unread activity
//
type ListPlannerItemsUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StartDate    time.Time `json:"start_date" url:"start_date,omitempty"`       //  (Optional)
		EndDate      time.Time `json:"end_date" url:"end_date,omitempty"`           //  (Optional)
		ContextCodes []string  `json:"context_codes" url:"context_codes,omitempty"` //  (Optional)
		Filter       string    `json:"filter" url:"filter,omitempty"`               //  (Optional) . Must be one of new_activity
	} `json:"query"`
}

func (t *ListPlannerItemsUsers) GetMethod() string {
	return "GET"
}

func (t *ListPlannerItemsUsers) GetURLPath() string {
	path := "users/{user_id}/planner/items"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListPlannerItemsUsers) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListPlannerItemsUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPlannerItemsUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPlannerItemsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Query.Filter != "" && !string_utils.Include([]string{"new_activity"}, t.Query.Filter) {
		errs = append(errs, "Filter must be one of new_activity")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPlannerItemsUsers) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
