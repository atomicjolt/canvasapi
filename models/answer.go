package models

type Answer struct {
	ID                             int64  `json:"id" url:"id,omitempty"`                                                               // The unique identifier for the answer.  Do not supply if this answer is part of a new question.Example: 6656
	AnswerText                     string `json:"answer_text" url:"answer_text,omitempty"`                                             // The text of the answer..Example: Constantinople
	AnswerWeight                   int64  `json:"answer_weight" url:"answer_weight,omitempty"`                                         // An integer to determine correctness of the answer. Incorrect answers should be 0, correct answers should be 100..Example: 100
	AnswerComments                 string `json:"answer_comments" url:"answer_comments,omitempty"`                                     // Specific contextual comments for a particular answer..Example: Remember to check your spelling prior to submitting this answer.
	TextAfterAnswers               string `json:"text_after_answers" url:"text_after_answers,omitempty"`                               // Used in missing word questions.  The text to follow the missing word.Example:  is the capital of Utah.
	AnswerMatchLeft                string `json:"answer_match_left" url:"answer_match_left,omitempty"`                                 // Used in matching questions.  The static value of the answer that will be displayed on the left for students to match for..Example: Salt Lake City
	AnswerMatchRight               string `json:"answer_match_right" url:"answer_match_right,omitempty"`                               // Used in matching questions. The correct match for the value given in answer_match_left.  Will be displayed in a dropdown with the other answer_match_right values...Example: Utah
	MatchingAnswerIncorrectMatches string `json:"matching_answer_incorrect_matches" url:"matching_answer_incorrect_matches,omitempty"` // Used in matching questions. A list of distractors, delimited by new lines (
	//) that will be seeded with all the answer_match_right values..Example: Nevada California Washington
	NumericalAnswerType string  `json:"numerical_answer_type" url:"numerical_answer_type,omitempty"` // Used in numerical questions.  Values can be 'exact_answer', 'range_answer', or 'precision_answer'..Example: exact_answer
	Exact               int64   `json:"exact" url:"exact,omitempty"`                                 // Used in numerical questions of type 'exact_answer'.  The value the answer should equal..Example: 42
	Margin              int64   `json:"margin" url:"margin,omitempty"`                               // Used in numerical questions of type 'exact_answer'. The margin of error allowed for the student's answer..Example: 4
	Approximate         float64 `json:"approximate" url:"approximate,omitempty"`                     // Used in numerical questions of type 'precision_answer'.  The value the answer should equal..Example: 1234600000.0
	Precision           int64   `json:"precision" url:"precision,omitempty"`                         // Used in numerical questions of type 'precision_answer'. The numerical precision that will be used when comparing the student's answer..Example: 4
	Start               int64   `json:"start" url:"start,omitempty"`                                 // Used in numerical questions of type 'range_answer'. The start of the allowed range (inclusive)..Example: 1
	End                 int64   `json:"end" url:"end,omitempty"`                                     // Used in numerical questions of type 'range_answer'. The end of the allowed range (inclusive)..Example: 10
	BlankID             int64   `json:"blank_id" url:"blank_id,omitempty"`                           // Used in fill in multiple blank and multiple dropdowns questions..Example: 1170
}

func (t *Answer) HasErrors() error {
	return nil
}
