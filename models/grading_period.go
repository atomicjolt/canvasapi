package models

type GradingPeriod struct {
	ID        int64  `json:"id" url:"id,omitempty"`                 // The unique identifier for the grading period..Example: 1023
	Title     string `json:"title" url:"title,omitempty"`           // The title for the grading period..Example: First Block
	StartDate string `json:"start_date" url:"start_date,omitempty"` // The start date of the grading period..Example: 2014-01-07T15:04:00Z
	EndDate   string `json:"end_date" url:"end_date,omitempty"`     // The end date of the grading period..Example: 2014-05-07T17:07:00Z
	CloseDate string `json:"close_date" url:"close_date,omitempty"` // Grades can only be changed before the close date of the grading period..Example: 2014-06-07T17:07:00Z
	Weight    int64  `json:"weight" url:"weight,omitempty"`         // A weight value that contributes to the overall weight of a grading period set which is used to calculate how much assignments in this period contribute to the total grade.Example: 33.33
	IsClosed  bool   `json:"is_closed" url:"is_closed,omitempty"`   // If true, the grading period's close_date has passed..Example: true
}

func (t *GradingPeriod) HasError() error {
	return nil
}
