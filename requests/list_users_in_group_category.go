package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListUsersInGroupCategory Returns a paginated list of users in the group category.
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # GroupCategoryID (Required) ID
//
// Query Parameters:
// # SearchTerm (Optional) The partial name or full ID of the users to match and return in the results
//    list. Must be at least 3 characters.
// # Unassigned (Optional) Set this value to true if you wish only to search unassigned users in the
//    group category.
//
type ListUsersInGroupCategory struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm string `json:"search_term"` //  (Optional)
		Unassigned bool   `json:"unassigned"`  //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListUsersInGroupCategory) GetBody() (string, error) {
	return "", nil
}

func (t *ListUsersInGroupCategory) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListUsersInGroupCategory) Do(c *canvasapi.Canvas) ([]*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
