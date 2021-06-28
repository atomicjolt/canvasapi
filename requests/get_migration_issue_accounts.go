package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetMigrationIssueAccounts Returns data on an individual migration issue
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ContentMigrationID (Required) ID
// # ID (Required) ID
//
type GetMigrationIssueAccounts struct {
	Path struct {
		AccountID          string `json:"account_id" url:"account_id,omitempty"`                     //  (Required)
		ContentMigrationID string `json:"content_migration_id" url:"content_migration_id,omitempty"` //  (Required)
		ID                 string `json:"id" url:"id,omitempty"`                                     //  (Required)
	} `json:"path"`
}

func (t *GetMigrationIssueAccounts) GetMethod() string {
	return "GET"
}

func (t *GetMigrationIssueAccounts) GetURLPath() string {
	path := "accounts/{account_id}/content_migrations/{content_migration_id}/migration_issues/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetMigrationIssueAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *GetMigrationIssueAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetMigrationIssueAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetMigrationIssueAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ContentMigrationID == "" {
		errs = append(errs, "'ContentMigrationID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetMigrationIssueAccounts) Do(c *canvasapi.Canvas) (*models.MigrationIssue, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.MigrationIssue{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
