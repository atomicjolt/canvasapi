package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// UsersUploadFile Upload a file to the user's personal files section.
//
// This API endpoint is the first step in uploading a file to a user's files.
// See the {file:file_uploads.html File Upload Documentation} for details on
// the file upload workflow.
//
// Note that typically users will only be able to upload files to their
// own files section. Passing a user_id of +self+ is an easy shortcut
// to specify the current user.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # UserID (Required) ID
//
type UsersUploadFile struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`
}

func (t *UsersUploadFile) GetMethod() string {
	return "POST"
}

func (t *UsersUploadFile) GetURLPath() string {
	path := "users/{user_id}/files"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *UsersUploadFile) GetQuery() (string, error) {
	return "", nil
}

func (t *UsersUploadFile) GetBody() (string, error) {
	return "", nil
}

func (t *UsersUploadFile) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UsersUploadFile) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
