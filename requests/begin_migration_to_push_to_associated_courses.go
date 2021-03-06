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

// BeginMigrationToPushToAssociatedCourses Begins a migration to push recently updated content to all associated courses.
// Only one migration can be running at a time.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TemplateID (Required) ID
//
// Form Parameters:
// # Form.Comment (Optional) An optional comment to be included in the sync history.
// # Form.SendNotification (Optional) Send a notification to the calling user when the sync completes.
// # Form.CopySettings (Optional) Whether course settings should be copied over to associated courses.
//    Defaults to true for newly associated courses.
// # Form.PublishAfterInitialSync (Optional) If set, newly associated courses will be automatically published after the sync completes
//
type BeginMigrationToPushToAssociatedCourses struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`     //  (Required)
		TemplateID string `json:"template_id" url:"template_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Comment                 string `json:"comment" url:"comment,omitempty"`                                       //  (Optional)
		SendNotification        bool   `json:"send_notification" url:"send_notification,omitempty"`                   //  (Optional)
		CopySettings            bool   `json:"copy_settings" url:"copy_settings,omitempty"`                           //  (Optional)
		PublishAfterInitialSync bool   `json:"publish_after_initial_sync" url:"publish_after_initial_sync,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *BeginMigrationToPushToAssociatedCourses) GetMethod() string {
	return "POST"
}

func (t *BeginMigrationToPushToAssociatedCourses) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}/migrations"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	return path
}

func (t *BeginMigrationToPushToAssociatedCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *BeginMigrationToPushToAssociatedCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *BeginMigrationToPushToAssociatedCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *BeginMigrationToPushToAssociatedCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'Path.TemplateID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *BeginMigrationToPushToAssociatedCourses) Do(c *canvasapi.Canvas) (*models.BlueprintMigration, error) {
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
