package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// EditSubmissionComment Edit the given submission comment.
// https://canvas.instructure.com/doc/api/submission_comments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
// # UserID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Comment (Optional) If this argument is present, edit the text of a comment.
//
type EditSubmissionComment struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
		UserID       string `json:"user_id"`       //  (Required)
		ID           string `json:"id"`            //  (Required)
	} `json:"path"`

	Form struct {
		Comment string `json:"comment"` //  (Optional)
	} `json:"form"`
}

func (t *EditSubmissionComment) GetMethod() string {
	return "PUT"
}

func (t *EditSubmissionComment) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions/{user_id}/comments/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *EditSubmissionComment) GetQuery() (string, error) {
	return "", nil
}

func (t *EditSubmissionComment) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *EditSubmissionComment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditSubmissionComment) Do(c *canvasapi.Canvas) (*models.SubmissionComment, error) {
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