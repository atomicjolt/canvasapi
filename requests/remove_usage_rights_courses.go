package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RemoveUsageRightsCourses Removes copyright and license information associated with one or more files
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.FileIDs (Required) List of ids of files to remove associated usage rights from.
// # Query.FolderIDs (Optional) List of ids of folders. Usage rights will be removed from all files in these folders.
//
type RemoveUsageRightsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		FileIDs   []string `json:"file_ids" url:"file_ids,omitempty"`     //  (Required)
		FolderIDs []string `json:"folder_ids" url:"folder_ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *RemoveUsageRightsCourses) GetMethod() string {
	return "DELETE"
}

func (t *RemoveUsageRightsCourses) GetURLPath() string {
	path := "courses/{course_id}/usage_rights"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *RemoveUsageRightsCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RemoveUsageRightsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveUsageRightsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveUsageRightsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Query.FileIDs == nil {
		errs = append(errs, "'Query.FileIDs' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveUsageRightsCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
