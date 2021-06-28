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
)

// ShowRevisionGroupsLatest Retrieve the metadata and optionally content of a revision of the page.
// Note that retrieving historic versions of pages requires edit rights.
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # GroupID (Required) ID
// # Url (Required) ID
//
// Query Parameters:
// # Summary (Optional) If set, exclude page content from results
//
type ShowRevisionGroupsLatest struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		Url     string `json:"url" url:"url,omitempty"`           //  (Required)
	} `json:"path"`

	Query struct {
		Summary bool `json:"summary" url:"summary,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ShowRevisionGroupsLatest) GetMethod() string {
	return "GET"
}

func (t *ShowRevisionGroupsLatest) GetURLPath() string {
	path := "groups/{group_id}/pages/{url}/revisions/latest"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{url}", fmt.Sprintf("%v", t.Path.Url))
	return path
}

func (t *ShowRevisionGroupsLatest) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ShowRevisionGroupsLatest) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowRevisionGroupsLatest) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowRevisionGroupsLatest) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Url' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowRevisionGroupsLatest) Do(c *canvasapi.Canvas) (*models.PageRevision, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PageRevision{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
