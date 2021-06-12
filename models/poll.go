package models

type Poll struct {
	ID           int64  `json:"id"`            // The unique identifier for the poll..Example: 1023
	Question     string `json:"question"`      // The question/title of the poll..Example: What do you consider most important to your learning in this course?
	Description  string `json:"description"`   // A short description of the poll..Example: This poll is to determine what priorities the students in the course have.
	CreatedAt    string `json:"created_at"`    // The time at which the poll was created..Example: 2014-01-07T15:16:18Z
	UserID       int64  `json:"user_id"`       // The unique identifier for the user that created the poll..Example: 105
	TotalResults string `json:"total_results"` // An aggregate of the results of all associated poll sessions, with the poll choice id as the key, and the aggregated submission count as the value..Example: 20, 5, 17
}

func (t *Poll) HasError() error {
	return nil
}
