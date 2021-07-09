package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetAssociatedCourseInformation Returns a list of courses that are configured to receive updates from this blueprint
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TemplateID (Required) ID
//
type GetAssociatedCourseInformation struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`     //  (Required)
		TemplateID string `json:"template_id" url:"template_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetAssociatedCourseInformation) GetMethod() string {
	return "GET"
}

func (t *GetAssociatedCourseInformation) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}/associated_courses"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	return path
}

func (t *GetAssociatedCourseInformation) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAssociatedCourseInformation) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAssociatedCourseInformation) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAssociatedCourseInformation) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'Path.TemplateID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAssociatedCourseInformation) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Course, *canvasapi.PagedResource, error) {
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
	ret := []*models.Course{}
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
