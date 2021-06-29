package models

type QuizAssignmentOverrideSetContainer struct {
	QuizAssignmentOverrides []*QuizAssignmentOverrideSet `json:"quiz_assignment_overrides" url:"quiz_assignment_overrides,omitempty"` // The QuizAssignmentOverrideSet.
}

func (t *QuizAssignmentOverrideSetContainer) HasErrors() error {
	return nil
}
