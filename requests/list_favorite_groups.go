package requests

import (
	"encoding/json"
	"io/ioutil"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListFavoriteGroups Retrieve the paginated list of favorite groups for the current user. If the user has not chosen
// any favorites, then a selection of groups that the user is a member of will be returned.
// https://canvas.instructure.com/doc/api/favorites.html
//
type ListFavoriteGroups struct {
}

func (t *ListFavoriteGroups) GetMethod() string {
	return "GET"
}

func (t *ListFavoriteGroups) GetURLPath() string {
	return ""
}

func (t *ListFavoriteGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListFavoriteGroups) GetBody() (string, error) {
	return "", nil
}

func (t *ListFavoriteGroups) HasErrors() error {
	return nil
}

func (t *ListFavoriteGroups) Do(c *canvasapi.Canvas) ([]*models.Group, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Group{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}