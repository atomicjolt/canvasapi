package models

import (
	"time"
)

type ProvisionalGrade struct {
	ProvisionalGradeID            int64     `json:"provisional_grade_id" url:"provisional_grade_id,omitempty"`                         // The identifier for the provisional grade.Example: 23
	Score                         int64     `json:"score" url:"score,omitempty"`                                                       // The numeric score.Example: 90
	Grade                         string    `json:"grade" url:"grade,omitempty"`                                                       // The grade.Example: A-
	GradeMatchesCurrentSubmission bool      `json:"grade_matches_current_submission" url:"grade_matches_current_submission,omitempty"` // Whether the grade was applied to the most current submission (false if the student resubmitted after grading).Example: true
	GradedAt                      time.Time `json:"graded_at" url:"graded_at,omitempty"`                                               // When the grade was given.Example: 2015-11-01T00:03:21-06:00
	Final                         bool      `json:"final" url:"final,omitempty"`                                                       // Whether this is the 'final' provisional grade created by the moderator.
	SpeedgraderUrl                string    `json:"speedgrader_url" url:"speedgrader_url,omitempty"`                                   // A link to view this provisional grade in SpeedGrader™.Example: http://www.example.com/courses/123/gradebook/speed_grader?.
}

func (t *ProvisionalGrade) HasErrors() error {
	return nil
}
