package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// ResetCourseFavorites Reset the current user's course favorites to the default
// automatically generated list of enrolled courses
// https://canvas.instructure.com/doc/api/favorites.html
//
type ResetCourseFavorites struct {
}

func (t *ResetCourseFavorites) GetMethod() string {
	return "DELETE"
}

func (t *ResetCourseFavorites) GetURLPath() string {
	return ""
}

func (t *ResetCourseFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *ResetCourseFavorites) GetBody() (string, error) {
	return "", nil
}

func (t *ResetCourseFavorites) HasErrors() error {
	return nil
}

func (t *ResetCourseFavorites) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
