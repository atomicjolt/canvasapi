package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # DestFolderID (Required) ID
//
// Form Parameters:
// # SourceFolderID (Required) The id of the source folder
//
type CopyFolder struct {
	Path struct {
		DestFolderID string `json:"dest_folder_id"` //  (Required)
	} `json:"path"`

	Form struct {
		SourceFolderID string `json:"source_folder_id"` //  (Required)
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

func (t *CopyFolder) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CopyFolder) HasErrors() error {
	errs := []string{}
	if t.Path.DestFolderID == "" {
		errs = append(errs, "'DestFolderID' is required")
	}
	if t.Form.SourceFolderID == "" {
		errs = append(errs, "'SourceFolderID' is required")
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