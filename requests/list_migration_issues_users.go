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

// ListMigrationIssuesUsers Returns paginated migration issues
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # UserID (Required) ID
// # ContentMigrationID (Required) ID
//
type ListMigrationIssuesUsers struct {
	Path struct {
		UserID             string `json:"user_id" url:"user_id,omitempty"`                           //  (Required)
		ContentMigrationID string `json:"content_migration_id" url:"content_migration_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListMigrationIssuesUsers) GetMethod() string {
	return "GET"
}

func (t *ListMigrationIssuesUsers) GetURLPath() string {
	path := "users/{user_id}/content_migrations/{content_migration_id}/migration_issues"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	return path
}

func (t *ListMigrationIssuesUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListMigrationIssuesUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMigrationIssuesUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMigrationIssuesUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ContentMigrationID == "" {
		errs = append(errs, "'ContentMigrationID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMigrationIssuesUsers) Do(c *canvasapi.Canvas) ([]*models.MigrationIssue, error) {
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
