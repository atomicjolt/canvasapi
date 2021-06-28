package models

type SubmissionHistory struct {
	SubmissionID int64                `json:"submission_id" url:"submission_id,omitempty"` // the id of the submission.Example: 4
	Versions     []*SubmissionVersion `json:"versions" url:"versions,omitempty"`           // an array of all the versions of this submission.
}

func (t *SubmissionHistory) HasError() error {
	return nil
}
