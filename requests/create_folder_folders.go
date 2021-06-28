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
)

// CreateFolderFolders Creates a folder in the specified context
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # FolderID (Required) ID
//
// Form Parameters:
// # Name (Required) The name of the folder
// # ParentFolderID (Optional) The id of the folder to store the file in. If this and parent_folder_path are sent an error will be returned. If neither is given, a default folder will be used.
// # ParentFolderPath (Optional) The path of the folder to store the new folder in. The path separator is the forward slash `/`, never a back slash. The parent folder will be created if it does not already exist. This parameter only applies to new folders in a context that has folders, such as a user, a course, or a group. If this and parent_folder_id are sent an error will be returned. If neither is given, a default folder will be used.
// # LockAt (Optional) The datetime to lock the folder at
// # UnlockAt (Optional) The datetime to unlock the folder at
// # Locked (Optional) Flag the folder as locked
// # Hidden (Optional) Flag the folder as hidden
// # Position (Optional) Set an explicit sort position for the folder
//
type CreateFolderFolders struct {
	Path struct {
		FolderID string `json:"folder_id" url:"folder_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Name             string    `json:"name" url:"name,omitempty"`                             //  (Required)
		ParentFolderID   string    `json:"parent_folder_id" url:"parent_folder_id,omitempty"`     //  (Optional)
		ParentFolderPath string    `json:"parent_folder_path" url:"parent_folder_path,omitempty"` //  (Optional)
		LockAt           time.Time `json:"lock_at" url:"lock_at,omitempty"`                       //  (Optional)
		UnlockAt         time.Time `json:"unlock_at" url:"unlock_at,omitempty"`                   //  (Optional)
		Locked           bool      `json:"locked" url:"locked,omitempty"`                         //  (Optional)
		Hidden           bool      `json:"hidden" url:"hidden,omitempty"`                         //  (Optional)
		Position         int64     `json:"position" url:"position,omitempty"`                     //  (Optional)
	} `json:"form"`
}

func (t *CreateFolderFolders) GetMethod() string {
	return "POST"
}

func (t *CreateFolderFolders) GetURLPath() string {
	path := "folders/{folder_id}/folders"
	path = strings.ReplaceAll(path, "{folder_id}", fmt.Sprintf("%v", t.Path.FolderID))
	return path
}

func (t *CreateFolderFolders) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateFolderFolders) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateFolderFolders) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateFolderFolders) HasErrors() error {
	errs := []string{}
	if t.Path.FolderID == "" {
		errs = append(errs, "'FolderID' is required")
	}
	if t.Form.Name == "" {
		errs = append(errs, "'Name' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateFolderFolders) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
