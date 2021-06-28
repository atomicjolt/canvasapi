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

// CreateUpdateProficiencyRatingsCourses Create or update account-level proficiency ratings. These ratings will apply to all
// sub-accounts, unless they have their own account-level proficiency ratings defined.
// https://canvas.instructure.com/doc/api/proficiency_ratings.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Ratings (Optional) The description of the rating level.
// # Ratings (Optional) The non-negative number of points of the rating level. Points across ratings should be strictly decreasing in value.
// # Ratings (Optional) Indicates the rating level where mastery is first achieved. Only one rating in a proficiency should be marked for mastery.
// # Ratings (Optional) The color associated with the rating level. Should be a hex color code like '00FFFF'.
//
type CreateUpdateProficiencyRatingsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Ratings struct {
			Description []string `json:"description" url:"description,omitempty"` //  (Optional)
			Points      []int64  `json:"points" url:"points,omitempty"`           //  (Optional)
			Mastery     []int64  `json:"mastery" url:"mastery,omitempty"`         //  (Optional)
			Color       []int64  `json:"color" url:"color,omitempty"`             //  (Optional)
		} `json:"ratings" url:"ratings,omitempty"`
	} `json:"form"`
}

func (t *CreateUpdateProficiencyRatingsCourses) GetMethod() string {
	return "POST"
}

func (t *CreateUpdateProficiencyRatingsCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_proficiency"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateUpdateProficiencyRatingsCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateUpdateProficiencyRatingsCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateUpdateProficiencyRatingsCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateUpdateProficiencyRatingsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateUpdateProficiencyRatingsCourses) Do(c *canvasapi.Canvas) (*models.Proficiency, error) {
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
