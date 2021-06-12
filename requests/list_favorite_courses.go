package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListFavoriteCourses Retrieve the paginated list of favorite courses for the current user. If the user has not chosen
// any favorites, then a selection of currently enrolled courses will be returned.
//
// See the {api:CoursesController#index List courses API} for details on accepted include[] parameters.
// https://canvas.instructure.com/doc/api/favorites.html
//
// Query Parameters:
// # ExcludeBlueprintCourses (Optional) When set, only return courses that are not configured as blueprint courses.
//
type ListFavoriteCourses struct {
	Query struct {
		ExcludeBlueprintCourses bool `json:"exclude_blueprint_courses"` //  (Optional)
	} `json:"query"`
}

func (t *ListFavoriteCourses) GetMethod() string {
	return "GET"
}

func (t *ListFavoriteCourses) GetURLPath() string {
	return ""
}

func (t *ListFavoriteCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListFavoriteCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListFavoriteCourses) HasErrors() error {
	return nil
}

func (t *ListFavoriteCourses) Do(c *canvasapi.Canvas) ([]*models.Course, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Course{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
