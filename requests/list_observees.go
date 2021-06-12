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
// # UserID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of avatar_url- "avatar_url": Optionally include avatar_url.
//
type ListObservees struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of avatar_url
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListObservees) GetBody() (string, error) {
	return "", nil
}

func (t *ListObservees) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
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

func (t *ListObservees) Do(c *canvasapi.Canvas) ([]*models.User, error) {
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
