package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// DeleteConcludeCourse Delete or conclude an existing course
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Event (Required) . Must be one of delete, concludeThe action to take on the course.
//
type DeleteConcludeCourse struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Event string `json:"event" url:"event,omitempty"` //  (Required) . Must be one of delete, conclude
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
	return v.Encode(), nil
}

func (t *DeleteConcludeCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteConcludeCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteConcludeCourse) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Query.Event == "" {
		errs = append(errs, "'Query.Event' is required")
	}
	if t.Query.Event != "" && !string_utils.Include([]string{"delete", "conclude"}, t.Query.Event) {
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
