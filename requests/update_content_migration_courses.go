package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateContentMigrationCourses Update a content migration. Takes same arguments as {api:ContentMigrationsController#create create} except that you
// can't change the migration type. However, changing most settings after the
// migration process has started will not do anything. Generally updating the
// content migration will be used when there is a file upload problem, or when
// importing content selectively. If the first upload has a problem you can
// supply new _pre_attachment_ values to start the process again.
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type UpdateContentMigrationCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *UpdateContentMigrationCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateContentMigrationCourses) GetURLPath() string {
	path := "courses/{course_id}/content_migrations/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateContentMigrationCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateContentMigrationCourses) GetBody() (string, error) {
	return "", nil
}

func (t *UpdateContentMigrationCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateContentMigrationCourses) Do(c *canvasapi.Canvas) (*models.ContentMigration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ContentMigration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
