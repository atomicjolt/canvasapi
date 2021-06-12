package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// DeleteConcludeCourse Delete or conclude an existing course
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # Event (Required) . Must be one of delete, concludeThe action to take on the course.
//
type DeleteConcludeCourse struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Query struct {
		Event string `json:"event"` //  (Required) . Must be one of delete, conclude
	} `json:"query"`
}

func (t *DeleteConcludeCourse) GetMethod() string {
	return "DELETE"
}

func (t *DeleteConcludeCourse) GetURLPath() string {
	path := "courses/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteConcludeCourse) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *DeleteConcludeCourse) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteConcludeCourse) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Query.Event == "" {
		errs = append(errs, "'Event' is required")
	}
	if !string_utils.Include([]string{"delete", "conclude"}, t.Query.Event) {
		errs = append(errs, "Event must be one of delete, conclude")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteConcludeCourse) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
