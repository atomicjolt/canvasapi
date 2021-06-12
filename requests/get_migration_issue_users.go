package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetMigrationIssueUsers Returns data on an individual migration issue
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # UserID (Required) ID
// # ContentMigrationID (Required) ID
// # ID (Required) ID
//
type GetMigrationIssueUsers struct {
	Path struct {
		UserID             string `json:"user_id"`              //  (Required)
		ContentMigrationID string `json:"content_migration_id"` //  (Required)
		ID                 string `json:"id"`                   //  (Required)
	} `json:"path"`
}

func (t *GetMigrationIssueUsers) GetMethod() string {
	return "GET"
}

func (t *GetMigrationIssueUsers) GetURLPath() string {
	path := "users/{user_id}/content_migrations/{content_migration_id}/migration_issues/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetMigrationIssueUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *GetMigrationIssueUsers) GetBody() (string, error) {
	return "", nil
}

func (t *GetMigrationIssueUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
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

func (t *GetMigrationIssueUsers) Do(c *canvasapi.Canvas) (*models.MigrationIssue, error) {
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
