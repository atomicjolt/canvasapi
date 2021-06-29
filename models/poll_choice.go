package models

type PollChoice struct {
	ID        int64  `json:"id" url:"id,omitempty"`                 // The unique identifier for the poll choice..Example: 1023
	PollID    int64  `json:"poll_id" url:"poll_id,omitempty"`       // The id of the poll this poll choice belongs to..Example: 1779
	IsCorrect bool   `json:"is_correct" url:"is_correct,omitempty"` // Specifies whether or not this poll choice is a 'correct' choice..Example: true
	Text      string `json:"text" url:"text,omitempty"`             // The text of the poll choice..Example: Choice A
	Position  int64  `json:"position" url:"position,omitempty"`     // The order of the poll choice in relation to it's sibling poll choices..Example: 1
}

func (t *PollChoice) HasErrors() error {
	return nil
}
