package models

type AssignmentExtension struct {
	AssignmentID  int64 `json:"assignment_id"`  // The ID of the Assignment the extension belongs to..Example: 2
	UserID        int64 `json:"user_id"`        // The ID of the Student that needs the assignment extension..Example: 3
	ExtraAttempts int64 `json:"extra_attempts"` // Number of times the student is allowed to re-submit the assignment.Example: 2
}

func (t *AssignmentExtension) HasError() error {
	return nil
}
