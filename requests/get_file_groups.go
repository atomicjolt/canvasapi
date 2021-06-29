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

// GetFileGroups Returns the standard attachment json object
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # GroupID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of userArray of additional information to include.
//
//    "user":: the user who uploaded the file or last edited its content
//    "usage_rights":: copyright and license information for the file (see UsageRights)
//
type GetFileGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		ID      string `json:"id" url:"id,omitempty"`             //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of user
	} `json:"query"`
}

func (t *GetFileGroups) GetMethod() string {
	return "GET"
}

func (t *GetFileGroups) GetURLPath() string {
	path := "groups/{group_id}/files/{id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetFileGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetFileGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetFileGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetFileGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"user"}, v) {
			errs = append(errs, "Include must be one of user")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFileGroups) Do(c *canvasapi.Canvas) (*models.File, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.File{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
