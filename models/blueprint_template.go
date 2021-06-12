package models

import (
	"time"
)

type BlueprintTemplate struct {
	ID                    int64               `json:"id"`                       // The ID of the template..Example: 1
	CourseID              int64               `json:"course_id"`                // The ID of the Course the template belongs to..Example: 2
	LastExportCompletedAt time.Time           `json:"last_export_completed_at"` // Time when the last export was completed.Example: 2013-08-28T23:59:00-06:00
	AssociatedCourseCount int64               `json:"associated_course_count"`  // Number of associated courses for the template.Example: 3
	LatestMigration       *BlueprintMigration `json:"latest_migration"`         // Details of the latest migration.
}

func (t *BlueprintTemplate) HasError() error {
	return nil
}
