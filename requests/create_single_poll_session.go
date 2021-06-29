package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateSinglePollSession Create a new poll session for this poll
// https://canvas.instructure.com/doc/api/poll_sessions.html
//
// Path Parameters:
// # Path.PollID (Required) ID
//
// Form Parameters:
// # Form.PollSessions.CourseID (Required) The id of the course this session is associated with.
// # Form.PollSessions.CourseSectionID (Optional) The id of the course section this session is associated with.
// # Form.PollSessions.HasPublicResults (Optional) Whether or not results are viewable by students.
//
type CreateSinglePollSession struct {
	Path struct {
		PollID string `json:"poll_id" url:"poll_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		PollSessions struct {
			CourseID         []string `json:"course_id" url:"course_id,omitempty"`                   //  (Required)
			CourseSectionID  []string `json:"course_section_id" url:"course_section_id,omitempty"`   //  (Optional)
			HasPublicResults []string `json:"has_public_results" url:"has_public_results,omitempty"` //  (Optional)
		} `json:"poll_sessions" url:"poll_sessions,omitempty"`
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

func (t *CreateSinglePollSession) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateSinglePollSession) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateSinglePollSession) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'Path.PollID' is required")
	}
	if t.Form.PollSessions.CourseID == nil {
		errs = append(errs, "'Form.PollSessions.CourseID' is required")
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
