package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSingleTopicCourses Returns data on an individual discussion topic. See the List action for the response formatting.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TopicID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of all_dates, sections, sections_user_count, overridesIf "all_dates" is passed, all dates associated with graded discussions'
//    assignments will be included.
//    if "sections" is passed, includes the course sections that are associated
//    with the topic, if the topic is specific to certain sections of the course.
//    If "sections_user_count" is passed, then:
//      (a) If sections were asked for *and* the topic is specific to certain
//          course sections, includes the number of users in each
//          section. (as part of the section json asked for above)
//      (b) Else, includes at the root level the total number of users in the
//          topic's context (group or course) that the topic applies to.
//    If "overrides" is passed, the overrides for the assignment will be included
//
type GetSingleTopicCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of all_dates, sections, sections_user_count, overrides
	} `json:"query"`
}

func (t *GetSingleTopicCourses) GetMethod() string {
	return "GET"
}

func (t *GetSingleTopicCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *GetSingleTopicCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSingleTopicCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleTopicCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleTopicCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'TopicID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"all_dates", "sections", "sections_user_count", "overrides"}, v) {
			errs = append(errs, "Include must be one of all_dates, sections, sections_user_count, overrides")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleTopicCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
