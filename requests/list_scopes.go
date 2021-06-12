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

// ListScopes A list of scopes that can be applied to developer keys and access tokens.
// https://canvas.instructure.com/doc/api/api_token_scopes.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # GroupBy (Optional) . Must be one of resource_nameThe attribute to group the scopes by. By default no grouping is done.
//
type ListScopes struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		GroupBy string `json:"group_by"` //  (Optional) . Must be one of resource_name
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListScopes) GetBody() (string, error) {
	return "", nil
}

func (t *ListScopes) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if !string_utils.Include([]string{"resource_name"}, t.Query.GroupBy) {
		errs = append(errs, "GroupBy must be one of resource_name")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListScopes) Do(c *canvasapi.Canvas) ([]*models.Scope, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Scope{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
