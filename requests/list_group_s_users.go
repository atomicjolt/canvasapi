package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # GroupID (Required) ID
//
// Query Parameters:
// # SearchTerm (Optional) The partial name or full ID of the users to match and return in the
//    results list. Must be at least 3 characters.
// # Include (Optional) . Must be one of avatar_url"avatar_url": Include users' avatar_urls.
// # ExcludeInactive (Optional) Whether to filter out inactive users from the results. Defaults to
//    false unless explicitly provided.
//
type ListGroupSUsers struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm      string   `json:"search_term"`      //  (Optional)
		Include         []string `json:"include"`          //  (Optional) . Must be one of avatar_url
		ExcludeInactive bool     `json:"exclude_inactive"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListGroupSUsers) GetBody() (string, error) {
	return "", nil
}

func (t *ListGroupSUsers) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"avatar_url"}, v) {
			errs = append(errs, "Include must be one of avatar_url")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupSUsers) Do(c *canvasapi.Canvas) ([]*models.User, error) {
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