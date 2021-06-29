package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateBookmark Creates a bookmark.
// https://canvas.instructure.com/doc/api/bookmarks.html
//
// Form Parameters:
// # Form.Name (Optional) The name of the bookmark
// # Form.Url (Optional) The url of the bookmark
// # Form.Position (Optional) The position of the bookmark. Defaults to the bottom.
// # Form.Data (Optional) The data associated with the bookmark
//
type CreateBookmark struct {
	Form struct {
		Name     string `json:"name" url:"name,omitempty"`         //  (Optional)
		Url      string `json:"url" url:"url,omitempty"`           //  (Optional)
		Position int64  `json:"position" url:"position,omitempty"` //  (Optional)
		Data     string `json:"data" url:"data,omitempty"`         //  (Optional)
	} `json:"form"`
}

func (t *CreateBookmark) GetMethod() string {
	return "POST"
}

func (t *CreateBookmark) GetURLPath() string {
	return ""
}

func (t *CreateBookmark) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateBookmark) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateBookmark) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateBookmark) HasErrors() error {
	return nil
}

func (t *CreateBookmark) Do(c *canvasapi.Canvas) (*models.Bookmark, error) {
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
