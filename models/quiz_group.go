package models

type QuizGroup struct {
	ID                       int64  `json:"id"`                          // The ID of the question group..Example: 1
	QuizID                   int64  `json:"quiz_id"`                     // The ID of the Quiz the question group belongs to..Example: 2
	Name                     string `json:"name"`                        // The name of the question group..Example: Fraction questions
	PickCount                int64  `json:"pick_count"`                  // The number of questions to pick from the group to display to the student..Example: 3
	QuestionPoints           int64  `json:"question_points"`             // The amount of points allotted to each question in the group..Example: 10
	AssessmentQuestionBankID int64  `json:"assessment_question_bank_id"` // The ID of the Assessment question bank to pull questions from..Example: 2
	Position                 int64  `json:"position"`                    // The order in which the question group will be retrieved and displayed..Example: 1
}

func (t *QuizGroup) HasError() error {
	return nil
}