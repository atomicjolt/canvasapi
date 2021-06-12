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
)

// CreateFolderUsers Creates a folder in the specified context
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # UserID (Required) ID
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
type CreateFolderUsers struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Name             string    `json:"name"`               //  (Required)
		ParentFolderID   string    `json:"parent_folder_id"`   //  (Optional)
		ParentFolderPath string    `json:"parent_folder_path"` //  (Optional)
		LockAt           time.Time `json:"lock_at"`            //  (Optional)
		UnlockAt         time.Time `json:"unlock_at"`          //  (Optional)
		Locked           bool      `json:"locked"`             //  (Optional)
		Hidden           bool      `json:"hidden"`             //  (Optional)
		Position         int64     `json:"position"`           //  (Optional)
	} `json:"form"`
}

func (t *CreateFolderUsers) GetMethod() string {
	return "POST"
}

func (t *CreateFolderUsers) GetURLPath() string {
	path := "users/{user_id}/folders"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *CreateFolderUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateFolderUsers) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateFolderUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Form.Name == "" {
		errs = append(errs, "'Name' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateFolderUsers) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
