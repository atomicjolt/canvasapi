package models

type SubmissionHistory struct {
	SubmissionID int64                `json:"submission_id"` // the id of the submission.Example: 4
	Versions     []*SubmissionVersion `json:"versions"`      // an array of all the versions of this submission.
}

func (t *SubmissionHistory) HasError() error {
	return nil
}
