package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListQuizzesInCourse Returns the paginated list of Quizzes in this course.
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # SearchTerm (Optional) The partial title of the quizzes to match and return.
//
type ListQuizzesInCourse struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm string `json:"search_term"` //  (Optional)
	} `json:"query"`
}

func (t *ListQuizzesInCourse) GetMethod() string {
	return "GET"
}

func (t *ListQuizzesInCourse) GetURLPath() string {
	path := "courses/{course_id}/quizzes"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListQuizzesInCourse) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListQuizzesInCourse) GetBody() (string, error) {
	return "", nil
}

func (t *ListQuizzesInCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListQuizzesInCourse) Do(c *canvasapi.Canvas) ([]*models.Quiz, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Quiz{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}