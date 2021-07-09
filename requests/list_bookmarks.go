package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

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

func (t *ListBookmarks) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListBookmarks) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListBookmarks) HasErrors() error {
	return nil
}

func (t *ListBookmarks) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Bookmark, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Bookmark{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
