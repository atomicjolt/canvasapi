package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeleteSubmissionComment Delete the given submission comment.
// https://canvas.instructure.com/doc/api/submission_comments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.UserID (Required) ID
// # Path.ID (Required) ID
//
type DeleteSubmissionComment struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		UserID       string `json:"user_id" url:"user_id,omitempty"`             //  (Required)
		ID           string `json:"id" url:"id,omitempty"`                       //  (Required)
	} `json:"path"`
}

func (t *DeleteSubmissionComment) GetMethod() string {
	return "DELETE"
}

func (t *DeleteSubmissionComment) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{user_id}/comments/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteSubmissionComment) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteSubmissionComment) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteSubmissionComment) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteSubmissionComment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteSubmissionComment) Do(c *canvasapi.Canvas) (*models.SubmissionComment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.SubmissionComment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
