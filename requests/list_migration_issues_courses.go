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

// ListMigrationIssuesCourses Returns paginated migration issues
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ContentMigrationID (Required) ID
//
type ListMigrationIssuesCourses struct {
	Path struct {
		CourseID           string `json:"course_id" url:"course_id,omitempty"`                       //  (Required)
		ContentMigrationID string `json:"content_migration_id" url:"content_migration_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListMigrationIssuesCourses) GetMethod() string {
	return "GET"
}

func (t *ListMigrationIssuesCourses) GetURLPath() string {
	path := "courses/{course_id}/content_migrations/{content_migration_id}/migration_issues"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	return path
}

func (t *ListMigrationIssuesCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListMigrationIssuesCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMigrationIssuesCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMigrationIssuesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ContentMigrationID == "" {
		errs = append(errs, "'ContentMigrationID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMigrationIssuesCourses) Do(c *canvasapi.Canvas) ([]*models.MigrationIssue, error) {
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
