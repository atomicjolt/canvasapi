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

// ListBlueprintImports Shows a paginated list of migrations imported into a course associated with a blueprint, starting with the most recent. See also
// {api:MasterCourses::MasterTemplatesController#migrations_index the blueprint course side}.
//
// Use 'default' as the subscription_id to use the currently active blueprint subscription.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.SubscriptionID (Required) ID
//
type ListBlueprintImports struct {
	Path struct {
		CourseID       string `json:"course_id" url:"course_id,omitempty"`             //  (Required)
		SubscriptionID string `json:"subscription_id" url:"subscription_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListBlueprintImports) GetMethod() string {
	return "GET"
}

func (t *ListBlueprintImports) GetURLPath() string {
	path := "courses/{course_id}/blueprint_subscriptions/{subscription_id}/migrations"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{subscription_id}", fmt.Sprintf("%v", t.Path.SubscriptionID))
	return path
}

func (t *ListBlueprintImports) GetQuery() (string, error) {
	return "", nil
}

func (t *ListBlueprintImports) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListBlueprintImports) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListBlueprintImports) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.SubscriptionID == "" {
		errs = append(errs, "'Path.SubscriptionID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListBlueprintImports) Do(c *canvasapi.Canvas) ([]*models.BlueprintMigration, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.BlueprintMigration{}
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
