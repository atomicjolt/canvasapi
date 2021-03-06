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

// ShowBlueprintImport Shows the status of an import into a course associated with a blueprint. See also
// {api:MasterCourses::MasterTemplatesController#migrations_show the blueprint course side}.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.SubscriptionID (Required) ID
// # Path.ID (Required) ID
//
type ShowBlueprintImport struct {
	Path struct {
		CourseID       string `json:"course_id" url:"course_id,omitempty"`             //  (Required)
		SubscriptionID string `json:"subscription_id" url:"subscription_id,omitempty"` //  (Required)
		ID             string `json:"id" url:"id,omitempty"`                           //  (Required)
	} `json:"path"`
}

func (t *ShowBlueprintImport) GetMethod() string {
	return "GET"
}

func (t *ShowBlueprintImport) GetURLPath() string {
	path := "courses/{course_id}/blueprint_subscriptions/{subscription_id}/migrations/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{subscription_id}", fmt.Sprintf("%v", t.Path.SubscriptionID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowBlueprintImport) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowBlueprintImport) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowBlueprintImport) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowBlueprintImport) HasErrors() error {
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

func (t *ShowBlueprintImport) Do(c *canvasapi.Canvas) (*models.BlueprintMigration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.BlueprintMigration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
