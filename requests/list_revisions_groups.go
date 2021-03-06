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

// ListRevisionsGroups A paginated list of the revisions of a page. Callers must have update rights on the page in order to see page history.
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.Url (Required) ID
//
type ListRevisionsGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		Url     string `json:"url" url:"url,omitempty"`           //  (Required)
	} `json:"path"`
}

func (t *ListRevisionsGroups) GetMethod() string {
	return "GET"
}

func (t *ListRevisionsGroups) GetURLPath() string {
	path := "groups/{group_id}/pages/{url}/revisions"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{url}", fmt.Sprintf("%v", t.Path.Url))
	return path
}

func (t *ListRevisionsGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListRevisionsGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListRevisionsGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListRevisionsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Path.Url' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRevisionsGroups) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.PageRevision, *canvasapi.PagedResource, error) {
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
	ret := []*models.PageRevision{}
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
