package models

type QuizSubmissionQuestion struct {
	ID      int64    `json:"id" url:"id,omitempty"`           // The ID of the QuizQuestion this answer is for..Example: 1
	Flagged bool     `json:"flagged" url:"flagged,omitempty"` // Whether this question is flagged..Example: true
	Answer  string   `json:"answer" url:"answer,omitempty"`   // The provided answer (if any) for this question. The format of this parameter depends on the type of the question, see the Appendix for more information..
	Answers []string `json:"answers" url:"answers,omitempty"` // The possible answers for this question when those possible answers are necessary.  The presence of this parameter is dependent on permissions..
}

func (t *QuizSubmissionQuestion) HasErrors() error {
	return nil
}
