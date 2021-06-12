package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListExternalToolsGroups Returns the paginated list of external tools for the current context.
// See the get request docs for a single tool for a list of properties on an external tool.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Query Parameters:
// # SearchTerm (Optional) The partial name of the tools to match and return.
// # Selectable (Optional) If true, then only tools that are meant to be selectable are returned
// # IncludeParents (Optional) If true, then include tools installed in all accounts above the current context
//
type ListExternalToolsGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm     string `json:"search_term"`     //  (Optional)
		Selectable     bool   `json:"selectable"`      //  (Optional)
		IncludeParents bool   `json:"include_parents"` //  (Optional)
	} `json:"query"`
}

func (t *ListExternalToolsGroups) GetMethod() string {
	return "GET"
}

func (t *ListExternalToolsGroups) GetURLPath() string {
	path := "groups/{group_id}/external_tools"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListExternalToolsGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListExternalToolsGroups) GetBody() (string, error) {
	return "", nil
}

func (t *ListExternalToolsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListExternalToolsGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
