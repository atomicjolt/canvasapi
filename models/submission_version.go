package models

import (
	"time"
)

type SubmissionVersion struct {
	AssignmentID                  int64     `json:"assignment_id"`                    // the id of the assignment this submissions is for.Example: 22604
	AssignmentName                string    `json:"assignment_name"`                  // the name of the assignment this submission is for.Example: some assignment
	Body                          string    `json:"body"`                             // the body text of the submission.Example: text from the submission
	CurrentGrade                  string    `json:"current_grade"`                    // the most up to date grade for the current version of this submission.Example: 100
	CurrentGradedAt               time.Time `json:"current_graded_at"`                // the latest time stamp for the grading of this submission.Example: 2013-01-31T18:16:31Z
	CurrentGrader                 string    `json:"current_grader"`                   // the name of the most recent grader for this submission.Example: Grader Name
	GradeMatchesCurrentSubmission bool      `json:"grade_matches_current_submission"` // boolean indicating whether the grade is equal to the current submission grade.Example: true
	GradedAt                      time.Time `json:"graded_at"`                        // time stamp for the grading of this version of the submission.Example: 2013-01-31T18:16:31Z
	Grader                        string    `json:"grader"`                           // the name of the user who graded this version of the submission.Example: Grader Name
	GraderID                      int64     `json:"grader_id"`                        // the user id of the user who graded this version of the submission.Example: 67379
	ID                            int64     `json:"id"`                               // the id of the submission of which this is a version.Example: 11607
	NewGrade                      string    `json:"new_grade"`                        // the updated grade provided in this version of the submission.Example: 100
	NewGradedAt                   time.Time `json:"new_graded_at"`                    // the timestamp for the grading of this version of the submission (alias for graded_at).Example: 2013-01-31T18:16:31Z
	NewGrader                     string    `json:"new_grader"`                       // alias for 'grader'.Example: Grader Name
	PreviousGrade                 string    `json:"previous_grade"`                   // the grade for the submission version immediately preceding this one.Example: 90
	PreviousGradedAt              time.Time `json:"previous_graded_at"`               // the timestamp for the grading of the submission version immediately preceding this one.Example: 2013-01-29T12:12:12Z
	PreviousGrader                string    `json:"previous_grader"`                  // the name of the grader who graded the version of this submission immediately preceding this one.Example: Graded on submission
	Score                         int64     `json:"score"`                            // the score for this version of the submission.Example: 100
	UserName                      string    `json:"user_name"`                        // the name of the student who created this submission.Example: student@example.com
	SubmissionType                string    `json:"submission_type"`                  // the type of submission.Example: online
	Url                           string    `json:"url"`                              // the url of the submission, if there is one.
	UserID                        int64     `json:"user_id"`                          // the user ID of the student who created this submission.Example: 67376
	WorkflowState                 string    `json:"workflow_state"`                   // the state of the submission at this version.Example: unsubmitted
}

func (t *SubmissionVersion) HasError() error {
	return nil
}
