package models

type QuizStatisticsAnswerPointBiserial struct {
	AnswerID      int64   `json:"answer_id"`      // ID of the answer the point biserial is for..Example: 3866
	PointBiserial float64 `json:"point_biserial"` // The point biserial value for this answer. Value ranges between -1 and 1..Example: -0.802955068546966
	Correct       bool    `json:"correct"`        // Convenience attribute that denotes whether this is the correct answer as opposed to being a distractor. This is mutually exclusive with the `distractor` value.Example: true
	Distractor    bool    `json:"distractor"`     // Convenience attribute that denotes whether this is a distractor answer and not the correct one. This is mutually exclusive with the `correct` value.
}

func (t *QuizStatisticsAnswerPointBiserial) HasError() error {
	return nil
}
