package models

type QuizQuestion struct {
	ID                int64     `json:"id"`                 // The ID of the quiz question..Example: 1
	QuizID            int64     `json:"quiz_id"`            // The ID of the Quiz the question belongs to..Example: 2
	Position          int64     `json:"position"`           // The order in which the question will be retrieved and displayed..Example: 1
	QuestionName      string    `json:"question_name"`      // The name of the question..Example: Prime Number Identification
	QuestionType      string    `json:"question_type"`      // The type of the question..Example: multiple_choice_question
	QuestionText      string    `json:"question_text"`      // The text of the question..Example: Which of the following is NOT a prime number?
	PointsPossible    int64     `json:"points_possible"`    // The maximum amount of points possible received for getting this question correct..Example: 5
	CorrectComments   string    `json:"correct_comments"`   // The comments to display if the student answers the question correctly..Example: That's correct!
	IncorrectComments string    `json:"incorrect_comments"` // The comments to display if the student answers incorrectly..Example: Unfortunately, that IS a prime number.
	NeutralComments   string    `json:"neutral_comments"`   // The comments to display regardless of how the student answered..Example: Goldbach's conjecture proposes that every even integer greater than 2 can be expressed as the sum of two prime numbers.
	Answers           []*Answer `json:"answers"`            // An array of available answers to display to the student..
}

func (t *QuizQuestion) HasError() error {
	return nil
}
