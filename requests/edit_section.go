package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # ID (Required) ID
//
// Form Parameters:
// # CourseSection (Optional) The name of the section
// # CourseSection (Optional) The sis ID of the section. Must have manage_sis permission to set.
// # CourseSection (Optional) The integration_id of the section. Must have manage_sis permission to set.
// # CourseSection (Optional) Section start date in ISO8601 format, e.g. 2011-01-01T01:00Z
// # CourseSection (Optional) Section end date in ISO8601 format. e.g. 2011-01-01T01:00Z
// # CourseSection (Optional) Set to true to restrict user enrollments to the start and end dates of the section.
//
type EditSection struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		CourseSection struct {
			Name                              string    `json:"name"`                                  //  (Optional)
			SISSectionID                      string    `json:"sis_section_id"`                        //  (Optional)
			IntegrationID                     string    `json:"integration_id"`                        //  (Optional)
			StartAt                           time.Time `json:"start_at"`                              //  (Optional)
			EndAt                             time.Time `json:"end_at"`                                //  (Optional)
			RestrictEnrollmentsToSectionDates bool      `json:"restrict_enrollments_to_section_dates"` //  (Optional)
		} `json:"course_section"`
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

func (t *EditSection) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *EditSection) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
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
