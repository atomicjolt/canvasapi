package models

type QuizAssignmentOverrideSetContainer struct {
	QuizAssignmentOverrides []*QuizAssignmentOverrideSet `json:"quiz_assignment_overrides"` // The QuizAssignmentOverrideSet.
}

func (t *QuizAssignmentOverrideSetContainer) HasError() error {
	return nil
}
