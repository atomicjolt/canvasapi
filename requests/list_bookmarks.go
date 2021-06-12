package requests

import (
	"encoding/json"
	"io/ioutil"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListBookmarks Returns the paginated list of bookmarks.
// https://canvas.instructure.com/doc/api/bookmarks.html
//
type ListBookmarks struct {
}

func (t *ListBookmarks) GetMethod() string {
	return "GET"
}

func (t *ListBookmarks) GetURLPath() string {
	return ""
}

func (t *ListBookmarks) GetQuery() (string, error) {
	return "", nil
}

func (t *ListBookmarks) GetBody() (string, error) {
	return "", nil
}

func (t *ListBookmarks) HasErrors() error {
	return nil
}

func (t *ListBookmarks) Do(c *canvasapi.Canvas) ([]*models.Bookmark, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Bookmark{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
