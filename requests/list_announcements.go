package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListAnnouncements Returns the paginated list of announcements for the given courses and date range.  Note that
// a +context_code+ field is added to the responses so you can tell which course each announcement
// belongs to.
// https://canvas.instructure.com/doc/api/announcements.html
//
// Query Parameters:
// # ContextCodes (Required) List of context_codes to retrieve announcements for (for example, +course_123+). Only courses
//    are presently supported. The call will fail unless the caller has View Announcements permission
//    in all listed courses.
// # StartDate (Optional) Only return announcements posted since the start_date (inclusive).
//    Defaults to 14 days ago. The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # EndDate (Optional) Only return announcements posted before the end_date (inclusive).
//    Defaults to 28 days from start_date. The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
//    Announcements scheduled for future posting will only be returned to course administrators.
// # ActiveOnly (Optional) Only return active announcements that have been published.
//    Applies only to requesting users that have permission to view
//    unpublished items.
//    Defaults to false for users with access to view unpublished items,
//    otherwise true and unmodifiable.
// # LatestOnly (Optional) Only return the latest announcement for each associated context.
//    The response will include at most one announcement for each
//    specified context in the context_codes[] parameter.
//    Defaults to false.
// # Include (Optional) Optional list of resources to include with the response. May include
//    a string of the name of the resource. Possible values are:
//    "sections", "sections_user_count"
//    if "sections" is passed, includes the course sections that are associated
//    with the topic, if the topic is specific to certain sections of the course.
//    If "sections_user_count" is passed, then:
//      (a) If sections were asked for *and* the topic is specific to certain
//          course sections sections, includes the number of users in each
//          section. (as part of the section json asked for above)
//      (b) Else, includes at the root level the total number of users in the
//          topic's context (group or course) that the topic applies to.
//
type ListAnnouncements struct {
	Query struct {
		ContextCodes []string  `json:"context_codes"` //  (Required)
		StartDate    time.Time `json:"start_date"`    //  (Optional)
		EndDate      time.Time `json:"end_date"`      //  (Optional)
		ActiveOnly   bool      `json:"active_only"`   //  (Optional)
		LatestOnly   bool      `json:"latest_only"`   //  (Optional)
		Include      string    `json:"include"`       //  (Optional)
	} `json:"query"`
}

func (t *ListAnnouncements) GetMethod() string {
	return "GET"
}

func (t *ListAnnouncements) GetURLPath() string {
	return ""
}

func (t *ListAnnouncements) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAnnouncements) GetBody() (string, error) {
	return "", nil
}

func (t *ListAnnouncements) HasErrors() error {
	errs := []string{}
	if t.Query.ContextCodes == nil {
		errs = append(errs, "'ContextCodes' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAnnouncements) Do(c *canvasapi.Canvas) ([]*models.DiscussionTopic, error) {
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
