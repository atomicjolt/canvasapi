package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// AddCourseToFavorites Add a course to the current user's favorites.  If the course is already
// in the user's favorites, nothing happens.
// https://canvas.instructure.com/doc/api/favorites.html
//
// Path Parameters:
// # ID (Required) The ID or SIS ID of the course to add.  The current user must be
//    registered in the course.
//
type AddCourseToFavorites struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *AddCourseToFavorites) GetMethod() string {
	return "POST"
}

func (t *AddCourseToFavorites) GetURLPath() string {
	path := "users/self/favorites/courses/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AddCourseToFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *AddCourseToFavorites) GetBody() (string, error) {
	return "", nil
}

func (t *AddCourseToFavorites) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddCourseToFavorites) Do(c *canvasapi.Canvas) (*models.Favorite, error) {
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
