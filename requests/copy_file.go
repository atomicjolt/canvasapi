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

// CopyFile Copy a file from elsewhere in Canvas into a folder.
//
// Copying a file across contexts (between courses and users) is permitted,
// but the source and destination must belong to the same institution.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # DestFolderID (Required) ID
//
// Form Parameters:
// # SourceFileID (Required) The id of the source file
// # OnDuplicate (Optional) . Must be one of overwrite, renameWhat to do if a file with the same name already exists at the destination.
//    If such a file exists and this parameter is not given, the call will fail.
//
//    "overwrite":: Replace an existing file with the same name
//    "rename":: Add a qualifier to make the new filename unique
//
type CopyFile struct {
	Path struct {
		DestFolderID string `json:"dest_folder_id"` //  (Required)
	} `json:"path"`

	Form struct {
		SourceFileID string `json:"source_file_id"` //  (Required)
		OnDuplicate  string `json:"on_duplicate"`   //  (Optional) . Must be one of overwrite, rename
	} `json:"form"`
}

func (t *CopyFile) GetMethod() string {
	return "POST"
}

func (t *CopyFile) GetURLPath() string {
	path := "folders/{dest_folder_id}/copy_file"
	path = strings.ReplaceAll(path, "{dest_folder_id}", fmt.Sprintf("%v", t.Path.DestFolderID))
	return path
}

func (t *CopyFile) GetQuery() (string, error) {
	return "", nil
}

func (t *CopyFile) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CopyFile) HasErrors() error {
	errs := []string{}
	if t.Path.DestFolderID == "" {
		errs = append(errs, "'DestFolderID' is required")
	}
	if t.Form.SourceFileID == "" {
		errs = append(errs, "'SourceFileID' is required")
	}
	if !string_utils.Include([]string{"overwrite", "rename"}, t.Form.OnDuplicate) {
		errs = append(errs, "OnDuplicate must be one of overwrite, rename")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CopyFile) Do(c *canvasapi.Canvas) (*models.File, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.File{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
