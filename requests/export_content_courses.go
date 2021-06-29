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
	"github.com/atomicjolt/string_utils"
)

// ExportContentCourses Begin a content export job for a course, group, or user.
//
// You can use the {api:ProgressController#show Progress API} to track the
// progress of the export. The migration's progress is linked to with the
// _progress_url_ value.
//
// When the export completes, use the {api:ContentExportsApiController#show Show content export} endpoint
// to retrieve a download URL for the exported content.
// https://canvas.instructure.com/doc/api/content_exports.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.ExportType (Required) . Must be one of common_cartridge, qti, zip"common_cartridge":: Export the contents of the course in the Common Cartridge (.imscc) format
//    "qti":: Export quizzes from a course in the QTI format
//    "zip":: Export files from a course, group, or user in a zip file
// # Form.SkipNotifications (Optional) Don't send the notifications about the export to the user. Default: false
// # Form.Select (Optional) . Must be one of folders, files, attachments, quizzes, assignments, announcements, calendar_events, discussion_topics, modules, module_items, pages, rubricsThe select parameter allows exporting specific data. The keys are object types like 'files',
//    'folders', 'pages', etc. The value for each key is a list of object ids. An id can be an
//    integer or a string.
//
//    Multiple object types can be selected in the same call. However, not all object types are
//    valid for every export_type. Common Cartridge supports all object types. Zip and QTI only
//    support the object types as described below.
//
//    "folders":: Also supported for zip export_type.
//    "files":: Also supported for zip export_type.
//    "quizzes":: Also supported for qti export_type.
//
type ExportContentCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ExportType        string                   `json:"export_type" url:"export_type,omitempty"`               //  (Required) . Must be one of common_cartridge, qti, zip
		SkipNotifications bool                     `json:"skip_notifications" url:"skip_notifications,omitempty"` //  (Optional)
		Select            map[string](interface{}) `json:"select" url:"select,omitempty"`                         //  (Optional) . Must be one of folders, files, attachments, quizzes, assignments, announcements, calendar_events, discussion_topics, modules, module_items, pages, rubrics
	} `json:"form"`
}

func (t *ExportContentCourses) GetMethod() string {
	return "POST"
}

func (t *ExportContentCourses) GetURLPath() string {
	path := "courses/{course_id}/content_exports"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ExportContentCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ExportContentCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ExportContentCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ExportContentCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.ExportType == "" {
		errs = append(errs, "'Form.ExportType' is required")
	}
	if t.Form.ExportType != "" && !string_utils.Include([]string{"common_cartridge", "qti", "zip"}, t.Form.ExportType) {
		errs = append(errs, "ExportType must be one of common_cartridge, qti, zip")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ExportContentCourses) Do(c *canvasapi.Canvas) (*models.ContentExport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ContentExport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
