package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListPagesGroups A paginated list of the wiki pages associated with a course or group
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
// Query Parameters:
// # Query.Sort (Optional) . Must be one of title, created_at, updated_atSort results by this field.
// # Query.Order (Optional) . Must be one of asc, descThe sorting order. Defaults to 'asc'.
// # Query.SearchTerm (Optional) The partial title of the pages to match and return.
// # Query.Published (Optional) If true, include only published paqes. If false, exclude published
//    pages. If not present, do not filter on published status.
//
type ListPagesGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Sort       string `json:"sort" url:"sort,omitempty"`               //  (Optional) . Must be one of title, created_at, updated_at
		Order      string `json:"order" url:"order,omitempty"`             //  (Optional) . Must be one of asc, desc
		SearchTerm string `json:"search_term" url:"search_term,omitempty"` //  (Optional)
		Published  bool   `json:"published" url:"published,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *ListPagesGroups) GetMethod() string {
	return "GET"
}

func (t *ListPagesGroups) GetURLPath() string {
	path := "groups/{group_id}/pages"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListPagesGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListPagesGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPagesGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPagesGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Query.Sort != "" && !string_utils.Include([]string{"title", "created_at", "updated_at"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of title, created_at, updated_at")
	}
	if t.Query.Order != "" && !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
		errs = append(errs, "Order must be one of asc, desc")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPagesGroups) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Page, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Page{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
