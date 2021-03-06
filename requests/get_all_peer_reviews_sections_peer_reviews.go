package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// GetAllPeerReviewsSectionsPeerReviews Get a list of all Peer Reviews for this assignment
// https://canvas.instructure.com/doc/api/peer_reviews.html
//
// Path Parameters:
// # Path.SectionID (Required) ID
// # Path.AssignmentID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of submission_comments, userAssociations to include with the peer review.
//
type GetAllPeerReviewsSectionsPeerReviews struct {
	Path struct {
		SectionID    string `json:"section_id" url:"section_id,omitempty"`       //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of submission_comments, user
	} `json:"query"`
}

func (t *GetAllPeerReviewsSectionsPeerReviews) GetMethod() string {
	return "GET"
}

func (t *GetAllPeerReviewsSectionsPeerReviews) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/peer_reviews"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *GetAllPeerReviewsSectionsPeerReviews) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetAllPeerReviewsSectionsPeerReviews) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllPeerReviewsSectionsPeerReviews) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllPeerReviewsSectionsPeerReviews) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'Path.SectionID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
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

func (t *GetAllPeerReviewsSectionsPeerReviews) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.PeerReview, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.PeerReview{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
