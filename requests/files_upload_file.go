package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// FilesUploadFile Upload a file to a folder.
//
// This API endpoint is the first step in uploading a file.
// See the {file:file_uploads.html File Upload Documentation} for details on
// the file upload workflow.
//
// Only those with the "Manage Files" permission on a course or group can
// upload files to a folder in that course or group.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # FolderID (Required) ID
//
type FilesUploadFile struct {
	Path struct {
		FolderID string `json:"folder_id" url:"folder_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *FilesUploadFile) GetMethod() string {
	return "POST"
}

func (t *FilesUploadFile) GetURLPath() string {
	path := "folders/{folder_id}/files"
	path = strings.ReplaceAll(path, "{folder_id}", fmt.Sprintf("%v", t.Path.FolderID))
	return path
}

func (t *FilesUploadFile) GetQuery() (string, error) {
	return "", nil
}

func (t *FilesUploadFile) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *FilesUploadFile) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *FilesUploadFile) HasErrors() error {
	errs := []string{}
	if t.Path.FolderID == "" {
		errs = append(errs, "'FolderID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *FilesUploadFile) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
