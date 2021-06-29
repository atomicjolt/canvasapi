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

// GetAllPeerReviewsCoursesPeerReviews Get a list of all Peer Reviews for this assignment
// https://canvas.instructure.com/doc/api/peer_reviews.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of submission_comments, userAssociations to include with the peer review.
//
type GetAllPeerReviewsCoursesPeerReviews struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of submission_comments, user
	} `json:"query"`
}

func (t *GetAllPeerReviewsCoursesPeerReviews) GetMethod() string {
	return "GET"
}

func (t *GetAllPeerReviewsCoursesPeerReviews) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/peer_reviews"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *GetAllPeerReviewsCoursesPeerReviews) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetAllPeerReviewsCoursesPeerReviews) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllPeerReviewsCoursesPeerReviews) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllPeerReviewsCoursesPeerReviews) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
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

func (t *GetAllPeerReviewsCoursesPeerReviews) Do(c *canvasapi.Canvas) ([]*models.PeerReview, error) {
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
