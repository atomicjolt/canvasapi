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

// AddGroupToFavorites Add a group to the current user's favorites.  If the group is already
// in the user's favorites, nothing happens.
// https://canvas.instructure.com/doc/api/favorites.html
//
// Path Parameters:
// # Path.ID (Required) The ID or SIS ID of the group to add.  The current user must be
//    a member of the group.
//
type AddGroupToFavorites struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *AddGroupToFavorites) GetMethod() string {
	return "POST"
}

func (t *AddGroupToFavorites) GetURLPath() string {
	path := "users/self/favorites/groups/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AddGroupToFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *AddGroupToFavorites) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *AddGroupToFavorites) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *AddGroupToFavorites) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddGroupToFavorites) Do(c *canvasapi.Canvas) (*models.Favorite, error) {
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
