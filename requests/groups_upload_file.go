package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GroupsUploadFile Upload a file to the group.
//
// This API endpoint is the first step in uploading a file to a group.
// See the {file:file_uploads.html File Upload Documentation} for details on
// the file upload workflow.
//
// Only those with the "Manage Files" permission on a group can upload files
// to the group. By default, this is anybody participating in the
// group, or any admin over the group.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type GroupsUploadFile struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GroupsUploadFile) GetMethod() string {
	return "POST"
}

func (t *GroupsUploadFile) GetURLPath() string {
	path := "groups/{group_id}/files"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GroupsUploadFile) GetQuery() (string, error) {
	return "", nil
}

func (t *GroupsUploadFile) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GroupsUploadFile) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GroupsUploadFile) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GroupsUploadFile) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
