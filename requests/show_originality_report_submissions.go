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

// ShowOriginalityReportSubmissions Get a single originality report
// https://canvas.instructure.com/doc/api/originality_reports.html
//
// Path Parameters:
// # Path.AssignmentID (Required) ID
// # Path.SubmissionID (Required) ID
// # Path.ID (Required) ID
//
type ShowOriginalityReportSubmissions struct {
	Path struct {
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		SubmissionID string `json:"submission_id" url:"submission_id,omitempty"` //  (Required)
		ID           string `json:"id" url:"id,omitempty"`                       //  (Required)
	} `json:"path"`
}

func (t *ShowOriginalityReportSubmissions) GetMethod() string {
	return "GET"
}

func (t *ShowOriginalityReportSubmissions) GetURLPath() string {
	path := "/lti/assignments/{assignment_id}/submissions/{submission_id}/originality_report/{id}"
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowOriginalityReportSubmissions) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowOriginalityReportSubmissions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowOriginalityReportSubmissions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowOriginalityReportSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.SubmissionID == "" {
		errs = append(errs, "'Path.SubmissionID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowOriginalityReportSubmissions) Do(c *canvasapi.Canvas) (*models.OriginalityReport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OriginalityReport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
