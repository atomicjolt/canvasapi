package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetMigrationIssueCourses Returns data on an individual migration issue
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ContentMigrationID (Required) ID
// # ID (Required) ID
//
type GetMigrationIssueCourses struct {
	Path struct {
		CourseID           string `json:"course_id"`            //  (Required)
		ContentMigrationID string `json:"content_migration_id"` //  (Required)
		ID                 string `json:"id"`                   //  (Required)
	} `json:"path"`
}

func (t *GetMigrationIssueCourses) GetMethod() string {
	return "GET"
}

func (t *GetMigrationIssueCourses) GetURLPath() string {
	path := "courses/{course_id}/content_migrations/{content_migration_id}/migration_issues/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetMigrationIssueCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetMigrationIssueCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetMigrationIssueCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
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

func (t *GetMigrationIssueCourses) Do(c *canvasapi.Canvas) (*models.MigrationIssue, error) {
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
