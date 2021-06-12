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

// CreateNewGradingStandardCourses Create a new grading standard
//
// If grading_scheme_entry arguments are omitted, then a default grading scheme
// will be set. The default scheme is as follows:
//      "A" : 94,
//      "A-" : 90,
//      "B+" : 87,
//      "B" : 84,
//      "B-" : 80,
//      "C+" : 77,
//      "C" : 74,
//      "C-" : 70,
//      "D+" : 67,
//      "D" : 64,
//      "D-" : 61,
//      "F" : 0,
// https://canvas.instructure.com/doc/api/grading_standards.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Title (Required) The title for the Grading Standard.
// # GradingSchemeEntry (Required) The name for an entry value within a GradingStandard that describes the range of the value
//    e.g. A-
// # GradingSchemeEntry (Required) The value for the name of the entry within a GradingStandard.
//    The entry represents the lower bound of the range for the entry.
//    This range includes the value up to the next entry in the GradingStandard,
//    or 100 if there is no upper bound. The lowest value will have a lower bound range of 0.
//    e.g. 93
//
type CreateNewGradingStandardCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Title              string `json:"title"` //  (Required)
		GradingSchemeEntry struct {
			Name  []string `json:"name"`  //  (Required)
			Value []int64  `json:"value"` //  (Required)
		} `json:"grading_scheme_entry"`
	} `json:"form"`
}

func (t *CreateNewGradingStandardCourses) GetMethod() string {
	return "POST"
}

func (t *CreateNewGradingStandardCourses) GetURLPath() string {
	path := "courses/{course_id}/grading_standards"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateNewGradingStandardCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateNewGradingStandardCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateNewGradingStandardCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Title == "" {
		errs = append(errs, "'Title' is required")
	}
	if t.Form.GradingSchemeEntry.Name == nil {
		errs = append(errs, "'GradingSchemeEntry' is required")
	}
	if t.Form.GradingSchemeEntry.Value == nil {
		errs = append(errs, "'GradingSchemeEntry' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateNewGradingStandardCourses) Do(c *canvasapi.Canvas) (*models.GradingStandard, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.GradingStandard{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
