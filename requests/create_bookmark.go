package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateBookmark Creates a bookmark.
// https://canvas.instructure.com/doc/api/bookmarks.html
//
// Form Parameters:
// # Name (Optional) The name of the bookmark
// # Url (Optional) The url of the bookmark
// # Position (Optional) The position of the bookmark. Defaults to the bottom.
// # Data (Optional) The data associated with the bookmark
//
type CreateBookmark struct {
	Form struct {
		Name     string `json:"name"`     //  (Optional)
		Url      string `json:"url"`      //  (Optional)
		Position int64  `json:"position"` //  (Optional)
		Data     string `json:"data"`     //  (Optional)
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

func (t *CreateBookmark) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
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
