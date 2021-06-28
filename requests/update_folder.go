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

// UpdateFolder Updates a folder
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Name (Optional) The new name of the folder
// # ParentFolderID (Optional) The id of the folder to move this folder into. The new folder must be in the same context as the original parent folder.
// # LockAt (Optional) The datetime to lock the folder at
// # UnlockAt (Optional) The datetime to unlock the folder at
// # Locked (Optional) Flag the folder as locked
// # Hidden (Optional) Flag the folder as hidden
// # Position (Optional) Set an explicit sort position for the folder
//
type UpdateFolder struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Name           string    `json:"name" url:"name,omitempty"`                         //  (Optional)
		ParentFolderID string    `json:"parent_folder_id" url:"parent_folder_id,omitempty"` //  (Optional)
		LockAt         time.Time `json:"lock_at" url:"lock_at,omitempty"`                   //  (Optional)
		UnlockAt       time.Time `json:"unlock_at" url:"unlock_at,omitempty"`               //  (Optional)
		Locked         bool      `json:"locked" url:"locked,omitempty"`                     //  (Optional)
		Hidden         bool      `json:"hidden" url:"hidden,omitempty"`                     //  (Optional)
		Position       int64     `json:"position" url:"position,omitempty"`                 //  (Optional)
	} `json:"form"`
}

func (t *UpdateFolder) GetMethod() string {
	return "PUT"
}

func (t *UpdateFolder) GetURLPath() string {
	path := "folders/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateFolder) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateFolder) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateFolder) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateFolder) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateFolder) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
