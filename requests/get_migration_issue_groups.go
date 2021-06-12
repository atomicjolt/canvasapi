package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetMigrationIssueGroups Returns data on an individual migration issue
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # GroupID (Required) ID
// # ContentMigrationID (Required) ID
// # ID (Required) ID
//
type GetMigrationIssueGroups struct {
	Path struct {
		GroupID            string `json:"group_id"`             //  (Required)
		ContentMigrationID string `json:"content_migration_id"` //  (Required)
		ID                 string `json:"id"`                   //  (Required)
	} `json:"path"`
}

func (t *GetMigrationIssueGroups) GetMethod() string {
	return "GET"
}

func (t *GetMigrationIssueGroups) GetURLPath() string {
	path := "groups/{group_id}/content_migrations/{content_migration_id}/migration_issues/{id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetMigrationIssueGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *GetMigrationIssueGroups) GetBody() (string, error) {
	return "", nil
}

func (t *GetMigrationIssueGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
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

func (t *GetMigrationIssueGroups) Do(c *canvasapi.Canvas) (*models.MigrationIssue, error) {
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
