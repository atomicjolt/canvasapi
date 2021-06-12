package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateSinglePollSession Update an existing poll session for this poll
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
// Path Parameters:
// # PollID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # PollSessions (Optional) The id of the course this session is associated with.
// # PollSessions (Optional) The id of the course section this session is associated with.
// # PollSessions (Optional) Whether or not results are viewable by students.
//
type UpdateSinglePollSession struct {
	Path struct {
		PollID string `json:"poll_id"` //  (Required)
		ID     string `json:"id"`      //  (Required)
	} `json:"path"`

	Form struct {
		PollSessions struct {
			CourseID         []int64 `json:"course_id"`          //  (Optional)
			CourseSectionID  []int64 `json:"course_section_id"`  //  (Optional)
			HasPublicResults []bool  `json:"has_public_results"` //  (Optional)
		} `json:"poll_sessions"`
	} `json:"form"`
}

func (t *UpdateSinglePollSession) GetMethod() string {
	return "PUT"
}

func (t *UpdateSinglePollSession) GetURLPath() string {
	path := "polls/{poll_id}/poll_sessions/{id}"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateSinglePollSession) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateSinglePollSession) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateSinglePollSession) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateSinglePollSession) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
