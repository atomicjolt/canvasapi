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

// GetFeatureFlagCourses Get the feature flag that applies to a given Account, Course, or User.
// The flag may be defined on the object, or it may be inherited from a parent
// account. You can look at the context_id and context_type of the returned object
// to determine which is the case. If these fields are missing, then the object
// is the global Canvas default.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.Feature (Required) ID
//
type GetFeatureFlagCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		Feature  string `json:"feature" url:"feature,omitempty"`     //  (Required)
	} `json:"path"`
}

func (t *GetFeatureFlagCourses) GetMethod() string {
	return "GET"
}

func (t *GetFeatureFlagCourses) GetURLPath() string {
	path := "courses/{course_id}/features/flags/{feature}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{feature}", fmt.Sprintf("%v", t.Path.Feature))
	return path
}

func (t *GetFeatureFlagCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetFeatureFlagCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetFeatureFlagCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetFeatureFlagCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.Feature == "" {
		errs = append(errs, "'Path.Feature' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFeatureFlagCourses) Do(c *canvasapi.Canvas) (*models.FeatureFlag, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.FeatureFlag{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
