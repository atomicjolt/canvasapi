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

// ListGroupsAvailableInContextAccounts Returns the paginated list of active groups in the given context that are visible to user.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.OnlyOwnGroups (Optional) Will only include groups that the user belongs to if this is set
// # Query.Include (Optional) . Must be one of tabs- "tabs": Include the list of tabs configured for each group.  See the
//      {api:TabsController#index List available tabs API} for more information.
//
type ListGroupsAvailableInContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		OnlyOwnGroups bool     `json:"only_own_groups" url:"only_own_groups,omitempty"` //  (Optional)
		Include       []string `json:"include" url:"include,omitempty"`                 //  (Optional) . Must be one of tabs
	} `json:"query"`
}

func (t *ListGroupsAvailableInContextAccounts) GetMethod() string {
	return "GET"
}

func (t *ListGroupsAvailableInContextAccounts) GetURLPath() string {
	path := "accounts/{account_id}/groups"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListGroupsAvailableInContextAccounts) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListGroupsAvailableInContextAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGroupsAvailableInContextAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGroupsAvailableInContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"tabs"}, v) {
			errs = append(errs, "Include must be one of tabs")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupsAvailableInContextAccounts) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Group, *canvasapi.PagedResource, error) {
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
	ret := []*models.Group{}
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
