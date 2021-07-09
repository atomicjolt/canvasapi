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

// GetImportDetails Show the changes that were propagated to a course associated with a blueprint.  See also
// {api:MasterCourses::MasterTemplatesController#migration_details the blueprint course side}.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.SubscriptionID (Required) ID
// # Path.ID (Required) ID
//
type GetImportDetails struct {
	Path struct {
		CourseID       string `json:"course_id" url:"course_id,omitempty"`             //  (Required)
		SubscriptionID string `json:"subscription_id" url:"subscription_id,omitempty"` //  (Required)
		ID             string `json:"id" url:"id,omitempty"`                           //  (Required)
	} `json:"path"`
}

func (t *GetImportDetails) GetMethod() string {
	return "GET"
}

func (t *GetImportDetails) GetURLPath() string {
	path := "courses/{course_id}/blueprint_subscriptions/{subscription_id}/migrations/{id}/details"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{subscription_id}", fmt.Sprintf("%v", t.Path.SubscriptionID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetImportDetails) GetQuery() (string, error) {
	return "", nil
}

func (t *GetImportDetails) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetImportDetails) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetImportDetails) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.SubscriptionID == "" {
		errs = append(errs, "'Path.SubscriptionID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetImportDetails) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.ChangeRecord, *canvasapi.PagedResource, error) {
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
	ret := []*models.ChangeRecord{}
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
