package models

type QuizStatisticsQuestionStatistics struct {
	Responses int64                           `json:"responses"` // Number of students who have provided an answer to this question. Blank or empty responses are not counted..Example: 3
	Answers   *QuizStatisticsAnswerStatistics `json:"answers"`   // Statistics related to each individual pre-defined answer..
}

func (t *QuizStatisticsQuestionStatistics) HasError() error {
	return nil
}
