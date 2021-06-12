package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// UpdateMigrationIssueUsers Update the workflow_state of a migration issue
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # UserID (Required) ID
// # ContentMigrationID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # WorkflowState (Required) . Must be one of active, resolvedSet the workflow_state of the issue.
//
type UpdateMigrationIssueUsers struct {
	Path struct {
		UserID             string `json:"user_id"`              //  (Required)
		ContentMigrationID string `json:"content_migration_id"` //  (Required)
		ID                 string `json:"id"`                   //  (Required)
	} `json:"path"`

	Form struct {
		WorkflowState string `json:"workflow_state"` //  (Required) . Must be one of active, resolved
	} `json:"form"`
}

func (t *UpdateMigrationIssueUsers) GetMethod() string {
	return "PUT"
}

func (t *UpdateMigrationIssueUsers) GetURLPath() string {
	path := "users/{user_id}/content_migrations/{content_migration_id}/migration_issues/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{content_migration_id}", fmt.Sprintf("%v", t.Path.ContentMigrationID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateMigrationIssueUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMigrationIssueUsers) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateMigrationIssueUsers) HasErrors() error {
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
	if t.Form.WorkflowState == "" {
		errs = append(errs, "'WorkflowState' is required")
	}
	if !string_utils.Include([]string{"active", "resolved"}, t.Form.WorkflowState) {
		errs = append(errs, "WorkflowState must be one of active, resolved")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMigrationIssueUsers) Do(c *canvasapi.Canvas) (*models.MigrationIssue, error) {
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
