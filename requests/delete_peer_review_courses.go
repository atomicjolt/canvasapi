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

// DeletePeerReviewCourses Delete a peer review for the assignment
// https://canvas.instructure.com/doc/api/peer_reviews.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
// # SubmissionID (Required) ID
//
// Query Parameters:
// # UserID (Required) user_id to delete as reviewer on this assignment
//
type DeletePeerReviewCourses struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		SubmissionID string `json:"submission_id" url:"submission_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		UserID int64 `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *DeletePeerReviewCourses) GetMethod() string {
	return "DELETE"
}

func (t *DeletePeerReviewCourses) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{submission_id}/peer_reviews"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	return path
}

func (t *DeletePeerReviewCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *DeletePeerReviewCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeletePeerReviewCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeletePeerReviewCourses) HasErrors() error {
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

func (t *DeletePeerReviewCourses) Do(c *canvasapi.Canvas) (*models.PeerReview, error) {
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
