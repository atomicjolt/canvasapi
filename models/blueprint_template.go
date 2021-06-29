package models

import (
	"time"
)

type BlueprintTemplate struct {
	ID                    int64               `json:"id" url:"id,omitempty"`                                             // The ID of the template..Example: 1
	CourseID              int64               `json:"course_id" url:"course_id,omitempty"`                               // The ID of the Course the template belongs to..Example: 2
	LastExportCompletedAt time.Time           `json:"last_export_completed_at" url:"last_export_completed_at,omitempty"` // Time when the last export was completed.Example: 2013-08-28T23:59:00-06:00
	AssociatedCourseCount int64               `json:"associated_course_count" url:"associated_course_count,omitempty"`   // Number of associated courses for the template.Example: 3
	LatestMigration       *BlueprintMigration `json:"latest_migration" url:"latest_migration,omitempty"`                 // Details of the latest migration.
}

func (t *BlueprintTemplate) HasErrors() error {
	return nil
}
