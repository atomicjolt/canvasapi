package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ExportContentGroups Begin a content export job for a course, group, or user.
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
// # GroupID (Required) ID
//
// Form Parameters:
// # ExportType (Required) . Must be one of common_cartridge, qti, zip"common_cartridge":: Export the contents of the course in the Common Cartridge (.imscc) format
//    "qti":: Export quizzes from a course in the QTI format
//    "zip":: Export files from a course, group, or user in a zip file
// # SkipNotifications (Optional) Don't send the notifications about the export to the user. Default: false
// # Select (Optional) . Must be one of folders, files, attachments, quizzes, assignments, announcements, calendar_events, discussion_topics, modules, module_items, pages, rubricsThe select parameter allows exporting specific data. The keys are object types like 'files',
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
type ExportContentGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Form struct {
		ExportType        string `json:"export_type"`        //  (Required) . Must be one of common_cartridge, qti, zip
		SkipNotifications bool   `json:"skip_notifications"` //  (Optional)
		Select            string `json:"select"`             //  (Optional) . Must be one of folders, files, attachments, quizzes, assignments, announcements, calendar_events, discussion_topics, modules, module_items, pages, rubrics
	} `json:"form"`
}

func (t *ExportContentGroups) GetMethod() string {
	return "POST"
}

func (t *ExportContentGroups) GetURLPath() string {
	path := "groups/{group_id}/content_exports"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ExportContentGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ExportContentGroups) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ExportContentGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Form.ExportType == "" {
		errs = append(errs, "'ExportType' is required")
	}
	if !string_utils.Include([]string{"common_cartridge", "qti", "zip"}, t.Form.ExportType) {
		errs = append(errs, "ExportType must be one of common_cartridge, qti, zip")
	}
	if !string_utils.Include([]string{"folders", "files", "attachments", "quizzes", "assignments", "announcements", "calendar_events", "discussion_topics", "modules", "module_items", "pages", "rubrics"}, t.Form.Select) {
		errs = append(errs, "Select must be one of folders, files, attachments, quizzes, assignments, announcements, calendar_events, discussion_topics, modules, module_items, pages, rubrics")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ExportContentGroups) Do(c *canvasapi.Canvas) (*models.ContentExport, error) {
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
