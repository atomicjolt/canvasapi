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
	"github.com/atomicjolt/string_utils"
)

// ListGroupSUsers Returns a paginated list of users in the group.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
// Query Parameters:
// # Query.SearchTerm (Optional) The partial name or full ID of the users to match and return in the
//    results list. Must be at least 3 characters.
// # Query.Include (Optional) . Must be one of avatar_url"avatar_url": Include users' avatar_urls.
// # Query.ExcludeInactive (Optional) Whether to filter out inactive users from the results. Defaults to
//    false unless explicitly provided.
//
type ListGroupSUsers struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm      string   `json:"search_term" url:"search_term,omitempty"`           //  (Optional)
		Include         []string `json:"include" url:"include,omitempty"`                   //  (Optional) . Must be one of avatar_url
		ExcludeInactive bool     `json:"exclude_inactive" url:"exclude_inactive,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListGroupSUsers) GetMethod() string {
	return "GET"
}

func (t *ListGroupSUsers) GetURLPath() string {
	path := "groups/{group_id}/users"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListGroupSUsers) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListGroupSUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGroupSUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGroupSUsers) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"avatar_url"}, v) {
			errs = append(errs, "Include must be one of avatar_url")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupSUsers) Do(c *canvasapi.Canvas) ([]*models.User, *canvasapi.PagedResource, error) {
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
