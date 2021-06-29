package models

type QuizAssignmentOverrideSet struct {
	QuizID   string                  `json:"quiz_id" url:"quiz_id,omitempty"`     // ID of the quiz those dates are for..Example: 1
	DueDates *QuizAssignmentOverride `json:"due_dates" url:"due_dates,omitempty"` // An array of quiz assignment overrides. For students, this array will always contain a single item which is the set of dates that apply to that student. For teachers and staff, it may contain more..
	AllDates *QuizAssignmentOverride `json:"all_dates" url:"all_dates,omitempty"` // An array of all assignment overrides active for the quiz. This is visible only to teachers and staff..
}

func (t *QuizAssignmentOverrideSet) HasErrors() error {
	return nil
}
