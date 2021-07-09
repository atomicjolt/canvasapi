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

// ListGroupMemberships A paginated list of the members of a group.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
// Query Parameters:
// # Query.FilterStates (Optional) . Must be one of accepted, invited, requestedOnly list memberships with the given workflow_states. By default it will
//    return all memberships.
//
type ListGroupMemberships struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		FilterStates []string `json:"filter_states" url:"filter_states,omitempty"` //  (Optional) . Must be one of accepted, invited, requested
	} `json:"query"`
}

func (t *ListGroupMemberships) GetMethod() string {
	return "GET"
}

func (t *ListGroupMemberships) GetURLPath() string {
	path := "groups/{group_id}/memberships"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListGroupMemberships) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListGroupMemberships) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGroupMemberships) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGroupMemberships) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	for _, v := range t.Query.FilterStates {
		if v != "" && !string_utils.Include([]string{"accepted", "invited", "requested"}, v) {
			errs = append(errs, "FilterStates must be one of accepted, invited, requested")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupMemberships) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.GroupMembership, *canvasapi.PagedResource, error) {
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
	ret := []*models.GroupMembership{}
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
