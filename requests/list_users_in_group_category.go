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
)

// ListUsersInGroupCategory Returns a paginated list of users in the group category.
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # Path.GroupCategoryID (Required) ID
//
// Query Parameters:
// # Query.SearchTerm (Optional) The partial name or full ID of the users to match and return in the results
//    list. Must be at least 3 characters.
// # Query.Unassigned (Optional) Set this value to true if you wish only to search unassigned users in the
//    group category.
//
type ListUsersInGroupCategory struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm string `json:"search_term" url:"search_term,omitempty"` //  (Optional)
		Unassigned bool   `json:"unassigned" url:"unassigned,omitempty"`   //  (Optional)
	} `json:"query"`
}

func (t *ListUsersInGroupCategory) GetMethod() string {
	return "GET"
}

func (t *ListUsersInGroupCategory) GetURLPath() string {
	path := "group_categories/{group_category_id}/users"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *ListUsersInGroupCategory) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListUsersInGroupCategory) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListUsersInGroupCategory) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListUsersInGroupCategory) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'Path.GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListUsersInGroupCategory) Do(c *canvasapi.Canvas) ([]*models.User, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.User{}
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
