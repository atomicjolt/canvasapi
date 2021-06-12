package models

type PollSession struct {
	ID               int64           `json:"id"`                 // The unique identifier for the poll session..Example: 1023
	PollID           int64           `json:"poll_id"`            // The id of the Poll this poll session is associated with.Example: 55
	CourseID         int64           `json:"course_id"`          // The id of the Course this poll session is associated with.Example: 1111
	CourseSectionID  int64           `json:"course_section_id"`  // The id of the Course Section this poll session is associated with.Example: 444
	IsPublished      bool            `json:"is_published"`       // Specifies whether or not this poll session has been published for students to participate in..Example: true
	HasPublicResults bool            `json:"has_public_results"` // Specifies whether the results are viewable by students..Example: true
	CreatedAt        string          `json:"created_at"`         // The time at which the poll session was created..Example: 2014-01-07T15:16:18Z
	Results          string          `json:"results"`            // The results of the submissions of the poll. Each key is the poll choice id, and the value is the count of submissions..Example: 10, 3, 27, 8
	PollSubmissions  *PollSubmission `json:"poll_submissions"`   // If the poll session has public results, this will return an array of all submissions, viewable by both students and teachers. If the results are not public, for students it will return their submission only..
}

func (t *PollSession) HasError() error {
	return nil
}
