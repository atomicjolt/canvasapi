package models

import (
	"time"
)

type CourseProgress struct {
	RequirementCount          int64     `json:"requirement_count" url:"requirement_count,omitempty"`                     // total number of requirements from all modules.Example: 10
	RequirementCompletedCount int64     `json:"requirement_completed_count" url:"requirement_completed_count,omitempty"` // total number of requirements the user has completed from all modules.Example: 1
	NextRequirementUrl        string    `json:"next_requirement_url" url:"next_requirement_url,omitempty"`               // url to next module item that has an unmet requirement. null if the user has completed the course or the current module does not require sequential progress.Example: http://localhost/courses/1/modules/items/2
	CompletedAt               time.Time `json:"completed_at" url:"completed_at,omitempty"`                               // date the course was completed. null if the course has not been completed by this user.Example: 2013-06-01T00:00:00-06:00
}

func (t *CourseProgress) HasError() error {
	return nil
}
