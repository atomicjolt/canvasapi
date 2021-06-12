package models

import (
	"time"
)

type ProvisionalGrade struct {
	ProvisionalGradeID            int64     `json:"provisional_grade_id"`             // The identifier for the provisional grade.Example: 23
	Score                         int64     `json:"score"`                            // The numeric score.Example: 90
	Grade                         string    `json:"grade"`                            // The grade.Example: A-
	GradeMatchesCurrentSubmission bool      `json:"grade_matches_current_submission"` // Whether the grade was applied to the most current submission (false if the student resubmitted after grading).Example: true
	GradedAt                      time.Time `json:"graded_at"`                        // When the grade was given.Example: 2015-11-01T00:03:21-06:00
	Final                         bool      `json:"final"`                            // Whether this is the 'final' provisional grade created by the moderator.
	SpeedgraderUrl                string    `json:"speedgrader_url"`                  // A link to view this provisional grade in SpeedGrader™.Example: http://www.example.com/courses/123/gradebook/speed_grader?.
}

func (t *ProvisionalGrade) HasError() error {
	return nil
}
