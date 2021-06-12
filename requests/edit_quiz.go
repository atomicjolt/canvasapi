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

// EditQuiz Modify an existing quiz. See the documentation for quiz creation.
//
// Additional arguments:
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Quiz (Optional) If true, notifies users that the quiz has changed.
//    Defaults to true
//
type EditQuiz struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Quiz struct {
			NotifyOfUpdate bool `json:"notify_of_update"` //  (Optional)
		} `json:"quiz"`
	} `json:"form"`
}

func (t *EditQuiz) GetMethod() string {
	return "PUT"
}

func (t *EditQuiz) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *EditQuiz) GetQuery() (string, error) {
	return "", nil
}

func (t *EditQuiz) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *EditQuiz) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditQuiz) Do(c *canvasapi.Canvas) (*models.Quiz, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Quiz{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
