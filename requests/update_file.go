package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// UpdateFile Update some settings on the specified file
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Name (Optional) The new display name of the file, with a limit of 255 characters.
// # ParentFolderID (Optional) The id of the folder to move this file into.
//    The new folder must be in the same context as the original parent folder.
//    If the file is in a context without folders this does not apply.
// # OnDuplicate (Optional) . Must be one of overwrite, renameIf the file is moved to a folder containing a file with the same name,
//    or renamed to a name matching an existing file, the API call will fail
//    unless this parameter is supplied.
//
//    "overwrite":: Replace the existing file with the same name
//    "rename":: Add a qualifier to make the new filename unique
// # LockAt (Optional) The datetime to lock the file at
// # UnlockAt (Optional) The datetime to unlock the file at
// # Locked (Optional) Flag the file as locked
// # Hidden (Optional) Flag the file as hidden
//
type UpdateFile struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Name           string    `json:"name"`             //  (Optional)
		ParentFolderID string    `json:"parent_folder_id"` //  (Optional)
		OnDuplicate    string    `json:"on_duplicate"`     //  (Optional) . Must be one of overwrite, rename
		LockAt         time.Time `json:"lock_at"`          //  (Optional)
		UnlockAt       time.Time `json:"unlock_at"`        //  (Optional)
		Locked         bool      `json:"locked"`           //  (Optional)
		Hidden         bool      `json:"hidden"`           //  (Optional)
	} `json:"form"`
}

func (t *UpdateFile) GetMethod() string {
	return "PUT"
}

func (t *UpdateFile) GetURLPath() string {
	path := "files/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateFile) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateFile) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateFile) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if !string_utils.Include([]string{"overwrite", "rename"}, t.Form.OnDuplicate) {
		errs = append(errs, "OnDuplicate must be one of overwrite, rename")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateFile) Do(c *canvasapi.Canvas) (*models.File, error) {
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
