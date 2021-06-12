package models

type PollChoice struct {
	ID        int64  `json:"id"`         // The unique identifier for the poll choice..Example: 1023
	PollID    int64  `json:"poll_id"`    // The id of the poll this poll choice belongs to..Example: 1779
	IsCorrect bool   `json:"is_correct"` // Specifies whether or not this poll choice is a 'correct' choice..Example: true
	Text      string `json:"text"`       // The text of the poll choice..Example: Choice A
	Position  int64  `json:"position"`   // The order of the poll choice in relation to it's sibling poll choices..Example: 1
}

func (t *PollChoice) HasError() error {
	return nil
}
