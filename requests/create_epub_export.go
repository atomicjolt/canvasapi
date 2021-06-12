package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateEpubExport Begin an ePub export for a course.
//
// You can use the {api:ProgressController#show Progress API} to track the
// progress of the export. The export's progress is linked to with the
// _progress_url_ value.
//
// When the export completes, use the {api:EpubExportsController#show Show content export} endpoint
// to retrieve a download URL for the exported content.
// https://canvas.instructure.com/doc/api/e_pub_exports.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type CreateEpubExport struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *CreateEpubExport) GetMethod() string {
	return "POST"
}

func (t *CreateEpubExport) GetURLPath() string {
	path := "courses/{course_id}/epub_exports"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateEpubExport) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateEpubExport) GetBody() (string, error) {
	return "", nil
}

func (t *CreateEpubExport) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateEpubExport) Do(c *canvasapi.Canvas) (*models.EpubExport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.EpubExport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
