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

// ListMigrationIssuesAccounts Returns paginated migration issues
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ContentMigrationID (Required) ID
//
type ListMigrationIssuesAccounts struct {
	Path struct {
		AccountID          string `json:"account_id" url:"account_id,omitempty"`                     //  (Required)
		ContentMigrationID string `json:"content_migration_id" url:"content_migration_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListMigrationIssuesAccounts) GetMethod() string {
	return "GET"
}

func (t *ListMigrationIssuesAccounts) GetURLPath() string {
	path := "accounts/{account_id}/content_migrations/{content_migration_id}/migration_issues"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	return path
}

func (t *ListMigrationIssuesAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListMigrationIssuesAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMigrationIssuesAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMigrationIssuesAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.ContentMigrationID == "" {
		errs = append(errs, "'Path.ContentMigrationID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMigrationIssuesAccounts) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.MigrationIssue, *canvasapi.PagedResource, error) {
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
