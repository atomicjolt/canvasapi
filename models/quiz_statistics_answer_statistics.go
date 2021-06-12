package models

type QuizStatisticsAnswerStatistics struct {
	ID        int64  `json:"id"`        // ID of the answer..Example: 3866
	Text      string `json:"text"`      // The text attached to the answer..Example: Blue.
	Weight    int64  `json:"weight"`    // An integer to determine correctness of the answer. Incorrect answers should be 0, correct answers should 100.Example: 100
	Responses int64  `json:"responses"` // Number of students who have chosen this answer..Example: 2
}

func (t *QuizStatisticsAnswerStatistics) HasError() error {
	return nil
}
