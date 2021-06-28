package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// PeerReviewsCreatePeerReviewCourses Create a peer review for the assignment
// https://canvas.instructure.com/doc/api/peer_reviews.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
// # SubmissionID (Required) ID
//
// Form Parameters:
// # UserID (Required) user_id to assign as reviewer on this assignment
//
type PeerReviewsCreatePeerReviewCourses struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		SubmissionID string `json:"submission_id" url:"submission_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		UserID int64 `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *PeerReviewsCreatePeerReviewCourses) GetMethod() string {
	return "POST"
}

func (t *PeerReviewsCreatePeerReviewCourses) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{submission_id}/peer_reviews"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	return path
}

func (t *PeerReviewsCreatePeerReviewCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *PeerReviewsCreatePeerReviewCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *PeerReviewsCreatePeerReviewCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *PeerReviewsCreatePeerReviewCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.SubmissionID == "" {
		errs = append(errs, "'SubmissionID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *PeerReviewsCreatePeerReviewCourses) Do(c *canvasapi.Canvas) (*models.PeerReview, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PeerReview{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
