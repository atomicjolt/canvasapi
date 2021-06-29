package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// EditSection Modify an existing section.
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.CourseSection.Name (Optional) The name of the section
// # Form.CourseSection.SISSectionID (Optional) The sis ID of the section. Must have manage_sis permission to set.
// # Form.CourseSection.IntegrationID (Optional) The integration_id of the section. Must have manage_sis permission to set.
// # Form.CourseSection.StartAt (Optional) Section start date in ISO8601 format, e.g. 2011-01-01T01:00Z
// # Form.CourseSection.EndAt (Optional) Section end date in ISO8601 format. e.g. 2011-01-01T01:00Z
// # Form.CourseSection.RestrictEnrollmentsToSectionDates (Optional) Set to true to restrict user enrollments to the start and end dates of the section.
//
type EditSection struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		CourseSection struct {
			Name                              string    `json:"name" url:"name,omitempty"`                                                                   //  (Optional)
			SISSectionID                      string    `json:"sis_section_id" url:"sis_section_id,omitempty"`                                               //  (Optional)
			IntegrationID                     string    `json:"integration_id" url:"integration_id,omitempty"`                                               //  (Optional)
			StartAt                           time.Time `json:"start_at" url:"start_at,omitempty"`                                                           //  (Optional)
			EndAt                             time.Time `json:"end_at" url:"end_at,omitempty"`                                                               //  (Optional)
			RestrictEnrollmentsToSectionDates bool      `json:"restrict_enrollments_to_section_dates" url:"restrict_enrollments_to_section_dates,omitempty"` //  (Optional)
		} `json:"course_section" url:"course_section,omitempty"`
	} `json:"form"`
}

func (t *EditSection) GetMethod() string {
	return "PUT"
}

func (t *EditSection) GetURLPath() string {
	path := "sections/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *EditSection) GetQuery() (string, error) {
	return "", nil
}

func (t *EditSection) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *EditSection) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *EditSection) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditSection) Do(c *canvasapi.Canvas) (*models.Section, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Section{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
