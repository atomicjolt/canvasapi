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

// DeletePeerReviewSections Delete a peer review for the assignment
// https://canvas.instructure.com/doc/api/peer_reviews.html
//
// Path Parameters:
// # Path.SectionID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.SubmissionID (Required) ID
//
// Query Parameters:
// # Query.UserID (Required) user_id to delete as reviewer on this assignment
//
type DeletePeerReviewSections struct {
	Path struct {
		SectionID    string `json:"section_id" url:"section_id,omitempty"`       //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		SubmissionID string `json:"submission_id" url:"submission_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		UserID int64 `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *DeletePeerReviewSections) GetMethod() string {
	return "DELETE"
}

func (t *DeletePeerReviewSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions/{submission_id}/peer_reviews"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	return path
}

func (t *DeletePeerReviewSections) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *DeletePeerReviewSections) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeletePeerReviewSections) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeletePeerReviewSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'Path.SectionID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.SubmissionID == "" {
		errs = append(errs, "'Path.SubmissionID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeletePeerReviewSections) Do(c *canvasapi.Canvas) (*models.PeerReview, error) {
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
