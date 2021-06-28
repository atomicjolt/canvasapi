package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// UpdateMigrationIssueCourses Update the workflow_state of a migration issue
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ContentMigrationID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # WorkflowState (Required) . Must be one of active, resolvedSet the workflow_state of the issue.
//
type UpdateMigrationIssueCourses struct {
	Path struct {
		CourseID           string `json:"course_id" url:"course_id,omitempty"`                       //  (Required)
		ContentMigrationID string `json:"content_migration_id" url:"content_migration_id,omitempty"` //  (Required)
		ID                 string `json:"id" url:"id,omitempty"`                                     //  (Required)
	} `json:"path"`

	Form struct {
		WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` //  (Required) . Must be one of active, resolved
	} `json:"form"`
}

func (t *UpdateMigrationIssueCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateMigrationIssueCourses) GetURLPath() string {
	path := "courses/{course_id}/content_migrations/{content_migration_id}/migration_issues/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateMigrationIssueCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMigrationIssueCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateMigrationIssueCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateMigrationIssueCourses) HasErrors() error {
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
	if t.Form.WorkflowState == "" {
		errs = append(errs, "'WorkflowState' is required")
	}
	if t.Form.WorkflowState != "" && !string_utils.Include([]string{"active", "resolved"}, t.Form.WorkflowState) {
		errs = append(errs, "WorkflowState must be one of active, resolved")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMigrationIssueCourses) Do(c *canvasapi.Canvas) (*models.MigrationIssue, error) {
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
