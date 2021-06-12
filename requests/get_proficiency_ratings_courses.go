package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetProficiencyRatingsCourses Get account-level proficiency ratings. If not defined for this account,
// it will return proficiency ratings for the nearest super-account with ratings defined.
// Will return 404 if none found.
//
//   Examples:
//     curl https://<canvas>/api/v1/accounts/<account_id>/outcome_proficiency \
//         -H 'Authorization: Bearer <token>'
// https://canvas.instructure.com/doc/api/proficiency_ratings.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetProficiencyRatingsCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *GetProficiencyRatingsCourses) GetMethod() string {
	return "GET"
}

func (t *GetProficiencyRatingsCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_proficiency"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetProficiencyRatingsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetProficiencyRatingsCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetProficiencyRatingsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetProficiencyRatingsCourses) Do(c *canvasapi.Canvas) (*models.Proficiency, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Proficiency{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
