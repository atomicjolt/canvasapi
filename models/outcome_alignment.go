package models

type OutcomeAlignment struct {
	ID              int64  `json:"id" url:"id,omitempty"`                             // the id of the aligned learning outcome..Example: 1
	AssignmentID    int64  `json:"assignment_id" url:"assignment_id,omitempty"`       // the id of the aligned assignment (null for live assessments)..Example: 2
	AssessmentID    int64  `json:"assessment_id" url:"assessment_id,omitempty"`       // the id of the aligned live assessment (null for assignments)..Example: 3
	SubmissionTypes string `json:"submission_types" url:"submission_types,omitempty"` // a string representing the different submission types of an aligned assignment..Example: online_text_entry,online_url
	Url             string `json:"url" url:"url,omitempty"`                           // the URL for the aligned assignment..Example: /courses/1/assignments/5
	Title           string `json:"title" url:"title,omitempty"`                       // the title of the aligned assignment..Example: Unit 1 test
}

func (t *OutcomeAlignment) HasError() error {
	return nil
}
