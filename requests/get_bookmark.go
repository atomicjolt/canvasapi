package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetBookmark Returns the details for a bookmark.
// https://canvas.instructure.com/doc/api/bookmarks.html
//
// Path Parameters:
// # ID (Required) ID
//
type GetBookmark struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *GetBookmark) GetMethod() string {
	return "GET"
}

func (t *GetBookmark) GetURLPath() string {
	path := "users/self/bookmarks/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetBookmark) GetQuery() (string, error) {
	return "", nil
}

func (t *GetBookmark) GetBody() (string, error) {
	return "", nil
}

func (t *GetBookmark) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetBookmark) Do(c *canvasapi.Canvas) (*models.Bookmark, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Bookmark{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
