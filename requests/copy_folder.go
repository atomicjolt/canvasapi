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

// CopyFolder Copy a folder (and its contents) from elsewhere in Canvas into a folder.
//
// Copying a folder across contexts (between courses and users) is permitted,
// but the source and destination must belong to the same institution.
// If the source and destination folders are in the same context, the
// source folder may not contain the destination folder. A folder will be
// renamed at its destination if another folder with the same name already
// exists.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.DestFolderID (Required) ID
//
// Form Parameters:
// # Form.SourceFolderID (Required) The id of the source folder
//
type CopyFolder struct {
	Path struct {
		DestFolderID string `json:"dest_folder_id" url:"dest_folder_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		SourceFolderID string `json:"source_folder_id" url:"source_folder_id,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *CopyFolder) GetMethod() string {
	return "POST"
}

func (t *CopyFolder) GetURLPath() string {
	path := "folders/{dest_folder_id}/copy_folder"
	path = strings.ReplaceAll(path, "{dest_folder_id}", fmt.Sprintf("%v", t.Path.DestFolderID))
	return path
}

func (t *CopyFolder) GetQuery() (string, error) {
	return "", nil
}

func (t *CopyFolder) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CopyFolder) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CopyFolder) HasErrors() error {
	errs := []string{}
	if t.Path.DestFolderID == "" {
		errs = append(errs, "'Path.DestFolderID' is required")
	}
	if t.Form.SourceFolderID == "" {
		errs = append(errs, "'Form.SourceFolderID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CopyFolder) Do(c *canvasapi.Canvas) (*models.Folder, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Folder{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
