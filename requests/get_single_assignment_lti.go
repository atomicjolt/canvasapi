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

// GetSingleAssignmentLti Get a single Canvas assignment by Canvas id or LTI id. Tool providers may only access
// assignments that are associated with their tool.
// https://canvas.instructure.com/doc/api/plagiarism_detection_platform_assignments.html
//
// Path Parameters:
// # AssignmentID (Required) ID
//
// Query Parameters:
// # UserID (Optional) The id of the user. Can be a Canvas or LTI id for the user.
//
type GetSingleAssignmentLti struct {
	Path struct {
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetSingleAssignmentLti) GetMethod() string {
	return "GET"
}

func (t *GetSingleAssignmentLti) GetURLPath() string {
	path := "/lti/assignments/{assignment_id}"
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *GetSingleAssignmentLti) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetSingleAssignmentLti) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleAssignmentLti) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleAssignmentLti) HasErrors() error {
	errs := []string{}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleAssignmentLti) Do(c *canvasapi.Canvas) (*models.LtiAssignment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.LtiAssignment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
