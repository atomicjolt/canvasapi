package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// ResetGroupFavorites Reset the current user's group favorites to the default
// automatically generated list of enrolled group
// https://canvas.instructure.com/doc/api/favorites.html
//
type ResetGroupFavorites struct {
}

func (t *ResetGroupFavorites) GetMethod() string {
	return "DELETE"
}

func (t *ResetGroupFavorites) GetURLPath() string {
	return ""
}

func (t *ResetGroupFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *ResetGroupFavorites) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ResetGroupFavorites) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ResetGroupFavorites) HasErrors() error {
	return nil
}

func (t *ResetGroupFavorites) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
