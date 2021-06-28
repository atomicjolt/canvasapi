package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSingleTopicGroups Returns data on an individual discussion topic. See the List action for the response formatting.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
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
type GetSingleTopicGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		TopicID string `json:"topic_id" url:"topic_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of all_dates, sections, sections_user_count, overrides
	} `json:"query"`
}

func (t *GetSingleTopicGroups) GetMethod() string {
	return "GET"
}

func (t *GetSingleTopicGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics/{topic_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *GetSingleTopicGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSingleTopicGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleTopicGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleTopicGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
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

func (t *GetSingleTopicGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
