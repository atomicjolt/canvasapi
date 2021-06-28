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

// ListBlueprintSubscriptions Returns a list of blueprint subscriptions for the given course. (Currently a course may have no more than one.)
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type ListBlueprintSubscriptions struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListBlueprintSubscriptions) GetMethod() string {
	return "GET"
}

func (t *ListBlueprintSubscriptions) GetURLPath() string {
	path := "courses/{course_id}/blueprint_subscriptions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListBlueprintSubscriptions) GetQuery() (string, error) {
	return "", nil
}

func (t *ListBlueprintSubscriptions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListBlueprintSubscriptions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListBlueprintSubscriptions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListBlueprintSubscriptions) Do(c *canvasapi.Canvas) ([]*models.BlueprintSubscription, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.BlueprintSubscription{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
