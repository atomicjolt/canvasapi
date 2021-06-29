package models

type QuizExtension struct {
	QuizID           int64  `json:"quiz_id" url:"quiz_id,omitempty"`                     // The ID of the Quiz the quiz extension belongs to..Example: 2
	UserID           int64  `json:"user_id" url:"user_id,omitempty"`                     // The ID of the Student that needs the quiz extension..Example: 3
	ExtraAttempts    int64  `json:"extra_attempts" url:"extra_attempts,omitempty"`       // Number of times the student is allowed to re-take the quiz over the multiple-attempt limit..Example: 1
	ExtraTime        int64  `json:"extra_time" url:"extra_time,omitempty"`               // Amount of extra time allowed for the quiz submission, in minutes..Example: 60
	ManuallyUnlocked bool   `json:"manually_unlocked" url:"manually_unlocked,omitempty"` // The student can take the quiz even if it's locked for everyone else.Example: true
	EndAt            string `json:"end_at" url:"end_at,omitempty"`                       // The time at which the quiz submission will be overdue, and be flagged as a late submission..Example: 2013-11-07T13:16:18Z
}

func (t *QuizExtension) HasErrors() error {
	return nil
}
