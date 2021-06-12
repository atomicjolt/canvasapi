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

// UpdateBookmark Updates a bookmark
// https://canvas.instructure.com/doc/api/bookmarks.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Name (Optional) The name of the bookmark
// # Url (Optional) The url of the bookmark
// # Position (Optional) The position of the bookmark. Defaults to the bottom.
// # Data (Optional) The data associated with the bookmark
//
type UpdateBookmark struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Name     string `json:"name"`     //  (Optional)
		Url      string `json:"url"`      //  (Optional)
		Position int64  `json:"position"` //  (Optional)
		Data     string `json:"data"`     //  (Optional)
	} `json:"form"`
}

func (t *UpdateBookmark) GetMethod() string {
	return "PUT"
}

func (t *UpdateBookmark) GetURLPath() string {
	path := "users/self/bookmarks/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateBookmark) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateBookmark) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateBookmark) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateBookmark) Do(c *canvasapi.Canvas) (*models.Folder, error) {
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
