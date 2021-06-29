package models

type QuizStatisticsQuestionStatistics struct {
	Responses int64                           `json:"responses" url:"responses,omitempty"` // Number of students who have provided an answer to this question. Blank or empty responses are not counted..Example: 3
	Answers   *QuizStatisticsAnswerStatistics `json:"answers" url:"answers,omitempty"`     // Statistics related to each individual pre-defined answer..
}

func (t *QuizStatisticsQuestionStatistics) HasErrors() error {
	return nil
}
