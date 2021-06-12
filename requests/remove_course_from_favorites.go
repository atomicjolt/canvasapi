package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// RemoveCourseFromFavorites Remove a course from the current user's favorites.
// https://canvas.instructure.com/doc/api/favorites.html
//
// Path Parameters:
// # ID (Required) the ID or SIS ID of the course to remove
//
type RemoveCourseFromFavorites struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *RemoveCourseFromFavorites) GetMethod() string {
	return "DELETE"
}

func (t *RemoveCourseFromFavorites) GetURLPath() string {
	path := "users/self/favorites/courses/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RemoveCourseFromFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *RemoveCourseFromFavorites) GetBody() (string, error) {
	return "", nil
}

func (t *RemoveCourseFromFavorites) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveCourseFromFavorites) Do(c *canvasapi.Canvas) (*models.Favorite, error) {
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
