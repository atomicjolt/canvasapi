package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SendMessageToUnsubmittedOrSubmittedUsersForQuiz {
//   "body": {
//     "type": "string",
//     "description": "message body of the conversation to be created",
//     "example": "Please take the quiz."
//   },
//   "recipients": {
//     "type": "string",
//     "description": "Who to send the message to. May be either 'submitted' or 'unsubmitted'",
//     "example": "submitted"
//   },
//   "subject": {
//     "type": "string",
//     "description": "Subject of the new Conversation created",
//     "example": "ATTN: Quiz 101 Students"
//   }
// }
// https://canvas.instructure.com/doc/api/quiz_submission_user_list.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Conversations (Optional) - Body and recipients to send the message to.
//
type SendMessageToUnsubmittedOrSubmittedUsersForQuiz struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		Conversations string `json:"conversations" url:"conversations,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *SendMessageToUnsubmittedOrSubmittedUsersForQuiz) GetMethod() string {
	return "POST"
}

func (t *SendMessageToUnsubmittedOrSubmittedUsersForQuiz) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{id}/submission_users/message"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *SendMessageToUnsubmittedOrSubmittedUsersForQuiz) GetQuery() (string, error) {
	return "", nil
}

func (t *SendMessageToUnsubmittedOrSubmittedUsersForQuiz) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SendMessageToUnsubmittedOrSubmittedUsersForQuiz) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SendMessageToUnsubmittedOrSubmittedUsersForQuiz) HasErrors() error {
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

func (t *SendMessageToUnsubmittedOrSubmittedUsersForQuiz) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
