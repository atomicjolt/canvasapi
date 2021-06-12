package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type MigrationIssue struct {
	ID                  int64     `json:"id"`                    // the unique identifier for the issue.Example: 370663
	ContentMigrationUrl string    `json:"content_migration_url"` // API url to the content migration.Example: https://example.com/api/v1/courses/1/content_migrations/1
	Description         string    `json:"description"`           // Description of the issue for the end-user.Example: Questions in this quiz couldn't be converted
	WorkflowState       string    `json:"workflow_state"`        // Current state of the issue: active, resolved.Example: active
	FixIssueHtmlUrl     string    `json:"fix_issue_html_url"`    // HTML Url to the Canvas page to investigate the issue.Example: https://example.com/courses/1/quizzes/2
	IssueType           string    `json:"issue_type"`            // Severity of the issue: todo, warning, error.Example: warning
	ErrorReportHtmlUrl  string    `json:"error_report_html_url"` // Link to a Canvas error report if present (If the requesting user has permissions).Example: https://example.com/error_reports/3
	ErrorMessage        string    `json:"error_message"`         // Site administrator error message (If the requesting user has permissions).Example: admin only message
	CreatedAt           time.Time `json:"created_at"`            // timestamp.Example: 2012-06-01T00:00:00-06:00
	UpdatedAt           time.Time `json:"updated_at"`            // timestamp.Example: 2012-06-01T00:00:00-06:00
}

func (t *MigrationIssue) HasError() error {
	var s []string
	s = []string{"active", "resolved"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	s = []string{"todo", "warning", "error"}
	if !string_utils.Include(s, t.IssueType) {
		return fmt.Errorf("expected 'issue_type' to be one of %v", s)
	}
	return nil
}
