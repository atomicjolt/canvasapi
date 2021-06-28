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

// ShowOriginalityReportFiles Get a single originality report
// https://canvas.instructure.com/doc/api/originality_reports.html
//
// Path Parameters:
// # AssignmentID (Required) ID
// # FileID (Required) ID
//
type ShowOriginalityReportFiles struct {
	Path struct {
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		FileID       string `json:"file_id" url:"file_id,omitempty"`             //  (Required)
	} `json:"path"`
}

func (t *ShowOriginalityReportFiles) GetMethod() string {
	return "GET"
}

func (t *ShowOriginalityReportFiles) GetURLPath() string {
	path := "/lti/assignments/{assignment_id}/files/{file_id}/originality_report"
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{file_id}", fmt.Sprintf("%v", t.Path.FileID))
	return path
}

func (t *ShowOriginalityReportFiles) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowOriginalityReportFiles) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowOriginalityReportFiles) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowOriginalityReportFiles) HasErrors() error {
	errs := []string{}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.FileID == "" {
		errs = append(errs, "'FileID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowOriginalityReportFiles) Do(c *canvasapi.Canvas) (*models.OriginalityReport, error) {
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
