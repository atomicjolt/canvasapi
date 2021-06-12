package models

type PollSubmission struct {
	ID           int64  `json:"id"`             // The unique identifier for the poll submission..Example: 1023
	PollChoiceID int64  `json:"poll_choice_id"` // The unique identifier of the poll choice chosen for this submission..Example: 155
	UserID       int64  `json:"user_id"`        // the unique identifier of the user who submitted this poll submission..Example: 4555
	CreatedAt    string `json:"created_at"`     // The date and time the poll submission was submitted..Example: 2013-11-07T13:16:18Z
}

func (t *PollSubmission) HasError() error {
	return nil
}
