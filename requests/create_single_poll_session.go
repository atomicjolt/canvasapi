package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateSinglePollSession Create a new poll session for this poll
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
// Path Parameters:
// # PollID (Required) ID
//
// Form Parameters:
// # PollSessions (Required) The id of the course this session is associated with.
// # PollSessions (Optional) The id of the course section this session is associated with.
// # PollSessions (Optional) Whether or not results are viewable by students.
//
type CreateSinglePollSession struct {
	Path struct {
		PollID string `json:"poll_id"` //  (Required)
	} `json:"path"`

	Form struct {
		PollSessions struct {
			CourseID         []int64 `json:"course_id"`          //  (Required)
			CourseSectionID  []int64 `json:"course_section_id"`  //  (Optional)
			HasPublicResults []bool  `json:"has_public_results"` //  (Optional)
		} `json:"poll_sessions"`
	} `json:"form"`
}

func (t *CreateSinglePollSession) GetMethod() string {
	return "POST"
}

func (t *CreateSinglePollSession) GetURLPath() string {
	path := "polls/{poll_id}/poll_sessions"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	return path
}

func (t *CreateSinglePollSession) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSinglePollSession) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateSinglePollSession) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if t.Form.PollSessions.CourseID == nil {
		errs = append(errs, "'PollSessions' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSinglePollSession) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
