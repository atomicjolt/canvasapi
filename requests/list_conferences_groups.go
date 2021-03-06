package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListConferencesGroups Retrieve the paginated list of conferences for this context
//
// This API returns a JSON object containing the list of conferences,
// the key for the list of conferences is "conferences"
// https://canvas.instructure.com/doc/api/conferences.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type ListConferencesGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListConferencesGroups) GetMethod() string {
	return "GET"
}

func (t *ListConferencesGroups) GetURLPath() string {
	path := "groups/{group_id}/conferences"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListConferencesGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListConferencesGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListConferencesGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListConferencesGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListConferencesGroups) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Conference, *canvasapi.PagedResource, error) {
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
	ret := []*models.Conference{}
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
