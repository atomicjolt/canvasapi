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

// CreateCourseSection Creates a new section for this course.
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # CourseSection (Optional) The name of the section
// # CourseSection (Optional) The sis ID of the section. Must have manage_sis permission to set. This is ignored if caller does not have permission to set.
// # CourseSection (Optional) The integration_id of the section. Must have manage_sis permission to set. This is ignored if caller does not have permission to set.
// # CourseSection (Optional) Section start date in ISO8601 format, e.g. 2011-01-01T01:00Z
// # CourseSection (Optional) Section end date in ISO8601 format. e.g. 2011-01-01T01:00Z
// # CourseSection (Optional) Set to true to restrict user enrollments to the start and end dates of the section.
// # EnableSISReactivation (Optional) When true, will first try to re-activate a deleted section with matching sis_section_id if possible.
//
type CreateCourseSection struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
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

		EnableSISReactivation bool `json:"enable_sis_reactivation"` //  (Optional)
	} `json:"form"`
}

func (t *CreateCourseSection) GetMethod() string {
	return "POST"
}

func (t *CreateCourseSection) GetURLPath() string {
	path := "courses/{course_id}/sections"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateCourseSection) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateCourseSection) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateCourseSection) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateCourseSection) Do(c *canvasapi.Canvas) (*models.Section, error) {
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
