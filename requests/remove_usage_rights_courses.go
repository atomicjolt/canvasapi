package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RemoveUsageRightsCourses Removes copyright and license information associated with one or more files
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # FileIDs (Required) List of ids of files to remove associated usage rights from.
// # FolderIDs (Optional) List of ids of folders. Usage rights will be removed from all files in these folders.
//
type RemoveUsageRightsCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		FileIDs   []string `json:"file_ids"`   //  (Required)
		FolderIDs []string `json:"folder_ids"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *RemoveUsageRightsCourses) GetBody() (string, error) {
	return "", nil
}

func (t *RemoveUsageRightsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Query.FileIDs == nil {
		errs = append(errs, "'FileIDs' is required")
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
