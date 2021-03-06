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

// ListMigrationIssuesGroups Returns paginated migration issues
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.ContentMigrationID (Required) ID
//
type ListMigrationIssuesGroups struct {
	Path struct {
		GroupID            string `json:"group_id" url:"group_id,omitempty"`                         //  (Required)
		ContentMigrationID string `json:"content_migration_id" url:"content_migration_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListMigrationIssuesGroups) GetMethod() string {
	return "GET"
}

func (t *ListMigrationIssuesGroups) GetURLPath() string {
	path := "groups/{group_id}/content_migrations/{content_migration_id}/migration_issues"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	return path
}

func (t *ListMigrationIssuesGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListMigrationIssuesGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMigrationIssuesGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMigrationIssuesGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.ContentMigrationID == "" {
		errs = append(errs, "'Path.ContentMigrationID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMigrationIssuesGroups) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.MigrationIssue, *canvasapi.PagedResource, error) {
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
	ret := []*models.MigrationIssue{}
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
