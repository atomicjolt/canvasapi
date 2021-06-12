package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateSinglePollSubmission Create a new poll submission for this poll session
// https://canvas.instructure.com/doc/api/poll_submissions.html
//
// Path Parameters:
// # PollID (Required) ID
// # PollSessionID (Required) ID
//
// Form Parameters:
// # PollSubmissions (Required) The chosen poll choice for this submission.
//
type CreateSinglePollSubmission struct {
	Path struct {
		PollID        string `json:"poll_id"`         //  (Required)
		PollSessionID string `json:"poll_session_id"` //  (Required)
	} `json:"path"`

	Form struct {
		PollSubmissions struct {
			PollChoiceID []int64 `json:"poll_choice_id"` //  (Required)
		} `json:"poll_submissions"`
	} `json:"form"`
}

func (t *CreateSinglePollSubmission) GetMethod() string {
	return "POST"
}

func (t *CreateSinglePollSubmission) GetURLPath() string {
	path := "polls/{poll_id}/poll_sessions/{poll_session_id}/poll_submissions"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{poll_session_id}", fmt.Sprintf("%v", t.Path.PollSessionID))
	return path
}

func (t *CreateSinglePollSubmission) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSinglePollSubmission) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateSinglePollSubmission) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if t.Path.PollSessionID == "" {
		errs = append(errs, "'PollSessionID' is required")
	}
	if t.Form.PollSubmissions.PollChoiceID == nil {
		errs = append(errs, "'PollSubmissions' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSinglePollSubmission) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
