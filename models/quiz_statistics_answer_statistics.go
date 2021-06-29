package models

type QuizStatisticsAnswerStatistics struct {
	ID        int64  `json:"id" url:"id,omitempty"`               // ID of the answer..Example: 3866
	Text      string `json:"text" url:"text,omitempty"`           // The text attached to the answer..Example: Blue.
	Weight    int64  `json:"weight" url:"weight,omitempty"`       // An integer to determine correctness of the answer. Incorrect answers should be 0, correct answers should 100.Example: 100
	Responses int64  `json:"responses" url:"responses,omitempty"` // Number of students who have chosen this answer..Example: 2
}

func (t *QuizStatisticsAnswerStatistics) HasErrors() error {
	return nil
}
