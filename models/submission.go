package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Submission struct {
	AssignmentID                  int64                `json:"assignment_id"`                    // The submission's assignment id.Example: 23
	Assignment                    *Assignment          `json:"assignment"`                       // The submission's assignment (see the assignments API) (optional).
	Course                        *Course              `json:"course"`                           // The submission's course (see the course API) (optional).
	Attempt                       int64                `json:"attempt"`                          // This is the submission attempt number..Example: 1
	Body                          string               `json:"body"`                             // The content of the submission, if it was submitted directly in a text field..Example: There are three factors too.
	Grade                         string               `json:"grade"`                            // The grade for the submission, translated into the assignment grading scheme (so a letter grade, for example)..Example: A-
	GradeMatchesCurrentSubmission bool                 `json:"grade_matches_current_submission"` // A boolean flag which is false if the student has re-submitted since the submission was last graded..Example: true
	HtmlUrl                       string               `json:"html_url"`                         // URL to the submission. This will require the user to log in..Example: http://example.com/courses/255/assignments/543/submissions/134
	PreviewUrl                    string               `json:"preview_url"`                      // URL to the submission preview. This will require the user to log in..Example: http://example.com/courses/255/assignments/543/submissions/134?preview=1
	Score                         float64              `json:"score"`                            // The raw score.Example: 13.5
	SubmissionComments            []*SubmissionComment `json:"submission_comments"`              // Associated comments for a submission (optional).
	SubmissionType                string               `json:"submission_type"`                  // The types of submission ex: ('online_text_entry'|'online_url'|'online_upload'|'media_recording'|'student_annotation').Example: online_text_entry
	SubmittedAt                   time.Time            `json:"submitted_at"`                     // The timestamp when the assignment was submitted.Example: 2012-01-01T01:00:00Z
	Url                           string               `json:"url"`                              // The URL of the submission (for 'online_url' submissions)..
	UserID                        int64                `json:"user_id"`                          // The id of the user who created the submission.Example: 134
	GraderID                      int64                `json:"grader_id"`                        // The id of the user who graded the submission. This will be null for submissions that haven't been graded yet. It will be a positive number if a real user has graded the submission and a negative number if the submission was graded by a process (e.g. Quiz autograder and autograding LTI tools).  Specifically autograded quizzes set grader_id to the negative of the quiz id.  Submissions autograded by LTI tools set grader_id to the negative of the tool id..Example: 86
	GradedAt                      time.Time            `json:"graded_at"`                        // Example: 2012-01-02T03:05:34Z
	User                          *User                `json:"user"`                             // The submissions user (see user API) (optional).
	Late                          bool                 `json:"late"`                             // Whether the submission was made after the applicable due date.
	AssignmentVisible             bool                 `json:"assignment_visible"`               // Whether the assignment is visible to the user who submitted the assignment. Submissions where `assignment_visible` is false no longer count towards the student's grade and the assignment can no longer be accessed by the student. `assignment_visible` becomes false for submissions that do not have a grade and whose assignment is no longer assigned to the student's section..Example: true
	Excused                       bool                 `json:"excused"`                          // Whether the assignment is excused.  Excused assignments have no impact on a user's grade..Example: true
	Missing                       bool                 `json:"missing"`                          // Whether the assignment is missing..Example: true
	LatePolicyStatus              string               `json:"late_policy_status"`               // The status of the submission in relation to the late policy. Can be late, missing, none, or null..Example: missing
	PointsDeducted                float64              `json:"points_deducted"`                  // The amount of points automatically deducted from the score by the missing/late policy for a late or missing assignment..Example: 12.3
	SecondsLate                   float64              `json:"seconds_late"`                     // The amount of time, in seconds, that an submission is late by..Example: 300
	WorkflowState                 string               `json:"workflow_state"`                   // The current state of the submission.Example: submitted
	ExtraAttempts                 float64              `json:"extra_attempts"`                   // Extra submission attempts allowed for the given user and assignment..Example: 10
	AnonymousID                   string               `json:"anonymous_id"`                     // A unique short ID identifying this submission without reference to the owning user. Only included if the caller has administrator access for the current account..Example: acJ4Q
	PostedAt                      time.Time            `json:"posted_at"`                        // The date this submission was posted to the student, or nil if it has not been posted..Example: 2020-01-02T11:10:30Z
	ReadStatus                    string               `json:"read_status"`                      // The read status of this submission for the given user (optional). Including read_status will mark submission(s) as read..Example: read
}

func (t *Submission) HasError() error {
	var s []string
	s = []string{"online_text_entry", "online_url", "online_upload", "media_recording", "student_annotation"}
	if !string_utils.Include(s, t.SubmissionType) {
		return fmt.Errorf("expected 'submission_type' to be one of %v", s)
	}
	s = []string{"graded", "submitted", "unsubmitted", "pending_review"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	s = []string{"read", "unread"}
	if !string_utils.Include(s, t.ReadStatus) {
		return fmt.Errorf("expected 'read_status' to be one of %v", s)
	}
	return nil
}
