package models

type Score struct {
	UserID           string  `json:"user_id"`           // The lti_user_id or the Canvas user_id.Example: 50 | 'abcasdf'
	ScoreGiven       float64 `json:"score_given"`       // The Current score received in the tool for this line item and user, scaled to the scoreMaximum.Example: 50
	ScoreMaximum     float64 `json:"score_maximum"`     // Maximum possible score for this result; it must be present if scoreGiven is present..Example: 50
	Comment          string  `json:"comment"`           // Comment visible to the student about this score..
	Timestamp        string  `json:"timestamp"`         // Date and time when the score was modified in the tool. Should use subsecond precision..Example: 2017-04-16T18:54:36.736+00:00
	ActivityProgress string  `json:"activity_progress"` // Indicate to Canvas the status of the user towards the activity's completion. Must be one of Initialized, Started, InProgress, Submitted, Completed.Example: Completed
	GradingProgress  string  `json:"grading_progress"`  // Indicate to Canvas the status of the grading process. A value of PendingManual will require intervention by a grader. Values of NotReady, Failed, and Pending will cause the scoreGiven to be ignored. FullyGraded values will require no action. Possible values are NotReady, Failed, Pending, PendingManual, FullyGraded.Example: FullyGraded
}

func (t *Score) HasError() error {
	return nil
}