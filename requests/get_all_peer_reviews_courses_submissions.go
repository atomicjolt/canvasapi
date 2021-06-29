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
	"github.com/atomicjolt/string_utils"
)

// GetAllPeerReviewsCoursesSubmissions Get a list of all Peer Reviews for this assignment
// https://canvas.instructure.com/doc/api/peer_reviews.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.SubmissionID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of submission_comments, userAssociations to include with the peer review.
//
type GetAllPeerReviewsCoursesSubmissions struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		SubmissionID string `json:"submission_id" url:"submission_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of submission_comments, user
	} `json:"query"`
}

func (t *GetAllPeerReviewsCoursesSubmissions) GetMethod() string {
	return "GET"
}

func (t *GetAllPeerReviewsCoursesSubmissions) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{submission_id}/peer_reviews"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	return path
}

func (t *GetAllPeerReviewsCoursesSubmissions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetAllPeerReviewsCoursesSubmissions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllPeerReviewsCoursesSubmissions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllPeerReviewsCoursesSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.SubmissionID == "" {
		errs = append(errs, "'Path.SubmissionID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"submission_comments", "user"}, v) {
			errs = append(errs, "Include must be one of submission_comments, user")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllPeerReviewsCoursesSubmissions) Do(c *canvasapi.Canvas) ([]*models.PeerReview, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.PeerReview{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
