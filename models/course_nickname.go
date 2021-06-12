package models

type CourseNickname struct {
	CourseID int64  `json:"course_id"` // the ID of the course.Example: 88
	Name     string `json:"name"`      // the actual name of the course.Example: S1048576 DPMS1200 Intro to Newtonian Mechanics
	Nickname string `json:"nickname"`  // the calling user's nickname for the course.Example: Physics
}

func (t *CourseNickname) HasError() error {
	return nil
}
