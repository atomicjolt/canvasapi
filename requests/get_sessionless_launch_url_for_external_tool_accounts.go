package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSessionlessLaunchUrlForExternalToolAccounts Returns a sessionless launch url for an external tool.
//
// NOTE: Either the id or url must be provided unless launch_type is assessment or module_item.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # ID (Optional) The external id of the tool to launch.
// # Url (Optional) The LTI launch url for the external tool.
// # AssignmentID (Optional) The assignment id for an assignment launch. Required if launch_type is set to "assessment".
// # ModuleItemID (Optional) The assignment id for a module item launch. Required if launch_type is set to "module_item".
// # LaunchType (Optional) . Must be one of assessment, module_itemThe type of launch to perform on the external tool. Placement names (eg. "course_navigation")
//    can also be specified to use the custom launch url for that placement; if done, the tool id
//    must be provided.
//
type GetSessionlessLaunchUrlForExternalToolAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		ID           string `json:"id" url:"id,omitempty"`                         //  (Optional)
		Url          string `json:"url" url:"url,omitempty"`                       //  (Optional)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"`   //  (Optional)
		ModuleItemID string `json:"module_item_id" url:"module_item_id,omitempty"` //  (Optional)
		LaunchType   string `json:"launch_type" url:"launch_type,omitempty"`       //  (Optional) . Must be one of assessment, module_item
	} `json:"query"`
}

func (t *GetSessionlessLaunchUrlForExternalToolAccounts) GetMethod() string {
	return "GET"
}

func (t *GetSessionlessLaunchUrlForExternalToolAccounts) GetURLPath() string {
	path := "accounts/{account_id}/external_tools/sessionless_launch"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetSessionlessLaunchUrlForExternalToolAccounts) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSessionlessLaunchUrlForExternalToolAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSessionlessLaunchUrlForExternalToolAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSessionlessLaunchUrlForExternalToolAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Query.LaunchType != "" && !string_utils.Include([]string{"assessment", "module_item"}, t.Query.LaunchType) {
		errs = append(errs, "LaunchType must be one of assessment, module_item")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSessionlessLaunchUrlForExternalToolAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
