package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
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
		PollID        string `json:"poll_id" url:"poll_id,omitempty"`                 //  (Required)
		PollSessionID string `json:"poll_session_id" url:"poll_session_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		PollSubmissions struct {
			PollChoiceID []int64 `json:"poll_choice_id" url:"poll_choice_id,omitempty"` //  (Required)
		} `json:"poll_submissions" url:"poll_submissions,omitempty"`
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

func (t *CreateSinglePollSubmission) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateSinglePollSubmission) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
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
