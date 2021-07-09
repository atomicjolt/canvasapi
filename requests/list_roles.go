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

// ListRoles A paginated list of the roles available to an account.
// https://canvas.instructure.com/doc/api/roles.html
//
// Path Parameters:
// # Path.AccountID (Required) The id of the account to retrieve roles for.
//
// Query Parameters:
// # Query.State (Optional) . Must be one of active, inactiveFilter by role state. If this argument is omitted, only 'active' roles are
//    returned.
// # Query.ShowInherited (Optional) If this argument is true, all roles inherited from parent accounts will
//    be included.
//
type ListRoles struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		State         []string `json:"state" url:"state,omitempty"`                   //  (Optional) . Must be one of active, inactive
		ShowInherited bool     `json:"show_inherited" url:"show_inherited,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListRoles) GetMethod() string {
	return "GET"
}

func (t *ListRoles) GetURLPath() string {
	path := "accounts/{account_id}/roles"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListRoles) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListRoles) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListRoles) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListRoles) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	for _, v := range t.Query.State {
		if v != "" && !string_utils.Include([]string{"active", "inactive"}, v) {
			errs = append(errs, "State must be one of active, inactive")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRoles) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Role, *canvasapi.PagedResource, error) {
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
	ret := []*models.Role{}
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
