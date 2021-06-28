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

// ListDiscussionTopicsGroups Returns the paginated list of discussion topics for this course or group.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # GroupID (Required) ID
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
// # OrderBy (Optional) . Must be one of position, recent_activity, titleDetermines the order of the discussion topic list. Defaults to "position".
// # Scope (Optional) . Must be one of locked, unlocked, pinned, unpinnedOnly return discussion topics in the given state(s). Defaults to including
//    all topics. Filtering is done after pagination, so pages
//    may be smaller than requested if topics are filtered.
//    Can pass multiple states as comma separated string.
// # OnlyAnnouncements (Optional) Return announcements instead of discussion topics. Defaults to false
// # FilterBy (Optional) . Must be one of all, unreadThe state of the discussion topic to return. Currently only supports unread state.
// # SearchTerm (Optional) The partial title of the discussion topics to match and return.
// # ExcludeContextModuleLockedTopics (Optional) For students, exclude topics that are locked by module progression.
//    Defaults to false.
//
type ListDiscussionTopicsGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include                          []string `json:"include" url:"include,omitempty"`                                                           //  (Optional) . Must be one of all_dates, sections, sections_user_count, overrides
		OrderBy                          string   `json:"order_by" url:"order_by,omitempty"`                                                         //  (Optional) . Must be one of position, recent_activity, title
		Scope                            string   `json:"scope" url:"scope,omitempty"`                                                               //  (Optional) . Must be one of locked, unlocked, pinned, unpinned
		OnlyAnnouncements                bool     `json:"only_announcements" url:"only_announcements,omitempty"`                                     //  (Optional)
		FilterBy                         string   `json:"filter_by" url:"filter_by,omitempty"`                                                       //  (Optional) . Must be one of all, unread
		SearchTerm                       string   `json:"search_term" url:"search_term,omitempty"`                                                   //  (Optional)
		ExcludeContextModuleLockedTopics bool     `json:"exclude_context_module_locked_topics" url:"exclude_context_module_locked_topics,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListDiscussionTopicsGroups) GetMethod() string {
	return "GET"
}

func (t *ListDiscussionTopicsGroups) GetURLPath() string {
	path := "groups/{group_id}/discussion_topics"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListDiscussionTopicsGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListDiscussionTopicsGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListDiscussionTopicsGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListDiscussionTopicsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"all_dates", "sections", "sections_user_count", "overrides"}, v) {
			errs = append(errs, "Include must be one of all_dates, sections, sections_user_count, overrides")
		}
	}
	if t.Query.OrderBy != "" && !string_utils.Include([]string{"position", "recent_activity", "title"}, t.Query.OrderBy) {
		errs = append(errs, "OrderBy must be one of position, recent_activity, title")
	}
	if t.Query.Scope != "" && !string_utils.Include([]string{"locked", "unlocked", "pinned", "unpinned"}, t.Query.Scope) {
		errs = append(errs, "Scope must be one of locked, unlocked, pinned, unpinned")
	}
	if t.Query.FilterBy != "" && !string_utils.Include([]string{"all", "unread"}, t.Query.FilterBy) {
		errs = append(errs, "FilterBy must be one of all, unread")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListDiscussionTopicsGroups) Do(c *canvasapi.Canvas) ([]*models.DiscussionTopic, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.DiscussionTopic{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
