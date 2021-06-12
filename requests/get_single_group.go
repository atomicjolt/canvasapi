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

// GetSingleGroup Returns the data for a single group, or a 401 if the caller doesn't have
// the rights to see it.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of permissions, tabs- "permissions": Include permissions the current user has
//      for the group.
//    - "tabs": Include the list of tabs configured for each group.  See the
//      {api:TabsController#index List available tabs API} for more information.
//
type GetSingleGroup struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of permissions, tabs
	} `json:"query"`
}

func (t *GetSingleGroup) GetMethod() string {
	return "GET"
}

func (t *GetSingleGroup) GetURLPath() string {
	path := "groups/{group_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GetSingleGroup) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSingleGroup) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleGroup) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"permissions", "tabs"}, v) {
			errs = append(errs, "Include must be one of permissions, tabs")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleGroup) Do(c *canvasapi.Canvas) (*models.Group, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Group{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
