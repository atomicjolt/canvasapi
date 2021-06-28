package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateExternalFeedCourses Create a new external feed for the course or group.
// https://canvas.instructure.com/doc/api/announcement_external_feeds.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Url (Required) The url to the external rss or atom feed
// # HeaderMatch (Optional) If given, only feed entries that contain this string in their title will be imported
// # Verbosity (Optional) . Must be one of full, truncate, link_onlyDefaults to "full"
//
type CreateExternalFeedCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Url         string `json:"url" url:"url,omitempty"`                   //  (Required)
		HeaderMatch bool   `json:"header_match" url:"header_match,omitempty"` //  (Optional)
		Verbosity   string `json:"verbosity" url:"verbosity,omitempty"`       //  (Optional) . Must be one of full, truncate, link_only
	} `json:"form"`
}

func (t *CreateExternalFeedCourses) GetMethod() string {
	return "POST"
}

func (t *CreateExternalFeedCourses) GetURLPath() string {
	path := "courses/{course_id}/external_feeds"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateExternalFeedCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateExternalFeedCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateExternalFeedCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateExternalFeedCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Url == "" {
		errs = append(errs, "'Url' is required")
	}
	if t.Form.Verbosity != "" && !string_utils.Include([]string{"full", "truncate", "link_only"}, t.Form.Verbosity) {
		errs = append(errs, "Verbosity must be one of full, truncate, link_only")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateExternalFeedCourses) Do(c *canvasapi.Canvas) (*models.ExternalFeed, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ExternalFeed{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
