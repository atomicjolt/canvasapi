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

// UpdateMigrationIssueUsers Update the workflow_state of a migration issue
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.UserID (Required) ID
// # Path.ContentMigrationID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.WorkflowState (Required) . Must be one of active, resolvedSet the workflow_state of the issue.
//
type UpdateMigrationIssueUsers struct {
	Path struct {
		UserID             string `json:"user_id" url:"user_id,omitempty"`                           //  (Required)
		ContentMigrationID string `json:"content_migration_id" url:"content_migration_id,omitempty"` //  (Required)
		ID                 string `json:"id" url:"id,omitempty"`                                     //  (Required)
	} `json:"path"`

	Form struct {
		WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` //  (Required) . Must be one of active, resolved
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

func (t *UpdateMigrationIssueUsers) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateMigrationIssueUsers) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateMigrationIssueUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.ContentMigrationID == "" {
		errs = append(errs, "'Path.ContentMigrationID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.WorkflowState == "" {
		errs = append(errs, "'Form.WorkflowState' is required")
	}
	if t.Form.WorkflowState != "" && !string_utils.Include([]string{"active", "resolved"}, t.Form.WorkflowState) {
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
