package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateErrorReport Create a new error report documenting an experienced problem
//
// Performs the same action as when a user uses the "help -> report a problem"
// dialog.
// https://canvas.instructure.com/doc/api/error_reports.html
//
// Form Parameters:
// # Error (Required) The summary of the problem
// # Error (Optional) URL from which the report was issued
// # Error (Optional) Email address for the reporting user
// # Error (Optional) The long version of the story from the user one what they experienced
// # Error (Optional) A collection of metadata about the users' environment.  If not provided,
//    canvas will collect it based on information found in the request.
//    (Doesn't have to be HTTPENV info, could be anything JSON object that can be
//    serialized as a hash, a mobile app might include relevant metadata for
//    itself)
//
type CreateErrorReport struct {
	Form struct {
		Error struct {
			Subject  string `json:"subject" url:"subject,omitempty"`   //  (Required)
			Url      string `json:"url" url:"url,omitempty"`           //  (Optional)
			Email    string `json:"email" url:"email,omitempty"`       //  (Optional)
			Comments string `json:"comments" url:"comments,omitempty"` //  (Optional)
			HttpEnv  string `json:"http_env" url:"http_env,omitempty"` //  (Optional)
		} `json:"error" url:"error,omitempty"`
	} `json:"form"`
}

func (t *CreateErrorReport) GetMethod() string {
	return "POST"
}

func (t *CreateErrorReport) GetURLPath() string {
	return ""
}

func (t *CreateErrorReport) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateErrorReport) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateErrorReport) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateErrorReport) HasErrors() error {
	errs := []string{}
	if t.Form.Error.Subject == "" {
		errs = append(errs, "'Error' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateErrorReport) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
