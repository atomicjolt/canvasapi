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

// ListObservees A paginated list of the users that the given user is observing.
//
// *Note:* all users are allowed to list their own observees. Administrators can list
// other users' observees.
//
// The returned observees will include an attribute "observation_link_root_account_ids", a list
// of ids for the root accounts the observer and observee are linked on. The observer will only be able to
// observe in courses associated with these root accounts.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of avatar_url- "avatar_url": Optionally include avatar_url.
//
type ListObservees struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of avatar_url
	} `json:"query"`
}

func (t *ListObservees) GetMethod() string {
	return "GET"
}

func (t *ListObservees) GetURLPath() string {
	path := "users/{user_id}/observees"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListObservees) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListObservees) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListObservees) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListObservees) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
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

func (t *ListObservees) Do(c *canvasapi.Canvas) ([]*models.User, *canvasapi.PagedResource, error) {
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
