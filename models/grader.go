package models

type Grader struct {
	ID          int64   `json:"id" url:"id,omitempty"`                   // the user_id of the user who graded the contained submissions.Example: 27
	Name        string  `json:"name" url:"name,omitempty"`               // the name of the user who graded the contained submissions.Example: Some User
	Assignments []int64 `json:"assignments" url:"assignments,omitempty"` // the assignment groups for all submissions in this response that were graded by this user.  The details are not nested inside here, but the fact that an assignment is present here means that the grader did grade submissions for this assignment on the contextual date. You can use the id of a grader and of an assignment to make another API call to find all submissions for a grader/assignment combination on a given date..Example: 1, 2, 3
}

func (t *Grader) HasError() error {
	return nil
}
