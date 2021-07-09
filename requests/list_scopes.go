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

// ListScopes A list of scopes that can be applied to developer keys and access tokens.
// https://canvas.instructure.com/doc/api/api_token_scopes.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.GroupBy (Optional) . Must be one of resource_nameThe attribute to group the scopes by. By default no grouping is done.
//
type ListScopes struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		GroupBy string `json:"group_by" url:"group_by,omitempty"` //  (Optional) . Must be one of resource_name
	} `json:"query"`
}

func (t *ListScopes) GetMethod() string {
	return "GET"
}

func (t *ListScopes) GetURLPath() string {
	path := "accounts/{account_id}/scopes"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListScopes) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListScopes) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListScopes) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListScopes) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Query.GroupBy != "" && !string_utils.Include([]string{"resource_name"}, t.Query.GroupBy) {
		errs = append(errs, "GroupBy must be one of resource_name")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListScopes) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Scope, *canvasapi.PagedResource, error) {
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
	ret := []*models.Scope{}
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
