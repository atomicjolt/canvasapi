package models

import (
	"time"
)

type Day struct {
	Date    time.Time `json:"date"`    // the date represented by this entry.Example: 1986-08-09
	Graders int64     `json:"graders"` // an array of the graders who were responsible for the submissions in this response. the submissions are grouped according to the person who graded them and the assignment they were submitted for..Example: []
}

func (t *Day) HasError() error {
	return nil
}
