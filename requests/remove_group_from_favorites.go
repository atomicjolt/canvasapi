package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// RemoveGroupFromFavorites Remove a group from the current user's favorites.
// https://canvas.instructure.com/doc/api/favorites.html
//
// Path Parameters:
// # Path.ID (Required) the ID or SIS ID of the group to remove
//
type RemoveGroupFromFavorites struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *RemoveGroupFromFavorites) GetMethod() string {
	return "DELETE"
}

func (t *RemoveGroupFromFavorites) GetURLPath() string {
	path := "users/self/favorites/groups/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RemoveGroupFromFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *RemoveGroupFromFavorites) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveGroupFromFavorites) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveGroupFromFavorites) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveGroupFromFavorites) Do(c *canvasapi.Canvas) (*models.Favorite, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Favorite{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
