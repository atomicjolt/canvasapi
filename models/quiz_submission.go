package models

type QuizSubmission struct {
	ID                        int64  `json:"id" url:"id,omitempty"`                                                     // The ID of the quiz submission..Example: 1
	QuizID                    int64  `json:"quiz_id" url:"quiz_id,omitempty"`                                           // The ID of the Quiz the quiz submission belongs to..Example: 2
	UserID                    int64  `json:"user_id" url:"user_id,omitempty"`                                           // The ID of the Student that made the quiz submission..Example: 3
	SubmissionID              int64  `json:"submission_id" url:"submission_id,omitempty"`                               // The ID of the Submission the quiz submission represents..Example: 1
	StartedAt                 string `json:"started_at" url:"started_at,omitempty"`                                     // The time at which the student started the quiz submission..Example: 2013-11-07T13:16:18Z
	FinishedAt                string `json:"finished_at" url:"finished_at,omitempty"`                                   // The time at which the student submitted the quiz submission..Example: 2013-11-07T13:16:18Z
	EndAt                     string `json:"end_at" url:"end_at,omitempty"`                                             // The time at which the quiz submission will be overdue, and be flagged as a late submission..Example: 2013-11-07T13:16:18Z
	Attempt                   int64  `json:"attempt" url:"attempt,omitempty"`                                           // For quizzes that allow multiple attempts, this field specifies the quiz submission attempt number..Example: 3
	ExtraAttempts             int64  `json:"extra_attempts" url:"extra_attempts,omitempty"`                             // Number of times the student was allowed to re-take the quiz over the multiple-attempt limit..Example: 1
	ExtraTime                 int64  `json:"extra_time" url:"extra_time,omitempty"`                                     // Amount of extra time allowed for the quiz submission, in minutes..Example: 60
	ManuallyUnlocked          bool   `json:"manually_unlocked" url:"manually_unlocked,omitempty"`                       // The student can take the quiz even if it's locked for everyone else.Example: true
	TimeSpent                 int64  `json:"time_spent" url:"time_spent,omitempty"`                                     // Amount of time spent, in seconds..Example: 300
	Score                     int64  `json:"score" url:"score,omitempty"`                                               // The score of the quiz submission, if graded..Example: 3
	ScoreBeforeRegrade        int64  `json:"score_before_regrade" url:"score_before_regrade,omitempty"`                 // The original score of the quiz submission prior to any re-grading..Example: 2
	KeptScore                 int64  `json:"kept_score" url:"kept_score,omitempty"`                                     // For quizzes that allow multiple attempts, this is the score that will be used, which might be the score of the latest, or the highest, quiz submission..Example: 5
	FudgePoints               int64  `json:"fudge_points" url:"fudge_points,omitempty"`                                 // Number of points the quiz submission's score was fudged by..Example: 1
	HasSeenResults            bool   `json:"has_seen_results" url:"has_seen_results,omitempty"`                         // Whether the student has viewed their results to the quiz..Example: true
	WorkflowState             string `json:"workflow_state" url:"workflow_state,omitempty"`                             // The current state of the quiz submission. Possible values: ['untaken'|'pending_review'|'complete'|'settings_only'|'preview']..Example: untaken
	OverdueAndNeedsSubmission bool   `json:"overdue_and_needs_submission" url:"overdue_and_needs_submission,omitempty"` // Indicates whether the quiz submission is overdue and needs submission.Example: false
}

func (t *QuizSubmission) HasErrors() error {
	return nil
}
