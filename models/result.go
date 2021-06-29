package models

type Result struct {
	ID            string  `json:"id" url:"id,omitempty"`                         // The fully qualified URL for showing the Result.Example: http://institution.canvas.com/api/lti/courses/5/line_items/2/results/1
	UserID        string  `json:"user_id" url:"user_id,omitempty"`               // The lti_user_id or the Canvas user_id.Example: 50 | 'abcasdf'
	ResultScore   float64 `json:"result_score" url:"result_score,omitempty"`     // The score of the result as defined by Canvas, scaled to the resultMaximum.Example: 50
	ResultMaximum float64 `json:"result_maximum" url:"result_maximum,omitempty"` // Maximum possible score for this result; 1 is the default value and will be assumed if not specified otherwise. Minimum value of 0 required..Example: 50
	Comment       string  `json:"comment" url:"comment,omitempty"`               // Comment visible to the student about the result..
	ScoreOf       string  `json:"score_of" url:"score_of,omitempty"`             // URL of the line item this belongs to.Example: http://institution.canvas.com/api/lti/courses/5/line_items/2
}

func (t *Result) HasErrors() error {
	return nil
}
