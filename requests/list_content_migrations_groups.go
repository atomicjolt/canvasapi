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

// ListContentMigrationsGroups Returns paginated content migrations
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type ListContentMigrationsGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListContentMigrationsGroups) GetMethod() string {
	return "GET"
}

func (t *ListContentMigrationsGroups) GetURLPath() string {
	path := "groups/{group_id}/content_migrations"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListContentMigrationsGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListContentMigrationsGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListContentMigrationsGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListContentMigrationsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListContentMigrationsGroups) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.ContentMigration, *canvasapi.PagedResource, error) {
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
	ret := []*models.ContentMigration{}
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
