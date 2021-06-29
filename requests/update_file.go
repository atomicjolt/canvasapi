package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
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
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Name (Optional) The new display name of the file, with a limit of 255 characters.
// # Form.ParentFolderID (Optional) The id of the folder to move this file into.
//    The new folder must be in the same context as the original parent folder.
//    If the file is in a context without folders this does not apply.
// # Form.OnDuplicate (Optional) . Must be one of overwrite, renameIf the file is moved to a folder containing a file with the same name,
//    or renamed to a name matching an existing file, the API call will fail
//    unless this parameter is supplied.
//
//    "overwrite":: Replace the existing file with the same name
//    "rename":: Add a qualifier to make the new filename unique
// # Form.LockAt (Optional) The datetime to lock the file at
// # Form.UnlockAt (Optional) The datetime to unlock the file at
// # Form.Locked (Optional) Flag the file as locked
// # Form.Hidden (Optional) Flag the file as hidden
//
type UpdateFile struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Name           string    `json:"name" url:"name,omitempty"`                         //  (Optional)
		ParentFolderID string    `json:"parent_folder_id" url:"parent_folder_id,omitempty"` //  (Optional)
		OnDuplicate    string    `json:"on_duplicate" url:"on_duplicate,omitempty"`         //  (Optional) . Must be one of overwrite, rename
		LockAt         time.Time `json:"lock_at" url:"lock_at,omitempty"`                   //  (Optional)
		UnlockAt       time.Time `json:"unlock_at" url:"unlock_at,omitempty"`               //  (Optional)
		Locked         bool      `json:"locked" url:"locked,omitempty"`                     //  (Optional)
		Hidden         bool      `json:"hidden" url:"hidden,omitempty"`                     //  (Optional)
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

func (t *UpdateFile) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateFile) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateFile) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.OnDuplicate != "" && !string_utils.Include([]string{"overwrite", "rename"}, t.Form.OnDuplicate) {
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
