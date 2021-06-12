package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListMigrationIssuesAccounts Returns paginated migration issues
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ContentMigrationID (Required) ID
//
type ListMigrationIssuesAccounts struct {
	Path struct {
		AccountID          string `json:"account_id"`           //  (Required)
		ContentMigrationID string `json:"content_migration_id"` //  (Required)
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

func (t *ListMigrationIssuesAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *ListMigrationIssuesAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ContentMigrationID == "" {
		errs = append(errs, "'ContentMigrationID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMigrationIssuesAccounts) Do(c *canvasapi.Canvas) ([]*models.MigrationIssue, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.MigrationIssue{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
